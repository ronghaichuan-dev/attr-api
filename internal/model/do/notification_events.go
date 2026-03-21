// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NotificationEvents is the golang structure of table notification_events for DAO operations like Where/Data.
type NotificationEvents struct {
	g.Meta         `orm:"table:notification_events, do:true"`
	Id             any         //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
	NotificationId any         //
	EventType      any         //
	EventData      any         //
	Status         any         //
}
