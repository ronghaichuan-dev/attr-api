package admin

import (
	"god-help-service/internal/model/entity"

	g "github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemSettings 系统设置模型结构体
type SystemSettings struct {
	Id        int         `json:"id"      gorm:"primaryKey" dc:"主键ID"`                         // 主键ID
	Key       string      `json:"key"     gorm:"size:255;unique;not null" dc:"设置键，唯一，最多255字符"` // 设置键，唯一
	Value     string      `json:"value"   gorm:"type:text" dc:"设置值"`                           // 设置值
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`                                        // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`                                        // 更新时间
	DeletedAt *gtime.Time `json:"deleted_at" dc:"删除时间（软删除）"`                                   // 删除时间（软删除）
}

// TableName 获取系统设置表名
func (s *SystemSettings) TableName() string {
	return "system_settings"
}

// SystemSettingListReq 系统设置列表请求参数结构体
type SystemSettingListReq struct {
	g.Meta `path:"/list" method:"get" tags:"系统设置" summary:"获取系统设置列表"`
	Page   int    `json:"page"     form:"page"     d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size   int    `json:"size"     form:"size"     d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	Key    string `json:"key"      form:"key"      d:""  v:"max-length:255" dc:"设置键搜索，可选"`            // 设置键搜索，可选
}

// SystemSettingListRes 系统设置列表响应参数结构体
type SystemSettingListRes struct {
	Total int64                    `json:"total" dc:"总记录数"`  // 总记录数
	List  []*entity.SystemSettings `json:"list" dc:"系统设置列表"` // 系统设置列表
}

// SystemSettingDetailReq 系统设置详情请求参数结构体
type SystemSettingDetailReq struct {
	g.Meta `path:"/detail" method:"get" tags:"系统设置" summary:"获取系统设置详情"`
	Id     int `json:"id"     form:"id"     v:"required#ID不能为空" dc:"设置ID，必填"` // 设置ID，必填
}

// SystemSettingDetailRes 系统设置详情响应参数结构体
type SystemSettingDetailRes struct {
	Setting *entity.SystemSettings `json:"setting" dc:"系统设置信息"`
}

// SystemSettingCreateReq 系统设置创建请求参数结构体
type SystemSettingCreateReq struct {
	g.Meta `path:"/create" method:"post" tags:"系统设置" summary:"创建系统设置"`
	Key    string `json:"key"    form:"key"    v:"required|max-length:255#设置键不能为空|设置键长度不能超过255个字符" dc:"设置键，必填"` // 设置键，必填
	Value  string `json:"value"  form:"value"  d:""  dc:"设置值，可选"`                                               // 设置值，可选
}

// SystemSettingCreateRes 系统设置创建响应参数结构体
type SystemSettingCreateRes struct {
	Id int `json:"id" dc:"新创建的系统设置ID"` // 新创建的系统设置ID
}

// SystemSettingUpdateReq 系统设置更新请求参数结构体
type SystemSettingUpdateReq struct {
	g.Meta `path:"/update" method:"put" tags:"系统设置" summary:"更新系统设置"`
	Id     int    `json:"id"     form:"id"     v:"required#ID不能为空" dc:"设置ID，必填"`                     // 设置ID，必填
	Key    string `json:"key"    form:"key"    d:""  v:"max-length:255#设置键长度不能超过255个字符" dc:"设置键，可选"` // 设置键，可选
	Value  string `json:"value"  form:"value"  d:""  dc:"设置值，可选"`                                    // 设置值，可选
}

// SystemSettingUpdateRes 系统设置更新响应参数结构体
type SystemSettingUpdateRes struct {
	Id int `json:"id" dc:"更新的系统设置ID"` // 更新的系统设置ID
}

// SystemSettingDeleteReq 系统设置删除请求参数结构体
type SystemSettingDeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"系统设置" summary:"删除系统设置"`
	Id     int `json:"id"     form:"id"     v:"required#ID不能为空" dc:"设置ID，必填"` // 设置ID，必填
}

// SystemSettingDeleteRes 系统设置删除响应参数结构体
type SystemSettingDeleteRes struct {
	Id int `json:"id" dc:"删除的系统设置ID"` // 删除的系统设置ID
}

// SystemSettingGetValueReq 根据key获取系统设置值请求参数结构体
type SystemSettingGetValueReq struct {
	g.Meta `path:"/get-value" method:"get" tags:"系统设置" summary:"根据key获取系统设置值"`
	Key    string `json:"key"    form:"key"    v:"required#设置键不能为空" dc:"设置键，必填"` // 设置键，必填
}

// SystemSettingGetValueRes 根据key获取系统设置值响应参数结构体
type SystemSettingGetValueRes struct {
	Value string `json:"value" dc:"系统设置值"` // 系统设置值
}
