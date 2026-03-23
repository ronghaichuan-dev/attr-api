package admin

import (
	g "github.com/gogf/gf/v2/frame/g"
)

// DashboardAnalyticsReq 仪表盘数据分析请求参数结构体
type DashboardAnalyticsReq struct {
	g.Meta `path:"/dashboard/analytics" method:"get" tags:"仪表盘数据分析" summary:"获取仪表盘数据分析"`
	AppIds string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，可选，多个ID用逗号分隔"` // 应用ID列表，可选（字符串格式，逗号分隔）
}

// DashboardAnalyticsRes 仪表盘数据分析响应参数结构体
type DashboardAnalyticsRes struct {
	TotalSubscriptionAmount float64                   `json:"total_subscription_amount" dc:"总订阅额"`
	TodaySubscriptionCount  int                       `json:"today_subscription_count" dc:"今日订阅量"`
	MonthSubscriptionCount  int                       `json:"month_subscription_count" dc:"当月订阅量"`
	MonthSubscriptionAmount float64                   `json:"month_subscription_amount" dc:"当月订阅总额"`
	DailyTrendData          []*DailySubscriptionTrend `json:"daily_trend_data" dc:"每日订阅趋势数据"`
}

// DailySubscriptionTrend 每日订阅趋势数据结构体
type DailySubscriptionTrend struct {
	Date   string  `json:"date" dc:"日期"`
	Count  int     `json:"count" dc:"订阅量"`
	Amount float64 `json:"amount" dc:"订阅金额"`
}

// AppDailyTrendReq 根据APP ID获取每日订阅量趋势请求参数结构体
type AppDailyTrendReq struct {
	g.Meta `path:"/app/daily/trend" method:"get" tags:"仪表盘数据分析" summary:"根据APP ID获取每日订阅量趋势数据"`
	AppIds string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，可选，多个ID用逗号分隔"`             // 应用ID列表，可选（字符串格式，逗号分隔）
	Days   int    `json:"days" form:"days" d:"30" v:"min:1|max:365#天数不能小于1|天数不能大于365" dc:"查询天数，默认30天，范围1-365"` // 查询天数，默认30天，范围1-365
}

// AppDailyTrendRes 根据APP ID获取每日订阅量趋势响应参数结构体
type AppDailyTrendRes struct {
	TrendData []*AppDailySubscriptionTrend `json:"trend_data" dc:"每日订阅趋势数据"`
}

// AppDailySubscriptionTrend 应用每日订阅趋势数据结构体
type AppDailySubscriptionTrend struct {
	AppId  string  `json:"app_id" dc:"应用ID"`
	Date   string  `json:"date" dc:"日期"`
	Count  int     `json:"count" dc:"订阅量"`
	Amount float64 `json:"amount" dc:"订阅金额"`
}

// AppSelectOption 应用下拉选项结构体
type AppSelectOption struct {
	AppId   string `json:"app_id" orm:"app_id" dc:"应用ID"`
	AppName string `json:"app_name" orm:"app_name" dc:"应用名称"`
}

// AppSelectListReq 获取应用下拉选项列表请求参数结构体
type AppSelectListReq struct {
	g.Meta `path:"/app/select/list" method:"get" tags:"仪表盘数据分析" summary:"获取应用下拉选项列表"`
}

// AppSelectListRes 获取应用下拉选项列表响应参数结构体
type AppSelectListRes struct {
	List []*AppSelectOption `json:"list" dc:"应用下拉选项列表"`
}

// ===== 渠道效果分析 =====

// ChannelOverviewReq 渠道效果概览请求
type ChannelOverviewReq struct {
	g.Meta    `path:"/channel/overview" method:"get" tags:"渠道分析" summary:"渠道效果概览"`
	AppIds    string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，逗号分隔"`
	StartDate string `json:"start_date" form:"start_date" v:"required#开始日期不能为空" dc:"开始日期 YYYY-MM-DD"`
	EndDate   string `json:"end_date" form:"end_date" v:"required#结束日期不能为空" dc:"结束日期 YYYY-MM-DD"`
}

// ChannelOverviewRes 渠道效果概览响应
type ChannelOverviewRes struct {
	List []*ChannelOverviewItem `json:"list" dc:"渠道效果列表"`
}

// ChannelOverviewItem 渠道效果数据项
type ChannelOverviewItem struct {
	TrackerNetwork     string  `json:"tracker_network" dc:"归因渠道"`
	InstallCount       int     `json:"install_count" dc:"安装量"`
	TrialCount         int     `json:"trial_count" dc:"试用量"`
	SubscribeCount     int     `json:"subscribe_count" dc:"订阅量"`
	RenewCount         int     `json:"renew_count" dc:"续订量"`
	RefundCount        int     `json:"refund_count" dc:"退款量"`
	Revenue            float64 `json:"revenue" dc:"收入（元）"`
	NetRevenue         float64 `json:"net_revenue" dc:"净收入（元）"`
	InstallToTrialRate float64 `json:"install_to_trial_rate" dc:"安装转试用率%"`
	TrialToPaidRate    float64 `json:"trial_to_paid_rate" dc:"试用转付费率%"`
}

// CampaignAnalysisReq 活动效果分析请求
type CampaignAnalysisReq struct {
	g.Meta         `path:"/campaign/analysis" method:"get" tags:"渠道分析" summary:"活动效果分析"`
	AppIds         string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，逗号分隔"`
	TrackerNetwork string `json:"tracker_network" form:"tracker_network" d:"" dc:"归因渠道筛选"`
	StartDate      string `json:"start_date" form:"start_date" v:"required#开始日期不能为空" dc:"开始日期 YYYY-MM-DD"`
	EndDate        string `json:"end_date" form:"end_date" v:"required#结束日期不能为空" dc:"结束日期 YYYY-MM-DD"`
}

// CampaignAnalysisRes 活动效果分析响应
type CampaignAnalysisRes struct {
	List []*CampaignAnalysisItem `json:"list" dc:"活动效果列表"`
}

// CampaignAnalysisItem 活动效果数据项
type CampaignAnalysisItem struct {
	CampaignId     string  `json:"campaign_id" dc:"推广活动ID"`
	TrackerNetwork string  `json:"tracker_network" dc:"归因渠道"`
	InstallCount   int     `json:"install_count" dc:"安装量"`
	SubscribeCount int     `json:"subscribe_count" dc:"订阅量"`
	Revenue        float64 `json:"revenue" dc:"收入（元）"`
	NetRevenue     float64 `json:"net_revenue" dc:"净收入（元）"`
}

// CountryRevenueReq 国家/地区收入分析请求
type CountryRevenueReq struct {
	g.Meta    `path:"/country/revenue" method:"get" tags:"渠道分析" summary:"国家/地区收入分析"`
	AppIds    string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，逗号分隔"`
	StartDate string `json:"start_date" form:"start_date" v:"required#开始日期不能为空" dc:"开始日期 YYYY-MM-DD"`
	EndDate   string `json:"end_date" form:"end_date" v:"required#结束日期不能为空" dc:"结束日期 YYYY-MM-DD"`
}

// CountryRevenueRes 国家/地区收入分析响应
type CountryRevenueRes struct {
	List []*CountryRevenueItem `json:"list" dc:"国家收入列表"`
}

// CountryRevenueItem 国家收入数据项
type CountryRevenueItem struct {
	Country        string  `json:"country" dc:"国家"`
	InstallCount   int     `json:"install_count" dc:"安装量"`
	SubscribeCount int     `json:"subscribe_count" dc:"订阅量"`
	Revenue        float64 `json:"revenue" dc:"收入（元）"`
	NetRevenue     float64 `json:"net_revenue" dc:"净收入（元）"`
}

// ===== 收入统计 =====

// RevenueTrendReq 收入趋势请求
type RevenueTrendReq struct {
	g.Meta    `path:"/revenue/trend" method:"get" tags:"收入统计" summary:"收入趋势"`
	AppIds    string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，逗号分隔"`
	StartDate string `json:"start_date" form:"start_date" v:"required#开始日期不能为空" dc:"开始日期 YYYY-MM-DD"`
	EndDate   string `json:"end_date" form:"end_date" v:"required#结束日期不能为空" dc:"结束日期 YYYY-MM-DD"`
	GroupBy   string `json:"group_by" form:"group_by" d:"day" dc:"分组方式: day/week/month"`
}

// RevenueTrendRes 收入趋势响应
type RevenueTrendRes struct {
	List []*RevenueTrendItem `json:"list" dc:"收入趋势列表"`
}

// RevenueTrendItem 收入趋势数据项
type RevenueTrendItem struct {
	Period         string  `json:"period" dc:"时间段"`
	Revenue        float64 `json:"revenue" dc:"总收入（元）"`
	NetRevenue     float64 `json:"net_revenue" dc:"净收入（元）"`
	SubscribeCount int     `json:"subscribe_count" dc:"订阅数"`
	RefundCount    int     `json:"refund_count" dc:"退款数"`
}

// RevenueSummaryReq 收入概览请求
type RevenueSummaryReq struct {
	g.Meta    `path:"/revenue/summary" method:"get" tags:"收入统计" summary:"收入概览KPI"`
	AppIds    string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，逗号分隔"`
	StartDate string `json:"start_date" form:"start_date" v:"required#开始日期不能为空" dc:"开始日期 YYYY-MM-DD"`
	EndDate   string `json:"end_date" form:"end_date" v:"required#结束日期不能为空" dc:"结束日期 YYYY-MM-DD"`
}

// RevenueSummaryRes 收入概览响应
type RevenueSummaryRes struct {
	TotalRevenue   float64 `json:"total_revenue" dc:"总收入（元）"`
	NetRevenue     float64 `json:"net_revenue" dc:"净收入（元）"`
	RefundAmount   float64 `json:"refund_amount" dc:"退款金额（元）"`
	SubscribeCount int     `json:"subscribe_count" dc:"订阅总数"`
	RenewCount     int     `json:"renew_count" dc:"续订总数"`
	RefundCount    int     `json:"refund_count" dc:"退款总数"`
	RefundRate     float64 `json:"refund_rate" dc:"退款率%"`
	RenewRate      float64 `json:"renew_rate" dc:"续订率%"`
	ARPU           float64 `json:"arpu" dc:"每用户平均收入（元）"`
}

// RevenueByProductReq 产品维度收入请求
type RevenueByProductReq struct {
	g.Meta    `path:"/revenue/by-product" method:"get" tags:"收入统计" summary:"产品维度收入"`
	AppIds    string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，逗号分隔"`
	StartDate string `json:"start_date" form:"start_date" v:"required#开始日期不能为空" dc:"开始日期 YYYY-MM-DD"`
	EndDate   string `json:"end_date" form:"end_date" v:"required#结束日期不能为空" dc:"结束日期 YYYY-MM-DD"`
}

// RevenueByProductRes 产品维度收入响应
type RevenueByProductRes struct {
	List []*RevenueByProductItem `json:"list" dc:"产品收入列表"`
}

// RevenueByProductItem 产品收入数据项
type RevenueByProductItem struct {
	ProductId      string  `json:"product_id" dc:"产品ID"`
	SubscribeCount int     `json:"subscribe_count" dc:"订阅数"`
	Revenue        float64 `json:"revenue" dc:"收入（元）"`
	Percentage     float64 `json:"percentage" dc:"收入占比%"`
}

// ===== LTV 分析 =====

// LtvByChannelReq 渠道LTV请求
type LtvByChannelReq struct {
	g.Meta           `path:"/ltv/by-channel" method:"get" tags:"LTV分析" summary:"渠道LTV分析"`
	AppIds           string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，逗号分隔"`
	InstallStartDate string `json:"install_start_date" form:"install_start_date" v:"required#安装开始日期不能为空" dc:"安装开始日期 YYYY-MM-DD"`
	InstallEndDate   string `json:"install_end_date" form:"install_end_date" v:"required#安装结束日期不能为空" dc:"安装结束日期 YYYY-MM-DD"`
}

// LtvByChannelRes 渠道LTV响应
type LtvByChannelRes struct {
	List []*LtvByChannelItem `json:"list" dc:"渠道LTV列表"`
}

// LtvByChannelItem 渠道LTV数据项
type LtvByChannelItem struct {
	TrackerNetwork string  `json:"tracker_network" dc:"归因渠道"`
	InstallCount   int     `json:"install_count" dc:"安装量"`
	PaidUserCount  int     `json:"paid_user_count" dc:"付费用户数"`
	TotalRevenue   float64 `json:"total_revenue" dc:"总收入（元）"`
	LTV            float64 `json:"ltv" dc:"人均LTV（元）"`
	PaidRate       float64 `json:"paid_rate" dc:"付费率%"`
}

// LtvCohortReq 安装队列LTV请求
type LtvCohortReq struct {
	g.Meta     `path:"/ltv/cohort" method:"get" tags:"LTV分析" summary:"安装队列LTV（Cohort）"`
	AppIds     string `json:"app_ids" form:"app_ids" d:"" dc:"应用ID列表，逗号分隔"`
	CohortType string `json:"cohort_type" form:"cohort_type" d:"day" dc:"队列类型: day/week/month"`
	StartDate  string `json:"start_date" form:"start_date" v:"required#开始日期不能为空" dc:"开始日期 YYYY-MM-DD"`
	EndDate    string `json:"end_date" form:"end_date" v:"required#结束日期不能为空" dc:"结束日期 YYYY-MM-DD"`
}

// LtvCohortRes 安装队列LTV响应
type LtvCohortRes struct {
	List []*LtvCohortItem `json:"list" dc:"队列LTV列表"`
}

// LtvCohortItem 安装队列LTV数据项
type LtvCohortItem struct {
	CohortDate   string  `json:"cohort_date" dc:"队列日期"`
	InstallCount int     `json:"install_count" dc:"安装量"`
	D7LTV        float64 `json:"d7_ltv" dc:"D7 LTV（元）"`
	D30LTV       float64 `json:"d30_ltv" dc:"D30 LTV（元）"`
	D60LTV       float64 `json:"d60_ltv" dc:"D60 LTV（元）"`
	D90LTV       float64 `json:"d90_ltv" dc:"D90 LTV（元）"`
}
