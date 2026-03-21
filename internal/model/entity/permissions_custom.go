// =================================================================================
// 自定义结构体，用于权限管理，使用驼峰命名规范
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionsCustom 是`permissions`表的自定义结构体，使用驼峰命名规范
type PermissionsCustom struct {
	Id             int         `json:"id"              orm:"id"              description:"权限ID，主键，自增"`            // 权限ID，主键，自增
	PermissionName string      `json:"permissionName"  orm:"permission_name" description:"权限名称，最多100字符，不能为空"`     // 权限名称，最多100字符，不能为空
	PermissionCode string      `json:"permissionCode"  orm:"permission_code" description:"权限代码，最多100字符，不能为空"`     // 权限代码，最多100字符，不能为空
	PermissionDesc string      `json:"permissionDesc"  orm:"permission_desc" description:"权限描述，最多500字符，可选"`       // 权限描述，最多500字符，可选
	Module         string      `json:"module"          orm:"module"          description:"所属模块，最多100字符，可选"`       // 所属模块，最多100字符，可选
	Status         int         `json:"status"          orm:"status"          description:"状态，0：禁用，1：启用，默认为1"`     // 状态，0：禁用，1：启用，默认为1
	CreatedAt      *gtime.Time `json:"createdAt"       orm:"created_at"      description:"创建时间"`                  // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"       orm:"updated_at"      description:"更新时间"`                  // 更新时间
	DeletedAt      *gtime.Time `json:"deletedAt"       orm:"deleted_at"      description:"删除时间（软删除）"`             // 删除时间（软删除）
	Route          string      `json:"route"           orm:"route"           description:"权限对应的路由"`               // 权限对应的路由
	ParentId       int         `json:"parentId"        orm:"parent_id"       description:"父级权限ID，0表示顶级权限"`        // 父级权限ID，0表示顶级权限
	Level          int         `json:"level"           orm:"level"           description:"权限级别，1表示顶级，2表示二级，以此类推"` // 权限级别，1表示顶级，2表示二级，以此类推
}
