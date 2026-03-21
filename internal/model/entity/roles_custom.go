// =================================================================================
// 自定义结构体，用于角色管理，使用驼峰命名规范
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RolesCustom 是`roles`表的自定义结构体，使用驼峰命名规范
type RolesCustom struct {
	Id        int         `json:"id"         orm:"id"         description:"角色ID，主键，自增"`        // 角色ID，主键，自增
	RoleName  string      `json:"roleName"   orm:"role_name"  description:"角色名称，最多100字符，不能为空"` // 角色名称，最多100字符，不能为空
	RoleCode  string      `json:"roleCode"   orm:"role_code"  description:"角色代码，最多100字符，不能为空"` // 角色代码，最多100字符，不能为空
	RoleDesc  string      `json:"roleDesc"   orm:"role_desc"  description:"角色描述，最多500字符，可选"`   // 角色描述，最多500字符，可选
	Status    int         `json:"status"     orm:"status"     description:"状态，1：启用，2：禁用，默认为1"` // 状态，1：启用，2：禁用，默认为1
	CreatedAt *gtime.Time `json:"createdAt"  orm:"created_at" description:"创建时间"`              // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt"  orm:"updated_at" description:"更新时间"`              // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt"  orm:"deleted_at" description:"删除时间（软删除）"`         // 删除时间（软删除）
}
