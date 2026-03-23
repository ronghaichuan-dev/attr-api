// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrDevice is the golang structure for table attr_device.
type AttrDevice struct {
	Id                 int64  `json:"id"                 orm:"id"                   description:"id"`           // id
	Rsid               string `json:"rsid"               orm:"rsid"                 description:"设备ID"`         // 设备ID
	Appid              string `json:"appid"              orm:"appid"                description:"应用ID"`         // 应用ID
	AttrSubscriptionId int64  `json:"attrSubscriptionId" orm:"attr_subscription_id" description:"归因订阅ID"`       // 归因订阅ID
	Country            string `json:"country"            orm:"country"              description:"国家"`           // 国家
	TrackerNetwork     string `json:"trackerNetwork"     orm:"tracker_network"      description:"归因渠道"`         // 归因渠道
	CampaignId         string `json:"campaignId"         orm:"campaign_id"          description:"推广活动ID"`       // 推广活动ID
	AdgroupId          string `json:"adgroupId"          orm:"adgroup_id"           description:"广告组ID"`        // 广告组ID
	AdId               string `json:"adId"               orm:"ad_id"                description:"广告ID"`         // 广告ID
	KeywordId          string `json:"keywordId"          orm:"keyword_id"           description:"关键词ID"`        // 关键词ID
	AttrInstallId      int64  `json:"attrInstallId"      orm:"attr_install_id"      description:"关联安装归因记录ID"`   // 关联安装归因记录ID
	IsFirstInstall     int    `json:"isFirstInstall"     orm:"is_first_install"     description:"是否首次安装"`       // 是否首次安装
	Channel            string `json:"channel"            orm:"channel"              description:"渠道来源"`         // 渠道来源
	IsRefund           int    `json:"isRefund"           orm:"is_refund"            description:"是否退款 1-是 2-否"` // 是否退款 1-是 2-否
	IsRenew            int    `json:"isRenew"            orm:"is_renew"             description:"是否续订"`         // 是否续订
	RenewCount         int    `json:"renewCount"         orm:"renew_count"          description:"续订次数"`         // 续订次数
	DeductionCount     int    `json:"deductionCount"     orm:"deduction_count"      description:"扣费次数"`         // 扣费次数
	CreatedAt          int64  `json:"createdAt"          orm:"created_at"           description:"创建时间"`         // 创建时间
	LastInstallAt      int64  `json:"lastInstallAt"      orm:"last_install_at"      description:"最后安装时间"`       // 最后安装时间
	LastTrialAt        int64  `json:"lastTrialAt"        orm:"last_trial_at"        description:"最后试用时间"`       // 最后试用时间
	LastSubscribeAt    int64  `json:"lastSubscribeAt"    orm:"last_subscribe_at"    description:"最后订阅时间"`       // 最后订阅时间
	LastRenewAt        int64  `json:"lastRenewAt"        orm:"last_renew_at"        description:"最后续费时间"`       // 最后续费时间
	LastRefundAt       int64  `json:"lastRefundAt"       orm:"last_refund_at"       description:"最后退款时间"`       // 最后退款时间
}
