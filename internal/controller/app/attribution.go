package app

import (
	"god-help-service/internal/service"
	"god-help-service/internal/util/response"

	appApi "god-help-service/api/v1/app"

	"github.com/gin-gonic/gin"
)

// AttributionController APP端归因控制器
type AttributionController struct{}

// Report 归因报告
func (c *AttributionController) Report(ctx *gin.Context) {
	req := &appApi.AttributionReportReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		response.Fail(ctx, err)
		return
	}
	err = service.Queue().PushAttributionToQueue(ctx, req.Attribution)
	if err != nil {
		response.Fail(ctx, err)
		return
	}
	response.Success(ctx, "ok")
}
