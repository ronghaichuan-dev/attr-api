package system

import (
	"context"
	"fmt"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/util"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GetUserList 获取用户列表
func (s *sSystem) GetUserList(ctx context.Context, req *adminApi.UserListReq) (*adminApi.UserListRes, error) {
	var (
		res = &adminApi.UserListRes{}
	)

	// 构建查询条件
	m := dao.SystemUsers.Ctx(ctx).Where("deleted_at IS NULL")

	if req.UserName != "" {
		m = m.WhereLike("username", "%"+req.UserName+"%")
	}
	if req.Email != "" {
		m = m.Where("email", req.Email)
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

// GetUserById 根据ID获取用户详情
func (s *sSystem) GetUserById(ctx context.Context, userId uint) (*entity.SystemUsers, error) {
	var user *entity.SystemUsers
	err := dao.SystemUsers.Ctx(ctx).Where("id", userId).Where("deleted_at IS NULL").Scan(&user)
	return user, err
}

// GetUserByUsername 根据用户名获取用户详情
func (s *sSystem) GetUserByUsername(ctx context.Context, username string) (*entity.SystemUsers, error) {
	var user *entity.SystemUsers
	err := dao.SystemUsers.Ctx(ctx).Where("username", username).Where("deleted_at IS NULL").Scan(&user)
	return user, err
}

// GetUsersByIds 批量根据用户ID获取用户详情
func (s *sSystem) GetUsersByIds(ctx context.Context, userIds []int) (map[int]*entity.SystemUsers, error) {
	var users []*entity.SystemUsers
	err := dao.SystemUsers.Ctx(ctx).Where("id IN (?)", userIds).Where("deleted_at IS NULL").Scan(&users)
	if err != nil {
		return nil, err
	}

	userMap := make(map[int]*entity.SystemUsers)
	for _, user := range users {
		userMap[int(user.Id)] = user
	}

	return userMap, nil
}

// CreateUser 创建用户
func (s *sSystem) CreateUser(ctx context.Context, req *adminApi.UserCreateReq) (*adminApi.UserCreateRes, error) {
	var (
		res = &adminApi.UserCreateRes{}
	)

	// 构建用户数据
	userDO := &entity.SystemUsers{
		Username: req.UserName,
		Password: util.HashPassword(req.Password),
		RoleId:   req.RoleId,
	}

	// 执行插入
	result, err := dao.SystemUsers.Ctx(ctx).Data(userDO).Insert()
	if err != nil {
		return nil, err
	}

	// 获取插入后的ID
	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	res.Id = uint(userId)
	return res, nil
}

// UpdateUser 更新用户信息
func (s *sSystem) UpdateUser(ctx context.Context, req *adminApi.UserUpdateReq) (*adminApi.UserUpdateRes, error) {
	var (
		res = &adminApi.UserUpdateRes{}
	)

	// 构建更新数据
	data := g.Map{}
	if req.UserName != "" {
		data["username"] = req.UserName
	}
	if req.Password != "" {
		data["password"] = util.HashPassword(req.Password)
	}
	if req.Email != "" {
		data["email"] = req.Email
	}
	if req.RoleId > 0 {
		data["role_id"] = req.RoleId
	}
	data["updated_at"] = gdb.Raw("NOW()")

	// 执行更新
	result, err := dao.SystemUsers.Ctx(ctx).Data(data).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被更新
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("用户不存在")
	}

	res.Id = req.Id
	return res, nil
}

// DeleteUser 删除用户（软删除）
func (s *sSystem) DeleteUser(ctx context.Context, req *adminApi.UserDeleteReq) (*adminApi.UserDeleteRes, error) {
	var (
		res = &adminApi.UserDeleteRes{}
	)

	// 执行软删除
	result, err := dao.SystemUsers.Ctx(ctx).Data("deleted_at", gdb.Raw("NOW()")).Where("id", req.Id).Where("deleted_at IS NULL").Update()
	if err != nil {
		return nil, err
	}

	// 检查是否有记录被删除
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, fmt.Errorf("用户不存在")
	}

	res.Id = req.Id
	return res, nil
}
