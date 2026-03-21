// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppSubscriptions is the golang structure for table app_subscriptions.
type AppSubscriptions struct {
	Id                    int64       `json:"id"                    orm:"id"                      description:"id"`                                        // id
	Environment           string      `json:"environment"           orm:"environment"             description:"环境"`                                        // 环境
	OrignialTransactionId string      `json:"orignialTransactionId" orm:"orignial_transaction_id" description:"原始交易ID"`                                    // 原始交易ID
	Uuid                  string      `json:"uuid"                  orm:"uuid"                    description:"用户ID"`                                      // 用户ID
	Appid                 string      `json:"appid"                 orm:"appid"                   description:"appID"`                                     // appID
	ProductId             string      `json:"productId"             orm:"product_id"              description:"产品ID"`                                      // 产品ID
	Status                string      `json:"status"                orm:"status"                  description:"订阅状态 ACTIVE-订阅中 EXPIRED-已过期  CANCELED-已取消"` // 订阅状态 ACTIVE-订阅中 EXPIRED-已过期  CANCELED-已取消
	AutoRenewStatus       int         `json:"autoRenewStatus"       orm:"auto_renew_status"       description:"自动续费状态"`                                    // 自动续费状态
	IsTrial               int         `json:"isTrial"               orm:"is_trial"                description:"是否试订"`                                      // 是否试订
	IsPaid                int         `json:"isPaid"                orm:"is_paid"                 description:"是否付费"`                                      // 是否付费
	LastEventAt           *gtime.Time `json:"lastEventAt"           orm:"last_event_at"           description:"上次事件时间"`                                    // 上次事件时间
	ExpiresAt             *gtime.Time `json:"expiresAt"             orm:"expires_at"              description:"过期时间"`                                      // 过期时间
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"              description:"创建时间"`                                      // 创建时间
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"              description:"更新时间"`                                      // 更新时间
}
