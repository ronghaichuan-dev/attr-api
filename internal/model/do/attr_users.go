// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrUsers is the golang structure of table attr_users for DAO operations like Where/Data.
type AttrUsers struct {
	g.Meta          `orm:"table:attr_users, do:true"`
	Id              any         // id
	Country         any         // 国家
	Appid           any         // 应用ID
	Uuid            any         // APP用户ID
	IsRefund        any         // 是否退款 1-是 2-否
	IsRenew         any         // 是否续订
	RenewCount      any         // 续订次数
	DeductionCount  any         // 扣费次数
	LastInstallAt   *gtime.Time // 上次安装时间
	LastSubscribeAt *gtime.Time // 上次订阅时间
	CreatedAt       *gtime.Time // 创建时间
}
