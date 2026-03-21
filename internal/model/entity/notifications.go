// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notifications is the golang structure for table notifications.
type Notifications struct {
	Id                    uint64      `json:"id"                    orm:"id"                      description:""` //
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"              description:""` //
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"              description:""` //
	DeletedAt             *gtime.Time `json:"deletedAt"             orm:"deleted_at"              description:""` //
	SignedPayload         string      `json:"signedPayload"         orm:"signed_payload"          description:""` //
	NotificationType      string      `json:"notificationType"      orm:"notification_type"       description:""` //
	Subtype               string      `json:"subtype"               orm:"subtype"                 description:""` //
	TransactionId         string      `json:"transactionId"         orm:"transaction_id"          description:""` //
	OriginalTransactionId string      `json:"originalTransactionId" orm:"original_transaction_id" description:""` //
	ProductId             string      `json:"productId"             orm:"product_id"              description:""` //
	Status                string      `json:"status"                orm:"status"                  description:""` //
}
