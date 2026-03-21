package attr

import (
	"context"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetEventByCode 根据事件编码获取事件详情
func (s *sAttr) GetEventByCode(ctx context.Context, eventCode string) (*entity.AttrEvent, error) {
	var event *entity.AttrEvent
	err := dao.AttrEvent.Ctx(ctx).Fields("event_code").Where("event_code", eventCode).Scan(&event)
	return event, err
}

// GetEventDropdownList 获取事件下拉选项列表
func (s *sAttr) GetEventDropdownList(ctx context.Context) (*adminApi.EventDropdownRes, error) {
	var entities []*entity.AttrEvent
	err := dao.AttrEvent.Ctx(ctx).Where("status", 1).OrderAsc("id").Scan(&entities)
	if err != nil {
		return nil, err
	}
	// 转换为响应格式
	res := &adminApi.EventDropdownRes{
		List: make([]*adminApi.EventDropdownItem, len(entities)),
	}

	for i, entity := range entities {
		res.List[i] = &adminApi.EventDropdownItem{
			Id:        int(entity.Id),
			EventName: entity.EventName,
		}
	}

	return res, nil
}

// GetEventList 获取事件列表
func (s *sAttr) GetEventList(ctx context.Context, req *adminApi.EventListReq) (*adminApi.EventListRes, error) {
	var (
		res      = &adminApi.EventListRes{}
		entities []*entity.AttrEvent
	)

	// 构建查询条件
	m := dao.AttrEvent.Ctx(ctx)

	if req.EventName != "" {
		m = m.WhereLike("event_name", "%"+req.EventName+"%")
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

	// 转换为响应格式
	res.List = make([]*adminApi.EventListItem, len(entities))
	for i, entity := range entities {
		createdAt := ""
		updatedAt := ""
		if entity.CreatedAt != nil {
			createdAt = entity.CreatedAt.Format("2006-01-02 15:04:05")
		}
		if entity.UpdatedAt != nil {
			updatedAt = entity.UpdatedAt.Format("2006-01-02 15:04:05")
		}
		res.List[i] = &adminApi.EventListItem{
			Id:        int(entity.Id),
			EventName: entity.EventName,
			Status:    entity.Status,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	}

	return res, nil
}

// GetEventDetailById 根据ID获取事件详情
func (s *sAttr) GetEventDetailById(ctx context.Context, eventId int64) (*adminApi.EventDetailItem, error) {
	var entity *entity.AttrEvent
	err := dao.AttrEvent.Ctx(ctx).Where("id", eventId).Scan(&entity)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, gerror.New("事件不存在")
	}

	// 转换为响应格式
	createdAt := ""
	updatedAt := ""
	if entity.CreatedAt != nil {
		createdAt = entity.CreatedAt.Format("2006-01-02 15:04:05")
	}
	if entity.UpdatedAt != nil {
		updatedAt = entity.UpdatedAt.Format("2006-01-02 15:04:05")
	}
	return &adminApi.EventDetailItem{
		Id:        int(entity.Id),
		EventName: entity.EventName,
		Status:    entity.Status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

// CreateEvent 创建事件
func (s *sAttr) CreateEvent(ctx context.Context, req *adminApi.EventCreateReq) (int64, error) {
	// 构建事件数据
	eventDO := &entity.AttrEvent{
		EventName: req.EventName,
		Status:    req.Status,
	}

	// 执行插入
	lastInsertId, err := dao.AttrEvent.Ctx(ctx).Data(eventDO).InsertAndGetId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

// UpdateEvent 更新事件
func (s *sAttr) UpdateEvent(ctx context.Context, req *adminApi.EventUpdateReq) error {
	// 构建更新数据
	data := g.Map{}
	if req.EventName != "" {
		data["event_name"] = req.EventName
	}
	if req.Status > 0 {
		data["status"] = req.Status
	}
	data["updated_at"] = gdb.Raw("NOW()")

	// 执行更新
	result, err := dao.AttrEvent.Ctx(ctx).Data(data).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return gerror.New("事件不存在")
	}

	return nil
}

// DeleteEvent 删除事件（软删除）
func (s *sAttr) DeleteEvent(ctx context.Context, eventId int64) error {
	// 执行软删除
	result, err := dao.AttrEvent.Ctx(ctx).Where("id", eventId).Where("deleted_at IS NULL").Update("deleted_at", gdb.Raw("NOW()"))
	if err != nil {
		return err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return gerror.New("事件不存在")
	}

	return nil
}
