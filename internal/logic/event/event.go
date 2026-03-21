package event

import (
	"context"
	appApi "god-help-service/api/v1/app"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
)

type sEvent struct {
}

func init() {
	service.RegisterEvent(NewEvent())
}

func NewEvent() *sEvent {
	return &sEvent{}
}

func (s *sEvent) Report(ctx context.Context, req *appApi.EventReportReq) {
	// 异步处理：将请求推送到消息队列
	err := service.Queue().PushEventToQueue(ctx, req)
	if err != nil {
		logger.Errorf("推送事件到队列失败:%s", err.Error())
	}
}
