// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrUsers is the golang structure for table attr_users.
type AttrUsers struct {
	Id              int64       `json:"id"              orm:"id"                description:"id"`           // id
	Country         string      `json:"country"         orm:"country"           description:"国家"`           // 国家
	Appid           string      `json:"appid"           orm:"appid"             description:"应用ID"`         // 应用ID
	Uuid            string      `json:"uuid"            orm:"uuid"              description:"APP用户ID"`      // APP用户ID
	IsRefund        int         `json:"isRefund"        orm:"is_refund"         description:"是否退款 1-是 2-否"` // 是否退款 1-是 2-否
	IsRenew         int         `json:"isRenew"         orm:"is_renew"          description:"是否续订"`         // 是否续订
	RenewCount      int         `json:"renewCount"      orm:"renew_count"       description:"续订次数"`         // 续订次数
	DeductionCount  int         `json:"deductionCount"  orm:"deduction_count"   description:"扣费次数"`         // 扣费次数
	LastInstallAt   *gtime.Time `json:"lastInstallAt"   orm:"last_install_at"   description:"上次安装时间"`       // 上次安装时间
	LastSubscribeAt *gtime.Time `json:"lastSubscribeAt" orm:"last_subscribe_at" description:"上次订阅时间"`       // 上次订阅时间
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:"创建时间"`         // 创建时间
}
