// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrDevice is the golang structure of table attr_device for DAO operations like Where/Data.
type AttrDevice struct {
	g.Meta             `orm:"table:attr_device, do:true"`
	Id                 any // id
	Rsid               any // 设备ID
	Appid              any // 应用ID
	AttrSubscriptionId any // 归因订阅ID
	Country            any // 国家
	TrackerNetwork     any // 归因渠道
	CampaignId         any // 推广活动ID
	AdgroupId          any // 广告组ID
	AdId               any // 广告ID
	KeywordId          any // 关键词ID
	AttrInstallId      any // 关联安装归因记录ID
	IsFirstInstall     any // 是否首次安装
	Channel            any // 渠道来源
	IsRefund           any // 是否退款 1-是 2-否
	IsRenew            any // 是否续订
	RenewCount         any // 续订次数
	DeductionCount     any // 扣费次数
	CreatedAt          any // 创建时间
	LastInstallAt      any // 最后安装时间
	LastTrialAt        any // 最后试用时间
	LastSubscribeAt    any // 最后订阅时间
	LastRenewAt        any // 最后续费时间
	LastRefundAt       any // 最后退款时间
}
