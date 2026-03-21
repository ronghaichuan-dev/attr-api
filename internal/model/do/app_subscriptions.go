// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppSubscriptions is the golang structure of table app_subscriptions for DAO operations like Where/Data.
type AppSubscriptions struct {
	g.Meta                `orm:"table:app_subscriptions, do:true"`
	Id                    any         // id
	Environment           any         // 环境
	OrignialTransactionId any         // 原始交易ID
	Uuid                  any         // 用户ID
	Appid                 any         // appID
	ProductId             any         // 产品ID
	Status                any         // 订阅状态 ACTIVE-订阅中 EXPIRED-已过期  CANCELED-已取消
	AutoRenewStatus       any         // 自动续费状态
	IsTrial               any         // 是否试订
	IsPaid                any         // 是否付费
	LastEventAt           *gtime.Time // 上次事件时间
	ExpiresAt             *gtime.Time // 过期时间
	CreatedAt             *gtime.Time // 创建时间
	UpdatedAt             *gtime.Time // 更新时间
}
