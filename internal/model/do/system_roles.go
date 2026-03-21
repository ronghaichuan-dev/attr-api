// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemRoles is the golang structure of table system_roles for DAO operations like Where/Data.
type SystemRoles struct {
	g.Meta    `orm:"table:system_roles, do:true"`
	Id        any         // 角色ID，主键，自增
	RoleName  any         // 角色名称，最多100字符，不能为空
	RoleCode  any         // 角色代码，最多100字符，不能为空
	RoleDesc  any         // 角色描述，最多500字符，可选
	Status    any         // 状态，0：禁用，1：启用，默认为1
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间（软删除）
}
