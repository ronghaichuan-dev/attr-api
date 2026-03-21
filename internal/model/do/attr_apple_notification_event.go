// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrAppleNotificationEvent is the golang structure of table attr_apple_notification_event for DAO operations like Where/Data.
type AttrAppleNotificationEvent struct {
	g.Meta                `orm:"table:attr_apple_notification_event, do:true"`
	Id                    any // id
	Envirment             any // 环境
	Version               any // 应用版本
	NotificationUuid      any // 通知唯一ID
	SignedPayload         any // 加密数据
	NotificationType      any // 通知类型
	Subtype               any // 通知子类型
	OriginalTransactionId any // 原始交易ID
	TransactionId         any // 交易ID
	ResponseText          any // 解密后的数据
	ReceivedAt            any // 通知接收时间
	ProcessedAt           any // 解密处理时间
}
