package admin

import (
	"context"
	"god-help-service/internal/dao"
	"god-help-service/internal/util/logger"

	apiadmin "god-help-service/api/v1/admin"
	"god-help-service/internal/model/entity"
)

// NoticeController 通知管理控制器
type NoticeController struct{}

// List 获取通知列表
func (c *NoticeController) List(ctx context.Context, req *apiadmin.NoticeListReq) (res *apiadmin.NoticeListRes, err error) {
	logger.Debugf("获取通知列表请求参数:%v", req)

	// 构建查询条件
	m := dao.AttrAppleNotificationEvent.Ctx(ctx)
	if req.Uuid != "" {
		m = m.Where("uuid = ?", req.Uuid)
	}
	if req.NoticeType != "" {
		m = m.Where("notice_type = ?", req.NoticeType)
	}
	if req.RenewalStatus != "" {
		m = m.Where("renewal_status = ?", req.RenewalStatus)
	}

	// 查询总记录数
	var total int
	var countErr error
	total, countErr = m.Count()
	if countErr != nil {
		logger.Errorf("查询通知总数失败:%d", countErr)
		return
	}

	// 查询列表数据
	var list []*entity.AttrNotifications
	if total > 0 {
		page := req.Page
		if page <= 0 {
			page = 1
		}
		pageSize := req.PageSize
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize

		if err = m.
			OrderDesc("created_at").
			Limit(pageSize, offset).
			Scan(&list); err != nil {
			logger.Errorf("查询通知列表失败:%v", err)
			return
		}
	}

	// 构建响应
	res = &apiadmin.NoticeListRes{
		Total: int64(total),
		List:  list,
	}

	return
}
