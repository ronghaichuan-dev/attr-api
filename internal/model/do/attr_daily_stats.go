// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrDailyStats is the golang structure of table attr_daily_stats for DAO operations like Where/Data.
type AttrDailyStats struct {
	g.Meta             `orm:"table:attr_daily_stats, do:true"`
	Id                 any //
	StatDate           any // 统计日期 YYYY-MM-DD
	AppId              any // 应用ID
	Country            any // 国家
	TrackerNetwork     any // 归因渠道
	CampaignId         any // 推广活动ID
	InstallCount       any // 安装量
	TrialCount         any // 试用量
	SubscribeCount     any // 订阅量（付费）
	RenewCount         any // 续订量
	RefundCount        any // 退款量
	Revenue            any // 收入（分）
	RefundAmount       any // 退款金额（分）
	NetRevenue         any // 净收入（分）
	InstallToTrialRate any // 安装转试用率%
	TrialToPaidRate    any // 试用转付费率%
	CreatedAt          any //
	UpdatedAt          any //
}
