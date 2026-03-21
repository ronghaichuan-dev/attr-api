package admin

// HelloReq 演示接口请求参数结构体
type HelloReq struct {
	Name string `json:"name" form:"name" dc:"姓名，可选"` // 姓名，可选
}

// HelloRes 演示接口响应参数结构体
type HelloRes struct {
	Message string `json:"message" dc:"欢迎消息"` // 欢迎消息
}
