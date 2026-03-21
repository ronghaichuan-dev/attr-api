// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Subscriptions is the golang structure of table subscriptions for DAO operations like Where/Data.
type Subscriptions struct {
	g.Meta          `orm:"table:subscriptions, do:true"`
	Id              any         // 主键ID
	AppId           any         // APP ID
	SubscriptionFee any         // 订阅费用
	EventId         any         // 事件ID
	Country         any         // 国家
	UserId          any         // 用户ID
	DeviceId        any         // 设备ID
	SubscribedAt    *gtime.Time // 订阅时间（当地时间）
	SubscribedStamp *gtime.Time // 订阅时间戳
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
