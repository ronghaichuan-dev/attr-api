package admin

import (
	"god-help-service/internal/model/entity"

	g "github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User 用户模型结构体（保持向后兼容性）
type User struct {
	Id        uint        `json:"id"  gorm:"primaryKey;autoIncrement" dc:"用户ID，主键，自增"`     // 用户ID，主键，自增
	UserName  string      `json:"username" gorm:"size:100;not null" dc:"用户名，最多100字符，不能为空"` // 用户名，最多100字符，不能为空
	Password  string      `json:"password" gorm:"size:255;not null" dc:"密码"`               // 密码
	Email     string      `json:"email"    gorm:"size:255" dc:"邮箱"`                        // 邮箱
	RoleId    uint        `json:"role_id"    gorm:"not null;default:0" dc:"角色ID，外键关联角色表"`  // 角色ID，外键关联角色表
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`                                    // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`                                    // 更新时间
	DeletedAt *gtime.Time `json:"deleted_at" dc:"删除时间（软删除）"`                               // 删除时间（软删除）
}

// TableName 获取用户表名
func (u *User) TableName() string {
	return "users"
}

// UserListReq 用户列表请求参数结构体
type UserListReq struct {
	g.Meta   `path:"/list" method:"get" tags:"用户管理" summary:"获取用户列表"`
	Page     int    `json:"page"     form:"page"     d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size     int    `json:"size"     form:"size"     d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	UserName string `json:"username" form:"username" d:""  v:"max-length:100" dc:"用户名搜索，可选"`            // 用户名搜索，可选
	Email    string `json:"email"    form:"email"    d:""  v:"max-length:255" dc:"邮箱搜索，可选"`             // 邮箱搜索，可选
}

// UserListRes 用户列表响应参数结构体
type UserListRes struct {
	Total int64                 `json:"total" dc:"总记录数"` // 总记录数
	List  []*entity.SystemUsers `json:"list" dc:"用户列表"`  // 用户列表
}

// UserDetailReq 用户详情请求参数结构体
type UserDetailReq struct {
	g.Meta `path:"/detail" method:"get" tags:"用户管理" summary:"获取用户详情"`
	Id     uint `json:"id" form:"id" v:"required|min:1#用户ID不能为空|用户ID必须大于0" dc:"用户ID，必填"` // 用户ID，必填
}

// UserDetailRes 用户详情响应参数结构体
type UserDetailRes struct {
	Users *entity.SystemUsers `json:"user" dc:"用户信息"`
}

// UserCreateReq 用户创建请求参数结构体
type UserCreateReq struct {
	g.Meta   `path:"/create" method:"post" tags:"用户管理" summary:"创建用户"`
	UserName string `json:"username" form:"username" v:"required|max-length:100#用户名不能为空|用户名长度不能超过100个字符" dc:"用户名，必填"` // 用户名，必填
	Password string `json:"password" form:"password" v:"required|max-length:255#密码不能为空|密码长度不能超过255个字符" dc:"密码，必填"`    // 密码，必填
	Email    string `json:"email"    form:"email"    v:"max-length:255#邮箱长度不能超过255个字符" dc:"邮箱，可选"`                    // 邮箱，可选
	RoleId   uint64 `json:"role_id"   form:"role_id"   v:"min:0#角色ID不能小于0" dc:"角色ID，可选，默认0"`                          // 角色ID，可选，默认0
}

// UserCreateRes 用户创建响应参数结构体
type UserCreateRes struct {
	Id uint `json:"id" dc:"新创建的用户ID"` // 新创建的用户ID
}

// UserUpdateReq 用户更新请求参数结构体
type UserUpdateReq struct {
	g.Meta   `path:"/update" method:"put" tags:"用户管理" summary:"更新用户信息"`
	Id       uint   `json:"id"  form:"id"  v:"required|min:1#用户ID不能为空|用户ID必须大于0" dc:"用户ID，必填"`             // 用户ID，必填
	UserName string `json:"username" form:"username" d:""  v:"max-length:100#用户名长度不能超过100个字符" dc:"用户名，可选"` // 用户名，可选
	Password string `json:"password" form:"password" d:""  v:"max-length:255#密码长度不能超过255个字符" dc:"密码，可选"`   // 密码，可选
	Email    string `json:"email"    form:"email"    d:""  v:"max-length:255#邮箱长度不能超过255个字符" dc:"邮箱，可选"`   // 邮箱，可选
	RoleId   uint   `json:"role_id"   form:"role_id"   d:""  v:"min:0#角色ID不能小于0" dc:"角色ID，可选"`             // 角色ID，可选
}

// UserUpdateRes 用户更新响应参数结构体
type UserUpdateRes struct {
	Id uint `json:"id" dc:"更新的用户ID"` // 更新的用户ID
}

// UserDeleteReq 用户删除请求参数结构体
type UserDeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"用户管理" summary:"删除用户"`
	Id     uint `json:"id" form:"id" v:"required|min:1#用户ID不能为空|用户ID必须大于0" dc:"用户ID，必填"` // 用户ID，必填
}

// UserDeleteRes 用户删除响应参数结构体
type UserDeleteRes struct {
	Id uint `json:"id" dc:"删除的用户ID"` // 删除的用户ID
}

// UserCreateTableReq 用户表创建请求参数结构体
type UserCreateTableReq struct {
	g.Meta `path:"/createTable" method:"get" tags:"用户管理" summary:"创建用户表"`
}

// UserCreateTableRes 用户表创建响应参数结构体
type UserCreateTableRes struct {
	Success bool `json:"success" dc:"是否成功"` // 是否成功
}
