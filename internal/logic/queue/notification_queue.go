package queue

import (
	"context"
	"encoding/json"
	"god-help-service/api/v1/api"
	"god-help-service/internal/consts"
	"god-help-service/internal/service"
	"god-help-service/internal/service/appleapi"
	"god-help-service/internal/util"
	"god-help-service/internal/util/logger"
	sqsutil "god-help-service/internal/util/queue"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"go.uber.org/zap"
)

func (s *sQueue) PushNotificationToQueue(ctx context.Context, req *api.NotificationReq) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := appleapi.CreateFromRawNotification(req.SignedPayload)
	if err != nil {
		return err
	}

	err = service.Attr().SaveNotificationEvent(ctx, req.SignedPayload, resp)
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	queueName := sqsutil.GetQueueName(consts.NotificationQueueSubject)
	err = sqsutil.PublishToQueue(queueName, marshal)
	if err != nil {
		return err
	}

	logger.Infof("通知成功推送到SQS队列: %s, 通知唯一ID: %s 通知类型: %s", queueName, resp.NotificationUUID, resp.NotificationType)
	return nil
}

func (s *sQueue) StartNotificationConsumer() {
	ctx := context.Background()
	logger.Info("通知消费协程已启动（使用SQS）")

	queueName := sqsutil.GetQueueName(consts.NotificationQueueSubject)
	util.SafeGoWithContext(ctx, func(ctx context.Context) {
		handler := func(data []byte) error {
			msgCtx := context.Background()
			params := &appleapi.ResponseBodyV2{}
			err := json.Unmarshal(data, params)
			if err != nil {
				logger.Errorf("反序列化通知失败: %v", err)
				return nil
			}

			s.handleNotification(msgCtx, params)
			return nil
		}

		logger.Infof("启动通知类型消费者: queue=%s", queueName)
		err := sqsutil.ConsumeQueue(queueName, handler)
		if err != nil {
			logger.Errorf("订阅SQS队列失败: %s, 错误: %v", queueName, err)
		}
	})
}

func (s *sQueue) handleNotification(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	switch resp.NotificationType {
	case consts.NotificationTypeConsumptionRequest, consts.NotificationTypeDidRenew, consts.NotificationTypeOneTimeCharge:
		if resp.NotificationType == consts.NotificationTypeOneTimeCharge {
			err := service.Attr().HandleTrialFreeSubscribe(ctx, "", "", "", nil, resp)
			if err != nil {
				logger.Error(err.Error())
				logger.Error("处理一次性收费通知失败", zap.Error(err))
			}
		}
		err := service.Attr().SaveAttrSubscriptionTransaction(ctx, resp.NotificationType, resp)
		if err != nil {
			logger.Error(err.Error())
			logger.Error("保存订阅交易失败", zap.Error(err))
			return
		}
	case consts.NotificationTypeDidChangeRenewalPref:
		logger.Infof("用户更改了续订偏好")
		if resp.Subtype != "" {
			switch resp.Subtype {
			case consts.SubtypeDowngrade:
				logger.Infof("用户降级了订阅")
			case consts.SubtypeUpgrade:
				logger.Infof("用户升级了订阅")
			}
		}
	case consts.NotificationTypeDidChangeRenewalStatus:
		logger.Infof("续订状态更改通知: %s", resp.NotificationType)
		updateFiledData := g.Map{}
		if resp.Subtype != "" {
			switch resp.Subtype {
			case consts.SubtypeAutoRenewDisabled:
				logger.Info("用户禁用了订阅自动续订")
				updateFiledData["auto_renew_status"] = consts.AutoRenewStatusNo
			case consts.SubtypeAutoRenewEnabled:
				logger.Info("用户启用了订阅自动续订")
				updateFiledData["auto_renew_status"] = consts.AutoRenewStatusYes
			}
		}
		_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFiledData)
		if err != nil {
			logger.Error(err.Error())
			return
		}
	case consts.NotificationTypeDidFailToRenew:
		logger.Infof("续订失败通知")
		if resp.Subtype == consts.SubtypeGracePeriod {
			logger.Infof("订阅进入宽限期")
		}
	case consts.NotificationTypeExpired:
		logger.Infof("订阅过期通知")
		if resp.Data == nil {
			logger.Errorf("订阅过期通知：Data 为空")
			return
		}
		if resp.Data.RenewalInfo == nil {
			logger.Errorf("订阅过期通知：RenewalInfo 为空")
			return
		}
		if resp.Data.TransactionInfo == nil {
			logger.Errorf("订阅过期通知：TransactionInfo 为空")
			return
		}
		updateFiledData := g.Map{}
		updateFiledData["expires_at"] = resp.Data.RenewalInfo.GracePeriodExpiresDate
		if resp.Subtype != "" {
			switch resp.Subtype {
			case consts.SubtypeBillingRetry:
				updateFiledData["expires_reason"] = consts.ExpiresReasonRetryFinished
			case consts.SubtypePriceIncrease:
				updateFiledData["expires_reason"] = consts.ExpiresReasonPriceUpgrade
			case consts.SubtypeProductNotForSale:
				updateFiledData["expires_reason"] = consts.ExpiresReasonUnavailableProduct
			case consts.SubtypeVoluntary:
				updateFiledData["expires_reason"] = consts.ExpiresReasonUserCancelSubscribe
			}
		}
		_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFiledData)
		if err != nil {
			logger.Errorf("更新订阅字段失败: %v", err)
			return
		}
	case consts.NotificationTypeGracePeriodExpired:
		if resp.Data == nil {
			logger.Errorf("宽限期过期通知：Data 为空")
			return
		}
		if resp.Data.RenewalInfo == nil {
			logger.Errorf("宽限期过期通知：RenewalInfo 为空")
			return
		}
		if resp.Data.TransactionInfo == nil {
			logger.Errorf("宽限期过期通知：TransactionInfo 为空")
			return
		}
		updateFiledData := g.Map{}
		updateFiledData["expires_at"] = resp.Data.RenewalInfo.GracePeriodExpiresDate
		_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFiledData)
		if err != nil {
			logger.Errorf("更新订阅字段失败: %v", err)
			return
		}
	case consts.NotificationTypeOfferRedeemed:
		logger.Infof("优惠兑换通知")
	case consts.NotificationTypePriceIncrease:
		logger.Infof("价格上涨通知")
		if resp.Subtype != "" {
			switch resp.Subtype {
			case consts.SubtypeAccepted:
				logger.Infof("用户接受了订阅价格上涨")
			case consts.SubtypePending:
				logger.Infof("用户尚未接受订阅价格上涨")
			}
		}
	case consts.NotificationTypeRefund:
		logger.Infof("退款通知")
		if resp.Data == nil {
			logger.Errorf("退款通知：Data 为空")
			return
		}
		if resp.Data.TransactionInfo == nil {
			logger.Errorf("退款通知：TransactionInfo 为空")
			return
		}
		err := service.Attr().UpdateAttrDeviceField(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, nil)
		if err != nil {
			logger.Errorf("更新设备字段失败: %v", err)
			return
		}
		logger.Warnf("退款通知处理逻辑待实现，OriginalTransactionID: %s", resp.Data.TransactionInfo.OriginalTransactionID)
	case consts.NotificationTypeRefundDeclined:
		logger.Infof("退款拒绝通知")
	case consts.NotificationTypeRefundReversed:
		logger.Infof("退款撤销通知")
	case consts.NotificationTypeRenewalExtended:
		logger.Infof("续订延长通知")
	case consts.NotificationTypeRenewalExtension:
		logger.Infof("续订延长请求通知")
		if resp.Subtype != "" {
			switch resp.Subtype {
			case consts.SubtypeFailure:
				logger.Infof("订阅续订日期延长失败")
			case consts.SubtypeSummary:
				logger.Infof("订阅续订日期延长请求完成")
			}
		}
	case consts.NotificationTypeRevoke:
		logger.Infof("撤销通知")
	case consts.NotificationTypeSubscribed:
		logger.Infof("订阅通知")
		if resp.Subtype != "" {
			switch resp.Subtype {
			case consts.SubtypeInitialBuy:
			case consts.SubtypeResubscribe:
				logger.Infof("用户重新订阅")
			}
		}
	case consts.NotificationTypeTest:
		logger.Infof("测试通知")
	default:
		logger.Infof("未知通知类型:%s", resp.NotificationType)
	}
}
