// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemPermissions is the golang structure of table system_permissions for DAO operations like Where/Data.
type SystemPermissions struct {
	g.Meta         `orm:"table:system_permissions, do:true"`
	Id             any         // 权限ID，主键，自增
	PermissionName any         // 权限名称，最多100字符，不能为空
	PermissionCode any         // 权限代码，最多100字符，不能为空
	PermissionDesc any         // 权限描述，最多500字符，可选
	Module         any         // 所属模块，最多100字符，可选
	Status         any         // 状态，0：禁用，1：启用，默认为1
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间（软删除）
	Route          any         // 权限对应的路由
	ParentId       any         // 父级权限ID，0表示顶级权限
	Level          any         // 权限级别，1表示顶级，2表示二级，以此类推
}
