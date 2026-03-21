package admin

import (
	"god-help-service/internal/model/entity"

	g "github.com/gogf/gf/v2/frame/g"
)

// PermissionTreeItem 包含子权限的权限树形结构项
type PermissionTreeItem struct {
	Id             int                   `json:"id"              dc:"权限ID"`
	PermissionName string                `json:"permission_name" dc:"权限名称"`
	PermissionCode string                `json:"permission_code" dc:"权限代码"`
	PermissionDesc string                `json:"permission_desc" dc:"权限描述"`
	Module         string                `json:"module"          dc:"所属模块"`
	Status         int                   `json:"status"          dc:"状态"`
	CreatedAt      string                `json:"created_at"      dc:"创建时间"`
	UpdatedAt      string                `json:"updated_at"      dc:"更新时间"`
	DeletedAt      string                `json:"deleted_at"      dc:"删除时间"`
	Route          string                `json:"route"           dc:"路由地址"`
	ParentId       int                   `json:"parent_id"       dc:"父级权限ID"`
	Level          int                   `json:"level"           dc:"权限级别"`
	Children       []*PermissionTreeItem `json:"children"        dc:"子权限列表"`
}

// PermissionListReq 权限列表请求参数结构体
type PermissionListReq struct {
	g.Meta         `path:"/list" method:"get" tags:"权限管理" summary:"获取权限列表"`
	Page           int    `json:"page"            form:"page"            d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size           int    `json:"size"            form:"size"            d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	PermissionName string `json:"permission_name" form:"permission_name" d:""  v:"max-length:100" dc:"权限名称搜索，可选"`           // 权限名称搜索，可选
	PermissionCode string `json:"permission_code" form:"permission_code" d:""  v:"max-length:100" dc:"权限代码搜索，可选"`           // 权限代码搜索，可选
	Module         string `json:"module"          form:"module"          d:""  v:"max-length:100" dc:"模块名称搜索，可选"`           // 模块名称搜索，可选
}

// PermissionListRes 权限列表响应参数结构体
type PermissionListRes struct {
	Total int64                       `json:"total" dc:"总记录数"` // 总记录数
	List  []*entity.PermissionsCustom `json:"list" dc:"权限列表"`  // 权限列表
}

// PermissionDetailReq 权限详情请求参数结构体
type PermissionDetailReq struct {
	g.Meta `path:"/detail" method:"get" tags:"权限管理" summary:"获取权限详情"`
	Id     uint `json:"id" form:"id" v:"required|min:1#权限ID不能为空|权限ID必须大于0" dc:"权限ID，必填"` // 权限ID，必填
}

// PermissionDetailRes 权限详情响应参数结构体
type PermissionDetailRes struct {
	Permission *entity.PermissionsCustom `json:"permission" dc:"权限信息"`
}

// PermissionCreateReq 权限创建请求参数结构体
type PermissionCreateReq struct {
	g.Meta         `path:"/create" method:"post" tags:"权限管理" summary:"创建权限"`
	PermissionName string `json:"permission_name" form:"permission_name" v:"required|max-length:100#权限名称不能为空|权限名称长度不能超过100个字符" dc:"权限名称，必填"` // 权限名称，必填
	PermissionCode string `json:"permission_code" form:"permission_code" v:"required|max-length:100#权限代码不能为空|权限代码长度不能超过100个字符" dc:"权限代码，必填"` // 权限代码，必填
	PermissionDesc string `json:"permission_desc" form:"permission_desc" d:""  v:"max-length:500#权限描述长度不能超过500个字符" dc:"权限描述，可选"`             // 权限描述，可选
	Module         string `json:"module"          form:"module"          d:""  v:"max-length:100#模块名称长度不能超过100个字符" dc:"所属模块，可选"`             // 所属模块，可选
	Route          string `json:"route"          form:"route"          d:""  v:"max-length:200#路由地址长度不能超过200个字符" dc:"路由地址，可选"`               // 路由地址，可选
	ParentId       uint   `json:"parent_id"      form:"parent_id"      d:"0"    v:"min:0" dc:"父级权限ID，0表示顶级权限"`                               // 父级权限ID，0表示顶级权限
	Status         int    `json:"status"          form:"status"          d:"1"    v:"in:1,2" dc:"状态，1：启用，2：禁用，默认为1"`                         // 状态，1：启用，2：禁用，默认为1
}

// PermissionCreateRes 权限创建响应参数结构体
type PermissionCreateRes struct {
	Id uint `json:"id" dc:"新创建的权限ID"` // 新创建的权限ID
}

// PermissionUpdateReq 权限更新请求参数结构体
type PermissionUpdateReq struct {
	g.Meta         `path:"/update" method:"put" tags:"权限管理" summary:"更新权限信息"`
	Id             uint   `json:"id"  form:"id"  v:"required|min:1#权限ID不能为空|权限ID必须大于0" dc:"权限ID，必填"`                             // 权限ID，必填
	PermissionName string `json:"permission_name" form:"permission_name" d:""  v:"max-length:100#权限名称长度不能超过100个字符" dc:"权限名称，可选"` // 权限名称，可选
	PermissionCode string `json:"permission_code" form:"permission_code" d:""  v:"max-length:100#权限代码长度不能超过100个字符" dc:"权限代码，可选"` // 权限代码，可选
	PermissionDesc string `json:"permission_desc" form:"permission_desc" d:""  v:"max-length:500#权限描述长度不能超过500个字符" dc:"权限描述，可选"` // 权限描述，可选
	Module         string `json:"module"          form:"module"          d:""  v:"max-length:100#模块名称长度不能超过100个字符" dc:"所属模块，可选"` // 所属模块，可选
	Route          string `json:"route"          form:"route"          d:""  v:"max-length:200#路由地址长度不能超过200个字符" dc:"路由地址，可选"`   // 路由地址，可选
	ParentId       uint   `json:"parent_id"      form:"parent_id"      d:"-1"    v:"min:0" dc:"父级权限ID，-1：不修改，0：顶级权限"`            // 父级权限ID，-1：不修改，0：顶级权限
	Status         int    `json:"status"          form:"status"          d:"-1"    v:"in:-1,1,2" dc:"状态，-1：不修改，1：启用，2：禁用"`       // 状态，-1：不修改，1：启用，2：禁用
}

// PermissionUpdateRes 权限更新响应参数结构体
type PermissionUpdateRes struct {
	Id uint `json:"id" dc:"更新的权限ID"` // 更新的权限ID
}

// PermissionDeleteReq 权限删除请求参数结构体
type PermissionDeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"权限管理" summary:"删除权限"`
	Id     uint `json:"id" form:"id" v:"required|min:1#权限ID不能为空|权限ID必须大于0" dc:"权限ID，必填"` // 权限ID，必填
}

// PermissionDeleteRes 权限删除响应参数结构体
type PermissionDeleteRes struct {
	Id uint `json:"id" dc:"删除的权限ID"` // 删除的权限ID
}

// PermissionTreeReq 权限树形结构请求参数结构体
type PermissionTreeReq struct {
	g.Meta `path:"/tree" method:"get" tags:"权限管理" summary:"获取权限树形结构"`
}

// PermissionTreeRes 权限树形结构响应参数结构体
type PermissionTreeRes struct {
	Tree []*PermissionTreeItem `json:"tree" dc:"权限树形结构"` // 权限树形结构
}

// PermissionCreateTableReq 权限表创建请求参数结构体
type PermissionCreateTableReq struct {
	g.Meta `path:"/createTable" method:"get" tags:"权限管理" summary:"创建权限表"`
}

// PermissionCreateTableRes 权限表创建响应参数结构体
type PermissionCreateTableRes struct {
	Success bool `json:"success" dc:"是否成功"` // 是否成功
}

// PermissionUserReq 根据用户ID获取拥有的权限请求参数结构体
type PermissionUserReq struct {
	g.Meta `path:"/user-permissions" method:"get" tags:"权限管理" summary:"根据用户ID获取拥有的权限"`
	UserId uint `json:"userId" form:"userId" v:"required|min:1#用户ID不能为空|用户ID必须大于0" dc:"用户ID，必填"` // 用户ID，必填
}

// PermissionUserRes 根据用户ID获取拥有的权限响应参数结构体
type PermissionUserRes struct {
	PermissionCodes []string `json:"permissionCodes" dc:"用户拥有的权限代码列表"` // 用户拥有的权限代码列表
}

// PermissionEnableReq 权限启用请求参数结构体
type PermissionEnableReq struct {
	g.Meta `path:"/enable" method:"put" tags:"权限管理" summary:"启用权限"`
	Id     uint `json:"id" form:"id" v:"required|min:1#权限ID不能为空|权限ID必须大于0" dc:"权限ID，必填"` // 权限ID，必填
}

// PermissionEnableRes 权限启用响应参数结构体
type PermissionEnableRes struct {
	Id uint `json:"id" dc:"启用的权限ID"` // 启用的权限ID
}

// PermissionDisableReq 权限禁用请求参数结构体
type PermissionDisableReq struct {
	g.Meta `path:"/disable" method:"put" tags:"权限管理" summary:"禁用权限"`
	Id     uint `json:"id" form:"id" v:"required|min:1#权限ID不能为空|权限ID必须大于0" dc:"权限ID，必填"` // 权限ID，必填
}

// PermissionDisableRes 权限禁用响应参数结构体
type PermissionDisableRes struct {
	Id uint `json:"id" dc:"禁用的权限ID"` // 禁用的权限ID
}

// PermissionUpdateStatusReq 权限状态更新请求参数结构体
type PermissionUpdateStatusReq struct {
	g.Meta `path:"/status" method:"put" tags:"权限管理" summary:"更新权限状态"`
	Id     uint `json:"id" form:"id" v:"required|min:1#权限ID不能为空|权限ID必须大于0" dc:"权限ID，必填"`             // 权限ID，必填
	Status int  `json:"status" form:"status" v:"required|in:1,2#状态不能为空|状态值必须为1或2" dc:"状态，1：启用，2：禁用"` // 状态，1：启用，2：禁用，必填
}

// PermissionUpdateStatusRes 权限状态更新响应参数结构体
type PermissionUpdateStatusRes struct {
	Id     uint `json:"id" dc:"更新状态的权限ID"` // 更新状态的权限ID
	Status int  `json:"status" dc:"新的状态值"` // 新的状态值
}
