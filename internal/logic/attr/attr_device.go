package attr

import (
	"context"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sAttr) UpdateAttrDeviceField(ctx context.Context, tx gdb.TX, uuid string, updateData g.Map) error {
	mod := dao.AttrDevice.Ctx(ctx)
	if tx != nil {
		mod = mod.TX(tx)
	}
	_, err := mod.Where(dao.AttrDevice.Columns().AttrSubscriptionId, uuid).Data(updateData).Update()
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

func (s *sAttr) CreateAttrDeviceOrUpdate(ctx context.Context, appid, uuid, country string) error {
	count, err := dao.AttrDevice.Ctx(ctx).Where("appid", appid).Where("uuid", uuid).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		data := entity.AttrDevice{
			Rsid:               uuid,
			Appid:              appid,
			AttrSubscriptionId: 0,
			Country:            country,
			CreatedAt:          time.Now().Unix(),
		}
		_, err = dao.AttrDevice.Ctx(ctx).Data(data).Insert()
		if err != nil {
			return err
		}
	} else {

	}
	return nil
}
