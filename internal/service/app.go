// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"god-help-service/internal/model/entity"
)

type (
	IAppSubscriptions interface {
		// GetAppSubscriptionByDate 根据AppId和日期获取应用订阅统计信息
		GetAppSubscriptionByDate(ctx context.Context, appId string, date string) (*entity.AppSubscriptions, error)
		// UpdateAppSubscription 更新应用订阅统计信息
		UpdateAppSubscription(ctx context.Context, id int64, count int64, amount float64) error
		// CreateTable 创建应用订阅统计表（如果不存在）
		CreateTable(ctx context.Context) error
	}
)

var (
	localAppSubscriptions IAppSubscriptions
)

func AppSubscriptions() IAppSubscriptions {
	if localAppSubscriptions == nil {
		panic("implement not found for interface IAppSubscriptions, forgot register?")
	}
	return localAppSubscriptions
}

func RegisterAppSubscriptions(i IAppSubscriptions) {
	localAppSubscriptions = i
}
