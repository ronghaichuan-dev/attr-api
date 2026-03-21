package admin

import (
	"god-help-service/internal/model/entity"

	g "github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermission 角色权限关联模型结构体（保持向后兼容性）
type RolePermission struct {
	Id           uint        `json:"id"            gorm:"primaryKey;autoIncrement" dc:"主键ID，自增"` // 主键ID，自增
	RoleId       uint        `json:"role_id"       gorm:"not null" dc:"角色ID"`                    // 角色ID
	PermissionId uint        `json:"permission_id" gorm:"not null" dc:"权限ID"`                    // 权限ID
	CreatedAt    *gtime.Time `json:"created_at"    dc:"创建时间"`                                    // 创建时间
	UpdatedAt    *gtime.Time `json:"updated_at"    dc:"更新时间"`                                    // 更新时间
	DeletedAt    *gtime.Time `json:"deleted_at"    dc:"删除时间（软删除）"`                               // 删除时间（软删除）
}

// TableName 获取角色权限关联表表名
func (rp *RolePermission) TableName() string {
	return "role_permissions"
}

// RolePermissionListReq 角色权限关联列表请求参数结构体
type RolePermissionListReq struct {
	g.Meta       `path:"/list" method:"get" tags:"角色权限关联管理" summary:"获取角色权限关联列表"`
	Page         int  `json:"page"        form:"page"        d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size         int  `json:"size"        form:"size"        d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	RoleId       uint `json:"role_id"     form:"role_id"     d:"0"    v:"min:0" dc:"角色ID，可选"`                   // 角色ID，可选
	PermissionId uint `json:"permission_id" form:"permission_id" d:"0"   v:"min:0" dc:"权限ID，可选"`                // 权限ID，可选
}

// RolePermissionListRes 角色权限关联列表响应参数结构体
type RolePermissionListRes struct {
	Total int64                           `json:"total" dc:"总记录数"`    // 总记录数
	List  []*entity.SystemRolePermissions `json:"list" dc:"角色权限关联列表"` // 角色权限关联列表
}

// RolePermissionDetailReq 角色权限关联详情请求参数结构体
type RolePermissionDetailReq struct {
	g.Meta `path:"/detail" method:"get" tags:"角色权限关联管理" summary:"获取角色权限关联详情"`
	Id     uint `json:"id" form:"id" v:"required|min:1#ID不能为空|ID必须大于0" dc:"ID，必填"` // ID，必填
}

// RolePermissionDetailRes 角色权限关联详情响应参数结构体
type RolePermissionDetailRes struct {
	RolePermission *entity.SystemRolePermissions `json:"role_permission" dc:"角色权限关联信息"`
}

// RolePermissionAssignReq 角色分配权限请求参数结构体
type RolePermissionAssignReq struct {
	g.Meta        `path:"/assign" method:"post" tags:"角色权限关联管理" summary:"给角色分配权限"`
	RoleId        uint   `json:"role_id"         form:"role_id"         v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID，必填"` // 角色ID，必填
	PermissionIds []uint `json:"permission_ids"  form:"permission_ids"  v:"required#权限ID列表不能为空" dc:"权限ID列表，必填"`             // 权限ID列表，必填
}

// RolePermissionAssignRes 角色分配权限响应参数结构体
type RolePermissionAssignRes struct {
	RoleId        uint   `json:"role_id"        dc:"角色ID"`   // 角色ID
	PermissionIds []uint `json:"permission_ids" dc:"权限ID列表"` // 权限ID列表
}

// RolePermissionRemoveReq 移除角色权限请求参数结构体
type RolePermissionRemoveReq struct {
	g.Meta       `path:"/remove" method:"delete" tags:"角色权限关联管理" summary:"移除角色权限"`
	RoleId       uint `json:"role_id"        form:"role_id"        v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID，必填"` // 角色ID，必填
	PermissionId uint `json:"permission_id"  form:"permission_id"  v:"required|min:1#权限ID不能为空|权限ID必须大于0" dc:"权限ID，必填"` // 权限ID，必填
}

// RolePermissionRemoveRes 移除角色权限响应参数结构体
type RolePermissionRemoveRes struct {
	RoleId       uint `json:"role_id"       dc:"角色ID"` // 角色ID
	PermissionId uint `json:"permission_id" dc:"权限ID"` // 权限ID
}

// RolePermissionGetPermissionsReq 根据角色ID获取权限列表请求参数结构体
type RolePermissionGetPermissionsReq struct {
	g.Meta `path:"/permissions" method:"get" tags:"角色权限关联管理" summary:"根据角色ID获取权限列表"`
	RoleId uint `json:"role_id" form:"role_id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID，必填"` // 角色ID，必填
}

// RolePermissionGetPermissionsRes 根据角色ID获取权限列表响应参数结构体
type RolePermissionGetPermissionsRes struct {
	RoleId      uint                        `json:"role_id"      dc:"角色ID"` // 角色ID
	Permissions []*entity.PermissionsCustom `json:"permissions"  dc:"权限列表"` // 权限列表
}

// RolePermissionGetRolesReq 根据权限ID获取角色列表请求参数结构体
type RolePermissionGetRolesReq struct {
	g.Meta       `path:"/roles" method:"get" tags:"角色权限关联管理" summary:"根据权限ID获取角色列表"`
	PermissionId uint `json:"permission_id" form:"permission_id" v:"required|min:1#权限ID不能为空|权限ID必须大于0" dc:"权限ID，必填"` // 权限ID，必填
}

// RolePermissionGetRolesRes 根据权限ID获取角色列表响应参数结构体
type RolePermissionGetRolesRes struct {
	PermissionId uint                  `json:"permission_id" dc:"权限ID"` // 权限ID
	Roles        []*entity.SystemRoles `json:"roles"         dc:"角色列表"` // 角色列表
}
