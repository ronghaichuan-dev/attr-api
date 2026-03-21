package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result 统一响应结构体
type Result struct {
	Status int         `json:"status"` // 状态码，200表示成功，500表示失败
	Msg    string      `json:"msg"`    // 提示信息
	Data   interface{} `json:"data"`   // 响应数据，成功时返回具体数据，失败时返回nil
}

type ResultNoData struct {
	Status int    `json:"status"` // 状态码，200表示成功，500表示失败
	Msg    string `json:"msg"`    // 提示信息
}

// Success 返回成功响应
func Success(c *gin.Context, msg string, data ...interface{}) {
	if data == nil {
		c.JSON(http.StatusOK, ResultNoData{
			Status: 200,
			Msg:    msg,
		})
	} else {
		c.JSON(http.StatusOK, Result{
			Status: 200,
			Msg:    msg,
			Data:   data,
		})
	}
}

// Fail 返回失败响应
func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusOK, ResultNoData{
		Status: 500,
		Msg:    err.Error(),
	})
}

// Response 返回响应
func Response(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Status: code,
		Msg:    msg,
		Data:   data,
	})
}
