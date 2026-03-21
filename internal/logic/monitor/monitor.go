package monitor

import (
	"context"
	"god-help-service/api/v1/api"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"runtime"
	"sync"
	"time"
)

// MemoryMonitor 内存监控器
type sMemoryMonitor struct {
	mu       sync.Mutex
	stats    []*api.MemoryStats
	maxStats int
}

func init() {
	service.RegisterMemoryMonitor(NewMemoryMonitor())
}

var (
	monitorOnce sync.Once
)

// NewMemoryMonitor 获取内存监控器实例
func NewMemoryMonitor() *sMemoryMonitor {
	var memoryMonitor *sMemoryMonitor
	monitorOnce.Do(func() {
		memoryMonitor = &sMemoryMonitor{
			maxStats: 10080, // 保存7天的数据，每分钟一条
		}
	})
	return memoryMonitor
}

// StartMonitoring 开始监控内存使用情况
func (m *sMemoryMonitor) StartMonitoring(ctx context.Context, interval time.Duration) {

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				m.collectStats()
			case <-ctx.Done():
				return
			}
		}
	}()
}

// collectStats 收集内存使用统计数据
func (m *sMemoryMonitor) collectStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// 计算内存使用率
	usagePercent := 0.0
	if memStats.Sys > 0 {
		usagePercent = float64(memStats.Alloc) / float64(memStats.Sys) * 100
	}

	stats := &api.MemoryStats{
		Timestamp:    time.Now(),
		Alloc:        memStats.Alloc,
		TotalAlloc:   memStats.TotalAlloc,
		Sys:          memStats.Sys,
		Mallocs:      memStats.Mallocs,
		Frees:        memStats.Frees,
		LiveObjects:  memStats.Mallocs - memStats.Frees,
		PauseTotalNs: memStats.PauseTotalNs,
		NumGC:        memStats.NumGC,
		UsagePercent: usagePercent,
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// 添加新的统计数据
	m.stats = append(m.stats, stats)

	// 如果数据量超过最大值，删除最早的数据
	if len(m.stats) > m.maxStats {
		m.stats = m.stats[len(m.stats)-m.maxStats:]
	}

	logger.Debugf("内存使用情况:%+v", stats)
}

// GetStats 获取内存使用统计数据
func (m *sMemoryMonitor) GetStats() []*api.MemoryStats {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 返回数据的副本
	stats := make([]*api.MemoryStats, len(m.stats))
	copy(stats, m.stats)
	logger.Debugf("GetStats被调用，返回数据量:%d", len(stats))
	return stats
}

// GetLatestStats 获取最新的内存使用统计数据
func (m *sMemoryMonitor) GetLatestStats() *api.MemoryStats {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.stats) == 0 {
		return nil
	}

	return m.stats[len(m.stats)-1]
}
