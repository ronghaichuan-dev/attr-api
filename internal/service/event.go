// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	appApi "god-help-service/api/v1/app"
)

type (
	IEvent interface {
		Report(ctx context.Context, req *appApi.EventReportReq)
	}
)

var (
	localEvent IEvent
)

func Event() IEvent {
	if localEvent == nil {
		panic("implement not found for interface IEvent, forgot register?")
	}
	return localEvent
}

func RegisterEvent(i IEvent) {
	localEvent = i
}
