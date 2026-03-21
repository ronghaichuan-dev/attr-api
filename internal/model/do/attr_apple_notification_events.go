// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrAppleNotificationEvents is the golang structure of table attr_apple_notification_events for DAO operations like Where/Data.
type AttrAppleNotificationEvents struct {
	g.Meta                `orm:"table:attr_apple_notification_events, do:true"`
	Id                    any //
	NotificationUuid      any //
	SignedPayload         any //
	NotificationType      any //
	Subtype               any //
	OriginalTransactionId any //
	TransactionId         any //
	ResponseText          any //
	ReceivedAt            any //
	ProcessedAt           any //
}
