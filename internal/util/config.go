package util

import (
	"context"

	g "github.com/gogf/gf/v2/frame/g"
)

// GetAppStoreEnv 获取App Store环境配置
func GetAppStoreEnv(ctx context.Context) (string, error) {
	envVar, err := g.Cfg().Get(ctx, "server.env")
	if err != nil {
		return "", err
	}
	return envVar.String(), nil
}

// GetConfigString 获取字符串类型的配置值
func GetConfigString(ctx context.Context, key string, defaultValue string) string {
	val, err := g.Cfg().Get(ctx, key)
	if err != nil {
		return defaultValue
	}
	return val.String()
}

// GetConfigInt 获取整数类型的配置值
func GetConfigInt(ctx context.Context, key string, defaultValue int) int {
	val, err := g.Cfg().Get(ctx, key)
	if err != nil {
		return defaultValue
	}
	return val.Int()
}

// GetConfigBool 获取布尔类型的配置值
func GetConfigBool(ctx context.Context, key string, defaultValue bool) bool {
	val, err := g.Cfg().Get(ctx, key)
	if err != nil {
		return defaultValue
	}
	return val.Bool()
}
