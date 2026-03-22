package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"god-help-service/api/v1/app"
	"god-help-service/internal/consts"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"god-help-service/internal/util"
	"god-help-service/internal/util/logger"
	sqsutil "god-help-service/internal/util/queue"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/redis/go-redis/v9"
)

var eventQueueMap = map[string]string{
	consts.EventCodeScreenUnlock: consts.EventQueueScreenUnlock,
	consts.EventCodeSubscribe:    consts.EventQueueSubscribe,
	consts.EventCodeSubscribeFix: consts.EventQueueSubscribeFix,
	consts.EventCodeStartUp:      consts.EventQueueStartUp,
}

func (s *sQueue) GetEventQueueSubject(eventCode string) string {
	if subject, ok := eventQueueMap[eventCode]; ok {
		return subject
	}
	return consts.EventQueueStartUp
}

func (s *sQueue) PushEventToQueue(ctx context.Context, req *app.EventReportReq) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// 使用Redis SetNX防止相同EventUuid并发重复处理，锁TTL设为5分钟
	lockKey := "event:dedup:" + req.EventUuid
	locked, err := util.GetRedisClient().SetNX(ctx, lockKey, 1, 5*time.Minute).Result()
	if err != nil {
		return err
	}
	if !locked {
		logger.Infof("EventUuid重复请求，跳过处理: %s", req.EventUuid)
		return nil
	}
	responseText, err := json.Marshal(req)
	if err != nil {
		return err
	}
	logger.Infof("事件上报参数:%s", string(responseText))
	for _, payload := range req.Payload {
		data := make(map[string]interface{})
		data["appid"] = req.Appid
		data["event_uuid"] = req.EventUuid
		data["country"] = req.Country
		data["region"] = req.Region
		data["city"] = req.City
		data["rsid"] = req.Rsid
		data["event_code"] = payload.EventCode
		data["response_text"] = string(responseText)
		data["created_at"] = gtime.Now().UnixMilli()
		data["sent_at"] = req.SentAt
		_, err = dao.AttrEventLog.Ctx(ctx).Data(data).Insert()
		if err != nil {
			return err
		}
		obj := app.EventParam{
			Environment:           req.Environment,
			AppVersion:            req.AppVersion,
			Appid:                 req.Appid,
			Country:               req.Country,
			EventCode:             payload.EventCode,
			EventUuid:             payload.EventUuid,
			TransactionId:         payload.Params.TransactionId,
			OriginalTransactionId: payload.Params.OriginalTransactionId,
			Rsid:                  req.Rsid,
			CreatedAt:             req.CreatedAt,
			SentAt:                req.SentAt,
		}
		marshal, err := json.Marshal(obj)
		if err != nil {
			logger.Errorf("序列化事件对象失败: %v", err)
			return err
		}

		queueName := sqsutil.GetQueueName(consts.EventQueueSubject)
		err = sqsutil.PublishToQueue(queueName, marshal)
		if err != nil {
			return err
		}
		logger.Infof("事件成功推送到SQS队列:%s EventUuid:%s Queue:%s", req.Appid, payload.EventUuid, queueName)
	}

	return nil
}

func (s *sQueue) SetAppInfoCache(ctx context.Context, appId string, appInfo interface{}, duration time.Duration) error {
	key := consts.SubscriptionCachePrefix + appId
	data, err := json.Marshal(appInfo)
	if err != nil {
		logger.Errorf("序列化应用信息失败: %v", err)
		return err
	}

	err = util.GetRedisClient().Do(ctx, "SETEX", key, int(duration.Seconds()), data).Err()
	if err != nil {
		logger.Errorf("设置应用信息缓存失败: %v", err)
		return err
	}

	return nil
}

func (s *sQueue) GetAppInfoCache(ctx context.Context, appId string) (*entity.SystemApps, error) {
	key := consts.SubscriptionCachePrefix + appId
	result, err := util.GetRedisClient().Do(ctx, "GET", key).Result()

	if result != nil && !errors.Is(err, redis.Nil) {
		var appInfo entity.SystemApps
		err = json.Unmarshal([]byte(result.(string)), &appInfo)
		if err != nil {
			logger.Errorf("反序列化应用信息失败:%s", err.Error())
			return nil, err
		}
		return &appInfo, nil
	}

	logger.Infof("应用信息缓存未命中，从数据库获取:%s", appId)

	one, err := service.System().GetAppById(ctx, appId)
	if err != nil {
		logger.Errorf("从数据库获取应用信息失败:%s", err.Error())
		return nil, err
	}

	if one == nil {
		logger.Infof("应用不存在:%s", appId)
		return nil, nil
	}

	err = s.SetAppInfoCache(ctx, appId, one, 1*time.Hour)
	if err != nil {
		logger.Errorf("设置应用信息缓存失败:%s", err.Error())
	}

	return one, nil
}

func (s *sQueue) GetEventInfoCache(ctx context.Context, eventCode string) (*entity.AttrEvent, error) {
	key := fmt.Sprintf("event:info:%s", eventCode)
	result, err := util.GetRedisClient().Do(ctx, "GET", key).Result()

	if result != nil && !errors.Is(err, redis.Nil) {
		var eventInfo entity.AttrEvent
		err = json.Unmarshal([]byte(result.(string)), &eventInfo)
		if err != nil {
			logger.Errorf("反序列化事件信息失败:%s", err.Error())
			return nil, err
		}
		return &eventInfo, nil
	}

	logger.Infof("事件信息缓存未命中，从数据库获取:%s", eventCode)

	event, err := service.Attr().GetEventByCode(ctx, eventCode)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, nil
	}

	eventData, err := json.Marshal(event)
	if err != nil {
		return event, nil
	}

	err = util.GetRedisClient().Do(ctx, "SETEX", key, int(1*time.Hour.Seconds()), eventData).Err()
	if err != nil {
		logger.Errorf("设置事件信息缓存失败: %v", err)
	}

	return event, nil
}

func (s *sQueue) StartEventConsumer() {
	logger.Info("事件消费协程已启动（使用SQS）")

	queueName := sqsutil.GetQueueName(consts.EventQueueSubject)

	util.SafeGoWithContext(context.Background(), func(ctx context.Context) {
		handler := func(data []byte) error {
			req := &app.EventParam{}
			err := json.Unmarshal(data, req)
			if err != nil {
				logger.Errorf("反序列化事件失败: %v", err)
				return nil
			}

			logger.Infof("从SQS接收到事件: AppId=%s, EventUuid=%s, Queue=%s", req.Appid, req.EventUuid, queueName)
			s.handleEvent(ctx, req)
			return nil
		}

		logger.Infof("启动事件上报消费者: queue=%s", queueName)
		err := sqsutil.ConsumeQueue(queueName, handler)
		if err != nil {
			logger.Errorf("订阅SQS队列失败: queue=%s, error=%v", queueName, err)
		}
	})
}

func (s *sQueue) handleEvent(ctx context.Context, req *app.EventParam) {
	logger.Infof("开始处理事件: EventUuid=%s, AppId=%s", req.EventUuid, req.Appid)

	appInfo, err := s.GetAppInfoCache(ctx, req.Appid)
	if err != nil {
		logger.Errorf("获取应用信息失败: %v", err)
		s.retryEvent(ctx, req)
		return
	}
	if appInfo == nil {
		logger.Warnf("应用不存在，过滤该事件: AppId=%s", req.Appid)
		return
	}

	event, err := s.GetEventInfoCache(ctx, req.EventCode)
	if err != nil {
		logger.Errorf("获取事件信息失败: %v", err)
		s.retryEvent(ctx, req)
		return
	}
	if event == nil {
		logger.Warnf("事件不存在，过滤该事件: EventCode=%s", req.EventCode)
		return
	}
	switch req.EventCode {
	case consts.EventCodeTrialFree:
		err = s.handleTrialFreeEvent(ctx, req.Appid, req.Rsid, req.Country, req.Environment, req)
		if err != nil {
			logger.Errorf("处理试用免费事件失败: %v", err)
			s.retryEvent(ctx, req)
			return
		}
	case consts.EventCodeInstall:
		s.handleInstallEvent(ctx, req)
	case consts.EventCodeScreenUnlock:
		s.handleScreenUnlockEvent(ctx, req)
	case consts.EventCodeSubscribe:
		s.handleSubscribeEvent(ctx, req)
	case consts.EventCodeSubscribeFix:
		s.handleSubscribeFixEvent(ctx, req.OriginalTransactionId, req.Rsid)
	case consts.EventCodeStartUp:
		s.handleStartUpEvent(ctx, req)
	default:
		logger.Warnf("未知事件类型，使用默认处理: EventCode=%s", req.EventCode)
		s.handleDefaultEvent(ctx, req)
	}
	logger.Infof("事件处理成功: EventUuid=%s, AppId=%s, EventCode=%s", req.EventUuid, req.Appid, req.EventCode)
}

func (s *sQueue) handleTrialFreeEvent(ctx context.Context, appid, uuid, country, environment string, payload *app.EventParam) error {
	err := service.Attr().HandleTrialFreeSubscribe(ctx, environment, appid, uuid, payload, nil)
	if err != nil {
		logger.Error(err.Error())
	}
	err = service.Attr().CreateAttrDeviceOrUpdate(ctx, appid, uuid, country)
	if err != nil {
		logger.Error(err.Error())
	}
	return nil
}

func (s *sQueue) handleInstallEvent(ctx context.Context, req *app.EventParam) {
	logger.Infof("处理安装事件: EventUuid=%s, AppId=%s", req.EventUuid, req.Appid)
	err := service.Attr().CreateAttrDeviceOrUpdate(ctx, req.Appid, req.Rsid, req.Country)
	if err != nil {
		logger.Error(err.Error())
	}
}

func (s *sQueue) handleScreenUnlockEvent(ctx context.Context, req *app.EventParam) {
	logger.Infof("处理开屏事件: EventUuid=%s, AppId=%s", req.EventUuid, req.Appid)
}

func (s *sQueue) handleSubscribeEvent(ctx context.Context, req *app.EventParam) {
	logger.Infof("处理订阅事件: EventUuid=%s, AppId=%s", req.EventUuid, req.Appid)
	err := service.Attr().CreateAttrDeviceOrUpdate(ctx, req.Appid, req.Rsid, req.Country)
	if err != nil {
		logger.Error(err.Error())
	}
}

func (s *sQueue) handleSubscribeFixEvent(ctx context.Context, originalTransactionId, uuid string) {
	info := entity.AttrAppSubscriptions{}
	err := dao.AttrAppSubscriptions.Ctx(ctx).Fields("uuid").Where(dao.AttrAppSubscriptions.Columns().OrignialTransactionId, originalTransactionId).Scan(&info)
	if err != nil {
		return
	}
	if info.Rsid == "" {
		updateFiledData := make(map[string]interface{})
		updateFiledData["uuid"] = uuid
		_, err = service.Attr().UpdateAppSubscriptionFieldByOriginalTransactionId(ctx, nil, originalTransactionId, updateFiledData)
		if err != nil {
			return
		}
	}
}

func (s *sQueue) handleStartUpEvent(ctx context.Context, req *app.EventParam) {
	logger.Infof("处理启动事件: EventUuid=%s, AppId=%s", req.EventUuid, req.Appid)
}

func (s *sQueue) handleDefaultEvent(ctx context.Context, req *app.EventParam) {
	logger.Infof("处理默认事件: EventUuid=%s, AppId=%s", req.EventUuid, req.Appid)
}

func (s *sQueue) retryEvent(ctx context.Context, req *app.EventParam) {
	retryKey := "event:retry:" + req.EventUuid
	result, err := util.GetRedisClient().Do(ctx, "INCR", retryKey).Result()
	if err != nil {
		logger.Errorf("获取重试次数失败: %v", err)
		return
	}
	retryCount, ok := result.(int64)
	if !ok {
		logger.Errorf("获取重试次数类型转换失败: %v", result)
		return
	}
	_ = util.GetRedisClient().Do(ctx, "EXPIRE", retryKey, 86400).Err()

	if retryCount > 5 {
		logger.Errorf("事件重试次数超过限制，放弃重试:%s", req.EventUuid)
		return
	}

	delay := time.Duration(1<<(retryCount-1)) * time.Minute
	if delay > 10*time.Minute {
		delay = 10 * time.Minute
	}

	logger.Infof("事件处理失败，准备重试:%s 重试次数:%d 延迟:%d", req.EventUuid, retryCount, delay)

	util.SafeGoWithContext(context.Background(), func(ctx context.Context) {
		time.Sleep(delay)
		count, countErr := dao.AttrEventLog.Ctx(ctx).Where("event_uuid", req.EventUuid).Count()
		if countErr == nil && count > 0 {
			logger.Infof("事件已成功处理，取消重试:%s", req.EventUuid)
			return
		}

		data, marshalErr := json.Marshal(req)
		if marshalErr != nil {
			logger.Errorf("序列化重试事件失败:%s", marshalErr.Error())
			return
		}

		queueName := sqsutil.GetQueueName(consts.EventQueueSubject)
		err = sqsutil.PublishToQueue(queueName, data)
		if err != nil {
			logger.Errorf("重试事件发布到SQS队列失败:%s", err.Error())
		} else {
			logger.Infof("事件重试已重新入队:%s 重试次数:%d 队列:%s", req.EventUuid, retryCount, queueName)
		}
	})
}
