package admin

import (
	"god-help-service/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppListReq 应用列表请求参数结构体
type AppListReq struct {
	g.Meta  `path:"/list" method:"get" tags:"APP管理" summary:"获取应用列表"`
	Page    int    `json:"page"     form:"page"     d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size    int    `json:"size"     form:"size"     d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	AppName string `json:"app_name" form:"app_name" d:""  v:"max-length:100" dc:"应用名称搜索，可选"`           // 应用名称搜索，可选
	AppId   string `json:"appid"    form:"appid"    d:""  v:"max-length:255" dc:"应用ID搜索，可选"`           // 应用ID搜索，可选
}

// AppListItem 应用列表项响应参数结构体
type AppListItem struct {
	AppId           string      `json:"appid"      dc:"应用ID，主键"`                   // 应用ID，主键
	AppName         string      `json:"app_name"  dc:"应用名称，最多100字符，不能为空"`          // 应用名称，最多100字符，不能为空
	Icon            string      `json:"icon"      dc:"应用图标，最多255字符"`               // 应用图标，最多255字符
	Creator         int         `json:"creator"   dc:"创建人ID"`                      // 创建人ID
	CreatorName     string      `json:"creator_name" dc:"创建人名称"`                   // 创建人名称
	Modifier        int         `json:"modifier"  dc:"修改人ID"`                      // 修改人ID
	ModifierName    string      `json:"modifier_name" dc:"修改人名称"`                  // 修改人名称
	SubscriptionFee float64     `json:"subscription_fee"  dc:"订阅费用，浮点类型，支持小数点后两位"` // 订阅费用，浮点类型，支持小数点后两位
	CreatedAt       *gtime.Time `json:"created_at" dc:"创建时间"`                      // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at" dc:"更新时间"`                      // 更新时间
	DeletedAt       *gtime.Time `json:"deleted_at" dc:"删除时间（软删除）"`                 // 删除时间（软删除）
}

// AppListRes 应用列表响应参数结构体
type AppListRes struct {
	Total int64          `json:"total" dc:"总记录数"` // 总记录数
	List  []*AppListItem `json:"list" dc:"应用列表"`  // 应用列表
}

// AppDetailItem 应用详情响应参数结构体
type AppDetailItem struct {
	AppId           string      `json:"appid"      dc:"应用ID，主键"`                   // 应用ID，主键
	AppName         string      `json:"app_name"  dc:"应用名称，最多100字符，不能为空"`          // 应用名称，最多100字符，不能为空
	Icon            string      `json:"icon"      dc:"应用图标，最多255字符"`               // 应用图标，最多255字符
	Creator         int         `json:"creator"   dc:"创建人ID"`                      // 创建人ID
	CreatorName     string      `json:"creator_name" dc:"创建人名称"`                   // 创建人名称
	Modifier        int         `json:"modifier"  dc:"修改人ID"`                      // 修改人ID
	ModifierName    string      `json:"modifier_name" dc:"修改人名称"`                  // 修改人名称
	SubscriptionFee float64     `json:"subscription_fee"  dc:"订阅费用，浮点类型，支持小数点后两位"` // 订阅费用，浮点类型，支持小数点后两位
	CreatedAt       *gtime.Time `json:"created_at" dc:"创建时间"`                      // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at" dc:"更新时间"`                      // 更新时间
	DeletedAt       *gtime.Time `json:"deleted_at" dc:"删除时间（软删除）"`                 // 删除时间（软删除）
}

// AppDetailReq 应用详情请求参数结构体
type AppDetailReq struct {
	g.Meta `path:"/detail" method:"get" tags:"APP管理" summary:"获取应用详情"`
	AppId  string `json:"appid" form:"appid" v:"required#应用ID不能为空" dc:"应用ID，必填"` // 应用ID，必填
}

// AppDetailRes 应用详情响应参数结构体
type AppDetailRes struct {
	App    *AppDetailItem `json:"app" dc:"应用信息"`
	Events []int          `json:"events" dc:"绑定的事件ID列表"`
}

// AppCreateReq 应用创建请求参数结构体
type AppCreateReq struct {
	g.Meta          `path:"/create" method:"post" tags:"APP管理" summary:"创建应用"`
	AppName         string  `json:"app_name" form:"app_name" v:"required|max-length:100#应用名称不能为空|应用名称长度不能超过100个字符" dc:"应用名称，必填"`                                // 应用名称，必填
	Icon            string  `json:"icon"     form:"icon"     d:""  v:"max-length:255#应用图标长度不能超过255个字符" dc:"应用图标，可选"`                                            // 应用图标，可选
	SubscriptionFee float64 `json:"subscription_fee" form:"subscription_fee" d:"0.00" v:"min:0|max:999999.99#订阅费用不能小于0|订阅费用不能大于999999.99" dc:"订阅费用，可选，默认为0.00"` // 订阅费用，可选，默认为0.00
	Events          []int   `json:"events"   form:"events"   d:""  dc:"事件ID列表，用于绑定应用事件，可选"`                                                                     // 事件ID列表，用于绑定应用事件，可选
}

// AppCreateRes 应用创建响应参数结构体
type AppCreateRes struct {
	AppId string `json:"appid" dc:"新创建的应用ID"` // 新创建的应用ID
}

// AppUpdateReq 应用更新请求参数结构体
type AppUpdateReq struct {
	g.Meta          `path:"/update" method:"put" tags:"APP管理" summary:"更新应用信息"`
	AppId           string  `json:"appid"    form:"appid"    v:"required#应用ID不能为空" dc:"应用ID，必填"`                                                     // 应用ID，必填
	AppName         string  `json:"app_name" form:"app_name" d:""  v:"max-length:100#应用名称长度不能超过100个字符" dc:"应用名称，可选"`                                 // 应用名称，可选
	Icon            string  `json:"icon"     form:"icon"     d:""  v:"max-length:255#应用图标长度不能超过255个字符" dc:"应用图标，可选"`                                 // 应用图标，可选
	SubscriptionFee float64 `json:"subscription_fee" form:"subscription_fee" d:""  v:"min:0|max:999999.99#订阅费用不能小于0|订阅费用不能大于999999.99" dc:"订阅费用，可选"` // 订阅费用，可选
	Events          []int   `json:"events"   form:"events"   d:""  dc:"事件ID列表，用于绑定应用事件，可选"`                                                          // 事件ID列表，用于绑定应用事件，可选
}

// AppUpdateRes 应用更新响应参数结构体
type AppUpdateRes struct {
	AppId string `json:"appid" dc:"更新的应用ID"` // 更新的应用ID
}

// AppDeleteReq 应用删除请求参数结构体
type AppDeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"APP管理" summary:"删除应用"`
	AppId  string `json:"appid" form:"appid" v:"required#应用ID不能为空" dc:"应用ID，必填"` // 应用ID，必填
}

// AppDeleteRes 应用删除响应参数结构体
type AppDeleteRes struct {
	AppId string `json:"appid" dc:"删除的应用ID"` // 删除的应用ID
}

// AppSubscriptionTrendReq 获取应用订阅趋势请求参数结构体
type AppSubscriptionTrendReq struct {
	g.Meta `path:"/subscription/trend" method:"get" tags:"APP管理" summary:"获取应用订阅趋势数据"`
	AppId  string `json:"appid" form:"appid" v:"required#应用ID不能为空" dc:"应用ID，必填"`                               // 应用ID，必填
	Days   int    `json:"days" form:"days" d:"30" v:"min:1|max:365#天数不能小于1|天数不能大于365" dc:"查询天数，默认30天，范围1-365"` // 查询天数，默认30天，范围1-365
}

// AppSubscriptionTrendRes 获取应用订阅趋势响应参数结构体
type AppSubscriptionTrendRes struct {
	Trend []*entity.AttrAppSubscriptions `json:"trend" dc:"订阅趋势数据"`
}
