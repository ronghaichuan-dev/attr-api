package queue

import (
	"context"
	"encoding/json"
	"god-help-service/api/v1/app"
	"god-help-service/internal/consts"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"god-help-service/internal/util"
	"god-help-service/internal/util/logger"
	sqsutil "god-help-service/internal/util/queue"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sQueue) PushAttributionToQueue(ctx context.Context, req *app.Attribution) error {

	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	logger.Infof("开始保存归因报告:%s", string(bytes))
	data := entity.AttrInstall{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return err
	}

	systemApp, err := service.System().GetAppByAppId(ctx, data.AppId)
	if err != nil {
		return err
	}
	if systemApp == nil {
		return nil
	}
	logger.Infof("systemApp:%s req App:%s", systemApp.AppToken, req.AppToken)
	if systemApp.AppToken != req.AppToken {
		return gerror.New("app token不存在")
	}
	// 使用Redis SetNX防止相同AttrUuid并发重复处理，锁TTL设为5分钟
	lockKey := "attr:dedup:" + req.AttrUuid
	locked, err := util.GetRedisClient().SetNX(ctx, lockKey, 1, 5*time.Minute).Result()
	if err != nil {
		return err
	}
	if !locked {
		logger.Infof("AttrUuid重复请求，跳过处理: %s", req.AttrUuid)
		return nil
	}

	data.TokenResponseText = "{}"
	_, err = dao.AttrInstall.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}
	queueName := sqsutil.GetQueueName(consts.AttributionQueueSubject)
	err = sqsutil.PublishToQueue(queueName, marshal)
	if err != nil {
		return err
	}
	return nil
}

func (s *sQueue) StartAttributionConsumer() {
	logger.Info("归因消费协程已启动（使用SQS）")

	queueName := sqsutil.GetQueueName(consts.AttributionQueueSubject)

	util.SafeGoWithContext(context.Background(), func(ctx context.Context) {
		handler := func(data []byte) error {
			req := &app.Attribution{}
			err := json.Unmarshal(data, req)
			if err != nil {
				logger.Errorf("反序列化事件失败: %v", err)
				return nil
			}

			logger.Infof("从SQS接收到归因数据: queue=%s", queueName)
			s.handleAttributionEvent(ctx, req)
			return nil
		}

		logger.Infof("启动归因消费者: queue=%s", queueName)
		err := sqsutil.ConsumeQueue(queueName, handler)
		if err != nil {
			logger.Errorf("订阅SQS队列失败: queue=%s, error=%v", queueName, err)
		}
	})
}

func (s *sQueue) handleAttributionEvent(ctx context.Context, req *app.Attribution) {
	err := service.Attr().HandleAttribution(ctx, req)
	if err != nil {
		logger.Error(err.Error())
	}
	err = service.Attr().CreateAttrDeviceOrUpdate(ctx, req.AppId, req.Rsid, req.Country)
	if err != nil {
		logger.Error(err.Error())
	}
}
