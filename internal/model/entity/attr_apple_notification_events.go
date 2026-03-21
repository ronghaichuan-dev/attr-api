// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrAppleNotificationEvents is the golang structure for table attr_apple_notification_events.
type AttrAppleNotificationEvents struct {
	Id                    int64  `json:"id"                    orm:"id"                      description:""` //
	NotificationUuid      string `json:"notificationUuid"      orm:"notification_uuid"       description:""` //
	SignedPayload         string `json:"signedPayload"         orm:"signed_payload"          description:""` //
	NotificationType      string `json:"notificationType"      orm:"notification_type"       description:""` //
	Subtype               string `json:"subtype"               orm:"subtype"                 description:""` //
	OriginalTransactionId string `json:"originalTransactionId" orm:"original_transaction_id" description:""` //
	TransactionId         string `json:"transactionId"         orm:"transaction_id"          description:""` //
	ResponseText          string `json:"responseText"          orm:"response_text"           description:""` //
	ReceivedAt            int64  `json:"receivedAt"            orm:"received_at"             description:""` //
	ProcessedAt           int64  `json:"processedAt"           orm:"processed_at"            description:""` //
}
