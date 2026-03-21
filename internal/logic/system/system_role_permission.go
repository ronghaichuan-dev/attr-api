package system

import (
	"context"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GetPermissionsByRoleId 根据角色ID获取权限列表
func (s *sSystem) GetPermissionsByRoleId(ctx context.Context, roleId uint) ([]*entity.PermissionsCustom, error) {
	var permissions []*entity.PermissionsCustom
	err := dao.SystemRolePermissions.Ctx(ctx).
		Where("role_id", roleId).
		Where("deleted_at IS NULL").
		LeftJoin("system_permissions", "system_permissions.id = system_role_permissions.permission_id").
		Where("system_permissions.deleted_at IS NULL").
		Fields("system_permissions.*").
		Scan(&permissions)
	return permissions, err
}

// GetRolesByPermissionId 根据权限ID获取角色列表
func (s *sSystem) GetRolesByPermissionId(ctx context.Context, permissionId uint) ([]*entity.SystemRoles, error) {
	var roles []*entity.SystemRoles
	err := dao.SystemRolePermissions.Ctx(ctx).
		Where("permission_id", permissionId).
		Where("deleted_at IS NULL").
		LeftJoin("system_roles", "system_roles.id = system_role_permissions.role_id").
		Where("system_roles.deleted_at IS NULL").
		Fields("system_roles.*").
		Scan(&roles)
	return roles, err
}

// AssignPermissionsToRole 给角色分配权限
func (s *sSystem) AssignPermissionsToRole(ctx context.Context, roleId uint, permissionIds []uint) error {
	// 先删除该角色的所有权限关联
	_, err := dao.SystemRolePermissions.Ctx(ctx).
		Where("role_id", roleId).
		Where("deleted_at IS NULL").
		Data("deleted_at", gdb.Raw("NOW()")).
		Update()
	if err != nil {
		return err
	}

	// 重新添加权限关联
	if len(permissionIds) > 0 {
		var dataList []g.Map
		for _, permissionId := range permissionIds {
			dataList = append(dataList, g.Map{
				"role_id":       roleId,
				"permission_id": permissionId,
			})
		}
		_, err = dao.SystemRolePermissions.Ctx(ctx).Data(dataList).Insert()
		if err != nil {
			return err
		}
	}

	return nil
}

// RemovePermissionFromRole 移除角色的权限
func (s *sSystem) RemovePermissionFromRole(ctx context.Context, roleId uint, permissionId uint) error {
	_, err := dao.SystemRolePermissions.Ctx(ctx).
		Where("role_id", roleId).
		Where("permission_id", permissionId).
		Where("deleted_at IS NULL").
		Data("deleted_at", gdb.Raw("NOW()")).
		Update()
	return err
}

// CheckPermission 检查角色是否有某个权限
func (s *sSystem) CheckPermission(ctx context.Context, roleId uint64, permissionCode string) (bool, error) {
	count, err := dao.SystemRolePermissions.Ctx(ctx).
		Where("system_role_permissions.role_id", roleId).
		Where("system_role_permissions.deleted_at IS NULL").
		LeftJoin("system_permissions", "system_permissions.id = system_role_permissions.permission_id").
		Where("system_permissions.permission_code", permissionCode).
		Where("system_permissions.deleted_at IS NULL").
		Where("system_permissions.status", 1).
		Count()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
