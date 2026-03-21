package admin

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	adminApi "god-help-service/api/v1/admin"
	"god-help-service/internal/util/logger"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadController 文件上传控制器
type UploadController struct{}

// UploadImage 图片上传
func (c *UploadController) UploadImage(ctx context.Context, req *adminApi.UploadImageReq) (*adminApi.UploadImageRes, error) {
	// 获取上传的文件
	file := g.RequestFromCtx(ctx).GetUploadFile("file")
	if file == nil {
		return nil, gerror.New("请选择要上传的图片")
	}

	// 检查文件类型
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}
	ext := filepath.Ext(file.Filename)
	ext = strings.ToLower(ext)
	isAllowed := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return nil, gerror.New("只允许上传图片格式文件（jpg、jpeg、png、gif、bmp、webp）")
	}

	// 检查文件大小（限制10MB）
	const maxSize = 10 * 1024 * 1024 // 10MB
	fileSize := file.Size
	if fileSize > maxSize {
		return nil, gerror.New("图片大小不能超过10MB")
	}

	// 创建上传目录
	uploadDir := fmt.Sprintf("upload/images/%s", gtime.Now().Format("20060102"))
	if !gfile.Exists(uploadDir) {
		if err := gfile.Mkdir(uploadDir); err != nil {
			logger.Errorf("创建上传目录失败: %v", err)
			return nil, gerror.New("创建上传目录失败")
		}
	}

	// 生成唯一文件名
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, fileName)

	// 保存文件
	if _, err := file.Save(filePath); err != nil {
		logger.Errorf("保存图片失败: %v", err)
		return nil, gerror.New("保存图片失败")
	}

	// 生成访问URL
	accessUrl := fmt.Sprintf("/%s", filePath)

	return &adminApi.UploadImageRes{
		Url: accessUrl,
	}, nil
}
