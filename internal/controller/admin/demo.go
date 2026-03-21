package admin

import (
	"context"

	adminApi "god-help-service/api/v1/admin"
)

type ControllerDemo struct{}

// Hello 演示接口
func (c *ControllerDemo) Hello(ctx context.Context, req *adminApi.HelloReq) (*adminApi.HelloRes, error) {
	name := req.Name
	if name == "" {
		name = "World"
	}
	return &adminApi.HelloRes{
		Message: "Hello, " + name + "!",
	}, nil
}
