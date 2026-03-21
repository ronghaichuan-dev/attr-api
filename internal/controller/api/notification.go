package api

import (
	"encoding/json"
	"god-help-service/api/v1/api"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"god-help-service/internal/util/response"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
}

func (c *NotificationController) Callback(ctx *gin.Context) {
	// 解析请求体
	req := &api.NotificationReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		response.Fail(ctx, err)
		return
	}
	marshal, _ := json.Marshal(req)

	logger.Infof("收到请求参数:%s", string(marshal))
	// 处理通知回调
	err = service.Notification().HandleCallback(ctx, req)
	if err != nil {
		response.Fail(ctx, err)
		return
	}
	response.Success(ctx, "ok")
	return
}

// GetMemoryStats 获取内存监控数据
func (c *NotificationController) GetMemoryStats(ctx *gin.Context) {
	logger.Info("GetMemoryStats被调用")
	// 获取内存监控数据
	utilStats := service.MemoryMonitor().GetStats()
	logger.Infof("获取到内存监控数据: %d", len(utilStats))

	// 返回成功响应
	response.Success(ctx, "ok", utilStats)
}
