package appleapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPRequest 处理HTTP请求
type HTTPRequest struct {
	client *http.Client
}

// NewHTTPRequest 创建一个新的HTTPRequest实例
func NewHTTPRequest() *HTTPRequest {
	return &HTTPRequest{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// PerformRequest 执行HTTP请求
func (hr *HTTPRequest) PerformRequest(ctx context.Context, request AbstractRequest) ([]byte, error) {
	// 创建JWT令牌
	token, err := request.GenerateToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token: %w", err)
	}

	// 构建URL
	url, err := request.ComposeURL()
	if err != nil {
		return nil, fmt.Errorf("failed to compose URL: %w", err)
	}

	// 创建请求体
	var body io.Reader
	if request.GetHTTPMethod() == "POST" || request.GetHTTPMethod() == "PUT" {
		if request.GetBody() != nil {
			content, err := request.GetBody().GetEncodedContent()
			if err != nil {
				return nil, fmt.Errorf("failed to get encoded content: %w", err)
			}
			body = bytes.NewBuffer(content)
		}
	}

	// 创建HTTP请求
	httpRequest, err := http.NewRequestWithContext(ctx, request.GetHTTPMethod(), url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// 添加授权头
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// 添加Content-Type头
	if request.GetHTTPMethod() == "POST" || request.GetHTTPMethod() == "PUT" {
		if request.GetBody() != nil {
			httpRequest.Header.Set("Content-Type", request.GetBody().GetContentType())
		}
	}

	// 发送请求
	response, err := hr.client.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer response.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// 检查状态码
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return nil, &HTTPRequestFailedError{
			Method:       request.GetHTTPMethod(),
			URL:          url,
			StatusCode:   response.StatusCode,
			ResponseBody: string(responseBody),
			Message:      fmt.Sprintf("HTTP request failed with status code: %d", response.StatusCode),
		}
	}

	return responseBody, nil
}

// HTTPRequestFailedError HTTP请求失败错误
type HTTPRequestFailedError struct {
	Method       string
	URL          string
	StatusCode   int
	ResponseBody string
	Message      string
}

// Error 实现error接口
func (e *HTTPRequestFailedError) Error() string {
	return fmt.Sprintf("%s: %s %s", e.Message, e.Method, e.URL)
}

// HTTPRequestAbortedError HTTP请求中断错误
type HTTPRequestAbortedError struct {
	Message string
	Err     error
}

// Error 实现error接口
func (e *HTTPRequestAbortedError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap 实现errors.Unwrap接口
func (e *HTTPRequestAbortedError) Unwrap() error {
	return e.Err
}

// AbstractRequest HTTP请求接口
type AbstractRequest interface {
	// GetKey 获取密钥
	GetKey() *Key

	// GetPayload 获取负载
	GetPayload() *Payload

	// GetHTTPMethod 获取HTTP方法
	GetHTTPMethod() string

	// GetBody 获取请求体
	GetBody() AbstractRequestBody

	// ComposeURL 组合URL
	ComposeURL() (string, error)

	// GenerateToken 生成JWT令牌
	GenerateToken(ctx context.Context) (string, error)

	// SetURLVars 设置URL变量
	SetURLVars(vars map[string]string)
}

// AbstractRequestBody 请求体接口
type AbstractRequestBody interface {
	// GetContentType 获取内容类型
	GetContentType() string

	// GetEncodedContent 获取编码后的内容
	GetEncodedContent() ([]byte, error)
}
