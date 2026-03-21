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
