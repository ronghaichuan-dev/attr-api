// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrTransactionOrder is the golang structure of table attr_transaction_order for DAO operations like Where/Data.
type AttrTransactionOrder struct {
	g.Meta           `orm:"table:attr_transaction_order, do:true"`
	Id               any         // id
	AppId            any         // 应用ID
	TransactionId    any         // 交易ID
	SubTransactionId any         // 子交易ID
	Uuid             any         // 用户ID
	SkuId            any         // sku
	Amount           any         // 订阅金额
	SubscribeStatus  any         // 订阅状态
	CreatedAt        *gtime.Time // 创建时间
}
