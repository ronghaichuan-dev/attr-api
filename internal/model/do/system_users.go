// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUsers is the golang structure of table system_users for DAO operations like Where/Data.
type SystemUsers struct {
	g.Meta    `orm:"table:system_users, do:true"`
	Id        any         //
	Username  any         //
	Password  any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time // 删除时间（软删除）
	RoleId    any         // 角色ID，关联角色表
}
