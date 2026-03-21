// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NotificationEvents is the golang structure for table notification_events.
type NotificationEvents struct {
	Id             uint64      `json:"id"             orm:"id"              description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""` //
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""` //
	NotificationId uint64      `json:"notificationId" orm:"notification_id" description:""` //
	EventType      string      `json:"eventType"      orm:"event_type"      description:""` //
	EventData      string      `json:"eventData"      orm:"event_data"      description:""` //
	Status         string      `json:"status"         orm:"status"          description:""` //
}
