package notification

import (
	"context"
	"god-help-service/api/v1/api"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
)

type sNotification struct {
}

func init() {
	service.RegisterNotification(NewNotification())
}

func NewNotification() *sNotification {
	return &sNotification{}
}

func (s *sNotification) HandleCallback(ctx context.Context, req *api.NotificationReq) error {
	// 将通知推送到队列
	err := service.Queue().PushNotificationToQueue(ctx, req)
	if err != nil {
		logger.Errorf("推送通知到队列失败:%s", err.Error())
		return err
	}
	return nil
}
