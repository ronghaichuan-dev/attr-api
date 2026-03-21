package system

import (
	"context"
	"fmt"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GetSystemSettingList 获取系统设置列表
func (s *sSystem) GetSystemSettingList(ctx context.Context, req *adminApi.SystemSettingListReq) (*adminApi.SystemSettingListRes, error) {
	var (
		res = &adminApi.SystemSettingListRes{}
	)
	// 构建查询条件
	m := dao.SystemSettings.Ctx(ctx)

	if req.Key != "" {
		m = m.WhereLike("key", "%"+req.Key+"%")
	}

	// 获取总记录数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = int64(total)

	// 分页查询
	offset := (req.Page - 1) * req.Size
	err = m.Offset(offset).Limit(req.Size).OrderDesc("created_at").Scan(&res.List)

	return res, err
}

// GetSystemSettingById 根据ID获取系统设置详情
func (s *sSystem) GetSystemSettingById(ctx context.Context, id int) (*entity.SystemSettings, error) {
	var setting *entity.SystemSettings
	err := dao.SystemSettings.Ctx(ctx).Where("id", id).Scan(&setting)
	return setting, err
}

// GetSystemSettingByKey 根据Key获取系统设置详情
func (s *sSystem) GetSystemSettingByKey(ctx context.Context, key string) (*entity.SystemSettings, error) {
	var setting *entity.SystemSettings
	err := dao.SystemSettings.Ctx(ctx).Where("key", key).Scan(&setting)
	return setting, err
}

// CreateSystemSetting 创建系统设置
func (s *sSystem) CreateSystemSetting(ctx context.Context, req *adminApi.SystemSettingCreateReq) (*adminApi.SystemSettingCreateRes, error) {
	var (
		res = &adminApi.SystemSettingCreateRes{}
	)

	// 检查key是否已存在
	existing, err := s.GetSystemSettingByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("设置键已存在")
	}

	// 构建系统设置数据
	settingDO := &entity.SystemSettings{
		Key:   req.Key,
		Value: req.Value,
	}

	// 执行插入
	result, err := dao.SystemSettings.Ctx(ctx).Data(settingDO).Insert()
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	res.Id = int(id)

	return res, nil
}

// UpdateSystemSetting 更新系统设置
func (s *sSystem) UpdateSystemSetting(ctx context.Context, req *adminApi.SystemSettingUpdateReq) (*adminApi.SystemSettingUpdateRes, error) {
	var (
		res = &adminApi.SystemSettingUpdateRes{}
	)

	// 构建更新数据
	data := g.Map{}
	if req.Key != "" {
		// 检查key是否已被其他记录使用
		existing, err := s.GetSystemSettingByKey(ctx, req.Key)
		if err != nil {
			return nil, err
		}
		if existing != nil && existing.Id != req.Id {
			return nil, fmt.Errorf("设置键已存在")
		}
		data["key"] = req.Key
	}
	if req.Value != "" {
		data["value"] = req.Value
	}
	data["updated_at"] = gdb.Raw("NOW()")

	// 执行更新
	result, err := dao.SystemSettings.Ctx(ctx).Data(data).Where("id", req.Id).Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("系统设置不存在")
	}

	res.Id = req.Id
	return res, nil
}

// DeleteSystemSetting 删除系统设置（软删除）
func (s *sSystem) DeleteSystemSetting(ctx context.Context, req *adminApi.SystemSettingDeleteReq) (*adminApi.SystemSettingDeleteRes, error) {
	var (
		res = &adminApi.SystemSettingDeleteRes{}
	)

	// 执行软删除
	result, err := dao.SystemSettings.Ctx(ctx).Data("deleted_at", gdb.Raw("NOW()")).Where("id", req.Id).Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被删除
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("系统设置不存在")
	}

	res.Id = req.Id
	return res, nil
}

// GetSystemSettingValueByKey 根据Key获取系统设置值
func (s *sSystem) GetSystemSettingValueByKey(ctx context.Context, key string) (string, error) {
	setting, err := s.GetSystemSettingByKey(ctx, key)
	if err != nil {
		return "", err
	}
	if setting == nil {
		return "", fmt.Errorf("系统设置不存在")
	}
	return setting.Value, nil
}
