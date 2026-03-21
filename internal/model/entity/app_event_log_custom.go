// =================================================================================
// 自定义结构体，用于事件日志管理，使用驼峰命名规范
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppEventLogCustom 是`app_event_log`表的自定义结构体，使用驼峰命名规范
type AppEventLogCustom struct {
	Id           int64       `json:"id"           orm:"id"            description:"id"`     // id
	Appid        string      `json:"appid"        orm:"appid"         description:"APP ID"` // APP ID
	EventCode    string      `json:"eventCode"      orm:"event_id"      description:"事件ID"` // 事件ID
	UserId       string      `json:"userId"       orm:"user_id"       description:"用户ID"`   // 用户ID
	ResponseText string      `json:"responseText" orm:"response_text" description:"事件内容"`   // 事件内容
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`   // 创建时间
}
