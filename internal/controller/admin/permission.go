package admin

import (
	"context"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"

	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// PermissionController 权限控制器
type PermissionController struct{}

// List 获取权限列表
func (c *PermissionController) List(ctx context.Context, req *adminApi.PermissionListReq) (*adminApi.PermissionListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行查询
	res, err := service.System().GetPermissionList(ctx, req)
	if err != nil {
		logger.Errorf("获取权限列表失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Detail 获取权限详情
func (c *PermissionController) Detail(ctx context.Context, req *adminApi.PermissionDetailReq) (*adminApi.PermissionDetailRes, error) {
	// 执行查询
	permission, err := service.System().GetPermissionById(ctx, req.Id)
	if err != nil {
		logger.Errorf("获取权限详情失败: %v", err)
		return nil, err
	}

	if permission == nil {
		return nil, gerror.New("权限不存在")
	}

	return &adminApi.PermissionDetailRes{Permission: permission}, nil
}

// Create 创建权限
func (c *PermissionController) Create(ctx context.Context, req *adminApi.PermissionCreateReq) (*adminApi.PermissionCreateRes, error) {
	// 执行创建

	res, err := service.System().CreatePermission(ctx, req)
	if err != nil {
		logger.Errorf("创建权限失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Update 更新权限
func (c *PermissionController) Update(ctx context.Context, req *adminApi.PermissionUpdateReq) (*adminApi.PermissionUpdateRes, error) {
	// 执行更新

	res, err := service.System().UpdatePermission(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "权限不存在" {
			return nil, gerror.New("权限不存在或已被删除")
		}

		logger.Errorf("更新权限失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Delete 删除权限
func (c *PermissionController) Delete(ctx context.Context, req *adminApi.PermissionDeleteReq) (*adminApi.PermissionDeleteRes, error) {
	// 执行删除

	res, err := service.System().DeletePermission(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "权限不存在" {
			return nil, gerror.New("权限不存在或已被删除")
		}

		logger.Errorf("删除权限失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Tree 获取权限树形结构
func (c *PermissionController) Tree(ctx context.Context, req *adminApi.PermissionTreeReq) (*adminApi.PermissionTreeRes, error) {
	// 执行查询

	tree, err := service.System().GetPermissionTree(ctx)
	if err != nil {
		logger.Errorf("获取权限树形结构失败: %v", err)
		return nil, err
	}

	return &adminApi.PermissionTreeRes{Tree: tree}, nil
}

// UserPermissions 根据用户ID获取拥有的权限
func (c *PermissionController) UserPermissions(ctx context.Context, req *adminApi.PermissionUserReq) (*adminApi.PermissionUserRes, error) {
	// 执行查询

	permissions, err := service.System().GetPermissionsByUserId(ctx, req.UserId)
	if err != nil {
		logger.Errorf("获取用户权限失败: %v", err)
		return nil, err
	}

	// 提取权限代码列表
	var permissionCodes []string
	for _, permission := range permissions {
		// 递归提取所有权限代码
		extractPermissionCodes(permission, &permissionCodes)
	}

	return &adminApi.PermissionUserRes{PermissionCodes: permissionCodes}, nil
}

// Enable 启用权限
func (c *PermissionController) Enable(ctx context.Context, req *adminApi.PermissionEnableReq) (*adminApi.PermissionEnableRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行启用

	res, err := service.System().EnablePermission(ctx, req.Id)
	if err != nil {
		if err != nil && err.Error() == "权限不存在" {
			return nil, gerror.New("权限不存在或已被删除")
		}

		logger.Errorf("启用权限失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Disable 禁用权限
func (c *PermissionController) Disable(ctx context.Context, req *adminApi.PermissionDisableReq) (*adminApi.PermissionDisableRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行禁用

	res, err := service.System().DisablePermission(ctx, req.Id)
	if err != nil {
		if err != nil && err.Error() == "权限不存在" {
			return nil, gerror.New("权限不存在或已被删除")
		}

		logger.Errorf("禁用权限失败: %v", err)
		return nil, err
	}

	return res, nil
}

// UpdateStatus 更新权限状态
func (c *PermissionController) UpdateStatus(ctx context.Context, req *adminApi.PermissionUpdateStatusReq) (*adminApi.PermissionUpdateStatusRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行状态更新

	res, err := service.System().UpdatePermissionStatus(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "权限不存在" {
			return nil, gerror.New("权限不存在或已被删除")
		}

		logger.Errorf("更新权限状态失败: %v", err)
		return nil, err
	}

	return res, nil
}

// extractPermissionCodes 提取权限代码（非递归，因为返回的是扁平的权限列表）
func extractPermissionCodes(permission *entity.PermissionsCustom, codes *[]string) {
	if permission != nil {
		*codes = append(*codes, permission.PermissionCode)
	}
}
