package admin

import (
	"context"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"

	adminApi "god-help-service/api/v1/admin"

	"github.com/gogf/gf/v2/errors/gerror"
)

// RoleController 角色控制器
type RoleController struct{}

// List 获取角色列表
func (c *RoleController) List(ctx context.Context, req *adminApi.RoleListReq) (*adminApi.RoleListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行查询

	res, err := service.System().GetRoleList(ctx, req)
	if err != nil {
		logger.Errorf("获取角色列表失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Detail 获取角色详情
func (c *RoleController) Detail(ctx context.Context, req *adminApi.RoleDetailReq) (*adminApi.RoleDetailRes, error) {
	// 执行查询

	role, err := service.System().GetRoleById(ctx, req.Id)
	if err != nil {
		logger.Errorf("获取角色详情失败: %v", err)
		return nil, err
	}

	if role == nil {
		return nil, gerror.New("角色不存在")
	}

	return &adminApi.RoleDetailRes{Role: role}, nil
}

// Create 创建角色
func (c *RoleController) Create(ctx context.Context, req *adminApi.RoleCreateReq) (*adminApi.RoleCreateRes, error) {
	// 执行创建

	res, err := service.System().CreateRole(ctx, req)
	if err != nil {
		logger.Errorf("创建角色失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Update 更新角色
func (c *RoleController) Update(ctx context.Context, req *adminApi.RoleUpdateReq) (*adminApi.RoleUpdateRes, error) {
	// 执行更新

	res, err := service.System().UpdateRole(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "角色不存在" {
			return nil, gerror.New("角色不存在或已被删除")
		}

		logger.Errorf("更新角色失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Delete 删除角色
func (c *RoleController) Delete(ctx context.Context, req *adminApi.RoleDeleteReq) (*adminApi.RoleDeleteRes, error) {
	// 执行删除

	res, err := service.System().DeleteRole(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "角色不存在" {
			return nil, gerror.New("角色不存在或已被删除")
		}

		logger.Errorf("删除角色失败: %v", err)
		return nil, err
	}

	return res, nil
}

// SelectList 获取角色选择列表（用于用户选择角色，无分页，只返回启用状态的角色）
func (c *RoleController) SelectList(ctx context.Context, req *adminApi.RoleSelectListReq) (*adminApi.RoleSelectListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行查询

	res, err := service.System().GetRoleSelectList(ctx)
	if err != nil {
		logger.Errorf("获取角色选择列表失败: %v", err)
		return nil, err
	}

	return res, nil
}

// UpdateStatus 更新角色状态
func (c *RoleController) UpdateStatus(ctx context.Context, req *adminApi.RoleUpdateStatusReq) (*adminApi.RoleUpdateStatusRes, error) {
	// 执行状态更新操作

	res, err := service.System().UpdateRoleStatus(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "角色不存在" {
			return nil, gerror.New("角色不存在或已被删除")
		}

		logger.Errorf("更新角色状态失败: %v", err)
		return nil, err
	}

	return res, nil
}
