package app

import (
	appApi "god-help-service/api/v1/app"
	"god-help-service/internal/service"
	"god-help-service/internal/util/response"

	"github.com/gin-gonic/gin"
)

// EventController APP端事件控制器
type EventController struct{}

// Report 事件上报通知
func (c *EventController) Report(ctx *gin.Context) {
	var req appApi.EventReportReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	service.Event().Report(ctx, &req)
	// 快速返回成功响应
	response.Success(ctx, "ok")
	return
}
