package util

import (
	"context"
	"god-help-service/internal/util/logger"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

// JWT claims 结构体
type JWTClaims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成JWT Token
func GenerateJWT(username string, userID int) (string, error) {
	// 从配置文件获取JWT配置
	var jwtConfig struct {
		Secret string `json:"secret" v:"required"`
		Expire string `json:"expire" v:"required"`
	}
	if err := g.Cfg().MustGet(context.Background(), "jwt").Scan(&jwtConfig); err != nil {
		logger.Errorf("获取JWT配置失败:%s", err)
		return "", err
	}

	// 解析过期时间
	expireDuration, err := time.ParseDuration(jwtConfig.Expire)
	if err != nil {
		logger.Errorf("解析JWT过期时间失败:%s", err)
		return "", err
	}

	// 创建Claims
	claims := JWTClaims{
		Username: username,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "god-help-service",
			Subject:   "user-auth",
		},
	}

	// 创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		logger.Errorf("生成JWT Token失败:", err)
		return "", err
	}

	return signedToken, nil
}

// ParseJWT 解析JWT Token
func ParseJWT(tokenString string) (*JWTClaims, error) {
	// 从配置文件获取JWT配置
	var jwtConfig struct {
		Secret string `json:"secret" v:"required"`
	}
	if err := g.Cfg().MustGet(context.Background(), "jwt").Scan(&jwtConfig); err != nil {
		logger.Errorf("获取JWT配置失败:", err)
		return nil, err
	}

	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		logger.Errorf("解析JWT Token失败:", err)
		return nil, err
	}

	// 验证Token
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// IsJWTValid 验证JWT Token是否有效
func IsJWTValid(tokenString string) bool {
	_, err := ParseJWT(tokenString)
	return err == nil
}
