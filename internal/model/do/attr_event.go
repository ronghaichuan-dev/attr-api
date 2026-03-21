// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrEvent is the golang structure of table attr_event for DAO operations like Where/Data.
type AttrEvent struct {
	g.Meta    `orm:"table:attr_event, do:true"`
	Id        any         // id
	EventName any         // 事件名称
	EventCode any         // 事件代码
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Status    any         // 状态 1-启用 2-禁用
}
