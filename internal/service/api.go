// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"god-help-service/api/v1/api"
)

type (
	IEcho interface {
		// Echo 处理echo测试接口
		Echo(ctx context.Context, req *api.EchoReq) (*api.EchoRes, error)
	}
)

var (
	localEcho IEcho
)

func Echo() IEcho {
	if localEcho == nil {
		panic("implement not found for interface IEcho, forgot register?")
	}
	return localEcho
}

func RegisterEcho(i IEcho) {
	localEcho = i
}
