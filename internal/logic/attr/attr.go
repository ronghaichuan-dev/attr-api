package attr

import (
	"context"
	"encoding/json"
	"god-help-service/api/v1/app"
	"god-help-service/internal/consts"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAttr struct {
}

func init() {
	service.RegisterAttr(NewAttr())
}

func NewAttr() *sAttr {
	return &sAttr{}
}

func (s *sAttr) Test() {

}

// HandleAttribution 处理归因信息
// 归因匹配优先级：设备ID匹配 > 第三方Tracker匹配 > Apple Ad Services Token匹配 > 概率匹配
func (s *sAttr) HandleAttribution(ctx context.Context, attr *app.Attribution) error {
	logger.Infof("处理归因记录 trackerUid:%s rsid: %s appid: %s", attr.TrackerUid, attr.Rsid, attr.AppId)

	updateFieldData := g.Map{}
	matchType := ""
	matchConfidence := ""

	// 优先级 1: 设备ID匹配（IDFA/GAID）— 查 attr_click 表精确匹配
	if matched, clickRecord := s.matchByDeviceId(ctx, attr); matched && clickRecord != nil {
		matchType = consts.MatchTypeDeviceId
		matchConfidence = consts.MatchConfidenceHigh
		updateFieldData["click_id"] = clickRecord.Id
		updateFieldData["click_to_install"] = int64(attr.InstallAt) - clickRecord.ClickAt
		// 从点击记录填充归因渠道信息
		updateFieldData["network"] = clickRecord.Network
		updateFieldData["channel"] = clickRecord.Network
		updateFieldData["campaign_id"] = clickRecord.CampaignId
		updateFieldData["adgroup_id"] = clickRecord.AdgroupId
		updateFieldData["ad_id"] = clickRecord.AdId
		logger.Infof("设备ID匹配成功，click_id:%d network:%s", clickRecord.Id, clickRecord.Network)
	} else if attr.Tracker != "" && attr.TrackerNetwork != "" {
		// 优先级 2: 第三方 Tracker 匹配（Adjust/Branch/AppsFlyer 直接传入归因数据）
		matchType = consts.MatchTypeTracker
		matchConfidence = consts.MatchConfidenceHigh
		updateFieldData["network"] = attr.TrackerNetwork
		updateFieldData["channel"] = attr.TrackerChannel
		updateFieldData["campaign_id"] = attr.TrackerCampaignId
		updateFieldData["adgroup_id"] = attr.TrackerAdgroupId
		updateFieldData["ad_id"] = attr.TrackerAdId
		logger.Infof("第三方Tracker匹配成功，tracker:%s network:%s", attr.Tracker, attr.TrackerNetwork)
	} else if attr.AdServicesToken != "" {
		// 优先级 3: Apple Ad Services Token 匹配
		matchType = consts.MatchTypeAdServices
		matchConfidence = consts.MatchConfidenceMedium

		appInfo, err := service.System().GetAppByAppId(ctx, attr.AppId)
		if err != nil {
			logger.Errorf("获取应用信息失败: %v", err)
			return err
		}
		if appInfo == nil {
			logger.Errorf("应用不存在: %s", attr.AppId)
			return gerror.New("应用不存在")
		}

		companyId := appInfo.CompanyId
		var systemAccount *entity.SystemAccount
		err = dao.SystemAccount.Ctx(ctx).WhereLike("app_id", "%"+attr.AppId+"%").Where("account_type", 4).Where("company_id", companyId).Scan(&systemAccount)
		if err != nil {
			logger.Errorf("获取账户信息失败: %v", err)
			return err
		}
		if systemAccount == nil {
			logger.Errorf("苹果广告服务账户不存在: appid=%s, company_id=%d", attr.AppId, companyId)
			return gerror.New("苹果广告服务账户不存在")
		}

		var accountInfo app.AccountInfo
		err = json.Unmarshal([]byte(systemAccount.AccountInfo), &accountInfo)
		if err != nil {
			logger.Errorf("解析账户信息失败: %v", err)
			return err
		}

		proxyURL := accountInfo.Socks5
		logger.Infof("归因来源为adServices，adServicesToken:%s", attr.AdServicesToken)

		attributionInfo, marshal, err := service.AppleServer().GetAttributionInfo(ctx, attr.AdServicesToken, proxyURL, accountInfo.Username, accountInfo.Password)
		if err != nil {
			logger.Errorf("调用苹果归因接口失败: %s", err)
			updateFieldData["is_handle_token"] = consts.IsHandleTokenNo
		} else {
			updateFieldData["is_handle_token"] = consts.IsHandleTokenYes
			updateFieldData["token_response_text"] = marshal
		}

		if attributionInfo != nil {
			updateFieldData["ad_id"] = gconv.String(attributionInfo.AdId)
			updateFieldData["campaign_id"] = gconv.String(attributionInfo.CampaignId)
			updateFieldData["adgroup_id"] = gconv.String(attributionInfo.AdGroupId)
			updateFieldData["keyword_id"] = gconv.String(attributionInfo.KeywordId)
		}
		logger.Infof("Apple Ad Services匹配完成")
	} else {
		// 无匹配来源，标记为已处理
		updateFieldData["is_handle_token"] = consts.IsHandleTokenYes
	}

	// 记录匹配方式和置信度
	if matchType != "" {
		updateFieldData["match_type"] = matchType
		updateFieldData["match_confidence"] = matchConfidence
	}

	// 保存归因结果到数据库
	_, err := dao.AttrInstall.Ctx(ctx).Where("attr_uuid", attr.AttrUuid).Data(updateFieldData).Update()
	if err != nil {
		return err
	}

	// 异步触发回传
	if matchType != "" {
		go s.triggerPostback(ctx, attr, matchType)
	}

	return nil
}

// matchByDeviceId 通过设备ID匹配点击记录
func (s *sAttr) matchByDeviceId(ctx context.Context, attr *app.Attribution) (bool, *entity.AttrClick) {
	now := time.Now().Unix()

	// 尝试 IDFA 匹配
	if attr.Idfa != "" {
		var click *entity.AttrClick
		err := dao.AttrClick.Ctx(ctx).
			Where("idfa", attr.Idfa).
			Where("app_id", attr.AppId).
			Where("click_at >", now-consts.ClickAttributionWindow).
			OrderDesc("click_at").
			Limit(1).
			Scan(&click)
		if err == nil && click != nil {
			return true, click
		}
	}

	// 尝试 GAID 匹配
	if attr.GpsAdid != "" {
		var click *entity.AttrClick
		err := dao.AttrClick.Ctx(ctx).
			Where("gps_adid", attr.GpsAdid).
			Where("app_id", attr.AppId).
			Where("click_at >", now-consts.ClickAttributionWindow).
			OrderDesc("click_at").
			Limit(1).
			Scan(&click)
		if err == nil && click != nil {
			return true, click
		}
	}

	return false, nil
}

// triggerPostback 触发归因回传（异步）
func (s *sAttr) triggerPostback(ctx context.Context, attr *app.Attribution, matchType string) {
	network := attr.TrackerNetwork
	if network == "" {
		network = attr.Tracker
	}
	if network == "" {
		return
	}

	err := s.SendPostback(ctx, &entity.AttrPostback{
		AppId:        attr.AppId,
		PostbackType: consts.PostbackTypeInstall,
		Network:      network,
		CreatedAt:    time.Now().Unix(),
	})
	if err != nil {
		logger.Errorf("触发归因回传失败: %v", err)
	}
}
