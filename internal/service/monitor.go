// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"god-help-service/api/v1/api"
	"time"
)

type (
	IMemoryMonitor interface {
		// StartMonitoring 开始监控内存使用情况
		StartMonitoring(ctx context.Context, interval time.Duration)
		// GetStats 获取内存使用统计数据
		GetStats() []*api.MemoryStats
		// GetLatestStats 获取最新的内存使用统计数据
		GetLatestStats() *api.MemoryStats
	}
)

var (
	localMemoryMonitor IMemoryMonitor
)

func MemoryMonitor() IMemoryMonitor {
	if localMemoryMonitor == nil {
		panic("implement not found for interface IMemoryMonitor, forgot register?")
	}
	return localMemoryMonitor
}

func RegisterMemoryMonitor(i IMemoryMonitor) {
	localMemoryMonitor = i
}
