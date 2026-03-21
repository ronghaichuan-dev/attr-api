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
func (s *sAttr) HandleAttribution(ctx context.Context, attr *app.Attribution) error {
	// 遍历所有归因记录
	logger.Infof("处理归因记录 trackerUid:%s  rsid: %s appid: %s", attr.TrackerUid, attr.Rsid, attr.AppId)

	// 3. 根据获取的信息去调用苹果广告服务归因接口
	updateFiledData := g.Map{}
	if attr.AdServicesToken != "" {
		// 1. 根据app_id获取应用信息（优先从Redis获取）
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
		logger.Infof("获取到公司ID: %d", companyId)

		// 2. 根据account_type、app_id和company_id在system_account表中获取account_info信息
		// 苹果广告服务的account_type为4
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

		logger.Infof("获取到账户信息: %s", systemAccount.AccountInfo)
		var accountInfo app.AccountInfo
		err = json.Unmarshal([]byte(systemAccount.AccountInfo), &accountInfo)
		if err != nil {
			logger.Errorf("解析账户信息失败: %v", err)
			return err
		}

		proxyURL := accountInfo.Socks5
		logger.Infof("获取到代理地址: %s", proxyURL)
		logger.Infof("归因来源为adServices，adServicesToken:%s", attr.AdServicesToken)

		attributionInfo, marshal, err := service.AppleServer().GetAttributionInfo(ctx, attr.AdServicesToken, proxyURL, accountInfo.Username, accountInfo.Password)
		if err != nil {
			logger.Errorf("调用苹果归因接口失败 : %s", err)
			attr.IsHandleToken = consts.IsHandleTokenNo
			updateFiledData["is_handle_token"] = consts.IsHandleTokenNo
		} else {
			updateFiledData["is_handle_token"] = consts.IsHandleTokenYes
			updateFiledData["token_response_text"] = marshal
		}

		// 重新赋值adId等字段
		if attributionInfo != nil {
			updateFiledData["ad_id"] = gconv.String(attributionInfo.AdId)
			// 可以根据需要赋值其他字段
			updateFiledData["campaign_id"] = gconv.String(attributionInfo.CampaignId)
			updateFiledData["adgroup_id"] = gconv.String(attributionInfo.AdGroupId)
			updateFiledData["keyword_id"] = gconv.String(attributionInfo.KeywordId)
		}
	} else {
		updateFiledData["is_handle_token"] = consts.IsHandleTokenYes
	}

	// 4. 将结果保存在数据库中
	_, err := dao.AttrInstall.Ctx(ctx).Where("attr_uuid", attr.AttrUuid).Data(updateFiledData).Update()
	if err != nil {
		return err
	}
	return nil
}
