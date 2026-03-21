// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrTransactionOrder is the golang structure for table attr_transaction_order.
type AttrTransactionOrder struct {
	Id               int64       `json:"id"               orm:"id"                 description:"id"`    // id
	AppId            string      `json:"appId"            orm:"app_id"             description:"应用ID"`  // 应用ID
	TransactionId    string      `json:"transactionId"    orm:"transaction_id"     description:"交易ID"`  // 交易ID
	SubTransactionId string      `json:"subTransactionId" orm:"sub_transaction_id" description:"子交易ID"` // 子交易ID
	Uuid             string      `json:"uuid"             orm:"uuid"               description:"用户ID"`  // 用户ID
	SkuId            string      `json:"skuId"            orm:"sku_id"             description:"sku"`   // sku
	Amount           float64     `json:"amount"           orm:"amount"             description:"订阅金额"`  // 订阅金额
	SubscribeStatus  string      `json:"subscribeStatus"  orm:" subscribe_status"  description:"订阅状态"`  // 订阅状态
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`  // 创建时间
}
