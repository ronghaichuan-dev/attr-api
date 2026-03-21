package admin

import (
	"context"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"strings"

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

	var trendData []*adminApi.AppDailySubscriptionTrend
	for _, trans := range transactions {
		amount := float64(trans.Price) / 100.0

		trendData = append(trendData, &adminApi.AppDailySubscriptionTrend{
			AppId:  trans.Appid,
			Date:   gtime.NewFromTimeStamp(trans.PurchaseAt).Format("Y-m-d"),
			Count:  1,
			Amount: amount,
		})
	}

	for i := 0; i < len(trendData); i++ {
		for j := i + 1; j < len(trendData); j++ {
			if trendData[i].AppId > trendData[j].AppId || (trendData[i].AppId == trendData[j].AppId && trendData[i].Date > trendData[j].Date) {
				trendData[i], trendData[j] = trendData[j], trendData[i]
			}
		}
	}

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
