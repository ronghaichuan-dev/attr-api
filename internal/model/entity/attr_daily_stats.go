// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrDailyStats is the golang structure for table attr_daily_stats.
type AttrDailyStats struct {
	Id                 int64   `json:"id"                 orm:"id"                    description:""`                //
	StatDate           string  `json:"statDate"           orm:"stat_date"             description:"统计日期 YYYY-MM-DD"` // 统计日期 YYYY-MM-DD
	AppId              string  `json:"appId"              orm:"app_id"                description:"应用ID"`            // 应用ID
	Country            string  `json:"country"            orm:"country"               description:"国家"`              // 国家
	TrackerNetwork     string  `json:"trackerNetwork"     orm:"tracker_network"       description:"归因渠道"`            // 归因渠道
	CampaignId         string  `json:"campaignId"         orm:"campaign_id"           description:"推广活动ID"`          // 推广活动ID
	InstallCount       int     `json:"installCount"       orm:"install_count"         description:"安装量"`             // 安装量
	TrialCount         int     `json:"trialCount"         orm:"trial_count"           description:"试用量"`             // 试用量
	SubscribeCount     int     `json:"subscribeCount"     orm:"subscribe_count"       description:"订阅量（付费）"`         // 订阅量（付费）
	RenewCount         int     `json:"renewCount"         orm:"renew_count"           description:"续订量"`             // 续订量
	RefundCount        int     `json:"refundCount"        orm:"refund_count"          description:"退款量"`             // 退款量
	Revenue            int64   `json:"revenue"            orm:"revenue"               description:"收入（分）"`           // 收入（分）
	RefundAmount       int64   `json:"refundAmount"       orm:"refund_amount"         description:"退款金额（分）"`         // 退款金额（分）
	NetRevenue         int64   `json:"netRevenue"         orm:"net_revenue"           description:"净收入（分）"`          // 净收入（分）
	InstallToTrialRate float64 `json:"installToTrialRate" orm:"install_to_trial_rate" description:"安装转试用率%"`         // 安装转试用率%
	TrialToPaidRate    float64 `json:"trialToPaidRate"    orm:"trial_to_paid_rate"    description:"试用转付费率%"`         // 试用转付费率%
	CreatedAt          int64   `json:"createdAt"          orm:"created_at"            description:""`                //
	UpdatedAt          int64   `json:"updatedAt"          orm:"updated_at"            description:""`                //
}
