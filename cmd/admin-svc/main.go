package main

import (
	"god-help-service/cmd/admin-svc/router"
	_ "god-help-service/internal/logic"
	"god-help-service/internal/util/logger"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "god-help-service/docs"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	// 创建 GoFrame 服务器
	s := g.Server()

	// 启用 OpenAPI 文档生成
	err := s.SetConfigWithMap(g.Map{
		"openapiPath": "/api.json", // OpenAPI 规范文件路径
		"swaggerPath": "/swagger",  // Swagger UI 路径
	})
	if err != nil {
		log.Fatalf("文档生成失败：%s", err.Error())
	}

	// 注册路由
	router.InitRouter(s)

	// 获取端口配置
	port := 8080
	s.SetPort(port)

	// 启动服务器
	go func() {
		logger.Infof("Admin service HTTP server starting on %d", port)
		s.Run()
	}()

	// 等待系统信号，优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down admin service...")

	// 关闭服务器
	if err = s.Shutdown(); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Info("Admin service stopped")
}
