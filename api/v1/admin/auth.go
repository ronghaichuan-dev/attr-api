package admin

import (
g "github.com/gogf/gf/v2/frame/g"
)

// CaptchaReq 获取验证码请求参数结构体
type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"系统管理" summary:"获取验证码"`
}

// CaptchaRes 获取验证码响应参数结构体
type CaptchaRes struct {
	Id     string `json:"id" dc:"验证码ID"`     // 验证码ID
	Base64 string `json:"base64" dc:"验证码图片base64编码"` // 验证码图片base64编码
}

// LoginReq 登录请求参数结构体
type LoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"系统管理" summary:"用户登录"`
	Username string `json:"username" form:"username" v:"required#用户名不能为空" dc:"用户名，必填"` // 用户名，必填
	Password string `json:"password" form:"password" v:"required#密码不能为空" dc:"密码，必填"`   // 密码，必填
	Captcha  string `json:"captcha" form:"captcha" dc:"验证码，可选"`  // 验证码，可选
	CaptchaId string `json:"captcha_id" form:"captcha_id" dc:"验证码ID，可选"` // 验证码ID，可选
}

// LoginRes 登录响应参数结构体
type LoginRes struct {
	Token string `json:"token" dc:"登录token"` // 登录token
	ExpireTime int64 `json:"expire_time" dc:"token过期时间戳"` // token过期时间戳
}

// LogoutReq 登出请求参数结构体
type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"系统管理" summary:"用户登出"`
}

// LogoutRes 登出响应参数结构体
type LogoutRes struct {
	Success bool `json:"success" dc:"是否成功"` // 是否成功
}
