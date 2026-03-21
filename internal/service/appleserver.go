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
	IAppleServer interface {
		GetAttributionInfo(ctx context.Context, token string, proxyURL string, username string, password string) (*api.AppleAttributionInfoResponse, string, error)
	}
)

var (
	localAppleServer IAppleServer
)

func AppleServer() IAppleServer {
	if localAppleServer == nil {
		panic("implement not found for interface IAppleServer, forgot register?")
	}
	return localAppleServer
}

func RegisterAppleServer(i IAppleServer) {
	localAppleServer = i
}
