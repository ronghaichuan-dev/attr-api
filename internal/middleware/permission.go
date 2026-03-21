package middleware

import (
	"god-help-service/internal/service"
	"strings"

	"god-help-service/internal/util"
	"god-help-service/internal/util/logger"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 不需要权限验证的路由
var noPermissionRoutes = map[string]bool{
	"/captcha":          true,
	"/login":            true,
	"/logout":           true,
	"/user/createTable": true,
	"/api.json":         true,
	"/swagger/*":        true,
	"/":                 true,
}

// PermissionMiddleware 权限控制中间件
func PermissionMiddleware(r *ghttp.Request) {
	// 直接判断是否是用户表初始化接口，跳过权限验证
	if r.URL.Path == "/user/createTable" || r.URL.Path == "/app/create" || noPermissionRoutes[r.URL.Path] || strings.HasPrefix(r.URL.Path, "/swagger/") {
		r.Middleware.Next()
		return
	}

	// 获取当前请求路径
	path := r.URL.Path

	// 从数据库查询该路由对应的权限信息
	permission, err := service.System().GetPermissionByRoute(r.Context(), path)
	if err != nil {
		logger.Errorf("查询权限信息失败:%s", err)
		// 查询权限信息失败时，直接通过权限验证，继续处理请求
		r.Middleware.Next()
		return
	}

	// 如果该路由不需要权限验证，直接通过
	if permission == nil {
		r.Middleware.Next()
		return
	}

	// 获取JWT Token
	token := r.Header.Get("Authorization")
	if token == "" {
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "未授权访问",
			"data":    nil,
		})
		return
	}

	// 去掉 Bearer 前缀
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	// 验证JWT Token并解析用户信息
	claims, err := util.ParseJWT(token)
	if err != nil {
		logger.Errorf("JWT Token解析失败:%s", err.Error())
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "token 无效",
			"data":    nil,
		})
		return
	}

	// 根据用户ID获取用户信息
	user, err := service.System().GetUserById(r.Context(), uint(claims.UserID))
	if err != nil {
		logger.Errorf("获取用户信息失败:%s", err.Error())
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "获取用户信息失败",
			"data":    nil,
		})
		return
	}

	if user == nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "用户不存在",
			"data":    nil,
		})
		return
	}

	// 检查用户是否有角色
	if user.RoleId == 0 {
		// 检查用户是否是admin用户（用户名或其他标识）
		if user.Username == "admin" {
			// admin用户不需要权限验证，可以访问所有接口
			r.Middleware.Next()
			return
		}
		r.Response.WriteJsonExit(g.Map{
			"code":    403,
			"message": "用户未分配角色，无法访问",
			"data":    nil,
		})
		return
	}

	if user.Username == "admin" {
		// 超级管理员角色不需要权限验证，可以访问所有接口
		r.Middleware.Next()
		return
	}

	// 检查用户角色是否有该权限
	hasPermission, err := service.System().CheckPermission(r.Context(), user.RoleId, permission.PermissionCode)
	if err != nil {
		logger.Errorf("检查权限失败:%s", err.Error())
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "检查权限失败",
			"data":    nil,
		})
		return
	}

	if !hasPermission {
		r.Response.WriteJsonExit(g.Map{
			"code":    403,
			"message": "您没有访问该资源的权限",
			"data":    nil,
		})
		return
	}

	// 权限验证成功，继续处理请求
	r.Middleware.Next()
}
