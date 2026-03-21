// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IAppStore interface {
		GetSkuPrice(ctx context.Context)
	}
)

var (
	localAppStore IAppStore
)

func AppStore() IAppStore {
	if localAppStore == nil {
		panic("implement not found for interface IAppStore, forgot register?")
	}
	return localAppStore
}

func RegisterAppStore(i IAppStore) {
	localAppStore = i
}
