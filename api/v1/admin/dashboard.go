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
