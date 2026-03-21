// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemRolePermissions is the golang structure of table system_role_permissions for DAO operations like Where/Data.
type SystemRolePermissions struct {
	g.Meta       `orm:"table:system_role_permissions, do:true"`
	Id           any         // 主键ID，自增
	RoleId       any         // 角色ID
	PermissionId any         // 权限ID
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间（软删除）
}
