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

// GetRoleList 获取角色列表
func (s *sSystem) GetRoleList(ctx context.Context, req *adminApi.RoleListReq) (*adminApi.RoleListRes, error) {
	var (
		res      = &adminApi.RoleListRes{}
		entities []*entity.SystemRoles
	)

	// 构建查询条件
	m := dao.SystemRoles.Ctx(ctx)

	if req.RoleName != "" {
		m = m.WhereLike("role_name", "%"+req.RoleName+"%")
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

	// 转换为自定义结构体
	res.List = make([]*entity.RolesCustom, len(entities))
	for i, role := range entities {
		res.List[i] = &entity.RolesCustom{
			Id:        role.Id,
			RoleName:  role.RoleName,
			RoleCode:  role.RoleCode,
			RoleDesc:  role.RoleDesc,
			Status:    role.Status,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
			DeletedAt: role.DeletedAt,
		}
	}

	return res, err
}

// GetRoleById 根据ID获取角色详情
func (s *sSystem) GetRoleById(ctx context.Context, roleId uint) (*entity.RolesCustom, error) {
	var role *entity.SystemRoles
	err := dao.SystemRoles.Ctx(ctx).Where("id", roleId).Where("deleted_at IS NULL").Scan(&role)
	if err != nil {
		return nil, err
	}

	// 转换为自定义结构体
	if role == nil {
		return nil, nil
	}

	return &entity.RolesCustom{
		Id:        role.Id,
		RoleName:  role.RoleName,
		RoleCode:  role.RoleCode,
		RoleDesc:  role.RoleDesc,
		Status:    role.Status,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}, nil
}

// CreateRole 创建角色
func (s *sSystem) CreateRole(ctx context.Context, req *adminApi.RoleCreateReq) (*adminApi.RoleCreateRes, error) {
	var (
		res = &adminApi.RoleCreateRes{}
	)

	// 构建角色数据
	roleDO := &entity.SystemRoles{
		RoleName: req.RoleName,
		RoleCode: req.RoleCode,
		RoleDesc: req.RoleDesc,
		Status:   req.Status,
	}

	// 执行插入
	lastInsertId, err := dao.SystemRoles.Ctx(ctx).Data(roleDO).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	res.Id = uint(lastInsertId)
	return res, nil
}

// UpdateRole 更新角色信息
func (s *sSystem) UpdateRole(ctx context.Context, req *adminApi.RoleUpdateReq) (*adminApi.RoleUpdateRes, error) {
	var (
		res = &adminApi.RoleUpdateRes{}
	)

	// 构建更新数据
	data := g.Map{}
	if req.RoleName != "" {
		data["role_name"] = req.RoleName
	}
	if req.RoleCode != "" {
		data["role_code"] = req.RoleCode
	}
	if req.RoleDesc != "" {
		data["role_desc"] = req.RoleDesc
	}
	if req.Status != -1 {
		data["status"] = req.Status
	}
	data["updated_at"] = gdb.Raw("NOW()")

	// 执行更新
	result, err := dao.SystemRoles.Ctx(ctx).Data(data).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("角色不存在")
	}

	res.Id = req.Id
	return res, nil
}

// DeleteRole 删除角色（软删除）
func (s *sSystem) DeleteRole(ctx context.Context, req *adminApi.RoleDeleteReq) (*adminApi.RoleDeleteRes, error) {
	var (
		res = &adminApi.RoleDeleteRes{}
	)

	// 执行软删除
	result, err := dao.SystemRoles.Ctx(ctx).Data("deleted_at", gdb.Raw("NOW()")).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被删除
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("角色不存在")
	}

	res.Id = req.Id
	return res, nil
}

// GetRoleSelectList 获取角色选择列表（用于用户选择角色，无分页，只返回启用状态的角色）
func (s *sSystem) GetRoleSelectList(ctx context.Context) (*adminApi.RoleSelectListRes, error) {
	var (
		res      = &adminApi.RoleSelectListRes{}
		entities []*entity.SystemRoles
	)

	// 查询所有启用状态的角色，按创建时间降序排序
	err := dao.SystemRoles.Ctx(ctx).Where("deleted_at IS NULL").Where("status = ?", 1).OrderDesc("created_at").Scan(&entities)
	if err != nil {
		return nil, err
	}

	// 转换为自定义结构体
	res.List = make([]*entity.RolesCustom, len(entities))
	for i, role := range entities {
		res.List[i] = &entity.RolesCustom{
			Id:        role.Id,
			RoleName:  role.RoleName,
			RoleCode:  role.RoleCode,
			RoleDesc:  role.RoleDesc,
			Status:    role.Status,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
			DeletedAt: role.DeletedAt,
		}
	}

	return res, err
}

// UpdateRoleStatus 更新角色状态
func (s *sSystem) UpdateRoleStatus(ctx context.Context, req *adminApi.RoleUpdateStatusReq) (*adminApi.RoleUpdateStatusRes, error) {
	var (
		res = &adminApi.RoleUpdateStatusRes{}
	)

	// 构建更新数据
	data := g.Map{
		"status":     req.Status,
		"updated_at": gdb.Raw("NOW()"),
	}

	// 执行更新
	result, err := dao.SystemRoles.Ctx(ctx).Data(data).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("角色不存在")
	}

	res.Id = req.Id
	res.Status = req.Status
	return res, nil
}
