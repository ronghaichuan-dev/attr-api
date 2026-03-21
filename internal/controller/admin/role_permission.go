package admin

import (
	"context"
	"god-help-service/internal/dao"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"

	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// RolePermissionController 角色权限关联控制器
type RolePermissionController struct{}

// List 获取角色权限关联列表
func (c *RolePermissionController) List(ctx context.Context, req *adminApi.RolePermissionListReq) (*adminApi.RolePermissionListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行查询
	res := &adminApi.RolePermissionListRes{}

	// 构建查询条件
	m := dao.SystemRoles.Ctx(ctx)

	if req.RoleId > 0 {
		m = m.Where("role_id", req.RoleId)
	}
	if req.PermissionId > 0 {
		m = m.Where("permission_id", req.PermissionId)
	}

	// 获取总记录数
	total, err := m.Count()
	if err != nil {
		logger.Errorf("获取角色权限关联列表失败: %v", err)
		return nil, err
	}
	res.Total = int64(total)

	// 分页查询
	offset := (req.Page - 1) * req.Size
	err = m.Offset(offset).Limit(req.Size).OrderDesc("created_at").Scan(&res.List)
	if err != nil {
		logger.Errorf("获取角色权限关联列表失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Detail 获取角色权限关联详情
func (c *RolePermissionController) Detail(ctx context.Context, req *adminApi.RolePermissionDetailReq) (*adminApi.RolePermissionDetailRes, error) {
	// 执行查询
	var rolePermission *entity.SystemRolePermissions
	err := dao.SystemPermissions.Ctx(ctx).
		Where("id", req.Id).
		Scan(&rolePermission)
	if err != nil {
		logger.Errorf("获取角色权限关联详情失败: %v", err)
		return nil, err
	}

	if rolePermission == nil {
		return nil, gerror.New("角色权限关联不存在")
	}

	return &adminApi.RolePermissionDetailRes{RolePermission: rolePermission}, nil
}

// Assign 给角色分配权限
func (c *RolePermissionController) Assign(ctx context.Context, req *adminApi.RolePermissionAssignReq) (*adminApi.RolePermissionAssignRes, error) {
	// 执行分配
	err := service.System().AssignPermissionsToRole(ctx, req.RoleId, req.PermissionIds)
	if err != nil {
		logger.Errorf("给角色分配权限失败: %v", err)
		return nil, err
	}

	return &adminApi.RolePermissionAssignRes{
		RoleId:        req.RoleId,
		PermissionIds: req.PermissionIds,
	}, nil
}

// Remove 移除角色权限
func (c *RolePermissionController) Remove(ctx context.Context, req *adminApi.RolePermissionRemoveReq) (*adminApi.RolePermissionRemoveRes, error) {
	// 执行移除
	err := service.System().RemovePermissionFromRole(ctx, req.RoleId, req.PermissionId)
	if err != nil {
		logger.Errorf("移除角色权限失败: %v", err)
		return nil, err
	}

	return &adminApi.RolePermissionRemoveRes{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	}, nil
}

// GetPermissions 根据角色ID获取权限列表
func (c *RolePermissionController) GetPermissions(ctx context.Context, req *adminApi.RolePermissionGetPermissionsReq) (*adminApi.RolePermissionGetPermissionsRes, error) {
	// 执行查询
	permissions, err := service.System().GetPermissionsByRoleId(ctx, req.RoleId)
	if err != nil {
		logger.Errorf("根据角色ID获取权限列表失败: %v", err)
		return nil, err
	}

	return &adminApi.RolePermissionGetPermissionsRes{
		RoleId:      req.RoleId,
		Permissions: permissions,
	}, nil
}

// GetRoles 根据权限ID获取角色列表
func (c *RolePermissionController) GetRoles(ctx context.Context, req *adminApi.RolePermissionGetRolesReq) (*adminApi.RolePermissionGetRolesRes, error) {
	// 执行查询
	roles, err := service.System().GetRolesByPermissionId(ctx, req.PermissionId)
	if err != nil {
		logger.Errorf("根据权限ID获取角色列表失败: %v", err)
		return nil, err
	}

	return &adminApi.RolePermissionGetRolesRes{
		PermissionId: req.PermissionId,
		Roles:        roles,
	}, nil
}
