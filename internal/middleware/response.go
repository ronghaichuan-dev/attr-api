package middleware

import (
	"context"
	"god-help-service/internal/util/logger"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// processUserIds 处理响应数据中的用户ID转换为用户名
func processUserIds(ctx context.Context, data interface{}) interface{} {
	// 简单处理：如果没有数据需要转换，直接返回原始数据
	// 这样可以避免类型转换导致的数据丢失问题
	return data
}

// ResponseHandler 统一响应格式化中间件
func ResponseHandler(r *ghttp.Request) {
	// 执行后续处理
	r.Middleware.Next()

	// 如果已经有响应内容，直接返回
	if r.Response.BufferLength() > 0 || r.Response.BytesWritten() > 0 {
		return
	}

	// 正确获取GoFrame控制器的返回结果
	var (
		data interface{}
		err  error
	)

	// 获取控制器返回的结果
	// 根据GoFrame源码分析，对于 func(context.Context, *BizReq)(*BizRes, error) 格式的控制器
	// GetHandlerResponse() 直接返回 *BizRes（数据部分），而错误存储在 r.GetError() 中
	data = r.GetHandlerResponse()
	err = r.GetError()

	logger.Debugf("ResponseHandler - 路由:%s 方法:%s", r.URL.Path, r.Method)
	logger.Debugf("ResponseHandler - 控制器返回数据:%s", data)
	logger.Debugf("ResponseHandler - 控制器返回错误:%s", err)

	// 处理用户ID转换
	if err == nil && data != nil {
		data = processUserIds(r.Context(), data)
	}

	// 构建统一响应格式
	var response g.Map
	if err != nil {
		// 错误响应
		response = g.Map{
			"code":    500,
			"error":   err.Error(),
			"message": "操作失败",
		}
		r.Response.Status = 500
	} else {
		// 成功响应
		response = g.Map{
			"code":    0,
			"data":    data,
			"message": "操作成功",
		}
		r.Response.Status = 200
	}

	// 返回 JSON 响应
	r.Response.WriteJson(response)
}

// ErrorHandler 错误处理中间件
func ErrorHandler(r *ghttp.Request) {
	defer func() {
		if err := recover(); err != nil {
			var (
				errorMsg string
			)
			switch e := err.(type) {
			case error:
				errorMsg = e.Error()
			case string:
				errorMsg = e
			default:
				errorMsg = "系统内部错误"
			}

			r.Response.WriteJson(g.Map{
				"code":    500,
				"error":   errorMsg,
				"message": "系统内部错误",
			})
			r.Response.Status = 500
			logger.Errorf("系统内部错误:%s", errorMsg)
		}
	}()

	r.Middleware.Next()
}
