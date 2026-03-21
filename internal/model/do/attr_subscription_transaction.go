// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrSubscriptionTransaction is the golang structure of table attr_subscription_transaction for DAO operations like Where/Data.
type AttrSubscriptionTransaction struct {
	g.Meta                `orm:"table:attr_subscription_transaction, do:true"`
	Id                    any // id
	TransactionType       any // 交易类型 RENEW / REFUND / TRIAL
	Envirment             any // 环境
	AppVersion            any // 应用版本
	Appid                 any // 应用ID
	OriginalTransactionId any // 原始交易ID
	TransactionId         any // 子交易ID
	InAppOwnership        any // 是否为用户购买 PURCHASED-购买 FAMILY_SHARED-家庭分享
	Uuid                  any // 用户ID
	ProductId             any // sku
	Price                 any // 订阅金额
	Currency              any // 币种
	SubscribeStatus       any // 订阅状态
	PurchaseAt            any // 购买时间
	CreatedAt             any // 创建时间
}
