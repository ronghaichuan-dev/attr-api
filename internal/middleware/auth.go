package middleware

import (
	"strings"

	"god-help-service/internal/util"
	"god-help-service/internal/util/logger"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// BasicAuthMiddleware Basic Auth 中间件
//func BasicAuthMiddleware(r *ghttp.Request) {
//	// 这里可以从配置文件或数据库中获取有效的用户名和密码
//
//
//	// 获取 Authorization 头部
//	authHeader := r.Header.Get("Authorization")
//	if authHeader == "" {
//		r.Response.Header().Set("WWW-Authenticate", `Basic realm="App API"`)
//		r.Response.WriteJsonExit(g.Map{
//			"code":    401,
//			"message": "需要Basic Auth认证",
//			"data":    nil,
//		})
//		return
//	}
//
//	// 解析 Basic Auth 信息
//	if !strings.HasPrefix(authHeader, "Basic ") {
//		r.Response.Header().Set("WWW-Authenticate", `Basic realm="App API"`)
//		r.Response.WriteJsonExit(g.Map{
//			"code":    401,
//			"message": "认证方式不正确，需要Basic Auth",
//			"data":    nil,
//		})
//		return
//	}
//
//	// 解码用户名和密码
//	encodedCreds := strings.TrimPrefix(authHeader, "Basic ")
//
//	decodedCreds, err := base64.StdEncoding.DecodeString(encodedCreds)
//	if err != nil {
//		r.Response.Header().Set("WWW-Authenticate", `Basic realm="App API"`)
//		r.Response.WriteJsonExit(g.Map{
//			"code":    401,
//			"message": "认证信息解析失败",
//			"data":    nil,
//		})
//		return
//	}
//
//	// 分割用户名和密码
//	creds := strings.SplitN(string(decodedCreds), ":", 2)
//	if len(creds) != 2 {
//		r.Response.Header().Set("WWW-Authenticate", `Basic realm="App API"`)
//		r.Response.WriteJsonExit(g.Map{
//			"code":    401,
//			"message": "认证信息格式不正确",
//			"data":    nil,
//		})
//		return
//	}
//
//	// 验证用户名和密码
//	username, password := creds[0], creds[1]
//	if username != validUsername || password != validPassword {
//		r.Response.Header().Set("WWW-Authenticate", `Basic realm="App API"`)
//		r.Response.WriteJsonExit(g.Map{
//			"code":    401,
//			"message": "用户名或密码错误",
//			"data":    nil,
//		})
//		return
//	}
//
//	// 认证成功，继续处理请求
//	r.Middleware.Next()
//}

// AuthMiddleware 用户认证中间件
func AuthMiddleware(r *ghttp.Request) {
	// 定义不需要认证的路由
	noAuthRoutes := map[string]bool{
		"/captcha":   true,
		"/login":     true,
		"/logout":    true,
		"/api.json":  true,
		"/swagger/*": true,
		"/":          true,
	}

	// 检查当前路由是否需要认证
	path := r.URL.Path
	if noAuthRoutes[path] || strings.HasPrefix(path, "/swagger/") {
		r.Middleware.Next()
		return
	}

	// 获取请求头中的 token
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

	// 验证JWT Token
	claims, err := util.ParseJWT(token)
	if err != nil {
		logger.Errorf("JWT Token解析失败: %v", err)
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "token 无效",
			"data":    nil,
		})
		return
	}

	// 检查Token是否在Redis中存在（支持强制登出）
	if !util.CheckTokenExists(r.Context(), token) {
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "token 已过期或已失效",
			"data":    nil,
		})
		return
	}

	// 将用户信息存储到上下文
	r.SetCtxVar("user", claims.Username)
	r.SetCtxVar("user_id", claims.UserID)

	// 继续处理请求
	r.Middleware.Next()
}
