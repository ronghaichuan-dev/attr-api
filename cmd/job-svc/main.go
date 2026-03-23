package main

import (
	"context"
	"fmt"
	_ "god-help-service/internal/logic"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/robfig/cron/v3"
)

func main() {
	logger.Info("Job service starting...")

	initCronJobs()

	router := setupRouter()

	srv := &http.Server{
		Addr:    ":8803",
		Handler: router,
	}

	go func() {
		logger.Info("Job service HTTP server starting on :8803")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("HTTP server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down job service...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Info("Job service stopped")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	r.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
		})
	})

	return r
}

func initCronJobs() {
	c := cron.New()

	ctx := context.Background()

	_, err := c.AddFunc("@every 10s", func() {
		service.Crontab().Test(ctx)
		fmt.Println("Test cron job executed")
	})
	if err != nil {
		logger.Errorf("注册测试任务失败: %v", err)
	}

	_, err = c.AddFunc("@every 1m", func() {
		service.Crontab().HandleAttributionTokens(ctx)
	})
	if err != nil {
		logger.Errorf("注册处理归因Token任务失败: %v", err)
	}

	// 每天凌晨 01:05 执行每日聚合统计
	_, err = c.AddFunc("5 1 * * *", func() {
		service.Crontab().AggregateDailyStats(ctx)
	})
	if err != nil {
		logger.Errorf("注册每日聚合统计任务失败: %v", err)
	}

	c.Start()
	logger.Info("Cron jobs started")
}
