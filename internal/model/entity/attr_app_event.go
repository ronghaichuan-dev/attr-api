// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrAppEvent is the golang structure for table attr_app_event.
type AttrAppEvent struct {
	Id        int64       `json:"id"        orm:"id"         description:"id"`    // id
	Appid     string      `json:"appid"     orm:"appid"      description:"appid"` // appid
	EventId   int         `json:"eventId"   orm:"event_id"   description:"事件ID"`  // 事件ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`  // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`  // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间（软删除）"`  // 删除时间（软删除）
}
