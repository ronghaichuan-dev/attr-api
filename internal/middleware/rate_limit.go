package middleware

import (
	"context"
	"errors"
	"fmt"
	"god-help-service/internal/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// IPBasedRateLimit 基于IP的速率限制中间件
func IPBasedRateLimit(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit:ip:%s", ip)

		ctx := context.Background()
		var count int

		// 获取当前计数
		val, err := util.GetRedisClient().Get(ctx, key).Result()
		if err != nil {
			// 如果key不存在，继续执行
			if errors.Is(err, redis.Nil) {
				count = 0
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  "Internal Server Error",
				})
				c.Abort()
				return
			}
		} else {
			// 尝试将字符串转换为整数
			count = 0
			fmt.Sscanf(val, "%d", &count)
		}

		if count >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": 429,
				"msg":  "Too Many Requests",
			})
			c.Abort()
			return
		}

		// 增加计数并设置过期时间
		if count == 0 {
			util.GetRedisClient().Set(ctx, key, 1, window)
		} else {
			util.GetRedisClient().Incr(ctx, key)
		}

		c.Next()
	}
}

// APIRateLimit 为API路由设置的速率限制
func APIRateLimit(ctx *gin.Context) {
	// API路由限制：每分钟60次请求
	IPBasedRateLimit(60, time.Minute)
}
