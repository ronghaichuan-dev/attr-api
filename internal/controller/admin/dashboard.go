package admin

import (
	"context"
	"fmt"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"sort"
	"strings"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

func isStatsMapEmpty(statsMap map[string]map[string]interface{}) bool {
	if len(statsMap) == 0 {
		return true
	}

	for _, dateStats := range statsMap {
		if len(dateStats) > 0 {
			for _, stat := range dateStats {
				statMap, ok := stat.(map[string]interface{})
				if ok && len(statMap) > 0 {
					return false
				}
			}
		}
	}

	return true
}

type DashboardController struct{}

func (c *DashboardController) DashboardAnalytics(ctx context.Context, req *adminApi.DashboardAnalyticsReq) (*adminApi.DashboardAnalyticsRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	var appIds []string
	if req.AppIds != "" {
		appIds = strings.Split(strings.TrimSpace(req.AppIds), ",")
		filteredAppIds := make([]string, 0, len(appIds))
		for _, id := range appIds {
			if id = strings.TrimSpace(id); id != "" {
				filteredAppIds = append(filteredAppIds, id)
			}
		}
		appIds = filteredAppIds
	}
	logger.Debugf("解析后的AppIds数组: %v", appIds)

	today := gtime.Now().Format("Y-m-d")
	month := gtime.Now().Format("Y-m")
	startDate := gtime.Now().AddDate(0, 0, -29).Format("Y-m-d")
	todayTime := gtime.NewFromStr(today)
	startDateTime := gtime.NewFromStr(startDate)

	res := &adminApi.DashboardAnalyticsRes{}

	m := dao.AttrSubscriptionTransaction.Ctx(ctx)

	if len(appIds) > 0 {
		m = m.WhereIn("appid", appIds)
	}

	m = m.Where("purchase_at >=", startDateTime.Timestamp())
	m = m.Where("purchase_at <=", todayTime.Timestamp())

	total, err := m.Count()
	if err != nil {
		logger.Errorf("查询订阅统计失败: %v", err)
		return res, nil
	}

	if total == 0 {
		return res, nil
	}

	var transactions []*entity.AttrSubscriptionTransaction
	err = m.Order("purchase_at ASC").Scan(&transactions)
	if err != nil {
		logger.Errorf("查询订阅记录失败: %v", err)
		return res, nil
	}

	var totalAmount float64
	var todayCount int
	var todayAmount float64
	var monthCount int
	var monthAmount float64
	dateMap := make(map[string]*adminApi.DailySubscriptionTrend)

	for _, trans := range transactions {
		amount := float64(trans.Price) / 100.0

		totalAmount += amount

		transDate := gtime.NewFromTimeStamp(trans.PurchaseAt).Format("Y-m-d")
		if transDate == today {
			todayCount++
			todayAmount += amount
		}

		if transDate[:7] == month {
			monthCount++
			monthAmount += amount
		}

		if _, ok := dateMap[transDate]; !ok {
			dateMap[transDate] = &adminApi.DailySubscriptionTrend{
				Date:   transDate,
				Count:  0,
				Amount: 0,
			}
		}
		dateMap[transDate].Count++
		dateMap[transDate].Amount += amount
	}

	res.TotalSubscriptionAmount = totalAmount
	res.TodaySubscriptionCount = todayCount
	res.MonthSubscriptionCount = monthCount
	res.MonthSubscriptionAmount = monthAmount

	res.DailyTrendData = make([]*adminApi.DailySubscriptionTrend, 0, len(dateMap))
	for _, trend := range dateMap {
		res.DailyTrendData = append(res.DailyTrendData, trend)
	}

	for i := 0; i < len(res.DailyTrendData); i++ {
		for j := i + 1; j < len(res.DailyTrendData); j++ {
			if res.DailyTrendData[i].Date > res.DailyTrendData[j].Date {
				res.DailyTrendData[i], res.DailyTrendData[j] = res.DailyTrendData[j], res.DailyTrendData[i]
			}
		}
	}

	return res, nil
}

func (c *DashboardController) AppDailyTrend(ctx context.Context, req *adminApi.AppDailyTrendReq) (*adminApi.AppDailyTrendRes, error) {
	logger.Debugf("=== 开始执行 AppDailyTrend 函数 ===")
	logger.Debugf("解析后的请求参数: %v", req)

	var appIds []string
	if req.AppIds != "" {
		appIds = strings.Split(strings.TrimSpace(req.AppIds), ",")
		filteredAppIds := make([]string, 0, len(appIds))
		for _, id := range appIds {
			if id = strings.TrimSpace(id); id != "" {
				filteredAppIds = append(filteredAppIds, id)
			}
		}
		appIds = filteredAppIds
	}
	logger.Debugf("解析后的AppIds数组: %v", appIds)

	startDate := gtime.Now().AddDate(0, 0, -req.Days+1).Format("Y-m-d")
	endDate := gtime.Now().Format("Y-m-d")
	startDateTime := gtime.NewFromStr(startDate)
	endDateTime := gtime.NewFromStr(endDate)
	logger.Debugf("查询日期范围: startDate=%s endDate=%s", startDate, endDate)

	m := dao.AttrSubscriptionTransaction.Ctx(ctx).
		Where("purchase_at >=", startDateTime.Timestamp()).
		Where("purchase_at <=", endDateTime.Timestamp())

	if len(appIds) > 0 {
		m = m.WhereIn("appid", appIds)
	}

	var transactions []*entity.AttrSubscriptionTransaction
	err := m.Order("appid ASC, purchase_at ASC").Scan(&transactions)
	if err != nil {
		logger.Errorf("查询订阅趋势数据失败: %v", err)
		return &adminApi.AppDailyTrendRes{TrendData: nil}, nil
	}

	// 按 app_id + date 聚合
	type aggKey struct {
		AppId string
		Date  string
	}
	aggMap := make(map[aggKey]*adminApi.AppDailySubscriptionTrend)

	for _, trans := range transactions {
		amount := float64(trans.Price) / 100.0
		transDate := gtime.NewFromTimeStamp(trans.PurchaseAt).Format("Y-m-d")
		key := aggKey{AppId: trans.Appid, Date: transDate}

		if _, ok := aggMap[key]; !ok {
			aggMap[key] = &adminApi.AppDailySubscriptionTrend{
				AppId: trans.Appid,
				Date:  transDate,
			}
		}
		aggMap[key].Count++
		aggMap[key].Amount += amount
	}

	trendData := make([]*adminApi.AppDailySubscriptionTrend, 0, len(aggMap))
	for _, v := range aggMap {
		trendData = append(trendData, v)
	}

	sort.Slice(trendData, func(i, j int) bool {
		if trendData[i].AppId != trendData[j].AppId {
			return trendData[i].AppId < trendData[j].AppId
		}
		return trendData[i].Date < trendData[j].Date
	})

	logger.Debugf("最终返回数据: %v", trendData)
	return &adminApi.AppDailyTrendRes{TrendData: trendData}, nil
}

func (c *DashboardController) AppSelectList(ctx context.Context, req *adminApi.AppSelectListReq) (*adminApi.AppSelectListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	list, err := service.System().GetAppSelectList(ctx)
	if err != nil {
		logger.Errorf("获取应用下拉选项列表失败: %v", err)
		return nil, err
	}

	logger.Debugf("DAO返回的列表长度: %d", len(list))
	if len(list) > 0 {
		logger.Debugf("DAO返回的第一条记录: %v", list[0])
	}

	res := &adminApi.AppSelectListRes{List: list}
	logger.Debugf("返回的响应: %v", res)

	return res, nil
}

// parseAppIds 解析逗号分隔的 AppIds 字符串
func parseAppIds(appIdsStr string) []string {
	if appIdsStr == "" {
		return nil
	}
	parts := strings.Split(strings.TrimSpace(appIdsStr), ",")
	result := make([]string, 0, len(parts))
	for _, id := range parts {
		if id = strings.TrimSpace(id); id != "" {
			result = append(result, id)
		}
	}
	return result
}

// ===== 渠道效果分析 =====

// ChannelOverview 渠道效果概览
func (c *DashboardController) ChannelOverview(ctx context.Context, req *adminApi.ChannelOverviewReq) (*adminApi.ChannelOverviewRes, error) {
	appIds := parseAppIds(req.AppIds)

	type channelRow struct {
		TrackerNetwork     string  `json:"tracker_network"`
		InstallCount       int     `json:"install_count"`
		TrialCount         int     `json:"trial_count"`
		SubscribeCount     int     `json:"subscribe_count"`
		RenewCount         int     `json:"renew_count"`
		RefundCount        int     `json:"refund_count"`
		Revenue            int64   `json:"revenue"`
		RefundAmount       int64   `json:"refund_amount"`
		NetRevenue         int64   `json:"net_revenue"`
		InstallToTrialRate float64 `json:"install_to_trial_rate"`
		TrialToPaidRate    float64 `json:"trial_to_paid_rate"`
	}

	m := dao.AttrDailyStats.Ctx(ctx).
		Fields("tracker_network, SUM(install_count) as install_count, SUM(trial_count) as trial_count, SUM(subscribe_count) as subscribe_count, SUM(renew_count) as renew_count, SUM(refund_count) as refund_count, SUM(revenue) as revenue, SUM(refund_amount) as refund_amount, SUM(net_revenue) as net_revenue").
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate)

	if len(appIds) > 0 {
		m = m.WhereIn("app_id", appIds)
	}

	var rows []channelRow
	err := m.Group("tracker_network").OrderDesc("revenue").Scan(&rows)
	if err != nil {
		logger.Errorf("查询渠道效果失败: %v", err)
		return &adminApi.ChannelOverviewRes{}, nil
	}

	list := make([]*adminApi.ChannelOverviewItem, 0, len(rows))
	for _, r := range rows {
		item := &adminApi.ChannelOverviewItem{
			TrackerNetwork: r.TrackerNetwork,
			InstallCount:   r.InstallCount,
			TrialCount:     r.TrialCount,
			SubscribeCount: r.SubscribeCount,
			RenewCount:     r.RenewCount,
			RefundCount:    r.RefundCount,
			Revenue:        float64(r.Revenue) / 100.0,
			NetRevenue:     float64(r.NetRevenue) / 100.0,
		}
		if r.InstallCount > 0 {
			item.InstallToTrialRate = float64(r.TrialCount) / float64(r.InstallCount) * 100
		}
		if r.TrialCount > 0 {
			item.TrialToPaidRate = float64(r.SubscribeCount) / float64(r.TrialCount) * 100
		}
		list = append(list, item)
	}

	return &adminApi.ChannelOverviewRes{List: list}, nil
}

// CampaignAnalysis 活动效果分析
func (c *DashboardController) CampaignAnalysis(ctx context.Context, req *adminApi.CampaignAnalysisReq) (*adminApi.CampaignAnalysisRes, error) {
	appIds := parseAppIds(req.AppIds)

	type campaignRow struct {
		CampaignId     string `json:"campaign_id"`
		TrackerNetwork string `json:"tracker_network"`
		InstallCount   int    `json:"install_count"`
		SubscribeCount int    `json:"subscribe_count"`
		Revenue        int64  `json:"revenue"`
		NetRevenue     int64  `json:"net_revenue"`
	}

	m := dao.AttrDailyStats.Ctx(ctx).
		Fields("campaign_id, tracker_network, SUM(install_count) as install_count, SUM(subscribe_count) as subscribe_count, SUM(revenue) as revenue, SUM(net_revenue) as net_revenue").
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate)

	if len(appIds) > 0 {
		m = m.WhereIn("app_id", appIds)
	}
	if req.TrackerNetwork != "" {
		m = m.Where("tracker_network", req.TrackerNetwork)
	}

	var rows []campaignRow
	err := m.Group("campaign_id, tracker_network").OrderDesc("revenue").Scan(&rows)
	if err != nil {
		logger.Errorf("查询活动效果失败: %v", err)
		return &adminApi.CampaignAnalysisRes{}, nil
	}

	list := make([]*adminApi.CampaignAnalysisItem, 0, len(rows))
	for _, r := range rows {
		list = append(list, &adminApi.CampaignAnalysisItem{
			CampaignId:     r.CampaignId,
			TrackerNetwork: r.TrackerNetwork,
			InstallCount:   r.InstallCount,
			SubscribeCount: r.SubscribeCount,
			Revenue:        float64(r.Revenue) / 100.0,
			NetRevenue:     float64(r.NetRevenue) / 100.0,
		})
	}

	return &adminApi.CampaignAnalysisRes{List: list}, nil
}

// CountryRevenue 国家/地区收入分析
func (c *DashboardController) CountryRevenue(ctx context.Context, req *adminApi.CountryRevenueReq) (*adminApi.CountryRevenueRes, error) {
	appIds := parseAppIds(req.AppIds)

	type countryRow struct {
		Country        string `json:"country"`
		InstallCount   int    `json:"install_count"`
		SubscribeCount int    `json:"subscribe_count"`
		Revenue        int64  `json:"revenue"`
		NetRevenue     int64  `json:"net_revenue"`
	}

	m := dao.AttrDailyStats.Ctx(ctx).
		Fields("country, SUM(install_count) as install_count, SUM(subscribe_count) as subscribe_count, SUM(revenue) as revenue, SUM(net_revenue) as net_revenue").
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate)

	if len(appIds) > 0 {
		m = m.WhereIn("app_id", appIds)
	}

	var rows []countryRow
	err := m.Group("country").OrderDesc("revenue").Scan(&rows)
	if err != nil {
		logger.Errorf("查询国家收入失败: %v", err)
		return &adminApi.CountryRevenueRes{}, nil
	}

	list := make([]*adminApi.CountryRevenueItem, 0, len(rows))
	for _, r := range rows {
		list = append(list, &adminApi.CountryRevenueItem{
			Country:        r.Country,
			InstallCount:   r.InstallCount,
			SubscribeCount: r.SubscribeCount,
			Revenue:        float64(r.Revenue) / 100.0,
			NetRevenue:     float64(r.NetRevenue) / 100.0,
		})
	}

	return &adminApi.CountryRevenueRes{List: list}, nil
}

// ===== 收入统计 =====

// RevenueTrend 收入趋势
func (c *DashboardController) RevenueTrend(ctx context.Context, req *adminApi.RevenueTrendReq) (*adminApi.RevenueTrendRes, error) {
	appIds := parseAppIds(req.AppIds)

	// 根据 group_by 确定日期分组表达式
	var dateExpr string
	switch req.GroupBy {
	case "week":
		dateExpr = "DATE_FORMAT(stat_date, '%x-W%v')"
	case "month":
		dateExpr = "DATE_FORMAT(stat_date, '%Y-%m')"
	default:
		dateExpr = "stat_date"
	}

	type trendRow struct {
		Period         string `json:"period"`
		Revenue        int64  `json:"revenue"`
		NetRevenue     int64  `json:"net_revenue"`
		SubscribeCount int    `json:"subscribe_count"`
		RefundCount    int    `json:"refund_count"`
	}

	m := dao.AttrDailyStats.Ctx(ctx).
		Fields(fmt.Sprintf("%s as period, SUM(revenue) as revenue, SUM(net_revenue) as net_revenue, SUM(subscribe_count) as subscribe_count, SUM(refund_count) as refund_count", dateExpr)).
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate)

	if len(appIds) > 0 {
		m = m.WhereIn("app_id", appIds)
	}

	var rows []trendRow
	err := m.Group("period").Order("period ASC").Scan(&rows)
	if err != nil {
		logger.Errorf("查询收入趋势失败: %v", err)
		return &adminApi.RevenueTrendRes{}, nil
	}

	list := make([]*adminApi.RevenueTrendItem, 0, len(rows))
	for _, r := range rows {
		list = append(list, &adminApi.RevenueTrendItem{
			Period:         r.Period,
			Revenue:        float64(r.Revenue) / 100.0,
			NetRevenue:     float64(r.NetRevenue) / 100.0,
			SubscribeCount: r.SubscribeCount,
			RefundCount:    r.RefundCount,
		})
	}

	return &adminApi.RevenueTrendRes{List: list}, nil
}

// RevenueSummary 收入概览KPI
func (c *DashboardController) RevenueSummary(ctx context.Context, req *adminApi.RevenueSummaryReq) (*adminApi.RevenueSummaryRes, error) {
	appIds := parseAppIds(req.AppIds)

	type summaryRow struct {
		Revenue        int64 `json:"revenue"`
		RefundAmount   int64 `json:"refund_amount"`
		NetRevenue     int64 `json:"net_revenue"`
		SubscribeCount int   `json:"subscribe_count"`
		RenewCount     int   `json:"renew_count"`
		RefundCount    int   `json:"refund_count"`
		InstallCount   int   `json:"install_count"`
	}

	m := dao.AttrDailyStats.Ctx(ctx).
		Fields("SUM(revenue) as revenue, SUM(refund_amount) as refund_amount, SUM(net_revenue) as net_revenue, SUM(subscribe_count) as subscribe_count, SUM(renew_count) as renew_count, SUM(refund_count) as refund_count, SUM(install_count) as install_count").
		Where("stat_date >= ?", req.StartDate).
		Where("stat_date <= ?", req.EndDate)

	if len(appIds) > 0 {
		m = m.WhereIn("app_id", appIds)
	}

	var row summaryRow
	err := m.Scan(&row)
	if err != nil {
		logger.Errorf("查询收入概览失败: %v", err)
		return &adminApi.RevenueSummaryRes{}, nil
	}

	res := &adminApi.RevenueSummaryRes{
		TotalRevenue:   float64(row.Revenue) / 100.0,
		NetRevenue:     float64(row.NetRevenue) / 100.0,
		RefundAmount:   float64(row.RefundAmount) / 100.0,
		SubscribeCount: row.SubscribeCount,
		RenewCount:     row.RenewCount,
		RefundCount:    row.RefundCount,
	}

	totalSubs := row.SubscribeCount + row.RenewCount
	if totalSubs > 0 {
		res.RefundRate = float64(row.RefundCount) / float64(totalSubs) * 100
	}
	if row.SubscribeCount > 0 {
		res.RenewRate = float64(row.RenewCount) / float64(row.SubscribeCount) * 100
	}
	if row.InstallCount > 0 {
		res.ARPU = float64(row.NetRevenue) / 100.0 / float64(row.InstallCount)
	}

	return res, nil
}

// RevenueByProduct 产品维度收入
func (c *DashboardController) RevenueByProduct(ctx context.Context, req *adminApi.RevenueByProductReq) (*adminApi.RevenueByProductRes, error) {
	appIds := parseAppIds(req.AppIds)

	startDate, _ := time.Parse("2006-01-02", req.StartDate)
	endDate, _ := time.Parse("2006-01-02", req.EndDate)
	startTs := startDate.Unix()
	endTs := endDate.AddDate(0, 0, 1).Unix()

	type productRow struct {
		ProductId      string `json:"product_id"`
		SubscribeCount int    `json:"subscribe_count"`
		Revenue        int64  `json:"revenue"`
	}

	m := dao.AttrSubscriptionTransaction.Ctx(ctx).
		Fields("product_id, COUNT(*) as subscribe_count, SUM(price) as revenue").
		Where("created_at >= ?", startTs).
		Where("created_at < ?", endTs)

	if len(appIds) > 0 {
		m = m.WhereIn("appid", appIds)
	}

	var rows []productRow
	err := m.Group("product_id").OrderDesc("revenue").Scan(&rows)
	if err != nil {
		logger.Errorf("查询产品收入失败: %v", err)
		return &adminApi.RevenueByProductRes{}, nil
	}

	var totalRevenue int64
	for _, r := range rows {
		totalRevenue += r.Revenue
	}

	list := make([]*adminApi.RevenueByProductItem, 0, len(rows))
	for _, r := range rows {
		var pct float64
		if totalRevenue > 0 {
			pct = float64(r.Revenue) / float64(totalRevenue) * 100
		}
		list = append(list, &adminApi.RevenueByProductItem{
			ProductId:      r.ProductId,
			SubscribeCount: r.SubscribeCount,
			Revenue:        float64(r.Revenue) / 100.0,
			Percentage:     pct,
		})
	}

	return &adminApi.RevenueByProductRes{List: list}, nil
}

// ===== LTV 分析 =====

// LtvByChannel 渠道LTV分析
func (c *DashboardController) LtvByChannel(ctx context.Context, req *adminApi.LtvByChannelReq) (*adminApi.LtvByChannelRes, error) {
	appIds := parseAppIds(req.AppIds)

	startDate, _ := time.Parse("2006-01-02", req.InstallStartDate)
	endDate, _ := time.Parse("2006-01-02", req.InstallEndDate)
	startTs := startDate.Unix()
	endTs := endDate.AddDate(0, 0, 1).Unix()

	// 从 attr_device 获取各渠道安装量
	type channelInstall struct {
		TrackerNetwork string `json:"tracker_network"`
		InstallCount   int    `json:"install_count"`
	}

	installM := dao.AttrDevice.Ctx(ctx).
		Fields("IFNULL(tracker_network,'') as tracker_network, COUNT(*) as install_count").
		Where("created_at >= ?", startTs).
		Where("created_at < ?", endTs)

	if len(appIds) > 0 {
		installM = installM.WhereIn("appid", appIds)
	}

	var installRows []channelInstall
	err := installM.Group("tracker_network").Scan(&installRows)
	if err != nil {
		logger.Errorf("查询渠道安装量失败: %v", err)
		return &adminApi.LtvByChannelRes{}, nil
	}

	// 从 attr_subscription_transaction 获取各渠道收入
	type channelRevenue struct {
		TrackerNetwork string `json:"tracker_network"`
		PaidUserCount  int    `json:"paid_user_count"`
		TotalRevenue   int64  `json:"total_revenue"`
	}

	revenueM := dao.AttrSubscriptionTransaction.Ctx(ctx).
		Fields("IFNULL(tracker_network,'') as tracker_network, COUNT(DISTINCT rsid) as paid_user_count, SUM(price) as total_revenue").
		Where("created_at >= ?", startTs).
		Where("transaction_type != 'REFUND'")

	if len(appIds) > 0 {
		revenueM = revenueM.WhereIn("appid", appIds)
	}

	var revenueRows []channelRevenue
	err = revenueM.Group("tracker_network").Scan(&revenueRows)
	if err != nil {
		logger.Errorf("查询渠道收入失败: %v", err)
		return &adminApi.LtvByChannelRes{}, nil
	}

	// 合并
	revenueMap := make(map[string]*channelRevenue)
	for i := range revenueRows {
		revenueMap[revenueRows[i].TrackerNetwork] = &revenueRows[i]
	}

	list := make([]*adminApi.LtvByChannelItem, 0, len(installRows))
	for _, inst := range installRows {
		item := &adminApi.LtvByChannelItem{
			TrackerNetwork: inst.TrackerNetwork,
			InstallCount:   inst.InstallCount,
		}
		if rev, ok := revenueMap[inst.TrackerNetwork]; ok {
			item.PaidUserCount = rev.PaidUserCount
			item.TotalRevenue = float64(rev.TotalRevenue) / 100.0
			if inst.InstallCount > 0 {
				item.LTV = float64(rev.TotalRevenue) / 100.0 / float64(inst.InstallCount)
				item.PaidRate = float64(rev.PaidUserCount) / float64(inst.InstallCount) * 100
			}
		}
		list = append(list, item)
	}

	return &adminApi.LtvByChannelRes{List: list}, nil
}

// LtvCohort 安装队列LTV
func (c *DashboardController) LtvCohort(ctx context.Context, req *adminApi.LtvCohortReq) (*adminApi.LtvCohortRes, error) {
	appIds := parseAppIds(req.AppIds)

	startDate, _ := time.Parse("2006-01-02", req.StartDate)
	endDate, _ := time.Parse("2006-01-02", req.EndDate)
	startTs := startDate.Unix()
	endTs := endDate.AddDate(0, 0, 1).Unix()

	// 确定日期格式
	var dateExpr string
	switch req.CohortType {
	case "week":
		dateExpr = "DATE_FORMAT(FROM_UNIXTIME(created_at), '%x-W%v')"
	case "month":
		dateExpr = "DATE_FORMAT(FROM_UNIXTIME(created_at), '%Y-%m')"
	default:
		dateExpr = "DATE_FORMAT(FROM_UNIXTIME(created_at), '%Y-%m-%d')"
	}

	// 获取各队列安装量
	type cohortInstall struct {
		CohortDate   string `json:"cohort_date"`
		InstallCount int    `json:"install_count"`
		MinCreatedAt int64  `json:"min_created_at"`
	}

	installM := dao.AttrDevice.Ctx(ctx).
		Fields(fmt.Sprintf("%s as cohort_date, COUNT(*) as install_count, MIN(created_at) as min_created_at", dateExpr)).
		Where("created_at >= ?", startTs).
		Where("created_at < ?", endTs)

	if len(appIds) > 0 {
		installM = installM.WhereIn("appid", appIds)
	}

	var installRows []cohortInstall
	err := installM.Group("cohort_date").Order("cohort_date ASC").Scan(&installRows)
	if err != nil {
		logger.Errorf("查询队列安装量失败: %v", err)
		return &adminApi.LtvCohortRes{}, nil
	}

	// 为每个队列计算 D7/D30/D60/D90 LTV
	now := time.Now().Unix()
	list := make([]*adminApi.LtvCohortItem, 0, len(installRows))

	for _, inst := range installRows {
		item := &adminApi.LtvCohortItem{
			CohortDate:   inst.CohortDate,
			InstallCount: inst.InstallCount,
		}

		cohortStart := inst.MinCreatedAt
		periods := []struct {
			days int
			ltv  *float64
		}{
			{7, &item.D7LTV},
			{30, &item.D30LTV},
			{60, &item.D60LTV},
			{90, &item.D90LTV},
		}

		for _, p := range periods {
			periodEnd := cohortStart + int64(p.days*86400)
			if now < periodEnd {
				// 还没到该周期，跳过
				continue
			}

			type revenueResult struct {
				TotalRevenue int64 `json:"total_revenue"`
			}
			var rev revenueResult

			revM := dao.AttrSubscriptionTransaction.Ctx(ctx).
				Fields("IFNULL(SUM(price),0) as total_revenue").
				Where("created_at >= ?", cohortStart).
				Where("created_at < ?", periodEnd).
				Where("transaction_type != 'REFUND'")

			if len(appIds) > 0 {
				revM = revM.WhereIn("appid", appIds)
			}

			err = revM.Scan(&rev)
			if err == nil && inst.InstallCount > 0 {
				*p.ltv = float64(rev.TotalRevenue) / 100.0 / float64(inst.InstallCount)
			}
		}

		list = append(list, item)
	}

	return &adminApi.LtvCohortRes{List: list}, nil
}
