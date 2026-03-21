package admin

import (
	"context"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"time"

	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/util"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type AuthController struct{}

func (c *AuthController) Login(ctx context.Context, req *adminApi.LoginReq) (*adminApi.LoginRes, error) {
	user, err := service.System().GetUserByUsername(ctx, req.Username)
	if err != nil {
		logger.Errorf("查询用户信息失败: %v", err)
		return nil, gerror.New("登录失败")
	}

	if user == nil {
		return nil, gerror.New("账号不存在")
	}
	hashedPassword := util.HashPassword(req.Password)
	if !util.VerifyPassword(hashedPassword, req.Password) {
		return nil, gerror.New("密码错误")
	}

	token, err := util.GenerateJWT(req.Username, user.Id)
	if err != nil {
		logger.Errorf("生成JWT Token失败: %v", err)
		return nil, gerror.New("登录失败")
	}

	var jwtConfig struct {
		Expire string `json:"expire" v:"required"`
	}
	if err = g.Cfg().MustGet(ctx, "jwt").Scan(&jwtConfig); err != nil {
		logger.Errorf("获取JWT配置失败: %v", err)
		return nil, gerror.New("登录失败")
	}

	expireDuration, err := time.ParseDuration(jwtConfig.Expire)
	if err != nil {
		logger.Errorf("解析JWT过期时间失败: %v", err)
		return nil, gerror.New("登录失败")
	}

	expireTime := time.Now().Add(expireDuration).Unix()

	if err := util.SetToken(token, req.Username, expireDuration); err != nil {
		logger.Errorf("存储Token到Redis失败: %v", err)
		return nil, gerror.New("登录失败")
	}

	res := &adminApi.LoginRes{
		Token:      token,
		ExpireTime: expireTime,
	}
	logger.Debugf("Login方法返回结果: %v", res)
	return res, nil
}

func (c *AuthController) Logout(ctx context.Context, req *adminApi.LogoutReq) (*adminApi.LogoutRes, error) {
	token := g.RequestFromCtx(ctx).Header.Get("Authorization")
	if token == "" {
		return &adminApi.LogoutRes{Success: true}, nil
	}

	if tokenPrefixIndex := len("Bearer "); len(token) > tokenPrefixIndex && token[:tokenPrefixIndex] == "Bearer " {
		token = token[tokenPrefixIndex:]
	}

	claims, err := util.ParseJWT(token)
	if err == nil {
		if err = util.DeleteUserOnline(ctx, claims.Username); err != nil {
			logger.Errorf("删除用户在线状态失败: %v", err)
		}
	}

	if err = util.DeleteToken(ctx, token); err != nil {
		logger.Errorf("删除token失败: %v", err)
	}

	return &adminApi.LogoutRes{Success: true}, nil
}

func (c *AuthController) Captcha(ctx context.Context, req *adminApi.CaptchaReq) (*adminApi.CaptchaRes, error) {
	id, code := util.GenerateCaptcha()
	util.CaptchaStoreHandler.Set(id, code)
	return &adminApi.CaptchaRes{
		Id:     id,
		Base64: code,
	}, nil
}
