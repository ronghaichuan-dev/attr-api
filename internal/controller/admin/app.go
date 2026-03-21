package admin

import (
	"context"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"

	adminApi "god-help-service/api/v1/admin"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// AppController 应用控制器
type AppController struct{}

// List 获取应用列表
func (c *AppController) List(ctx context.Context, req *adminApi.AppListReq) (*adminApi.AppListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行查询
	res, err := service.System().GetAppList(ctx, req)
	if err != nil {
		logger.Errorf("获取应用列表失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Detail 获取应用详情
func (c *AppController) Detail(ctx context.Context, req *adminApi.AppDetailReq) (*adminApi.AppDetailRes, error) {
	// 执行查询
	appDetail, events, err := service.System().View(ctx, req.AppId)
	if err != nil {
		logger.Errorf("获取应用详情失败: %v", err)
		return nil, err
	}

	if appDetail == nil {
		return nil, gerror.New("应用不存在")
	}

	return &adminApi.AppDetailRes{
		App:    appDetail,
		Events: events,
	}, nil
}

// Create 创建应用
func (c *AppController) Create(ctx context.Context, req *adminApi.AppCreateReq) (*adminApi.AppCreateRes, error) {
	// 从上下文中获取用户ID（如果没有，使用默认值）
	userId := g.RequestFromCtx(ctx).GetCtxVar("user_id").Int()
	// 如果是不需要认证的路由访问，允许userId为0
	// if userId <= 0 {
	// 	return nil, gerror.New("获取用户信息失败")
	// }

	// 执行创建
	res, err := service.System().CreateApp(ctx, req, userId)
	if err != nil {
		logger.Errorf("创建应用失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Update 更新应用
func (c *AppController) Update(ctx context.Context, req *adminApi.AppUpdateReq) (*adminApi.AppUpdateRes, error) {
	// 从上下文中获取用户ID
	userId := g.RequestFromCtx(ctx).GetCtxVar("user_id").Int()
	if userId <= 0 {
		return nil, gerror.New("获取用户信息失败")
	}

	// 执行更新
	res, err := service.System().UpdateApp(ctx, req, userId)
	if err != nil {
		if err.Error() == "应用不存在" {
			return nil, gerror.New("应用不存在或已被删除")
		}

		logger.Errorf("更新应用失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Delete 删除应用
func (c *AppController) Delete(ctx context.Context, req *adminApi.AppDeleteReq) (*adminApi.AppDeleteRes, error) {
	// 执行删除
	res, err := service.System().DeleteApp(ctx, req)
	if err != nil {
		if err.Error() == "应用不存在" {
			return nil, gerror.New("应用不存在或已被删除")
		}

		logger.Errorf("删除应用失败: %v", err)
		return nil, err
	}

	return res, nil
}

// SubscriptionTrend 获取应用订阅趋势数据
func (c *AppController) SubscriptionTrend(ctx context.Context, req *adminApi.AppSubscriptionTrendReq) (*adminApi.AppSubscriptionTrendRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	return &adminApi.AppSubscriptionTrendRes{Trend: nil}, nil
}
