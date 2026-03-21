// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/model/entity"
)

type (
	ISystem interface {
		// GetAppById 根据ID获取应用详情
		GetAppById(ctx context.Context, appid string) (*entity.SystemApps, error)
		// GetAppList 获取应用列表
		GetAppList(ctx context.Context, req *adminApi.AppListReq) (*adminApi.AppListRes, error)
		// View 根据ID获取应用详情
		View(ctx context.Context, appid string) (*adminApi.AppDetailItem, []int, error)
		// CreateApp 创建应用
		CreateApp(ctx context.Context, req *adminApi.AppCreateReq, userId int) (*adminApi.AppCreateRes, error)
		// UpdateApp 更新应用信息
		UpdateApp(ctx context.Context, req *adminApi.AppUpdateReq, userId int) (*adminApi.AppUpdateRes, error)
		// DeleteApp 删除应用（软删除）
		DeleteApp(ctx context.Context, req *adminApi.AppDeleteReq) (*adminApi.AppDeleteRes, error)
		// GetAppSelectList 获取应用下拉选项列表
		GetAppSelectList(ctx context.Context) ([]*adminApi.AppSelectOption, error)
		// GetAppByAppId 统一获取系统应用信息（优先从Redis获取）
		GetAppByAppId(ctx context.Context, appid string) (*entity.SystemApps, error)
		// GetPermissionList 获取权限列表
		GetPermissionList(ctx context.Context, req *adminApi.PermissionListReq) (*adminApi.PermissionListRes, error)
		// GetPermissionById 根据ID获取权限详情
		GetPermissionById(ctx context.Context, permissionId uint) (*entity.PermissionsCustom, error)
		// GetPermissionByRoute 根据路由获取权限信息
		GetPermissionByRoute(ctx context.Context, route string) (*entity.PermissionsCustom, error)
		// CreatePermission 创建权限
		CreatePermission(ctx context.Context, req *adminApi.PermissionCreateReq) (*adminApi.PermissionCreateRes, error)
		// UpdatePermission 更新权限信息
		UpdatePermission(ctx context.Context, req *adminApi.PermissionUpdateReq) (*adminApi.PermissionUpdateRes, error)
		// DeletePermission 删除权限（软删除）
		DeletePermission(ctx context.Context, req *adminApi.PermissionDeleteReq) (*adminApi.PermissionDeleteRes, error)
		// GetPermissionTree 获取权限树形结构
		GetPermissionTree(ctx context.Context) ([]*adminApi.PermissionTreeItem, error)
		// GetPermissionsByUserId 根据用户ID获取拥有的权限
		GetPermissionsByUserId(ctx context.Context, userId uint) ([]*entity.PermissionsCustom, error)
		// EnablePermission 启用权限
		EnablePermission(ctx context.Context, id uint) (*adminApi.PermissionEnableRes, error)
		// DisablePermission 禁用权限
		DisablePermission(ctx context.Context, id uint) (*adminApi.PermissionDisableRes, error)
		// UpdatePermissionStatus 更新权限状态
		UpdatePermissionStatus(ctx context.Context, req *adminApi.PermissionUpdateStatusReq) (*adminApi.PermissionUpdateStatusRes, error)
		// GetRoleList 获取角色列表
		GetRoleList(ctx context.Context, req *adminApi.RoleListReq) (*adminApi.RoleListRes, error)
		// GetRoleById 根据ID获取角色详情
		GetRoleById(ctx context.Context, roleId uint) (*entity.RolesCustom, error)
		// CreateRole 创建角色
		CreateRole(ctx context.Context, req *adminApi.RoleCreateReq) (*adminApi.RoleCreateRes, error)
		// UpdateRole 更新角色信息
		UpdateRole(ctx context.Context, req *adminApi.RoleUpdateReq) (*adminApi.RoleUpdateRes, error)
		// DeleteRole 删除角色（软删除）
		DeleteRole(ctx context.Context, req *adminApi.RoleDeleteReq) (*adminApi.RoleDeleteRes, error)
		// GetRoleSelectList 获取角色选择列表（用于用户选择角色，无分页，只返回启用状态的角色）
		GetRoleSelectList(ctx context.Context) (*adminApi.RoleSelectListRes, error)
		// UpdateRoleStatus 更新角色状态
		UpdateRoleStatus(ctx context.Context, req *adminApi.RoleUpdateStatusReq) (*adminApi.RoleUpdateStatusRes, error)
		// GetPermissionsByRoleId 根据角色ID获取权限列表
		GetPermissionsByRoleId(ctx context.Context, roleId uint) ([]*entity.PermissionsCustom, error)
		// GetRolesByPermissionId 根据权限ID获取角色列表
		GetRolesByPermissionId(ctx context.Context, permissionId uint) ([]*entity.SystemRoles, error)
		// AssignPermissionsToRole 给角色分配权限
		AssignPermissionsToRole(ctx context.Context, roleId uint, permissionIds []uint) error
		// RemovePermissionFromRole 移除角色的权限
		RemovePermissionFromRole(ctx context.Context, roleId uint, permissionId uint) error
		// CheckPermission 检查角色是否有某个权限
		CheckPermission(ctx context.Context, roleId uint64, permissionCode string) (bool, error)
		// GetSystemSettingList 获取系统设置列表
		GetSystemSettingList(ctx context.Context, req *adminApi.SystemSettingListReq) (*adminApi.SystemSettingListRes, error)
		// GetSystemSettingById 根据ID获取系统设置详情
		GetSystemSettingById(ctx context.Context, id int) (*entity.SystemSettings, error)
		// GetSystemSettingByKey 根据Key获取系统设置详情
		GetSystemSettingByKey(ctx context.Context, key string) (*entity.SystemSettings, error)
		// CreateSystemSetting 创建系统设置
		CreateSystemSetting(ctx context.Context, req *adminApi.SystemSettingCreateReq) (*adminApi.SystemSettingCreateRes, error)
		// UpdateSystemSetting 更新系统设置
		UpdateSystemSetting(ctx context.Context, req *adminApi.SystemSettingUpdateReq) (*adminApi.SystemSettingUpdateRes, error)
		// DeleteSystemSetting 删除系统设置（软删除）
		DeleteSystemSetting(ctx context.Context, req *adminApi.SystemSettingDeleteReq) (*adminApi.SystemSettingDeleteRes, error)
		// GetSystemSettingValueByKey 根据Key获取系统设置值
		GetSystemSettingValueByKey(ctx context.Context, key string) (string, error)
		// GetUserList 获取用户列表
		GetUserList(ctx context.Context, req *adminApi.UserListReq) (*adminApi.UserListRes, error)
		// GetUserById 根据ID获取用户详情
		GetUserById(ctx context.Context, userId uint) (*entity.SystemUsers, error)
		// GetUserByUsername 根据用户名获取用户详情
		GetUserByUsername(ctx context.Context, username string) (*entity.SystemUsers, error)
		// GetUsersByIds 批量根据用户ID获取用户详情
		GetUsersByIds(ctx context.Context, userIds []int) (map[int]*entity.SystemUsers, error)
		// CreateUser 创建用户
		CreateUser(ctx context.Context, req *adminApi.UserCreateReq) (*adminApi.UserCreateRes, error)
		// UpdateUser 更新用户信息
		UpdateUser(ctx context.Context, req *adminApi.UserUpdateReq) (*adminApi.UserUpdateRes, error)
		// DeleteUser 删除用户（软删除）
		DeleteUser(ctx context.Context, req *adminApi.UserDeleteReq) (*adminApi.UserDeleteRes, error)
	}
)

var (
	localSystem ISystem
)

func System() ISystem {
	if localSystem == nil {
		panic("implement not found for interface ISystem, forgot register?")
	}
	return localSystem
}

func RegisterSystem(i ISystem) {
	localSystem = i
}
