package queue

import (
	"context"
	"encoding/json"
	"god-help-service/api/v1/api"
	"god-help-service/internal/consts"
	"god-help-service/internal/dao"
	"god-help-service/internal/service"
	"god-help-service/internal/service/appleapi"
	"god-help-service/internal/util"
	"god-help-service/internal/util/logger"
	sqsutil "god-help-service/internal/util/queue"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
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
	case consts.NotificationTypeConsumptionRequest:
		logger.Infof("消费请求通知")
		s.handleConsumptionRequest(ctx, resp)

	case consts.NotificationTypeDidRenew:
		logger.Infof("续订成功通知")
		s.handleDidRenew(ctx, resp)

	case consts.NotificationTypeOneTimeCharge:
		logger.Infof("一次性收费通知")
		s.handleOneTimeCharge(ctx, resp)

	case consts.NotificationTypeDidChangeRenewalPref:
		logger.Infof("用户更改了续订偏好")
		s.handleDidChangeRenewalPref(ctx, resp)

	case consts.NotificationTypeDidChangeRenewalStatus:
		logger.Infof("续订状态更改通知: %s", resp.NotificationType)
		s.handleDidChangeRenewalStatus(ctx, resp)

	case consts.NotificationTypeDidFailToRenew:
		logger.Infof("续订失败通知")
		s.handleDidFailToRenew(ctx, resp)

	case consts.NotificationTypeExpired:
		logger.Infof("订阅过期通知")
		s.handleExpired(ctx, resp)

	case consts.NotificationTypeGracePeriodExpired:
		logger.Infof("宽限期过期通知")
		s.handleGracePeriodExpired(ctx, resp)

	case consts.NotificationTypeOfferRedeemed:
		logger.Infof("优惠兑换通知")
		s.handleOfferRedeemed(ctx, resp)

	case consts.NotificationTypePriceIncrease:
		logger.Infof("价格上涨通知")
		s.handlePriceIncrease(ctx, resp)

	case consts.NotificationTypeRefund:
		logger.Infof("退款通知")
		s.handleRefund(ctx, resp)

	case consts.NotificationTypeRefundDeclined:
		logger.Infof("退款拒绝通知")

	case consts.NotificationTypeRefundReversed:
		logger.Infof("退款撤销通知")
		s.handleRefundReversed(ctx, resp)

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
		s.handleRevoke(ctx, resp)

	case consts.NotificationTypeSubscribed:
		logger.Infof("订阅通知")
		s.handleSubscribed(ctx, resp)

	case consts.NotificationTypeTest:
		logger.Infof("测试通知")

	default:
		logger.Infof("未知通知类型:%s", resp.NotificationType)
	}
}

// handleConsumptionRequest 消费请求通知
func (s *sQueue) handleConsumptionRequest(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	err := service.Attr().SaveAttrSubscriptionTransaction(ctx, resp.NotificationType, resp)
	if err != nil {
		logger.Error("保存订阅交易失败", zap.Error(err))
	}
}

// handleDidRenew 续订成功通知
func (s *sQueue) handleDidRenew(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	// 保存交易记录
	err := service.Attr().SaveAttrSubscriptionTransaction(ctx, resp.NotificationType, resp)
	if err != nil {
		logger.Error("保存订阅交易失败", zap.Error(err))
		return
	}

	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		return
	}

	// 更新订阅状态为激活
	updateSubData := g.Map{
		"status":       consts.SubscriptionStatusActive,
		"is_paid":      consts.IsPaidYes,
		"last_event_at": time.Now().Unix(),
	}
	if resp.Data.RenewalInfo != nil {
		updateSubData["auto_renew_status"] = resp.Data.RenewalInfo.AutoRenewStatus
	}
	_, err = service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateSubData)
	if err != nil {
		logger.Errorf("更新订阅字段失败: %v", err)
	}

	// 更新设备表的续订信息
	rsid := s.getRsidByOriginalTransactionId(ctx, resp.Data.TransactionInfo.OriginalTransactionID)
	appid := s.getAppIdFromResp(resp)
	if rsid != "" && appid != "" {
		deviceUpdateData := g.Map{
			"is_renew":      1,
			"renew_count":   gdb.Raw("renew_count + 1"),
			"last_renew_at": time.Now().Unix(),
		}
		err = service.Attr().UpdateAttrDeviceSubscription(ctx, rsid, appid, deviceUpdateData)
		if err != nil {
			logger.Errorf("更新设备续订信息失败: %v", err)
		}
	}
}

// handleOneTimeCharge 一次性收费通知
func (s *sQueue) handleOneTimeCharge(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	err := service.Attr().HandleTrialFreeSubscribe(ctx, "", "", "", nil, resp)
	if err != nil {
		logger.Error("处理一次性收费通知失败", zap.Error(err))
	}
	err = service.Attr().SaveAttrSubscriptionTransaction(ctx, resp.NotificationType, resp)
	if err != nil {
		logger.Error("保存订阅交易失败", zap.Error(err))
	}
}

// handleDidChangeRenewalPref 续订偏好更改通知
func (s *sQueue) handleDidChangeRenewalPref(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		return
	}

	updateFieldData := g.Map{
		"last_event_at": time.Now().Unix(),
	}

	if resp.Subtype != "" {
		switch resp.Subtype {
		case consts.SubtypeDowngrade:
			logger.Infof("用户降级了订阅")
		case consts.SubtypeUpgrade:
			logger.Infof("用户升级了订阅")
		}
	}

	// 记录新的 product_id
	if resp.Data.RenewalInfo != nil && resp.Data.RenewalInfo.AutoRenewProductID != "" {
		updateFieldData["product_id"] = resp.Data.RenewalInfo.AutoRenewProductID
	}

	_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFieldData)
	if err != nil {
		logger.Errorf("更新订阅字段失败: %v", err)
	}
}

// handleDidChangeRenewalStatus 续订状态更改通知
func (s *sQueue) handleDidChangeRenewalStatus(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		return
	}

	updateFieldData := g.Map{
		"last_event_at": time.Now().Unix(),
	}
	if resp.Subtype != "" {
		switch resp.Subtype {
		case consts.SubtypeAutoRenewDisabled:
			logger.Info("用户禁用了订阅自动续订")
			updateFieldData["auto_renew_status"] = consts.AutoRenewStatusNo
		case consts.SubtypeAutoRenewEnabled:
			logger.Info("用户启用了订阅自动续订")
			updateFieldData["auto_renew_status"] = consts.AutoRenewStatusYes
		}
	}
	_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFieldData)
	if err != nil {
		logger.Errorf("更新订阅字段失败: %v", err)
	}
}

// handleDidFailToRenew 续订失败通知
func (s *sQueue) handleDidFailToRenew(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		return
	}

	updateFieldData := g.Map{
		"last_event_at": time.Now().Unix(),
	}

	if resp.Subtype == consts.SubtypeGracePeriod {
		logger.Infof("订阅进入宽限期")
		updateFieldData["status"] = consts.SubscriptionStatusGracePeriod
	} else {
		updateFieldData["status"] = consts.SubscriptionStatusBillingRetry
	}

	_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFieldData)
	if err != nil {
		logger.Errorf("更新订阅字段失败: %v", err)
	}
}

// handleExpired 订阅过期通知
func (s *sQueue) handleExpired(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.RenewalInfo == nil || resp.Data.TransactionInfo == nil {
		logger.Errorf("订阅过期通知：数据不完整")
		return
	}

	updateFieldData := g.Map{
		"status":       consts.SubscriptionStatusExpired,
		"expires_at":   resp.Data.RenewalInfo.GracePeriodExpiresDate,
		"last_event_at": time.Now().Unix(),
	}
	if resp.Subtype != "" {
		switch resp.Subtype {
		case consts.SubtypeBillingRetry:
			updateFieldData["expires_reason"] = consts.ExpiresReasonRetryFinished
		case consts.SubtypePriceIncrease:
			updateFieldData["expires_reason"] = consts.ExpiresReasonPriceUpgrade
		case consts.SubtypeProductNotForSale:
			updateFieldData["expires_reason"] = consts.ExpiresReasonUnavailableProduct
		case consts.SubtypeVoluntary:
			updateFieldData["expires_reason"] = consts.ExpiresReasonUserCancelSubscribe
		}
	}
	_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFieldData)
	if err != nil {
		logger.Errorf("更新订阅字段失败: %v", err)
	}
}

// handleGracePeriodExpired 宽限期过期通知
func (s *sQueue) handleGracePeriodExpired(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.RenewalInfo == nil || resp.Data.TransactionInfo == nil {
		logger.Errorf("宽限期过期通知：数据不完整")
		return
	}

	updateFieldData := g.Map{
		"status":       consts.SubscriptionStatusExpired,
		"expires_at":   resp.Data.RenewalInfo.GracePeriodExpiresDate,
		"last_event_at": time.Now().Unix(),
	}
	_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFieldData)
	if err != nil {
		logger.Errorf("更新订阅字段失败: %v", err)
	}
}

// handleOfferRedeemed 优惠兑换通知
func (s *sQueue) handleOfferRedeemed(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		return
	}

	updateFieldData := g.Map{
		"last_event_at": time.Now().Unix(),
	}

	// 记录优惠信息
	if resp.Data.TransactionInfo.OfferDiscountType != "" {
		updateFieldData["offer_type"] = resp.Data.TransactionInfo.OfferDiscountType
	}
	if resp.Data.TransactionInfo.OfferID != "" {
		updateFieldData["offer_id"] = resp.Data.TransactionInfo.OfferID
	}

	_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFieldData)
	if err != nil {
		logger.Errorf("更新订阅字段失败: %v", err)
	}
}

// handlePriceIncrease 价格上涨通知
func (s *sQueue) handlePriceIncrease(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Subtype != "" {
		switch resp.Subtype {
		case consts.SubtypeAccepted:
			logger.Infof("用户接受了订阅价格上涨")
		case consts.SubtypePending:
			logger.Infof("用户尚未接受订阅价格上涨")
		}
	}
}

// handleRefund 退款通知
func (s *sQueue) handleRefund(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		logger.Errorf("退款通知：数据不完整")
		return
	}

	originalTxId := resp.Data.TransactionInfo.OriginalTransactionID

	// 保存退款交易记录
	err := service.Attr().SaveAttrSubscriptionTransaction(ctx, resp.NotificationType, resp)
	if err != nil {
		logger.Errorf("保存退款交易记录失败: %v", err)
	}

	// 更新订阅状态
	updateSubData := g.Map{
		"last_event_at": time.Now().Unix(),
	}
	_, err = service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, originalTxId, updateSubData)
	if err != nil {
		logger.Errorf("更新订阅字段失败: %v", err)
	}

	// 更新设备表的退款信息
	rsid := s.getRsidByOriginalTransactionId(ctx, originalTxId)
	appid := s.getAppIdFromResp(resp)
	if rsid != "" && appid != "" {
		deviceUpdateData := g.Map{
			"is_refund":     1,
			"last_refund_at": time.Now().Unix(),
		}
		err = service.Attr().UpdateAttrDeviceSubscription(ctx, rsid, appid, deviceUpdateData)
		if err != nil {
			logger.Errorf("更新设备退款信息失败: %v", err)
		}
	}
}

// handleRefundReversed 退款撤销通知
func (s *sQueue) handleRefundReversed(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		return
	}

	// 退款撤销，恢复设备退款状态
	rsid := s.getRsidByOriginalTransactionId(ctx, resp.Data.TransactionInfo.OriginalTransactionID)
	appid := s.getAppIdFromResp(resp)
	if rsid != "" && appid != "" {
		deviceUpdateData := g.Map{
			"is_refund": 2, // 2-否
		}
		err := service.Attr().UpdateAttrDeviceSubscription(ctx, rsid, appid, deviceUpdateData)
		if err != nil {
			logger.Errorf("更新设备退款撤销信息失败: %v", err)
		}
	}
}

// handleRevoke 撤销通知
func (s *sQueue) handleRevoke(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		return
	}

	updateFieldData := g.Map{
		"status":            consts.SubscriptionStatusRevoked,
		"revocation_date":   resp.Data.TransactionInfo.RevocationDate,
		"revocation_reason": resp.Data.TransactionInfo.RevocationReason,
		"last_event_at":     time.Now().Unix(),
	}
	_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateFieldData)
	if err != nil {
		logger.Errorf("更新订阅撤销信息失败: %v", err)
	}
}

// handleSubscribed 订阅通知
func (s *sQueue) handleSubscribed(ctx context.Context, resp *appleapi.ResponseBodyV2) {
	if resp.Data == nil || resp.Data.TransactionInfo == nil {
		return
	}

	switch resp.Subtype {
	case consts.SubtypeInitialBuy:
		logger.Infof("用户首次订阅")
		// 创建订阅记录
		err := service.Attr().HandleTrialFreeSubscribe(ctx, "", "", "", nil, resp)
		if err != nil {
			logger.Errorf("处理初始订阅失败: %v", err)
		}
		// 保存交易记录
		err = service.Attr().SaveAttrSubscriptionTransaction(ctx, resp.NotificationType, resp)
		if err != nil {
			logger.Errorf("保存订阅交易记录失败: %v", err)
		}
		// 更新设备订阅时间
		rsid := s.getRsidByOriginalTransactionId(ctx, resp.Data.TransactionInfo.OriginalTransactionID)
		appid := s.getAppIdFromResp(resp)
		if rsid != "" && appid != "" {
			deviceUpdateData := g.Map{
				"last_subscribe_at": time.Now().Unix(),
			}
			err = service.Attr().UpdateAttrDeviceSubscription(ctx, rsid, appid, deviceUpdateData)
			if err != nil {
				logger.Errorf("更新设备订阅时间失败: %v", err)
			}
		}

	case consts.SubtypeResubscribe:
		logger.Infof("用户重新订阅")
		// 更新订阅状态为激活
		updateSubData := g.Map{
			"status":       consts.SubscriptionStatusActive,
			"last_event_at": time.Now().Unix(),
		}
		if resp.Data.RenewalInfo != nil {
			updateSubData["auto_renew_status"] = resp.Data.RenewalInfo.AutoRenewStatus
		}
		_, err := service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, resp.Data.TransactionInfo.OriginalTransactionID, updateSubData)
		if err != nil {
			logger.Errorf("更新订阅状态失败: %v", err)
		}
		// 保存交易记录
		err = service.Attr().SaveAttrSubscriptionTransaction(ctx, resp.NotificationType, resp)
		if err != nil {
			logger.Errorf("保存重新订阅交易记录失败: %v", err)
		}
		// 更新设备订阅时间
		rsid := s.getRsidByOriginalTransactionId(ctx, resp.Data.TransactionInfo.OriginalTransactionID)
		appid := s.getAppIdFromResp(resp)
		if rsid != "" && appid != "" {
			deviceUpdateData := g.Map{
				"last_subscribe_at": time.Now().Unix(),
			}
			err = service.Attr().UpdateAttrDeviceSubscription(ctx, rsid, appid, deviceUpdateData)
			if err != nil {
				logger.Errorf("更新设备订阅时间失败: %v", err)
			}
		}
	}
}

// getRsidByOriginalTransactionId 通过原始交易ID获取设备RSID
func (s *sQueue) getRsidByOriginalTransactionId(ctx context.Context, originalTransactionId string) string {
	var sub struct {
		Rsid string `json:"rsid"`
	}
	err := dao.AttrAppSubscriptions.Ctx(ctx).
		Fields("rsid").
		Where("orignial_transaction_id", originalTransactionId).
		Scan(&sub)
	if err != nil || sub.Rsid == "" {
		return ""
	}
	return sub.Rsid
}

// getAppIdFromResp 从通知响应中获取AppID
func (s *sQueue) getAppIdFromResp(resp *appleapi.ResponseBodyV2) string {
	if resp.Data != nil && resp.Data.TransactionInfo != nil && resp.Data.TransactionInfo.BundleID != "" {
		return resp.Data.TransactionInfo.BundleID
	}
	if resp.AppMetadata != nil && resp.AppMetadata.BundleID != "" {
		return resp.AppMetadata.BundleID
	}
	return ""
}
