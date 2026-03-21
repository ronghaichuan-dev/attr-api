package admin

import (
	"context"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"

	adminApi "god-help-service/api/v1/admin"

	"github.com/gogf/gf/v2/errors/gerror"
)

// AppEventLogController 事件日志控制器
type AppEventLogController struct{}

// List 获取事件日志列表
func (c *AppEventLogController) List(ctx context.Context, req *adminApi.AppEventLogListReq) (*adminApi.AppEventLogListRes, error) {
	logger.Debugf("解析后的请求参数:%s", req)

	// 执行查询
	res, err := service.Attr().GetAppEventLogList(ctx, req)
	if err != nil {
		logger.Errorf("获取事件日志列表失败:%s", err.Error())
		return nil, err
	}

	return res, nil
}

// Detail 获取事件日志详情
func (c *AppEventLogController) Detail(ctx context.Context, req *adminApi.AppEventLogDetailReq) (*adminApi.AppEventLogDetailRes, error) {
	// 执行查询
	log, err := service.Attr().GetAppEventLogById(ctx, req.Id)
	if err != nil {
		logger.Errorf("获取事件日志详情失败:%s", err.Error())
		return nil, err
	}

	if log == nil {
		return nil, gerror.New("事件日志不存在")
	}

	return &adminApi.AppEventLogDetailRes{EventLog: log}, nil
}
