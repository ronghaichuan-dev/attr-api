// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Subscriptions is the golang structure for table subscriptions.
type Subscriptions struct {
	Id              int         `json:"id"              orm:"id"               description:"主键ID"`       // 主键ID
	AppId           string      `json:"appId"           orm:"app_id"           description:"APP ID"`     // APP ID
	SubscriptionFee float64     `json:"subscriptionFee" orm:"subscription_fee" description:"订阅费用"`       // 订阅费用
	EventId         int         `json:"eventId"         orm:"event_id"         description:"事件ID"`       // 事件ID
	Country         string      `json:"country"         orm:"country"          description:"国家"`         // 国家
	UserId          string      `json:"userId"          orm:"user_id"          description:"用户ID"`       // 用户ID
	DeviceId        string      `json:"deviceId"        orm:"device_id"        description:"设备ID"`       // 设备ID
	SubscribedAt    *gtime.Time `json:"subscribedAt"    orm:"subscribed_at"    description:"订阅时间（当地时间）"` // 订阅时间（当地时间）
	SubscribedStamp *gtime.Time `json:"subscribedStamp" orm:"subscribed_stamp" description:"订阅时间戳"`      // 订阅时间戳
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`       // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`       // 更新时间
}
