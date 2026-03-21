// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrEventLog is the golang structure for table attr_event_log.
type AttrEventLog struct {
	Id           int64  `json:"id"           orm:"id"            description:"id"`     // id
	Country      string `json:"country"      orm:"country"       description:"国家"`     // 国家
	City         string `json:"city"         orm:"city"          description:"城市"`     // 城市
	Region       string `json:"region"       orm:"region"        description:"州/省"`    // 州/省
	EventUuid    string `json:"eventUuid"    orm:"event_uuid"    description:"事件唯一ID"` // 事件唯一ID
	Appid        string `json:"appid"        orm:"appid"         description:"APP ID"` // APP ID
	EventCode    string `json:"eventCode"    orm:"event_code"    description:"事件ID"`   // 事件ID
	Rsid         string `json:"rsid"         orm:"rsid"          description:"设备ID"`   // 设备ID
	ResponseText string `json:"responseText" orm:"response_text" description:"事件内容"`   // 事件内容
	SentAt       int64  `json:"sentAt"       orm:"sent_at"       description:"发送时间"`   // 发送时间
	CreatedAt    int64  `json:"createdAt"    orm:"created_at"    description:"创建时间"`   // 创建时间
}
