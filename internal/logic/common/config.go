package common

import (
	"context"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
)

type sConfig struct {
}

func init() {
	service.RegisterConfig(NewConfig())
}

func NewConfig() *sConfig {
	return &sConfig{}
}

func (s *sConfig) GetConfigs(ctx context.Context, keys []string) (map[string]string, error) {
	var list []*entity.SystemSettings
	err := dao.SystemSettings.Ctx(ctx).WhereIn("key", keys).Scan(&list)
	if err != nil {
		return nil, err
	}
	configs := make(map[string]string)
	for _, settings := range list {
		configs[settings.Key] = settings.Value
	}
	return configs, nil
}
