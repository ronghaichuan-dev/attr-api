package attr

import (
	"context"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// GetAppEventLogList 获取事件日志列表
func (s *sAttr) GetAppEventLogList(ctx context.Context, req *adminApi.AppEventLogListReq) (*adminApi.AppEventLogListRes, error) {
	var (
		res      = &adminApi.AppEventLogListRes{}
		entities []*entity.AppEventLogCustom
	)

	// 构建查询条件
	m := dao.AttrEventLog.Ctx(ctx)

	if req.Appid != "" {
		m = m.Where("appid", req.Appid)
	}
	if req.UserId != "" {
		m = m.Where("uuid", req.UserId)
	}

	// 获取总记录数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = int64(total)

	// 分页查询
	offset := (req.Page - 1) * req.Size
	err = m.Offset(offset).Limit(req.Size).OrderDesc("created_at").Scan(&entities)
	if err != nil {
		return nil, err
	}

	res.List = entities

	return res, nil
}

// GetAppEventLogById 根据ID获取事件日志详情
func (s *sAttr) GetAppEventLogById(ctx context.Context, logId int64) (*entity.AppEventLogCustom, error) {
	var entity *entity.AppEventLogCustom
	err := dao.AttrEventLog.Ctx(ctx).Where("id", logId).Scan(&entity)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, gerror.New("事件日志不存在")
	}

	return entity, nil
}
