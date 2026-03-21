package api

import (
	g "github.com/gogf/gf/v2/frame/g"
)

// EchoReq echo测试接口请求结构体
type EchoReq struct {
	g.Meta  `path:"/echo" method:"get" tags:"公共API" summary:"echo测试接口"`
	Message string `json:"message" form:"message" v:"required#消息不能为空" dc:"要echo的消息"`
}

// EchoRes echo测试接口响应结构体
type EchoRes struct {
	Message string `json:"message" dc:"echo的消息"`
}
