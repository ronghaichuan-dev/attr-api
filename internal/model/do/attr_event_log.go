// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrEventLog is the golang structure of table attr_event_log for DAO operations like Where/Data.
type AttrEventLog struct {
	g.Meta       `orm:"table:attr_event_log, do:true"`
	Id           any // id
	Country      any // 国家
	City         any // 城市
	Region       any // 州/省
	EventUuid    any // 事件唯一ID
	Appid        any // APP ID
	EventCode    any // 事件ID
	Rsid         any // 设备ID
	ResponseText any // 事件内容
	SentAt       any // 发送时间
	CreatedAt    any // 创建时间
}
