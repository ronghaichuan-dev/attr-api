package admin

import (
	"context"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"

	adminApi "god-help-service/api/v1/admin"

	"github.com/gogf/gf/v2/errors/gerror"
)

// SystemSettingController 系统设置控制器
type SystemSettingController struct{}

// List 获取系统设置列表
func (c *SystemSettingController) List(ctx context.Context, req *adminApi.SystemSettingListReq) (*adminApi.SystemSettingListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行查询

	res, err := service.System().GetSystemSettingList(ctx, req)
	if err != nil {
		logger.Errorf("获取系统设置列表失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Detail 获取系统设置详情
func (c *SystemSettingController) Detail(ctx context.Context, req *adminApi.SystemSettingDetailReq) (*adminApi.SystemSettingDetailRes, error) {
	// 执行查询

	setting, err := service.System().GetSystemSettingById(ctx, req.Id)
	if err != nil {
		logger.Errorf("获取系统设置详情失败: %v", err)
		return nil, err
	}

	if setting == nil {
		return nil, gerror.New("系统设置不存在")
	}

	return &adminApi.SystemSettingDetailRes{Setting: setting}, nil
}

// Create 创建系统设置
func (c *SystemSettingController) Create(ctx context.Context, req *adminApi.SystemSettingCreateReq) (*adminApi.SystemSettingCreateRes, error) {
	// 执行创建

	res, err := service.System().CreateSystemSetting(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "设置键已存在" {
			return nil, gerror.New("设置键已存在")
		}
		logger.Errorf("创建系统设置失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Update 更新系统设置
func (c *SystemSettingController) Update(ctx context.Context, req *adminApi.SystemSettingUpdateReq) (*adminApi.SystemSettingUpdateRes, error) {
	// 执行更新

	res, err := service.System().UpdateSystemSetting(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "系统设置不存在" {
			return nil, gerror.New("系统设置不存在或已被删除")
		}
		if err != nil && err.Error() == "设置键已存在" {
			return nil, gerror.New("设置键已存在")
		}
		logger.Errorf("更新系统设置失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Delete 删除系统设置
func (c *SystemSettingController) Delete(ctx context.Context, req *adminApi.SystemSettingDeleteReq) (*adminApi.SystemSettingDeleteRes, error) {
	// 执行删除

	res, err := service.System().DeleteSystemSetting(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "系统设置不存在" {
			return nil, gerror.New("系统设置不存在或已被删除")
		}
		logger.Errorf("删除系统设置失败: %v", err)
		return nil, err
	}

	return res, nil
}

// GetValue 根据key获取系统设置值
func (c *SystemSettingController) GetValue(ctx context.Context, req *adminApi.SystemSettingGetValueReq) (*adminApi.SystemSettingGetValueRes, error) {
	// 执行查询

	value, err := service.System().GetSystemSettingValueByKey(ctx, req.Key)
	if err != nil {
		if err != nil && err.Error() == "系统设置不存在" {
			return nil, gerror.New("系统设置不存在")
		}
		logger.Errorf("获取系统设置值失败: %v", err)
		return nil, err
	}

	return &adminApi.SystemSettingGetValueRes{Value: value}, nil
}
