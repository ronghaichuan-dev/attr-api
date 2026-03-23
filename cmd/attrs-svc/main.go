package main

import (
	"context"
	"encoding/json"
	"fmt"
	"god-help-service/cmd/notification-svc/router"
	"god-help-service/internal/dao"
	_ "god-help-service/internal/logic"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"god-help-service/internal/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

var (
	wg           sync.WaitGroup
	shutdownChan chan struct{}
)

func main() {
	shutdownChan = make(chan struct{})

	engine := router.InitRouter()
	pprof.Register(engine)
	port := util.GetConfigString(context.Background(), "server.port", "")

	server := &http.Server{
		Addr:    port,
		Handler: engine,
	}

	go func() {
		fmt.Println("启动HTTP服务器", "端口:", port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	time.Sleep(1 * time.Second)

	redisAddr := util.GetConfigString(context.Background(), "redis.addr", "localhost:6379")
	redisPassword := util.GetConfigString(context.Background(), "redis.password", "")
	redisDB := util.GetConfigInt(context.Background(), "redis.db", 0)
	util.InitRedisClient(redisAddr, redisPassword, redisDB)
	fmt.Println("Redis客户端初始化完成")

	// 加载系统应用信息到Redis
	loadSystemAppsToRedis()

	service.Queue().StartNotificationConsumer()
	service.Queue().StartEventConsumer()
	service.Queue().StartAttributionConsumer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("正在关闭服务器...")

	close(shutdownChan)
	wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("服务器强制关闭:", err)
	} else {
		fmt.Println("HTTP服务器已关闭")
	}

	fmt.Println("服务器已优雅退出")
}

// loadSystemAppsToRedis 加载系统应用信息到Redis
func loadSystemAppsToRedis() {
	ctx := context.Background()
	fmt.Println("开始加载系统应用信息到Redis...")

	// 从数据库获取所有系统应用
	var apps []entity.SystemApps
	err := dao.SystemApps.Ctx(ctx).Scan(&apps)
	if err != nil {
		fmt.Printf("获取系统应用信息失败: %v\n", err)
		return
	}

	redisClient := util.GetRedisClient()
	count := 0

	for _, app := range apps {
		// 将应用信息转换为JSON
		appJSON, err := json.Marshal(app)
		if err != nil {
			fmt.Printf("序列化应用信息失败: %v\n", err)
			continue
		}

		key := "system_app:" + app.Appid
		err = redisClient.Set(ctx, key, appJSON, 0).Err()
		if err != nil {
			fmt.Printf("存储应用信息到Redis失败: %v\n", err)
			continue
		}

		count++
	}

	fmt.Printf("系统应用信息加载完成，共加载 %d 个应用\n", count)
}
