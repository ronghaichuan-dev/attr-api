// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUsers is the golang structure for table system_users.
type SystemUsers struct {
	Id        int         `json:"id"        orm:"id"         description:""`           //
	Username  string      `json:"username"  orm:"username"   description:""`           //
	Password  string      `json:"password"  orm:"password"   description:""`           //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`           //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`           //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间（软删除）"`  // 删除时间（软删除）
	RoleId    uint64      `json:"roleId"    orm:"role_id"    description:"角色ID，关联角色表"` // 角色ID，关联角色表
}
