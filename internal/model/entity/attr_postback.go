// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrPostback is the golang structure for table attr_postback.
type AttrPostback struct {
	Id                    int64  `json:"id"                    orm:"id"                      description:"id"`                               // id
	AppId                 string `json:"appId"                 orm:"app_id"                  description:"应用ID"`                             // 应用ID
	PostbackType          string `json:"postbackType"          orm:"postback_type"           description:"回传类型: install/event/reengagement"` // 回传类型: install/event/reengagement
	Network               string `json:"network"               orm:"network"                 description:"渠道"`                               // 渠道
	OriginalTransactionId string `json:"originalTransactionId" orm:"original_transaction_id" description:"原始交易ID"`                           // 原始交易ID
	EventName             string `json:"eventName"             orm:"event_name"              description:"事件名"`                              // 事件名
	PostbackUrl           string `json:"postbackUrl"           orm:"postback_url"            description:"回传URL"`                            // 回传URL
	ResponseCode          int    `json:"responseCode"          orm:"response_code"           description:"响应码"`                              // 响应码
	ResponseBody          string `json:"responseBody"          orm:"response_body"           description:"响应内容"`                             // 响应内容
	Status                int    `json:"status"                orm:"status"                  description:"状态: 1-成功 2-失败 3-重试中"`              // 状态: 1-成功 2-失败 3-重试中
	RetryCount            int    `json:"retryCount"            orm:"retry_count"             description:"重试次数"`                             // 重试次数
	CreatedAt             int64  `json:"createdAt"             orm:"created_at"              description:"创建时间"`                             // 创建时间
}
