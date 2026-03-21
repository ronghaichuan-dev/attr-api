package api

import "time"

// MemoryStats 内存使用统计数据
type MemoryStats struct {
	Timestamp     time.Time `json:"timestamp"`
	Alloc         uint64    `json:"alloc"`         // 当前分配的内存量
	TotalAlloc    uint64    `json:"totalAlloc"`    // 累计分配的内存量
	Sys           uint64    `json:"sys"`           // 从系统获取的内存量
	Mallocs       uint64    `json:"mallocs"`       // 分配的内存块数
	Frees         uint64    `json:"frees"`         // 释放的内存块数
	LiveObjects   uint64    `json:"liveObjects"`   // 存活的对象数
	PauseTotalNs  uint64    `json:"pauseTotalNs"`  // GC暂停的总时间
	NumGC         uint32    `json:"numGC"`         // GC的次数
	UsagePercent  float64   `json:"usagePercent"`  // 内存使用率百分比
}

// MemoryStatsRes 内存监控响应
type MemoryStatsRes struct {
	Success bool           `json:"success"`
	Data    []*MemoryStats `json:"data"`
}
