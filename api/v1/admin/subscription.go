package admin

import (
	"god-help-service/internal/model/entity"

	g "github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Subscription 订阅事件模型结构体
type Subscription struct {
	ID              uint        `json:"id" gorm:"primaryKey;autoIncrement" dc:"订阅事件ID，主键，自增"`               // 订阅事件ID，主键，自增
	AppId           string      `json:"app_id" gorm:"size:255;not null" dc:"应用ID，最多255字符，不能为空"`             // 应用ID，最多255字符，不能为空
	EventType       string      `json:"event_type" gorm:"size:100;not null" dc:"事件类型，最多100字符，不能为空"`         // 事件类型，最多100字符，不能为空
	Country         string      `json:"country" gorm:"size:20;not null" dc:"国家，最多20字符，不能为空"`                // 国家，最多20字符，不能为空
	UserId          string      `json:"user_id" gorm:"size:255" dc:"用户ID，最多255字符"`                          // 用户ID，最多255字符
	DeviceId        string      `json:"device_id" gorm:"size:255" dc:"设备ID，最多255字符"`                        // 设备ID，最多255字符
	SubscriptionFee float64     `json:"subscription_fee" gorm:"type:decimal(10,2)" dc:"订阅费用，浮点类型，支持小数点后两位"` // 订阅费用，浮点类型，支持小数点后两位
	SubscribedAt    *gtime.Time `json:"subscribed_at" dc:"订阅时间"`                                            // 订阅时间
	CreatedAt       *gtime.Time `json:"created_at" dc:"创建时间"`                                               // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at" dc:"更新时间"`                                               // 更新时间
	DeletedAt       *gtime.Time `json:"deleted_at" dc:"删除时间（软删除）"`                                          // 删除时间（软删除）
}

// TableName 获取订阅事件表名
func (s *Subscription) TableName() string {
	return "subscriptions"
}

// SubscriptionListReq 订阅事件列表请求参数结构体
type SubscriptionListReq struct {
	g.Meta  `path:"/list" method:"get" tags:"订阅事件" summary:"获取订阅事件列表"`
	Page    int    `json:"page"     form:"page"     d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size    int    `json:"size"     form:"size"     d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	AppId   string `json:"app_id"    form:"app_id"    d:""  v:"max-length:255" dc:"应用ID搜索，可选"`         // 应用ID搜索，可选
	EventId string `json:"event_id" form:"event_id" d:""  v:"max-length:100" dc:"事件ID搜索，可选"`           // 事件ID搜索，可选
	Country string `json:"country"   form:"country"   d:""  v:"max-length:20" dc:"国家搜索，可选"`            // 国家搜索，可选
	UserId  string `json:"user_id"   form:"user_id"   d:""  v:"max-length:255" dc:"用户ID搜索，可选"`         // 用户ID搜索，可选
}

// SubscriptionListRes 订阅事件列表响应参数结构体
type SubscriptionListRes struct {
	Total int64                   `json:"total" dc:"总记录数"`  // 总记录数
	List  []*entity.Subscriptions `json:"list" dc:"订阅事件列表"` // 订阅事件列表
}

// SubscriptionDetailReq 订阅事件详情请求参数结构体
type SubscriptionDetailReq struct {
	g.Meta `path:"/detail" method:"get" tags:"订阅事件" summary:"获取订阅事件详情"`
	ID     uint `json:"id" form:"id" v:"required|min:1#订阅事件ID不能为空|订阅事件ID必须大于0" dc:"订阅事件ID，必填"` // 订阅事件ID，必填
}

// SubscriptionDetailRes 订阅事件详情响应参数结构体
type SubscriptionDetailRes struct {
	Subscription *entity.Subscriptions `json:"subscription" dc:"订阅事件信息"`
}

// SubscriptionCreateReq 订阅事件创建请求参数结构体
type SubscriptionCreateReq struct {
	g.Meta          `path:"/create" method:"post" tags:"订阅事件" summary:"创建订阅事件"`
	AppId           string  `json:"app_id"    v:"required|max-length:255#应用ID不能为空|应用ID长度不能超过255个字符" dc:"应用ID，必填"`                       // 应用ID，必填
	EventId         string  `json:"event_id" v:"required|max-length:100#事件ID不能为空|事件ID长度不能超过100个字符" dc:"事件ID，必填"`                        // 事件ID，必填
	Country         string  `json:"country"   v:"required|max-length:20#国家不能为空|国家长度不能超过20个字符" dc:"国家，必填"`                               // 国家，必填
	UserId          string  `json:"user_id"   d:""  v:"max-length:255" dc:"用户ID，可选"`                                                    // 用户ID，可选
	DeviceId        string  `json:"device_id" d:""  v:"max-length:255" dc:"设备ID，可选"`                                                    // 设备ID，可选
	SubscriptionFee float64 `json:"subscription_fee" d:"0.00" v:"min:0|max:999999.99#订阅费用不能小于0|订阅费用不能大于999999.99" dc:"订阅费用，可选，默认为0.00"` // 订阅费用，可选，默认为0.00
	SubscribedAt    string  `json:"subscribed_at" d:""  v:"date-format:Y-m-d H:i:s" dc:"订阅时间，格式：Y-m-d H:i:s，可选"`                        // 订阅时间，可选
}

// SubscriptionCreateRes 订阅事件创建响应参数结构体
type SubscriptionCreateRes struct {
	ID uint `json:"id" dc:"新创建的订阅事件ID"` // 新创建的订阅事件ID
}

// SubscriptionUpdateReq 订阅事件更新请求参数结构体
type SubscriptionUpdateReq struct {
	g.Meta          `path:"/update" method:"put" tags:"订阅事件" summary:"更新订阅事件信息"`
	ID              uint    `json:"id"          form:"id"          v:"required|min:1#订阅事件ID不能为空|订阅事件ID必须大于0" dc:"订阅事件ID，必填"`                         // 订阅事件ID，必填
	AppId           string  `json:"app_id"      form:"app_id"      d:""  v:"max-length:255#应用ID长度不能超过255个字符" dc:"应用ID，可选"`                           // 应用ID，可选
	EventId         string  `json:"event_id"  form:"event_id"  d:""  v:"max-length:100#事件ID长度不能超过100个字符" dc:"事件ID，可选"`                               // 事件ID，可选
	Country         string  `json:"country"     form:"country"     d:""  v:"max-length:20#国家长度不能超过20个字符" dc:"国家，可选"`                                 // 国家，可选
	UserId          string  `json:"user_id"     form:"user_id"     d:""  v:"max-length:255" dc:"用户ID，可选"`                                            // 用户ID，可选
	DeviceId        string  `json:"device_id"   form:"device_id"   d:""  v:"max-length:255" dc:"设备ID，可选"`                                            // 设备ID，可选
	SubscriptionFee float64 `json:"subscription_fee" form:"subscription_fee" d:""  v:"min:0|max:999999.99#订阅费用不能小于0|订阅费用不能大于999999.99" dc:"订阅费用，可选"` // 订阅费用，可选
	SubscribedAt    string  `json:"subscribed_at" form:"subscribed_at" d:""  v:"date-format:Y-m-d H:i:s" dc:"订阅时间，格式：Y-m-d H:i:s，可选"`                // 订阅时间，可选
}

// SubscriptionUpdateRes 订阅事件更新响应参数结构体
type SubscriptionUpdateRes struct {
	ID uint `json:"id" dc:"更新的订阅事件ID"` // 更新的订阅事件ID
}

// SubscriptionDeleteReq 订阅事件删除请求参数结构体
type SubscriptionDeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"订阅事件" summary:"删除订阅事件"`
	ID     uint `json:"id" form:"id" v:"required|min:1#订阅事件ID不能为空|订阅事件ID必须大于0" dc:"订阅事件ID，必填"` // 订阅事件ID，必填
}

// SubscriptionDeleteRes 订阅事件删除响应参数结构体
type SubscriptionDeleteRes struct {
	ID uint `json:"id" dc:"删除的订阅事件ID"` // 删除的订阅事件ID
}
