// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICrontab interface {
		// Test 测试定时任务
		Test(ctx context.Context)
		// HandleAttributionTokens 处理归因Token的定时任务
		HandleAttributionTokens(ctx context.Context)
	}
)

var (
	localCrontab ICrontab
)

func Crontab() ICrontab {
	if localCrontab == nil {
		panic("implement not found for interface ICrontab, forgot register?")
	}
	return localCrontab
}

func RegisterCrontab(i ICrontab) {
	localCrontab = i
}
