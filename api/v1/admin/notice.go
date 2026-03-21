package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeListReq 通知列表查询请求参数
type NoticeListReq struct {
	g.Meta        `path:"/notice/list" method:"get" summary:"获取通知列表" tags:"通知管理"`
	Uuid          string `json:"uuid" dc:"用户ID，用于查询该用户的通知数据"`
	Page          int    `json:"page" dc:"页码，默认值为1" d:"1"`
	PageSize      int    `json:"pageSize" dc:"每页大小，默认值为10" d:"10"`
	NoticeType    string `json:"noticeType" dc:"通知类型，可选参数"`
	RenewalStatus string `json:"renewalStatus" dc:"续费状态，可选参数"`
}

// NoticeListRes 通知列表查询响应参数
type NoticeListRes struct {
	Total int64       `json:"total" dc:"总记录数"`
	List  interface{} `json:"list" dc:"通知列表数据"`
}
