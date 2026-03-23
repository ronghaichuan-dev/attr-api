package attr

import (
	"context"
	"fmt"
	"god-help-service/internal/consts"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/util/logger"
	"io"
	"math"
	"net/http"
	"time"
)

// SendPostback 发送归因回传
func (s *sAttr) SendPostback(ctx context.Context, postback *entity.AttrPostback) error {
	if postback.PostbackUrl == "" {
		// 没有配置回传URL，只记录不发送
		postback.Status = consts.PostbackStatusSuccess
		postback.CreatedAt = time.Now().Unix()
		_, err := dao.AttrPostback.Ctx(ctx).Data(postback).Insert()
		return err
	}

	// 发送HTTP请求
	var lastErr error
	maxRetries := 5

	for i := 0; i <= maxRetries; i++ {
		if i > 0 {
			// 指数退避：2^(i-1) 秒
			backoff := time.Duration(math.Pow(2, float64(i-1))) * time.Second
			time.Sleep(backoff)
		}

		statusCode, respBody, err := s.doPostbackRequest(postback.PostbackUrl)
		if err != nil {
			lastErr = err
			postback.RetryCount = i + 1
			continue
		}

		postback.ResponseCode = statusCode
		postback.ResponseBody = respBody

		if statusCode >= 200 && statusCode < 300 {
			// 成功
			postback.Status = consts.PostbackStatusSuccess
			postback.RetryCount = i
			postback.CreatedAt = time.Now().Unix()
			_, insertErr := dao.AttrPostback.Ctx(ctx).Data(postback).Insert()
			return insertErr
		}

		if statusCode >= 500 {
			// 5xx 错误，重试
			lastErr = fmt.Errorf("postback returned %d", statusCode)
			postback.RetryCount = i + 1
			continue
		}

		// 4xx 错误，不重试
		postback.Status = consts.PostbackStatusFailed
		postback.RetryCount = i
		postback.CreatedAt = time.Now().Unix()
		_, insertErr := dao.AttrPostback.Ctx(ctx).Data(postback).Insert()
		if insertErr != nil {
			logger.Errorf("保存回传记录失败: %v", insertErr)
		}
		return fmt.Errorf("postback returned %d", statusCode)
	}

	// 所有重试都失败
	postback.Status = consts.PostbackStatusFailed
	postback.CreatedAt = time.Now().Unix()
	_, insertErr := dao.AttrPostback.Ctx(ctx).Data(postback).Insert()
	if insertErr != nil {
		logger.Errorf("保存回传记录失败: %v", insertErr)
	}
	return fmt.Errorf("postback failed after %d retries: %v", maxRetries, lastErr)
}

// doPostbackRequest 执行回传HTTP请求
func (s *sAttr) doPostbackRequest(url string) (int, string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}

	// 限制响应体长度
	respBody := string(body)
	if len(respBody) > 1024 {
		respBody = respBody[:1024]
	}

	return resp.StatusCode, respBody, nil
}
