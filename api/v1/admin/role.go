package admin

import (
	"god-help-service/internal/model/entity"

	g "github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role 角色模型结构体（保持向后兼容性）
type Role struct {
	Id        uint        `json:"id"          gorm:"primaryKey;autoIncrement" dc:"角色ID，主键，自增"` // 角色ID，主键，自增
	RoleName  string      `json:"role_name"  gorm:"size:100;not null" dc:"角色名称，最多100字符，不能为空"`  // 角色名称，最多100字符，不能为空
	RoleCode  string      `json:"role_code"  gorm:"size:100;not null" dc:"角色代码，最多100字符，不能为空"`  // 角色代码，最多100字符，不能为空
	RoleDesc  string      `json:"role_desc"  gorm:"size:500" dc:"角色描述，最多500字符，可选"`             // 角色描述，最多500字符，可选
	Status    int         `json:"status"      gorm:"default:1" dc:"状态，1：启用，2：禁用"`              // 状态，1：启用，2：禁用，默认为1
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`                                        // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`                                        // 更新时间
	DeletedAt *gtime.Time `json:"deleted_at" dc:"删除时间（软删除）"`                                   // 删除时间（软删除）
}

// TableName 获取角色表名
func (r *Role) TableName() string {
	return "roles"
}

// RoleListReq 角色列表请求参数结构体
type RoleListReq struct {
	g.Meta   `path:"/list" method:"get" tags:"角色管理" summary:"获取角色列表"`
	Page     int    `json:"page"      form:"page"      d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size     int    `json:"size"      form:"size"      d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	RoleName string `json:"role_name" form:"role_name" d:""  v:"max-length:100" dc:"角色名称搜索，可选"`           // 角色名称搜索，可选
}

// RoleListRes 角色列表响应参数结构体
type RoleListRes struct {
	Total int64                 `json:"total" dc:"总记录数"` // 总记录数
	List  []*entity.RolesCustom `json:"list" dc:"角色列表"`  // 角色列表
}

// RoleDetailReq 角色详情请求参数结构体
type RoleDetailReq struct {
	g.Meta `path:"/detail" method:"get" tags:"角色管理" summary:"获取角色详情"`
	Id     uint `json:"id" form:"id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID，必填"` // 角色ID，必填
}

// RoleDetailRes 角色详情响应参数结构体
type RoleDetailRes struct {
	Role *entity.RolesCustom `json:"role" dc:"角色信息"`
}

// RoleCreateReq 角色创建请求参数结构体
type RoleCreateReq struct {
	g.Meta   `path:"/create" method:"post" tags:"角色管理" summary:"创建角色"`
	RoleName string `json:"role_name" form:"role_name" v:"required|max-length:100#角色名称不能为空|角色名称长度不能超过100个字符" dc:"角色名称，必填"` // 角色名称，必填
	RoleCode string `json:"role_code" form:"role_code" v:"required|max-length:100#角色代码不能为空|角色代码长度不能超过100个字符" dc:"角色代码，必填"` // 角色代码，必填
	RoleDesc string `json:"role_desc" form:"role_desc" d:""  v:"max-length:500#角色描述长度不能超过500个字符" dc:"角色描述，可选"`             // 角色描述，可选
	Status   int    `json:"status"      form:"status"      d:"1"    v:"in:1,2" dc:"状态，1：启用，2：禁用，默认为1"`                     // 状态，1：启用，2：禁用，默认为1
}

// RoleCreateRes 角色创建响应参数结构体
type RoleCreateRes struct {
	Id uint `json:"id" dc:"新创建的角色ID"` // 新创建的角色ID
}

// RoleUpdateReq 角色更新请求参数结构体
type RoleUpdateReq struct {
	g.Meta   `path:"/update" method:"put" tags:"角色管理" summary:"更新角色信息"`
	Id       uint   `json:"id"         form:"id"         v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID，必填"`   // 角色ID，必填
	RoleName string `json:"role_name" form:"role_name" d:""  v:"max-length:100#角色名称长度不能超过100个字符" dc:"角色名称，可选"` // 角色名称，可选
	RoleCode string `json:"role_code" form:"role_code" d:""  v:"max-length:100#角色代码长度不能超过100个字符" dc:"角色代码，可选"` // 角色代码，可选
	RoleDesc string `json:"role_desc" form:"role_desc" d:""  v:"max-length:500#角色描述长度不能超过500个字符" dc:"角色描述，可选"` // 角色描述，可选
	Status   int    `json:"status"      form:"status"      d:"-1"    v:"in:-1,1,2" dc:"状态，-1：不修改，1：启用，2：禁用"`   // 状态，-1：不修改，1：启用，2：禁用
}

// RoleUpdateRes 角色更新响应参数结构体
type RoleUpdateRes struct {
	Id uint `json:"id" dc:"更新的角色ID"` // 更新的角色ID
}

// RoleDeleteReq 角色删除请求参数结构体
type RoleDeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"角色管理" summary:"删除角色"`
	Id     uint `json:"id" form:"id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID，必填"` // 角色ID，必填
}

// RoleDeleteRes 角色删除响应参数结构体
type RoleDeleteRes struct {
	Id uint `json:"id" dc:"删除的角色ID"` // 删除的角色ID
}

// RoleSelectListReq 角色选择列表请求参数结构体（用于用户选择角色，无分页）
type RoleSelectListReq struct {
	g.Meta `path:"/select/list" method:"get" tags:"角色管理" summary:"获取角色选择列表（用于用户选择角色）"`
}

// RoleSelectListRes 角色选择列表响应参数结构体
type RoleSelectListRes struct {
	List []*entity.RolesCustom `json:"list" dc:"角色列表"` // 角色列表
}

// RoleCreateTableReq 角色表创建请求参数结构体
type RoleCreateTableReq struct {
	g.Meta `path:"/createTable" method:"get" tags:"角色管理" summary:"创建角色表"`
}

// RoleCreateTableRes 角色表创建响应参数结构体
type RoleCreateTableRes struct {
	Success bool `json:"success" dc:"是否成功"` // 是否成功
}

// RoleUpdateStatusReq 角色状态更新请求参数结构体
type RoleUpdateStatusReq struct {
	g.Meta `path:"/status" method:"put" tags:"角色管理" summary:"更新角色状态"`
	Id     uint `json:"id" form:"id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID，必填"`             // 角色ID，必填
	Status int  `json:"status" form:"status" v:"required|in:1,2#状态不能为空|状态值必须为1或2" dc:"状态，1：启用，2：禁用"` // 状态，1：启用，2：禁用，必填
}

// RoleUpdateStatusRes 角色状态更新响应参数结构体
type RoleUpdateStatusRes struct {
	Id     uint `json:"id" dc:"更新状态的角色ID"` // 更新状态的角色ID
	Status int  `json:"status" dc:"新的状态值"` // 新的状态值
}
