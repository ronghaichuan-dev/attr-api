package system

import (
	"context"
	"fmt"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GetPermissionList 获取权限列表
func (s *sSystem) GetPermissionList(ctx context.Context, req *adminApi.PermissionListReq) (*adminApi.PermissionListRes, error) {
	var (
		res = &adminApi.PermissionListRes{}
	)
	// 构建查询条件
	m := dao.SystemPermissions.Ctx(ctx)

	if req.PermissionName != "" {
		m = m.WhereLike("permission_name", "%"+req.PermissionName+"%")
	}
	if req.PermissionCode != "" {
		m = m.WhereLike("permission_code", "%"+req.PermissionCode+"%")
	}
	if req.Module != "" {
		m = m.WhereLike("module", "%"+req.Module+"%")
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

// GetPermissionById 根据ID获取权限详情
func (s *sSystem) GetPermissionById(ctx context.Context, permissionId uint) (*entity.PermissionsCustom, error) {
	var permission *entity.PermissionsCustom
	err := dao.SystemPermissions.Ctx(ctx).Where("id", permissionId).Where("deleted_at IS NULL").Scan(&permission)
	return permission, err
}

// GetPermissionByRoute 根据路由获取权限信息
func (s *sSystem) GetPermissionByRoute(ctx context.Context, route string) (*entity.PermissionsCustom, error) {
	var permission *entity.PermissionsCustom
	err := dao.SystemPermissions.Ctx(ctx).Where("route", route).Where("deleted_at IS NULL").Where("status", 1).Scan(&permission)
	return permission, err
}

// CreatePermission 创建权限
func (s *sSystem) CreatePermission(ctx context.Context, req *adminApi.PermissionCreateReq) (*adminApi.PermissionCreateRes, error) {
	var (
		res = &adminApi.PermissionCreateRes{}
	)

	// 计算权限级别
	level := 1
	if req.ParentId > 0 {
		var parentPermission *entity.SystemPermissions
		err := dao.SystemPermissions.Ctx(ctx).Where("id", req.ParentId).Where("deleted_at IS NULL").Scan(&parentPermission)
		if err != nil {
			return nil, err
		}
		if parentPermission == nil {
			return nil, fmt.Errorf("父级权限不存在")
		}
		level = parentPermission.Level + 1
	}

	// 构建权限数据
	permissionDO := &entity.SystemPermissions{
		PermissionName: req.PermissionName,
		PermissionCode: req.PermissionCode,
		PermissionDesc: req.PermissionDesc,
		Module:         req.Module,
		Route:          req.Route,
		ParentId:       int(req.ParentId),
		Level:          level,
		Status:         req.Status,
	}

	// 执行插入
	lastInsertId, err := dao.SystemPermissions.Ctx(ctx).Data(permissionDO).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	res.Id = uint(lastInsertId)
	return res, nil
}

// UpdatePermission 更新权限信息
func (s *sSystem) UpdatePermission(ctx context.Context, req *adminApi.PermissionUpdateReq) (*adminApi.PermissionUpdateRes, error) {
	var (
		res = &adminApi.PermissionUpdateRes{}
	)

	// 构建更新数据
	data := g.Map{}
	if req.PermissionName != "" {
		data["permission_name"] = req.PermissionName
	}
	if req.PermissionCode != "" {
		data["permission_code"] = req.PermissionCode
	}
	if req.PermissionDesc != "" {
		data["permission_desc"] = req.PermissionDesc
	}
	if req.Module != "" {
		data["module"] = req.Module
	}
	if req.Route != "" {
		data["route"] = req.Route
	}
	if req.ParentId != 0 {
		if req.ParentId == req.Id {
			return nil, fmt.Errorf("父级权限不能是自己")
		}
		// 计算权限级别
		var parentPermission *entity.SystemPermissions
		err := dao.SystemPermissions.Ctx(ctx).Where("id", req.ParentId).Where("deleted_at IS NULL").Scan(&parentPermission)
		if err != nil {
			return nil, err
		}
		if parentPermission == nil {
			return nil, fmt.Errorf("父级权限不存在")
		}
		data["parent_id"] = req.ParentId
		data["level"] = parentPermission.Level + 1
	}
	if req.Status != -1 {
		data["status"] = req.Status
	}
	data["updated_at"] = gdb.Raw("NOW()")

	// 执行更新
	result, err := dao.SystemPermissions.Ctx(ctx).Data(data).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("权限不存在")
	}

	res.Id = req.Id
	return res, nil
}

// DeletePermission 删除权限（软删除）
func (s *sSystem) DeletePermission(ctx context.Context, req *adminApi.PermissionDeleteReq) (*adminApi.PermissionDeleteRes, error) {
	var (
		res = &adminApi.PermissionDeleteRes{}
	)

	// 执行软删除
	result, err := dao.SystemPermissions.Ctx(ctx).Data("deleted_at", gdb.Raw("NOW()")).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被删除
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("权限不存在")
	}

	res.Id = req.Id
	return res, nil
}

// GetPermissionTree 获取权限树形结构
func (s *sSystem) GetPermissionTree(ctx context.Context) ([]*adminApi.PermissionTreeItem, error) {
	var permissions []*entity.SystemPermissions
	err := dao.SystemPermissions.Ctx(ctx).Where("deleted_at IS NULL").Order("level").Order("created_at").Scan(&permissions)
	if err != nil {
		return nil, err
	}

	// 将实体转换为包含子权限的树状结构项
	var treeItems []*adminApi.PermissionTreeItem
	for _, permission := range permissions {
		createdAt := ""
		if permission.CreatedAt != nil {
			createdAt = permission.CreatedAt.String()
		}
		updatedAt := ""
		if permission.UpdatedAt != nil {
			updatedAt = permission.UpdatedAt.String()
		}
		deletedAt := ""
		if permission.DeletedAt != nil {
			deletedAt = permission.DeletedAt.String()
		}
		treeItems = append(treeItems, &adminApi.PermissionTreeItem{
			Id:             permission.Id,
			PermissionName: permission.PermissionName,
			PermissionCode: permission.PermissionCode,
			PermissionDesc: permission.PermissionDesc,
			Module:         permission.Module,
			Status:         permission.Status,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			DeletedAt:      deletedAt,
			Route:          permission.Route,
			ParentId:       permission.ParentId,
			Level:          permission.Level,
			Children:       []*adminApi.PermissionTreeItem{}, // 初始化子权限列表
		})
	}

	// 创建权限ID到权限对象的映射
	permissionMap := make(map[int]*adminApi.PermissionTreeItem)
	for _, item := range treeItems {
		permissionMap[item.Id] = item
	}

	// 构建树形结构
	var tree []*adminApi.PermissionTreeItem
	for _, item := range treeItems {
		if item.ParentId == 0 {
			tree = append(tree, item)
		} else {
			parent, exists := permissionMap[item.ParentId]
			if exists {
				parent.Children = append(parent.Children, item)
			} else {
				// 如果找不到父级权限，将其作为顶级权限处理
				tree = append(tree, item)
			}
		}
	}

	return tree, nil
}

// GetPermissionsByUserId 根据用户ID获取拥有的权限
func (s *sSystem) GetPermissionsByUserId(ctx context.Context, userId uint) ([]*entity.PermissionsCustom, error) {
	// 查询用户信息，获取role_id
	user, err := service.System().GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("用户不存在")
	}

	// 查询角色拥有的权限
	permissions, err := service.System().GetPermissionsByRoleId(ctx, uint(user.RoleId))
	if err != nil {
		return nil, err
	}

	// 权限去重
	seenIds := make(map[int]bool)
	var uniquePermissions []*entity.PermissionsCustom
	for _, permission := range permissions {
		if !seenIds[permission.Id] {
			seenIds[permission.Id] = true
			uniquePermissions = append(uniquePermissions, permission)
		}
	}

	return uniquePermissions, nil
}

// EnablePermission 启用权限
func (s *sSystem) EnablePermission(ctx context.Context, id uint) (*adminApi.PermissionEnableRes, error) {
	var res = &adminApi.PermissionEnableRes{}

	// 构建更新数据
	data := g.Map{
		"status":     1,
		"updated_at": gdb.Raw("NOW()"),
	}

	// 执行更新
	result, err := dao.SystemPermissions.Ctx(ctx).Data(data).Where("id", id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("权限不存在")
	}

	res.Id = id
	return res, nil
}

// DisablePermission 禁用权限
func (s *sSystem) DisablePermission(ctx context.Context, id uint) (*adminApi.PermissionDisableRes, error) {
	var res = &adminApi.PermissionDisableRes{}

	// 构建更新数据
	data := g.Map{
		"status":     2,
		"updated_at": gdb.Raw("NOW()"),
	}

	// 执行更新
	result, err := dao.SystemPermissions.Ctx(ctx).Data(data).Where("id", id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("权限不存在")
	}

	res.Id = id
	return res, nil
}

// UpdatePermissionStatus 更新权限状态
func (s *sSystem) UpdatePermissionStatus(ctx context.Context, req *adminApi.PermissionUpdateStatusReq) (*adminApi.PermissionUpdateStatusRes, error) {
	var (
		res = &adminApi.PermissionUpdateStatusRes{}
	)

	// 构建更新数据
	data := g.Map{
		"status":     req.Status,
		"updated_at": gdb.Raw("NOW()"),
	}

	// 执行更新
	result, err := dao.SystemPermissions.Ctx(ctx).Data(data).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("权限不存在")
	}

	res.Id = req.Id
	res.Status = req.Status
	return res, nil
}
