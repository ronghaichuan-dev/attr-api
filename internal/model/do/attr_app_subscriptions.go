// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrAppSubscriptions is the golang structure of table attr_app_subscriptions for DAO operations like Where/Data.
type AttrAppSubscriptions struct {
	g.Meta                `orm:"table:attr_app_subscriptions, do:true"`
	Id                    any // id
	Environment           any // 环境
	OrignialTransactionId any // 原始交易ID
	Rsid                  any // 用户设备ID
	Appid                 any // appID
	ProductId             any // 产品ID
	Status                any // 订阅状态 1-自动续订服务已激活 2-自动续订服务已过期自动  3-自动续订服务目前处于计费重试期 4-自动续订服务目前处于账单宽限期 5-自动续订订阅已取消
	AutoRenewStatus       any // 自动续费状态 1-启用 2-禁用
	IsTrial               any // 是否试订 1-是 2-否
	IsPaid                any // 是否付费 1-是 2-否
	LastEventAt           any // 上次事件时间
	ExpiresReason         any // 过期原因 1-无 2-订阅在计费重试期结束后过期 3-订阅因价格上涨过期 4-订阅因产品不可售过期 5-用户自愿取消订阅导致过期
	ExpiresAt             any // 过期时间
	OfferType             any // 优惠类型
	OfferId               any // 优惠ID
	RevocationDate        any // 撤销时间
	RevocationReason      any // 撤销原因
	CreatedAt             any // 创建时间
	UpdatedAt             any // 更新时间
}
