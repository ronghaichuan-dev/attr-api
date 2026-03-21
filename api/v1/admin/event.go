package admin

import (
	"god-help-service/internal/model/entity"

	g "github.com/gogf/gf/v2/frame/g"
)

// EventDropdownReq 事件下拉选项请求参数结构体
type EventDropdownReq struct {
	g.Meta `path:"/dropdown" method:"get" tags:"事件管理" summary:"获取事件下拉选项列表"`
}

// EventDropdownItem 事件下拉选项项
type EventDropdownItem struct {
	Id        int    `json:"id"         dc:"事件ID"` // 事件ID
	EventName string `json:"event_name" dc:"事件名称"` // 事件名称
}

// EventDropdownRes 事件下拉选项响应参数结构体
type EventDropdownRes struct {
	List []*EventDropdownItem `json:"list" dc:"事件下拉选项列表"` // 事件下拉选项列表
}

// EventListReq 事件列表请求参数结构体
type EventListReq struct {
	g.Meta    `path:"/list" method:"get" tags:"事件管理" summary:"获取事件列表"`
	Page      int    `json:"page"         form:"page"         d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size      int    `json:"size"         form:"size"         d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	EventName string `json:"event_name"   form:"event_name"   d:""  v:"max-length:100" dc:"事件名称搜索，可选"`           // 事件名称搜索，可选
	Status    int    `json:"status"       form:"status"       d:"0"  v:"min:0" dc:"状态搜索，0表示不搜索，可选"`              // 状态搜索，0表示不搜索，可选
}

// EventListItem 事件列表项
type EventListItem struct {
	Id        int    `json:"id"         dc:"事件ID"` // 事件ID
	EventName string `json:"event_name" dc:"事件名称"` // 事件名称
	Status    int    `json:"status"     dc:"状态"`   // 状态
	CreatedAt string `json:"created_at" dc:"创建时间"` // 创建时间
	UpdatedAt string `json:"updated_at" dc:"更新时间"` // 更新时间
}

// EventListRes 事件列表响应参数结构体
type EventListRes struct {
	Total int64            `json:"total" dc:"总记录数"` // 总记录数
	List  []*EventListItem `json:"list" dc:"事件列表"`  // 事件列表
}

// EventDetailReq 事件详情请求参数结构体
type EventDetailReq struct {
	g.Meta `path:"/detail" method:"get" tags:"事件管理" summary:"获取事件详情"`
	Id     int `json:"id" form:"id" v:"required|min:1#事件ID不能为空|事件ID必须大于0" dc:"事件ID，必填"` // 事件ID，必填
}

// EventDetailItem 事件详情项
type EventDetailItem struct {
	Id        int    `json:"id"         dc:"事件ID"` // 事件ID
	EventName string `json:"event_name" dc:"事件名称"` // 事件名称
	Status    int    `json:"status"     dc:"状态"`   // 状态
	CreatedAt string `json:"created_at" dc:"创建时间"` // 创建时间
	UpdatedAt string `json:"updated_at" dc:"更新时间"` // 更新时间
}

// EventDetailRes 事件详情响应参数结构体
type EventDetailRes struct {
	Event *EventDetailItem `json:"event" dc:"事件信息"`
}

// EventCreateReq 事件创建请求参数结构体
type EventCreateReq struct {
	g.Meta    `path:"/create" method:"post" tags:"事件管理" summary:"创建事件"`
	EventName string `json:"event_name"   form:"event_name"   v:"required|max-length:100#事件名称不能为空|事件名称长度不能超过100个字符" dc:"事件名称，必填"` // 事件名称，必填
	Status    int    `json:"status"       form:"status"       d:"1"  v:"min:1|max:2" dc:"状态，1启用，2禁用，默认1"`                         // 状态，1启用，2禁用，默认1
}

// EventCreateRes 事件创建响应参数结构体
type EventCreateRes struct {
	Id int `json:"id" dc:"事件ID"` // 事件ID
}

// EventUpdateReq 事件更新请求参数结构体
type EventUpdateReq struct {
	g.Meta    `path:"/update" method:"put" tags:"事件管理" summary:"更新事件"`
	Id        int    `json:"id"          form:"id"          v:"required|min:1#事件ID不能为空|事件ID必须大于0" dc:"事件ID，必填"` // 事件ID，必填
	EventName string `json:"event_name"   form:"event_name"   v:"max-length:100#事件名称长度不能超过100个字符" dc:"事件名称，可选"` // 事件名称，可选
	Status    int    `json:"status"       form:"status"      v:"min:1|max:2#状态必须是1或2" dc:"状态，1启用，2禁用，可选"`       // 状态，1启用，2禁用，可选
}

// EventUpdateRes 事件更新响应参数结构体
type EventUpdateRes struct {
	Id int `json:"id" dc:"事件ID"` // 事件ID
}

// EventDeleteReq 事件删除请求参数结构体
type EventDeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"事件管理" summary:"删除事件"`
	Id     int `json:"id" form:"id" v:"required|min:1#事件ID不能为空|事件ID必须大于0" dc:"事件ID，必填"` // 事件ID，必填
}

// EventDeleteRes 事件删除响应参数结构体
type EventDeleteRes struct {
	Id int `json:"id" dc:"事件ID"` // 事件ID
}

// EventLogListReq 事件日志列表请求参数结构体
type EventLogListReq struct {
	g.Meta    `path:"/event/log/list" method:"get" tags:"事件日志管理" summary:"获取事件日志列表"`
	Page      int    `json:"page"        form:"page"        d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size      int    `json:"size"        form:"size"        d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	Appid     string `json:"appid"       form:"appid"       d:""  v:"max-length:100" dc:"APP ID搜索，可选"`         // APP ID搜索，可选
	UserId    string `json:"user_id"     form:"user_id"     d:""  v:"max-length:100" dc:"用户ID搜索，可选"`           // 用户ID搜索，可选
	EventId   int    `json:"event_id"    form:"event_id"    d:"-1"  v:"min:-1" dc:"事件ID搜索，-1表示不搜索，可选"`         // 事件ID搜索，-1表示不搜索，可选
	StartTime string `json:"start_time"  form:"start_time"  d:""  v:"date-format:Y-m-d H:i:s" dc:"开始时间搜索，可选"`  // 开始时间搜索，可选
	EndTime   string `json:"end_time"    form:"end_time"    d:""  v:"date-format:Y-m-d H:i:s" dc:"结束时间搜索，可选"`  // 结束时间搜索，可选
}

// EventLogListItem 事件日志列表项
type EventLogListItem struct {
	Id           int64  `json:"id"            dc:"事件日志ID"` // 事件日志ID
	Appid        string `json:"appid"         dc:"应用ID"`   // 应用ID
	EventCode    string `json:"event_code"      dc:"事件ID"` // 事件ID
	UserId       string `json:"user_id"       dc:"用户ID"`   // 用户ID
	ResponseText string `json:"response_text" dc:"响应内容"`   // 响应内容
	CreatedAt    string `json:"created_at"    dc:"创建时间"`   // 创建时间
}

// EventLogListRes 事件日志列表响应参数结构体
type EventLogListRes struct {
	Total int64               `json:"total" dc:"总记录数"`  // 总记录数
	List  []*EventLogListItem `json:"list" dc:"事件日志列表"` // 事件日志列表
}

// EventLogDetailReq 事件日志详情请求参数结构体
type EventLogDetailReq struct {
	g.Meta `path:"/event/log/detail" method:"get" tags:"事件日志管理" summary:"获取事件日志详情"`
	Id     int64 `json:"id" form:"id" v:"required|min:1#事件日志ID不能为空|事件日志ID必须大于0" dc:"事件日志ID，必填"` // 事件日志ID，必填
}

// EventLogDetailRes 事件日志详情响应参数结构体
type EventLogDetailRes struct {
	Event_log *entity.AttrEventLog `json:"event_log" dc:"事件日志信息"`
}

// EventLogCreateTableReq 事件日志表初始化请求参数结构体
type EventLogCreateTableReq struct {
	g.Meta `path:"/event/log/create-table" method:"get" tags:"事件日志管理" summary:"初始化事件日志表"`
}

// EventLogCreateTableRes 事件日志表初始化响应参数结构体
type EventLogCreateTableRes struct {
	Success bool `json:"success" dc:"是否成功"` // 是否成功
}
