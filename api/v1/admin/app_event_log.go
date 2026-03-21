package admin

import (
	"god-help-service/internal/model/entity"

	g "github.com/gogf/gf/v2/frame/g"
)

// AppEventLogListReq 事件日志列表请求参数结构体
type AppEventLogListReq struct {
	g.Meta `path:"/app-event-log/list" method:"get" tags:"事件日志管理" summary:"获取事件日志列表"`
	Page   int    `json:"page"      form:"page"      d:"1"    v:"min:1" dc:"页码，默认1，最小1"`                // 页码，默认1，最小1
	Size   int    `json:"size"      form:"size"      d:"10"   v:"min:0|max:100" dc:"每页数量，默认10，范围0-100"` // 每页数量，默认10，范围0-100
	Appid  string `json:"appid"     form:"appid"     d:""    v:"max-length:100" dc:"应用ID搜索，可选"`         // 应用ID搜索，可选
	UserId string `json:"user_id"   form:"user_id"   d:""    v:"max-length:100" dc:"用户ID搜索，可选"`         // 用户ID搜索，可选
}

// AppEventLogListRes 事件日志列表响应参数结构体
type AppEventLogListRes struct {
	Total int64                       `json:"total" dc:"总记录数"`  // 总记录数
	List  []*entity.AppEventLogCustom `json:"list" dc:"事件日志列表"` // 事件日志列表
}

// AppEventLogDetailReq 事件日志详情请求参数结构体
type AppEventLogDetailReq struct {
	g.Meta `path:"/app-event-log/detail" method:"get" tags:"事件日志管理" summary:"获取事件日志详情"`
	Id     int64 `json:"id" form:"id" v:"required|min:1#事件日志ID不能为空|事件日志ID必须大于0" dc:"事件日志ID，必填"` // 事件日志ID，必填
}

// AppEventLogDetailRes 事件日志详情响应参数结构体
type AppEventLogDetailRes struct {
	EventLog *entity.AppEventLogCustom `json:"event_log" dc:"事件日志信息"`
}
