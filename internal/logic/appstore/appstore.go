package appstore

import (
	"context"
	"fmt"
	"god-help-service/internal/service"
	"god-help-service/internal/service/appleapi"
	"god-help-service/internal/util"
)

type sAppStore struct {
}

func init() {
	service.RegisterAppStore(NewAppStore())
}

func NewAppStore() *sAppStore {
	return &sAppStore{}
}

func (s *sAppStore) GetSkuPrice(ctx context.Context) {
	// 从配置文件中获取env值
	env, err := util.GetAppStoreEnv(ctx)
	if err != nil {
		return
	}
	// 根据配置的env值选择对应的Apple API环境
	appleEnv := appleapi.Sandbox
	if env == "production" {
		appleEnv = appleapi.Production
	}

	client, err := appleapi.NewAppStoreServerAPI(appleEnv, "", "", "", "")
	if err != nil {
		return
	}
	history, err := client.GetNotificationHistory(ctx, nil)
	if err != nil {
		return
	}
	fmt.Println(history)
}
