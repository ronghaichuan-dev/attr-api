// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemRolePermissions is the golang structure for table system_role_permissions.
type SystemRolePermissions struct {
	Id           int         `json:"id"           orm:"id"            description:"主键ID，自增"`   // 主键ID，自增
	RoleId       int         `json:"roleId"       orm:"role_id"       description:"角色ID"`      // 角色ID
	PermissionId int         `json:"permissionId" orm:"permission_id" description:"权限ID"`      // 权限ID
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`      // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`      // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间（软删除）"` // 删除时间（软删除）
}
