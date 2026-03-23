package router

import (
	"encoding/base64"
	"god-help-service/internal/controller/api"
	"god-help-service/internal/controller/app"
	"god-help-service/internal/middleware"
	"god-help-service/internal/util/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	validUsername = "god-help-service-auth"
	validPassword = "CsRsZJk#fkd2048fj"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	// 全局中间件 - 启用CORS支持
	r.Use(middleware.CorsMiddleware)
	//r.Static("./public", "./public")
	//// 静态文件服务 - 使用StaticFile为特定文件创建路由，避免与API路由冲突
	//// 首页路由
	//r.GET("/", func(c *gin.Context) {
	//	c.File("./public/monitor.html")
	//})
	// 内存监控路由
	//apiGroup.GET("/notification/memory", func(r *gin.Context) {
	//	logger.Print(r, "内存监控路由被调用")
	//	// 获取内存监控数据
	//	stats := service.MemoryMonitor().GetStats()
	//	logger.Infof(r, "获取到内存监控数据:%d", len(stats))
	//
	//	// 构建响应
	//	response := map[string]interface{}{
	//		"success": true,
	//		"data":    stats,
	//	}
	//
	//	logger.Infof(r, "返回内存监控数据:%v", response)
	//	r.JSON(http.StatusOK, response)
	//})
	r.GET("/healthy", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok"})
	})
	// API路由组
	apiGroup := r.Group("/api/v1")
	notification := new(api.NotificationController)
	// 通知回调路由
	apiGroup.POST("/notification/callback", notification.Callback)

	// 点击/展示追踪路由
	tracking := new(api.TrackingController)
	apiGroup.GET("/tracking/click", tracking.Click)
	apiGroup.GET("/tracking/impression", tracking.Impression)

	logger.Info("Basic " + base64.StdEncoding.EncodeToString([]byte(validUsername+":"+validPassword)))
	// APP路由组
	appGroup := r.Group("/app/v1", gin.BasicAuth(gin.Accounts{validUsername: validPassword}))
	appGroup.POST("/event/report", new(app.EventController).Report)
	appGroup.POST("/attribution/report", new(app.AttributionController).Report)

	return r
}
