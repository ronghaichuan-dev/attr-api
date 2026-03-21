// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notifications is the golang structure of table notifications for DAO operations like Where/Data.
type Notifications struct {
	g.Meta                `orm:"table:notifications, do:true"`
	Id                    any         //
	CreatedAt             *gtime.Time //
	UpdatedAt             *gtime.Time //
	DeletedAt             *gtime.Time //
	SignedPayload         any         //
	NotificationType      any         //
	Subtype               any         //
	TransactionId         any         //
	OriginalTransactionId any         //
	ProductId             any         //
	Status                any         //
}
