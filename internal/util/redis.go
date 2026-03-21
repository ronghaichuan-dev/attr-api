package util

import (
	"context"
	"god-help-service/internal/util/logger"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis客户端实例
var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

// InitRedisClient 初始化Redis客户端
func InitRedisClient(addr, password string, db int) {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
	})
}

// GetRedisClient 获取Redis客户端实例
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		// 默认连接配置
		InitRedisClient("localhost:6379", "", 0)
	}
	return redisClient
}

// SetToken 存储Token到Redis
func SetToken(token string, username string, expireDuration time.Duration) error {
	ctx := context.Background()
	// 存储Token和用户名关联
	_, err := GetRedisClient().Set(ctx, token, username, expireDuration).Result()
	if err != nil {
		logger.Errorf("存储Token到Redis失败:%s", err)
		return err
	}

	logger.Infof("Token存储到Redis成功:%s", token)
	return nil
}

// GetToken 获取Token对应的用户名
func GetToken(ctx context.Context, token string) (string, error) {
	val, err := GetRedisClient().Get(ctx, token).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		logger.Errorf("从Redis获取Token失败:%s", err.Error())
		return "", err
	}
	return val, nil
}

// DeleteToken 从Redis删除Token
func DeleteToken(ctx context.Context, token string) error {
	_, err := GetRedisClient().Del(ctx, token).Result()
	if err != nil {
		logger.Errorf("从Redis删除Token失败:%s", err.Error())
		return err
	}

	logger.Infof("Token从Redis删除成功:%s", token)
	return nil
}

// CheckTokenExists 检查Token是否存在于Redis
func CheckTokenExists(ctx context.Context, token string) bool {
	_, err := GetRedisClient().Get(ctx, token).Result()
	return err == nil
}

// DeleteUserOnline 删除用户在线状态
func DeleteUserOnline(ctx context.Context, username string) error {
	_, err := GetRedisClient().Del(ctx, "user:"+username).Result()
	if err != nil {
		logger.Errorf("删除用户在线状态失败:%s", err.Error())
		return err
	}

	return nil
}
