package router

import (
	"fmt"
	"god-help-service/internal/controller/admin"
	"god-help-service/internal/middleware"
	"god-help-service/internal/util"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouter(r *ghttp.Server) {
	// 使用响应中间件

	// 基础路由
	r.BindHandler("/", func(c *ghttp.Request) {
		c.Response.WriteString("Hello, Admin Service!")
	})

	r.BindHandler("/healthy", func(c *ghttp.Request) {
		c.Response.WriteJson(g.Map{"status": "ok"})
	})
	fmt.Println(util.HashPassword("123456"))
	r.BindHandler("/ready", func(c *ghttp.Request) {
		c.Response.WriteJson(g.Map{"status": "ready"})
	})
	// 管理后台路由组
	adminGroup := r.Group("/admin").Middleware(middleware.ResponseHandler)
	// 认证相关路由
	auth := new(admin.AuthController)
	adminGroup.POST("/login", auth.Login)
	adminGroup.POST("/logout", auth.Logout)
	adminGroup.GET("/captcha", auth.Captcha)

	// 用户管理路由
	user := new(admin.UserController)
	userGroup := adminGroup.Group("/user")
	userGroup.GET("/list", user.List)
	userGroup.GET("/detail", user.Detail)
	userGroup.POST("/create", user.Create)
	userGroup.PUT("/update", user.Update)
	userGroup.DELETE("/delete", user.Delete)

	// 角色管理路由
	role := new(admin.RoleController)
	roleGroup := adminGroup.Group("/role")
	roleGroup.GET("/list", role.List)
	roleGroup.GET("/detail", role.Detail)
	roleGroup.POST("/create", role.Create)
	roleGroup.PUT("/update", role.Update)
	roleGroup.DELETE("/delete", role.Delete)
	roleGroup.GET("/select/list", role.SelectList)
	roleGroup.PUT("/status", role.UpdateStatus)

	// 权限管理路由
	permission := new(admin.PermissionController)
	permissionGroup := adminGroup.Group("/permission")
	permissionGroup.GET("/list", permission.List)
	permissionGroup.GET("/tree", permission.Tree)
	permissionGroup.GET("/detail", permission.Detail)
	permissionGroup.POST("/create", permission.Create)
	permissionGroup.PUT("/update", permission.Update)
	permissionGroup.DELETE("/delete", permission.Delete)
	permissionGroup.PUT("/enable", permission.Enable)
	permissionGroup.PUT("/disable", permission.Disable)

	// 角色权限管理路由
	rolePermission := new(admin.RolePermissionController)
	rolePermissionGroup := adminGroup.Group("/role-permission")
	rolePermissionGroup.GET("/list", rolePermission.List)
	rolePermissionGroup.GET("/detail", rolePermission.Detail)
	rolePermissionGroup.POST("/assign", rolePermission.Assign)
	rolePermissionGroup.DELETE("/remove", rolePermission.Remove)
	rolePermissionGroup.GET("/permissions", rolePermission.GetPermissions)
	rolePermissionGroup.GET("/roles", rolePermission.GetRoles)

	// 应用管理路由
	app := new(admin.AppController)
	appGroup := adminGroup.Group("/app")
	appGroup.GET("/list", app.List)
	appGroup.GET("/detail", app.Detail)
	appGroup.POST("/create", app.Create)
	appGroup.PUT("/update", app.Update)
	appGroup.DELETE("/delete", app.Delete)
	appGroup.GET("/subscription/trend", app.SubscriptionTrend)

	// 仪表盘路由
	dashboard := new(admin.DashboardController)
	dashboardGroup := adminGroup.Group("/dashboard")
	dashboardGroup.GET("/analytics", dashboard.DashboardAnalytics)
	dashboardGroup.GET("/app/daily/trend", dashboard.AppDailyTrend)
	dashboardGroup.GET("/app/select/list", dashboard.AppSelectList)

	// 事件管理路由
	event := new(admin.EventController)
	eventGroup := adminGroup.Group("/event")
	eventGroup.GET("/list", event.List)
	eventGroup.GET("/detail", event.Detail)
	eventGroup.POST("/create", event.Create)
	eventGroup.PUT("/update", event.Update)
	eventGroup.DELETE("/delete", event.Delete)
	eventGroup.GET("/dropdown", event.Dropdown)

	// 应用事件日志路由
	appEventLog := new(admin.AppEventLogController)
	appEventLogGroup := adminGroup.Group("/app-event-log")
	appEventLogGroup.GET("/list", appEventLog.List)
	appEventLogGroup.GET("/detail", appEventLog.Detail)

	// 通知管理路由
	notification := new(admin.NoticeController)
	notificationGroup := adminGroup.Group("/notification")
	notificationGroup.GET("/list", notification.List)

	// 系统设置路由
	setting := new(admin.SystemSettingController)
	settingGroup := adminGroup.Group("/system-setting")
	settingGroup.GET("/list", setting.List)
	settingGroup.GET("/detail", setting.Detail)
	settingGroup.POST("/create", setting.Create)
	settingGroup.PUT("/update", setting.Update)
	settingGroup.DELETE("/delete", setting.Delete)

	// 文件上传路由
	upload := new(admin.UploadController)
	uploadGroup := adminGroup.Group("/upload")
	uploadGroup.POST("", upload.UploadImage)
}
