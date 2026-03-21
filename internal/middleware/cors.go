package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/net/ghttp"
)

// CorsMiddleware CORS中间件（Gin框架版本）
func CorsMiddleware(r *gin.Context) {
	// 设置CORS headers
	r.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	r.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	r.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	r.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24小时
	if r.Request.Method == "OPTIONS" {
		r.AbortWithStatus(http.StatusNoContent)
		return
	}
	r.Next()
}

// CorsMiddlewareGF CORS中间件（GF框架版本）
func CorsMiddlewareGF(r *ghttp.Request) {
	// 设置CORS headers
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	r.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	r.Response.Header().Set("Access-Control-Max-Age", "86400") // 24小时
	if r.Method == "OPTIONS" {
		r.Response.Status = http.StatusNoContent
		r.Exit()
		return
	}
	// GF框架的中间件不需要调用Next()，会自动继续处理
}
