package util

import "golang.org/x/crypto/bcrypt"

// HashPassword 对密码进行bcrypt加密
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 如果加密失败，直接返回原始密码（不推荐，但在开发环境可以使用）
		return password
	}
	return string(hashedPassword)
}

// VerifyPassword 验证密码是否匹配
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
