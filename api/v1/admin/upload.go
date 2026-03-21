package admin

import (
	g "github.com/gogf/gf/v2/frame/g"
)

// UploadImageReq 图片上传请求参数结构体
type UploadImageReq struct {
	g.Meta `path:"/upload/image" method:"post" tags:"文件上传" summary:"图片上传"`
	File   *g.Var `json:"file" form:"file" v:"required#请选择要上传的图片" dc:"图片文件，必填"` // 图片文件，必填
}

// UploadImageRes 图片上传响应参数结构体
type UploadImageRes struct {
	Url string `json:"url" dc:"图片访问URL"` // 图片访问URL
}
