package admin

import (
	"context"
	"fmt"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"

	adminApi "god-help-service/api/v1/admin"

	"github.com/gogf/gf/v2/errors/gerror"
)

// EventController 事件控制器
type EventController struct{}

// LogList 获取事件日志列表
func (c *EventController) LogList(ctx context.Context, req *adminApi.EventLogListReq) (*adminApi.EventLogListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)
	// 转换请求参数类型
	appEventLogReq := &adminApi.AppEventLogListReq{
		Appid:  req.Appid,
		UserId: req.UserId,
		Page:   req.Page,
		Size:   req.Size,
	}

	res, err := service.Attr().GetAppEventLogList(ctx, appEventLogReq)
	if err != nil {
		logger.Errorf("获取事件日志列表失败: %v", err)
		return nil, err
	}

	// 转换响应数据类型
	eventLogRes := &adminApi.EventLogListRes{
		Total: res.Total,
		List:  make([]*adminApi.EventLogListItem, len(res.List)),
	}
	for i, item := range res.List {
		var createdAt string
		if item.CreatedAt != nil {
			createdAt = item.CreatedAt.String()
		}
		eventLogRes.List[i] = &adminApi.EventLogListItem{
			Id:           item.Id,
			Appid:        item.Appid,
			EventCode:    item.EventCode,
			UserId:       item.UserId,
			ResponseText: item.ResponseText,
			CreatedAt:    createdAt,
		}
	}

	// 确保返回的数据结构完整，避免返回 null
	if eventLogRes == nil {
		eventLogRes = &adminApi.EventLogListRes{
			Total: 0,
			List:  []*adminApi.EventLogListItem{},
		}
	}
	if eventLogRes.List == nil {
		eventLogRes.List = []*adminApi.EventLogListItem{}
	}

	logger.Debugf("事件日志列表查询结果: %v", eventLogRes)
	return eventLogRes, nil
}

// Dropdown 获取事件下拉选项列表
func (c *EventController) Dropdown(ctx context.Context, req *adminApi.EventDropdownReq) (*adminApi.EventDropdownRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行查询
	res, err := service.Attr().GetEventDropdownList(ctx)
	if err != nil {
		logger.Errorf("获取事件下拉选项列表失败: %v", err)
		return nil, err
	}

	// 确保返回的数据结构完整，避免返回 null
	if res == nil {
		res = &adminApi.EventDropdownRes{
			List: []*adminApi.EventDropdownItem{},
		}
	}
	if res.List == nil {
		res.List = []*adminApi.EventDropdownItem{}
	}

	logger.Debugf("事件下拉选项查询结果: %v", res)
	return res, nil
}

// List 获取事件列表
func (c *EventController) List(ctx context.Context, req *adminApi.EventListReq) (*adminApi.EventListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)
	fmt.Println(111111)
	// 执行查询
	res, err := service.Attr().GetEventList(ctx, req)
	if err != nil {
		logger.Errorf("获取事件列表失败: %v", err)
		return nil, err
	}

	// 确保返回的数据结构完整，避免返回 null
	if res == nil {
		res = &adminApi.EventListRes{
			Total: 0,
			List:  []*adminApi.EventListItem{},
		}
	}
	if res.List == nil {
		res.List = []*adminApi.EventListItem{}
	}

	logger.Debugf("事件列表查询结果: %v", res)
	return res, nil
}

// Detail 获取事件详情
func (c *EventController) Detail(ctx context.Context, req *adminApi.EventDetailReq) (*adminApi.EventDetailRes, error) {
	// 执行查询
	event, err := service.Attr().GetEventDetailById(ctx, int64(req.Id))
	if err != nil {
		logger.Errorf("获取事件详情失败: %v", err)
		return nil, err
	}

	if event == nil {
		return nil, gerror.New("事件不存在")
	}

	return &adminApi.EventDetailRes{Event: event}, nil
}

// Create 创建事件
func (c *EventController) Create(ctx context.Context, req *adminApi.EventCreateReq) (*adminApi.EventCreateRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行创建
	id, err := service.Attr().CreateEvent(ctx, req)
	if err != nil {
		logger.Errorf("创建事件失败: %v", err)
		return nil, err
	}

	logger.Debugf("事件创建成功，ID: %v", id)
	res := &adminApi.EventCreateRes{Id: int(id)}
	logger.Debugf("准备返回的响应对象: %v", res)
	logger.Debugf("响应对象类型: %v", fmt.Sprintf("%T", res))

	return res, nil
}

// Update 更新事件
func (c *EventController) Update(ctx context.Context, req *adminApi.EventUpdateReq) (*adminApi.EventUpdateRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行更新
	err := service.Attr().UpdateEvent(ctx, req)
	if err != nil {
		logger.Errorf("更新事件失败: %v", err)
		return nil, err
	}

	logger.Debugf("事件更新成功，ID: %v", req.Id)
	return &adminApi.EventUpdateRes{Id: req.Id}, nil
}

// Delete 删除事件
func (c *EventController) Delete(ctx context.Context, req *adminApi.EventDeleteReq) (*adminApi.EventDeleteRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行删除
	err := service.Attr().DeleteEvent(ctx, int64(req.Id))
	if err != nil {
		logger.Errorf("删除事件失败: %v", err)
		return nil, err
	}

	logger.Debugf("事件删除成功，ID: %v", req.Id)
	return &adminApi.EventDeleteRes{Id: req.Id}, nil
}
