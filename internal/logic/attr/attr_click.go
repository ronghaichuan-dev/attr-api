package attr

import (
	"context"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"time"

	"github.com/google/uuid"
)

// RecordClick 记录广告点击/展示
func (s *sAttr) RecordClick(ctx context.Context, record *service.ClickRecord) error {
	clickUuid := uuid.New().String()
	now := time.Now().Unix()

	data := &entity.AttrClick{
		ClickUuid:    clickUuid,
		AppId:        record.AppId,
		ClickType:    record.ClickType,
		Idfa:         record.Idfa,
		Idfv:         record.Idfv,
		GpsAdid:      record.GpsAdid,
		Ip:           record.Ip,
		UserAgent:    record.UserAgent,
		Network:      record.Network,
		CampaignId:   record.CampaignId,
		CampaignName: record.CampaignName,
		AdgroupId:    record.AdgroupId,
		AdId:         record.AdId,
		KeywordId:    record.KeywordId,
		Creative:     record.Creative,
		ClickUrl:     record.ClickUrl,
		RedirectUrl:  record.RedirectUrl,
		ClickAt:      now,
		CreatedAt:    now,
	}

	_, err := dao.AttrClick.Ctx(ctx).Data(data).Insert()
	return err
}
