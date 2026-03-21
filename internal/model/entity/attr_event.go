// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrEvent is the golang structure for table attr_event.
type AttrEvent struct {
	Id        int64       `json:"id"        orm:"id"         description:"id"`           // id
	EventName string      `json:"eventName" orm:"event_name" description:"事件名称"`         // 事件名称
	EventCode string      `json:"eventCode" orm:"event_code" description:"事件代码"`         // 事件代码
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
	Status    int         `json:"status"    orm:"status"     description:"状态 1-启用 2-禁用"` // 状态 1-启用 2-禁用
}
