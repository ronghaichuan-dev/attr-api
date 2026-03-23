package attr

import (
	"context"
	"god-help-service/api/v1/app"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sAttr) UpdateAttrDeviceField(ctx context.Context, tx gdb.TX, originalTransactionId string, updateData g.Map) error {
	mod := dao.AttrDevice.Ctx(ctx)
	if tx != nil {
		mod = mod.TX(tx)
	}
	_, err := mod.Where(dao.AttrDevice.Columns().AttrSubscriptionId, originalTransactionId).Data(updateData).Update()
	if err != nil {
		return err
	}
	return nil
}

func (s *sAttr) GetAttrDevice(ctx context.Context, uuid []string) (map[string]struct{}, error) {
	var list []*entity.AttrDevice
	err := dao.AttrDevice.Ctx(ctx).Fields(dao.AttrDevice.Columns().Rsid, dao.AttrDevice.Columns().Id).WhereIn("uuid", uuid).Scan(&list)
	if err != nil {
		return nil, err
	}
	deviceMap := make(map[string]struct{})
	for _, device := range list {
		deviceMap[device.Rsid] = struct{}{}
	}
	return deviceMap, nil
}

// CreateAttrDeviceOrUpdate 创建或更新设备归因记录
func (s *sAttr) CreateAttrDeviceOrUpdate(ctx context.Context, appid, uuid, country string) error {
	count, err := dao.AttrDevice.Ctx(ctx).Where("appid", appid).Where("rsid", uuid).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		data := entity.AttrDevice{
			Rsid:           uuid,
			Appid:          appid,
			Country:        country,
			IsFirstInstall: 1,
			CreatedAt:      time.Now().Unix(),
			LastInstallAt:  time.Now().Unix(),
		}
		_, err = dao.AttrDevice.Ctx(ctx).Data(data).Insert()
		if err != nil {
			return err
		}
	} else {
		// 设备已存在，更新安装时间
		updateData := g.Map{
			"last_install_at":  time.Now().Unix(),
			"is_first_install": 0,
		}
		if country != "" {
			updateData["country"] = country
		}
		_, err = dao.AttrDevice.Ctx(ctx).Where("appid", appid).Where("rsid", uuid).Data(updateData).Update()
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateAttrDeviceOrUpdateWithAttribution 创建或更新设备归因记录（含归因信息）
func (s *sAttr) CreateAttrDeviceOrUpdateWithAttribution(ctx context.Context, attr *app.Attribution, installId int64) error {
	count, err := dao.AttrDevice.Ctx(ctx).Where("appid", attr.AppId).Where("rsid", attr.Rsid).Count()
	if err != nil {
		return err
	}

	updateData := g.Map{
		"last_install_at":  time.Now().Unix(),
		"tracker_network":  attr.TrackerNetwork,
		"campaign_id":      attr.TrackerCampaignId,
		"adgroup_id":       attr.TrackerAdgroupId,
		"ad_id":            attr.TrackerAdId,
		"keyword_id":       attr.TrackerKeywordId,
		"channel":          attr.TrackerChannel,
		"attr_install_id":  installId,
	}

	if count == 0 {
		data := entity.AttrDevice{
			Rsid:           attr.Rsid,
			Appid:          attr.AppId,
			Country:        attr.Country,
			TrackerNetwork: attr.TrackerNetwork,
			CampaignId:     attr.TrackerCampaignId,
			AdgroupId:      attr.TrackerAdgroupId,
			AdId:           attr.TrackerAdId,
			KeywordId:      attr.TrackerKeywordId,
			Channel:        attr.TrackerChannel,
			AttrInstallId:  installId,
			IsFirstInstall: 1,
			CreatedAt:      time.Now().Unix(),
			LastInstallAt:  time.Now().Unix(),
		}
		_, err = dao.AttrDevice.Ctx(ctx).Data(data).Insert()
		return err
	}

	if attr.Country != "" {
		updateData["country"] = attr.Country
	}
	_, err = dao.AttrDevice.Ctx(ctx).Where("appid", attr.AppId).Where("rsid", attr.Rsid).Data(updateData).Update()
	return err
}

// UpdateAttrDeviceSubscription 更新设备订阅相关字段
func (s *sAttr) UpdateAttrDeviceSubscription(ctx context.Context, rsid, appid string, updateData g.Map) error {
	if rsid == "" || appid == "" {
		return nil
	}
	_, err := dao.AttrDevice.Ctx(ctx).Where("rsid", rsid).Where("appid", appid).Data(updateData).Update()
	return err
}
