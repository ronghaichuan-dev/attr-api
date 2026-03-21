// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrAppleNotificationEvent is the golang structure for table attr_apple_notification_event.
type AttrAppleNotificationEvent struct {
	Id                    int64  `json:"id"                    orm:"id"                      description:"id"`     // id
	Envirment             string `json:"envirment"             orm:"envirment"               description:"环境"`     // 环境
	Version               string `json:"version"               orm:"version"                 description:"应用版本"`   // 应用版本
	NotificationUuid      string `json:"notificationUuid"      orm:"notification_uuid"       description:"通知唯一ID"` // 通知唯一ID
	SignedPayload         string `json:"signedPayload"         orm:"signed_payload"          description:"加密数据"`   // 加密数据
	NotificationType      string `json:"notificationType"      orm:"notification_type"       description:"通知类型"`   // 通知类型
	Subtype               string `json:"subtype"               orm:"subtype"                 description:"通知子类型"`  // 通知子类型
	OriginalTransactionId string `json:"originalTransactionId" orm:"original_transaction_id" description:"原始交易ID"` // 原始交易ID
	TransactionId         string `json:"transactionId"         orm:"transaction_id"          description:"交易ID"`   // 交易ID
	ResponseText          string `json:"responseText"          orm:"response_text"           description:"解密后的数据"` // 解密后的数据
	ReceivedAt            int64  `json:"receivedAt"            orm:"received_at"             description:"通知接收时间"` // 通知接收时间
	ProcessedAt           int64  `json:"processedAt"           orm:"processed_at"            description:"解密处理时间"` // 解密处理时间
}
