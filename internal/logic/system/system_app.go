package system

import (
	"context"
	"encoding/json"
	"fmt"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/util"
	"god-help-service/internal/util/logger"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// GetAppById 根据ID获取应用详情
func (s *sSystem) GetAppById(ctx context.Context, appid string) (*entity.SystemApps, error) {
	// 使用统一的获取应用信息方法（优先从Redis获取）
	return s.GetAppByAppId(ctx, appid)
}

// GetAppList 获取应用列表
func (s *sSystem) GetAppList(ctx context.Context, req *adminApi.AppListReq) (*adminApi.AppListRes, error) {
	var (
		res      = &adminApi.AppListRes{}
		entities []*entity.SystemApps
	)

	// 构建查询条件
	m := dao.SystemApps.Ctx(ctx)

	if req.AppName != "" {
		m = m.WhereLike("app_name", "%"+req.AppName+"%")
	}
	if req.AppId != "" {
		m = m.Where("appid", req.AppId)
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

	// 收集所有用户ID用于查询用户名称
	userIdSet := make(map[int]bool)
	for _, app := range entities {
		if app.Creator > 0 {
			userIdSet[app.Creator] = true
		}
		if app.Modifier > 0 {
			userIdSet[app.Modifier] = true
		}
	}

	// 批量查询用户信息
	userMap := make(map[int]*entity.SystemUsers)
	if len(userIdSet) > 0 {
		userIds := make([]int, 0, len(userIdSet))
		for id := range userIdSet {
			userIds = append(userIds, id)
		}
		userMap, err = s.GetUsersByIds(ctx, userIds)
		if err != nil {
			logger.Warnf("批量查询用户信息失败: %v", err)
		}
	}

	// 转换为响应格式
	res.List = make([]*adminApi.AppListItem, len(entities))
	for i, app := range entities {
		creatorName := ""
		modifierName := ""

		if creator, ok := userMap[app.Creator]; ok {
			creatorName = creator.Username
		}
		if modifier, ok := userMap[app.Modifier]; ok {
			modifierName = modifier.Username
		}

		res.List[i] = &adminApi.AppListItem{
			AppId:           app.Appid,
			AppName:         app.AppName,
			Icon:            app.Icon,
			Creator:         app.Creator,
			CreatorName:     creatorName,
			Modifier:        app.Modifier,
			ModifierName:    modifierName,
			SubscriptionFee: app.SubscriptionFee,
			CreatedAt:       app.CreatedAt,
			UpdatedAt:       app.UpdatedAt,
			DeletedAt:       app.DeletedAt,
		}
	}

	return res, nil
}

// View 根据ID获取应用详情
func (s *sSystem) View(ctx context.Context, appid string) (*adminApi.AppDetailItem, []int, error) {
	// 使用统一的获取应用信息方法（优先从Redis获取）
	app, err := s.GetAppByAppId(ctx, appid)
	if err != nil {
		return nil, nil, err
	}
	// 转换为响应格式
	detailItem := &adminApi.AppDetailItem{
		AppId:           app.Appid,
		AppName:         app.AppName,
		Icon:            app.Icon,
		Creator:         app.Creator,
		Modifier:        app.Modifier,
		SubscriptionFee: app.SubscriptionFee,
		CreatedAt:       app.CreatedAt,
		UpdatedAt:       app.UpdatedAt,
		DeletedAt:       app.DeletedAt,
	}

	// 这里可以添加获取应用事件列表的逻辑
	// 暂时返回空数组
	events := []int{}

	return detailItem, events, nil
}

// CreateApp 创建应用
func (s *sSystem) CreateApp(ctx context.Context, req *adminApi.AppCreateReq, userId int) (*adminApi.AppCreateRes, error) {
	var (
		res = &adminApi.AppCreateRes{}
	)
	// 生成UUID作为AppId
	appId := uuid.New().String()
	// 使用事务确保应用创建和事件绑定的原子性
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 构建应用数据
		appDO := &entity.SystemApps{
			Appid:           appId,
			AppName:         req.AppName,
			Icon:            req.Icon,
			Creator:         userId,
			SubscriptionFee: req.SubscriptionFee,
		}

		// 执行插入（使用事务）
		_, err := tx.Model("system_apps").Data(appDO).Insert()
		if err != nil {
			return err
		}

		return nil // 提交事务
	})

	if err != nil {
		return nil, err
	}

	res.AppId = appId
	return res, nil
}

// UpdateApp 更新应用信息
func (s *sSystem) UpdateApp(ctx context.Context, req *adminApi.AppUpdateReq, userId int) (*adminApi.AppUpdateRes, error) {
	var (
		res = &adminApi.AppUpdateRes{}
	)

	// 构建更新数据
	data := g.Map{}
	if req.AppName != "" {
		data["app_name"] = req.AppName
	}
	if req.Icon != "" {
		data["icon"] = req.Icon
	}
	// 使用用户ID作为修改人
	data["modifier"] = userId
	// 检查SubscriptionFee是否被设置（非0值也应该被更新）
	if req.SubscriptionFee >= 0 {
		data["subscription_fee"] = req.SubscriptionFee
	}
	data["updated_at"] = gtime.Now()

	// 执行更新
	result, err := dao.SystemApps.Ctx(ctx).Data(data).Where("appid", req.AppId).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("应用不存在")
	}

	res.AppId = req.AppId
	return res, nil
}

// DeleteApp 删除应用（软删除）
func (s *sSystem) DeleteApp(ctx context.Context, req *adminApi.AppDeleteReq) (*adminApi.AppDeleteRes, error) {
	var (
		res = &adminApi.AppDeleteRes{}
	)

	// 执行软删除
	result, err := dao.SystemApps.Ctx(ctx).Data("deleted_at", gdb.Raw("NOW()")).Where("appid", req.AppId).Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被删除
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("应用不存在")
	}

	res.AppId = req.AppId
	return res, nil
}

// GetAppSelectList 获取应用下拉选项列表
func (s *sSystem) GetAppSelectList(ctx context.Context) ([]*adminApi.AppSelectOption, error) {
	var entities []*entity.SystemApps
	logger.Debug("开始查询应用下拉选项列表")
	err := dao.SystemApps.Ctx(ctx).
		Fields("appid, app_name").
		Where("deleted_at IS NULL").
		OrderAsc("app_name").
		Scan(&entities)
	if err != nil {
		logger.Errorf("获取应用下拉选项列表失败:%s", err.Error())
		return nil, err
	}
	logger.Debugf("查询成功，结果数量:%d", len(entities))
	if len(entities) > 0 {
		logger.Debugf("第一条记录:%v", entities[0])
	}

	// 手动转换为响应类型
	var list []*adminApi.AppSelectOption
	for _, app := range entities {
		list = append(list, &adminApi.AppSelectOption{
			AppId:   app.Appid,
			AppName: app.AppName,
		})
	}

	return list, nil
}

// GetAppByAppId 统一获取系统应用信息（优先从Redis获取）
func (s *sSystem) GetAppByAppId(ctx context.Context, appid string) (*entity.SystemApps, error) {
	// 优先从Redis获取
	redisClient := util.GetRedisClient()
	key := "system_app:" + appid
	appJSONStr, err := redisClient.Get(ctx, key).Result()
	logger.Infof("app:%s", appJSONStr)
	if err == nil {
		// Redis中存在数据，直接解析返回
		var app entity.SystemApps
		err = json.Unmarshal([]byte(appJSONStr), &app)
		if err == nil {
			logger.Debugf("从Redis获取应用信息成功: %s", appid)
			return &app, nil
		}
		logger.Warnf("解析Redis中的应用信息失败: %v", err)
	} else if err != redis.Nil {
		// Redis操作出错，记录日志但继续从数据库获取
		logger.Warnf("从Redis获取应用信息失败: %v", err)
	}

	// Redis中不存在或解析失败，从数据库获取
	var app *entity.SystemApps
	err = dao.SystemApps.Ctx(ctx).Where("appid", appid).Scan(&app)
	if err != nil {
		logger.Errorf("从数据库获取应用信息失败: %v", err)
		return nil, err
	}

	// 将从数据库获取的数据存储到Redis
	if app != nil {
		appJSON, err := json.Marshal(app)
		if err == nil {
			err = redisClient.Set(ctx, key, appJSON, 0).Err()
			if err != nil {
				logger.Warnf("存储应用信息到Redis失败: %v", err)
			}
		} else {
			logger.Warnf("序列化应用信息失败: %v", err)
		}
	}

	return app, nil
}
