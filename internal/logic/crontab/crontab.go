package crontab

import (
	"context"
	"fmt"
	"god-help-service/internal/consts"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type sCrontab struct {
}

func init() {
	service.RegisterCrontab(NewCrontab())
}

func NewCrontab() *sCrontab {
	return &sCrontab{}
}

// Test 测试定时任务
func (s *sCrontab) Test(ctx context.Context) {
	fmt.Println(1111)
}

// HandleAttributionTokens 处理归因Token的定时任务
func (s *sCrontab) HandleAttributionTokens(ctx context.Context) {
	logger.Info("开始执行处理归因Token的定时任务")

	// 获取attr_record表中is_handle_token为2的数据，每次获取10条
	var records []*entity.AttrInstall
	if err := dao.AttrInstall.Ctx(ctx).Fields("id,app_token").Where("is_handle_token = ?", consts.IsHandleTokenNo).Limit(10).Scan(&records); err != nil {
		logger.Errorf("获取待处理的归因记录失败:%s", err.Error())
		return
	}

	logger.Infof("获取到待处理的归因记录数量:%d", len(records))

	// 循环处理每条记录
	for _, record := range records {
		// 调用GetAttributionInfo函数获取广告ID等数据
		attributionInfo, marshal, err := service.AppleServer().GetAttributionInfo(ctx, record.AppToken, "", "", "")
		if err != nil {
			logger.Errorf("获取归因信息失败:%s recordId: %d ", err.Error(), record.Id)
			continue
		}

		// 更新attr_record表
		updateData := g.Map{
			"is_handle_token": 1, // 更新为已调用
		}

		// 如果归因成功，更新广告相关数据
		if attributionInfo.Attribution {
			updateData["campaign_id"] = strconv.Itoa(attributionInfo.CampaignId)
			updateData["adgroup_id"] = strconv.Itoa(attributionInfo.AdGroupId)
			updateData["ad_id"] = strconv.Itoa(attributionInfo.AdId)
			updateData["keyword_id"] = strconv.Itoa(attributionInfo.KeywordId)
			updateData["country"] = attributionInfo.CountryOrRegion
			updateData["token_response_text"] = marshal
		}

		// 执行更新
		if _, err := dao.AttrInstall.Ctx(ctx).Where("id = ?", record.Id).Data(updateData).Update(); err != nil {
			logger.Errorf("更新归因记录失败:%s recordId:%d", err.Error(), record.Id)
			continue
		}

		logger.Infof("更新归因记录成功 recordId:%d", record.Id)
	}

	logger.Info("处理归因Token的定时任务执行完成")
}

// AggregateDailyStats 每日聚合统计任务 — 聚合前一天数据到 attr_daily_stats
func (s *sCrontab) AggregateDailyStats(ctx context.Context) {
	logger.Info("开始执行每日聚合统计任务")

	yesterday := time.Now().AddDate(0, 0, -1)
	statDate := yesterday.Format("2006-01-02")
	dayStart := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location()).Unix()
	dayEnd := dayStart + 86400

	// 1. 聚合安装量：按 app_id, country, tracker_network, campaign_id 分组
	type installStat struct {
		AppId          string `json:"app_id"`
		Country        string `json:"country"`
		TrackerNetwork string `json:"tracker_network"`
		CampaignId     string `json:"campaign_id"`
		InstallCount   int    `json:"install_count"`
	}
	var installStats []installStat
	err := dao.AttrInstall.Ctx(ctx).
		Fields("app_id, country, IFNULL(tracker_network,'') as tracker_network, IFNULL(campaign_id,'') as campaign_id, COUNT(*) as install_count").
		Where("install_at >= ?", dayStart).
		Where("install_at < ?", dayEnd).
		Group("app_id, country, tracker_network, campaign_id").
		Scan(&installStats)
	if err != nil {
		logger.Errorf("聚合安装量失败: %v", err)
		return
	}

	// 2. 聚合交易数据：按 appid, country, tracker_network, campaign_id, transaction_type 分组
	type txStat struct {
		Appid           string `json:"appid"`
		Country         string `json:"country"`
		TrackerNetwork  string `json:"tracker_network"`
		CampaignId      string `json:"campaign_id"`
		TransactionType string `json:"transaction_type"`
		TxCount         int    `json:"tx_count"`
		TotalAmount     int64  `json:"total_amount"`
	}
	var txStats []txStat
	err = dao.AttrSubscriptionTransaction.Ctx(ctx).
		Fields("appid, IFNULL(country,'') as country, IFNULL(tracker_network,'') as tracker_network, IFNULL(campaign_id,'') as campaign_id, transaction_type, COUNT(*) as tx_count, IFNULL(SUM(price),0) as total_amount").
		Where("created_at >= ?", dayStart).
		Where("created_at < ?", dayEnd).
		Group("appid, country, tracker_network, campaign_id, transaction_type").
		Scan(&txStats)
	if err != nil {
		logger.Errorf("聚合交易数据失败: %v", err)
		return
	}

	// 3. 合并数据到 map，key = app_id|country|tracker_network|campaign_id
	type statsAgg struct {
		AppId          string
		Country        string
		TrackerNetwork string
		CampaignId     string
		InstallCount   int
		TrialCount     int
		SubscribeCount int
		RenewCount     int
		RefundCount    int
		Revenue        int64
		RefundAmount   int64
	}
	aggMap := make(map[string]*statsAgg)

	makeKey := func(appId, country, network, campaignId string) string {
		return appId + "|" + country + "|" + network + "|" + campaignId
	}

	for _, s := range installStats {
		key := makeKey(s.AppId, s.Country, s.TrackerNetwork, s.CampaignId)
		if _, ok := aggMap[key]; !ok {
			aggMap[key] = &statsAgg{
				AppId: s.AppId, Country: s.Country,
				TrackerNetwork: s.TrackerNetwork, CampaignId: s.CampaignId,
			}
		}
		aggMap[key].InstallCount = s.InstallCount
	}

	for _, t := range txStats {
		key := makeKey(t.Appid, t.Country, t.TrackerNetwork, t.CampaignId)
		if _, ok := aggMap[key]; !ok {
			aggMap[key] = &statsAgg{
				AppId: t.Appid, Country: t.Country,
				TrackerNetwork: t.TrackerNetwork, CampaignId: t.CampaignId,
			}
		}
		agg := aggMap[key]
		switch t.TransactionType {
		case "TRIAL":
			agg.TrialCount = t.TxCount
		case "RENEW":
			agg.RenewCount = t.TxCount
			agg.Revenue += t.TotalAmount
		case "REFUND":
			agg.RefundCount = t.TxCount
			agg.RefundAmount += t.TotalAmount
		default:
			// 首次订阅等其他类型计入订阅量和收入
			agg.SubscribeCount += t.TxCount
			agg.Revenue += t.TotalAmount
		}
	}

	// 4. Upsert 到 attr_daily_stats
	now := time.Now().Unix()
	for _, agg := range aggMap {
		netRevenue := agg.Revenue - agg.RefundAmount

		var installToTrialRate, trialToPaidRate float64
		if agg.InstallCount > 0 {
			installToTrialRate = float64(agg.TrialCount) / float64(agg.InstallCount) * 100
		}
		if agg.TrialCount > 0 {
			trialToPaidRate = float64(agg.SubscribeCount) / float64(agg.TrialCount) * 100
		}

		// 尝试更新
		result, err := dao.AttrDailyStats.Ctx(ctx).
			Where("stat_date", statDate).
			Where("app_id", agg.AppId).
			Where("country", agg.Country).
			Where("tracker_network", agg.TrackerNetwork).
			Where("campaign_id", agg.CampaignId).
			Data(g.Map{
				"install_count":         agg.InstallCount,
				"trial_count":           agg.TrialCount,
				"subscribe_count":       agg.SubscribeCount,
				"renew_count":           agg.RenewCount,
				"refund_count":          agg.RefundCount,
				"revenue":               agg.Revenue,
				"refund_amount":         agg.RefundAmount,
				"net_revenue":           netRevenue,
				"install_to_trial_rate": installToTrialRate,
				"trial_to_paid_rate":    trialToPaidRate,
				"updated_at":            now,
			}).Update()
		if err != nil {
			logger.Errorf("更新每日统计失败: %v", err)
			continue
		}

		affected, _ := result.RowsAffected()
		if affected == 0 {
			// 不存在则插入
			_, err = dao.AttrDailyStats.Ctx(ctx).Data(g.Map{
				"stat_date":             statDate,
				"app_id":               agg.AppId,
				"country":              agg.Country,
				"tracker_network":      agg.TrackerNetwork,
				"campaign_id":          agg.CampaignId,
				"install_count":        agg.InstallCount,
				"trial_count":          agg.TrialCount,
				"subscribe_count":      agg.SubscribeCount,
				"renew_count":          agg.RenewCount,
				"refund_count":         agg.RefundCount,
				"revenue":              agg.Revenue,
				"refund_amount":        agg.RefundAmount,
				"net_revenue":          netRevenue,
				"install_to_trial_rate": installToTrialRate,
				"trial_to_paid_rate":   trialToPaidRate,
				"created_at":           now,
				"updated_at":           now,
			}).Insert()
			if err != nil {
				logger.Errorf("插入每日统计失败: %v", err)
			}
		}
	}

	logger.Infof("每日聚合统计任务执行完成，统计日期: %s，共处理 %d 组数据", statDate, len(aggMap))
}
