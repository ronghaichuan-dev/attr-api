// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrAppSubscriptions is the golang structure for table attr_app_subscriptions.
type AttrAppSubscriptions struct {
	Id                    int64  `json:"id"                    orm:"id"                      description:"id"`                                                                              // id
	Environment           string `json:"environment"           orm:"environment"             description:"环境"`                                                                              // 环境
	OrignialTransactionId string `json:"orignialTransactionId" orm:"orignial_transaction_id" description:"原始交易ID"`                                                                          // 原始交易ID
	Rsid                  string `json:"rsid"                  orm:"rsid"                    description:"用户设备ID"`                                                                          // 用户设备ID
	Appid                 string `json:"appid"                 orm:"appid"                   description:"appID"`                                                                           // appID
	ProductId             string `json:"productId"             orm:"product_id"              description:"产品ID"`                                                                            // 产品ID
	Status                int    `json:"status"                orm:"status"                  description:"订阅状态 1-自动续订服务已激活 2-自动续订服务已过期自动  3-自动续订服务目前处于计费重试期 4-自动续订服务目前处于账单宽限期 5-自动续订订阅已取消"` // 订阅状态 1-自动续订服务已激活 2-自动续订服务已过期自动  3-自动续订服务目前处于计费重试期 4-自动续订服务目前处于账单宽限期 5-自动续订订阅已取消
	AutoRenewStatus       int    `json:"autoRenewStatus"       orm:"auto_renew_status"       description:"自动续费状态 1-启用 2-禁用"`                                                                // 自动续费状态 1-启用 2-禁用
	IsTrial               int    `json:"isTrial"               orm:"is_trial"                description:"是否试订 1-是 2-否"`                                                                    // 是否试订 1-是 2-否
	IsPaid                int    `json:"isPaid"                orm:"is_paid"                 description:"是否付费 1-是 2-否"`                                                                    // 是否付费 1-是 2-否
	LastEventAt           int64  `json:"lastEventAt"           orm:"last_event_at"           description:"上次事件时间"`                                                                          // 上次事件时间
	ExpiresReason         int    `json:"expiresReason"         orm:"expires_reason"          description:"过期原因 1-无 2-订阅在计费重试期结束后过期 3-订阅因价格上涨过期 4-订阅因产品不可售过期 5-用户自愿取消订阅导致过期"`                // 过期原因 1-无 2-订阅在计费重试期结束后过期 3-订阅因价格上涨过期 4-订阅因产品不可售过期 5-用户自愿取消订阅导致过期
	ExpiresAt             int64  `json:"expiresAt"             orm:"expires_at"              description:"过期时间"`                                                                            // 过期时间
	CreatedAt             int64  `json:"createdAt"             orm:"created_at"              description:"创建时间"`                                                                            // 创建时间
	UpdatedAt             int64  `json:"updatedAt"             orm:"updated_at"              description:"更新时间"`                                                                            // 更新时间
}
