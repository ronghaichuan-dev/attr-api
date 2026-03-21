package admin

import (
	"context"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"

	adminApi "god-help-service/api/v1/admin"

	"github.com/gogf/gf/v2/errors/gerror"
)

// UserController 用户控制器
type UserController struct{}

// List 获取用户列表
func (c *UserController) List(ctx context.Context, req *adminApi.UserListReq) (*adminApi.UserListRes, error) {
	logger.Debugf("解析后的请求参数: %v", req)

	// 执行查询

	res, err := service.System().GetUserList(ctx, req)
	if err != nil {
		logger.Errorf("获取用户列表失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Detail 获取用户详情
func (c *UserController) Detail(ctx context.Context, req *adminApi.UserDetailReq) (*adminApi.UserDetailRes, error) {
	// 执行查询

	user, err := service.System().GetUserById(ctx, req.Id)
	if err != nil {
		logger.Errorf("获取用户详情失败: %v", err)
		return nil, err
	}

	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	return &adminApi.UserDetailRes{Users: user}, nil
}

// Create 创建用户
func (c *UserController) Create(ctx context.Context, req *adminApi.UserCreateReq) (*adminApi.UserCreateRes, error) {
	// 执行创建

	res, err := service.System().CreateUser(ctx, req)
	if err != nil {
		logger.Errorf("创建用户失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Update 更新用户
func (c *UserController) Update(ctx context.Context, req *adminApi.UserUpdateReq) (*adminApi.UserUpdateRes, error) {
	// 执行更新

	res, err := service.System().UpdateUser(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "用户不存在" {
			return nil, gerror.New("用户不存在或已被删除")
		}

		logger.Errorf("更新用户失败: %v", err)
		return nil, err
	}

	return res, nil
}

// Delete 删除用户
func (c *UserController) Delete(ctx context.Context, req *adminApi.UserDeleteReq) (*adminApi.UserDeleteRes, error) {
	// 执行删除

	res, err := service.System().DeleteUser(ctx, req)
	if err != nil {
		if err != nil && err.Error() == "用户不存在" {
			return nil, gerror.New("用户不存在或已被删除")
		}

		logger.Errorf("删除用户失败: %v", err)
		return nil, err
	}

	return res, nil
}
