// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"god-help-service/api/v1/api"
	"god-help-service/api/v1/app"
	"god-help-service/internal/model/entity"
	"time"
)

type (
	IQueue interface {
		PushAttributionToQueue(ctx context.Context, req *app.Attribution) error
		StartAttributionConsumer()
		GetEventQueueSubject(eventCode string) string
		PushEventToQueue(ctx context.Context, req *app.EventReportReq) error
		SetAppInfoCache(ctx context.Context, appId string, appInfo interface{}, duration time.Duration) error
		GetAppInfoCache(ctx context.Context, appId string) (*entity.SystemApps, error)
		GetEventInfoCache(ctx context.Context, eventCode string) (*entity.AttrEvent, error)
		StartEventConsumer()
		PushNotificationToQueue(ctx context.Context, req *api.NotificationReq) error
		StartNotificationConsumer()
	}
)

var (
	localQueue IQueue
)

func Queue() IQueue {
	if localQueue == nil {
		panic("implement not found for interface IQueue, forgot register?")
	}
	return localQueue
}

func RegisterQueue(i IQueue) {
	localQueue = i
}
