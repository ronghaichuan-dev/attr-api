package request

import (
	"context"
	"fmt"
	"regexp"

	"god-help-service/internal/service/appleapi"
)

// HTTPMethod 定义HTTP方法常量
const (
	HTTPMethodGET  = "GET"
	HTTPMethodPOST = "POST"
	HTTPMethodPUT  = "PUT"
)

// AbstractRequest 抽象请求结构体
type AbstractRequest struct {
	Key         *appleapi.Key
	Payload     *appleapi.Payload
	QueryParams AbstractRequestQueryParams
	Body        appleapi.AbstractRequestBody
	URLVars     map[string]string
}

// NewAbstractRequest 创建一个新的抽象请求实例
func NewAbstractRequest(key *appleapi.Key, payload *appleapi.Payload, queryParams AbstractRequestQueryParams, body appleapi.AbstractRequestBody) *AbstractRequest {
	return &AbstractRequest{
		Key:         key,
		Payload:     payload,
		QueryParams: queryParams,
		Body:        body,
		URLVars:     make(map[string]string),
	}
}

// GetKey 获取密钥
func (r *AbstractRequest) GetKey() *appleapi.Key {
	return r.Key
}

// GetPayload 获取负载
func (r *AbstractRequest) GetPayload() *appleapi.Payload {
	return r.Payload
}

// GetQueryParams 获取查询参数
func (r *AbstractRequest) GetQueryParams() AbstractRequestQueryParams {
	return r.QueryParams
}

// GetBody 获取请求体
func (r *AbstractRequest) GetBody() appleapi.AbstractRequestBody {
	return r.Body
}

// GetURLVars 获取URL变量
func (r *AbstractRequest) GetURLVars() map[string]string {
	return r.URLVars
}

// SetURLVars 设置URL变量
func (r *AbstractRequest) SetURLVars(vars map[string]string) {
	for k, v := range vars {
		r.URLVars[k] = v
	}
}

// ComposeURL 组合URL
func (r *AbstractRequest) ComposeURL() (string, error) {
	// 获取URL模式
	urlPattern, err := r.GetURLPattern()
	if err != nil {
		return "", err
	}

	// 替换URL变量
	urlString := regexp.MustCompile(`{([^}]+)}`).ReplaceAllStringFunc(urlPattern, func(match string) string {
		key := match[1 : len(match)-1]
		if value, exists := r.URLVars[key]; exists {
			return value
		}
		return match
	})

	// 添加查询参数
	if r.QueryParams != nil {
		queryString := r.QueryParams.GetQueryString()
		if queryString != "" {
			urlString += "?" + queryString
		}
	}

	return urlString, nil
}

// GenerateToken 生成JWT令牌
func (r *AbstractRequest) GenerateToken(ctx context.Context) (string, error) {
	// 这里需要使用JWT库生成令牌
	// 目前使用之前创建的util.JWT
	// 但是需要注意，我们需要使用ES256算法，而不是RSA
	return "", fmt.Errorf("ES256 JWT token generation not implemented yet")
}

// GetHTTPMethod 获取HTTP方法
func (r *AbstractRequest) GetHTTPMethod() string {
	// 默认使用GET方法
	return HTTPMethodGET
}

// GetURLPattern 获取URL模式
func (r *AbstractRequest) GetURLPattern() (string, error) {
	return "", fmt.Errorf("URL pattern not implemented")
}

// AbstractRequestQueryParams 抽象请求查询参数接口
type AbstractRequestQueryParams interface {
	// GetQueryString 获取查询字符串
	GetQueryString() string
}

// GetTransactionHistoryRequest 获取交易历史请求
type GetTransactionHistoryRequest struct {
	*AbstractRequest
	TransactionID string
}

// NewGetTransactionHistoryRequest 创建一个新的获取交易历史请求实例
func NewGetTransactionHistoryRequest(key *appleapi.Key, payload *appleapi.Payload, queryParams AbstractRequestQueryParams, body appleapi.AbstractRequestBody, transactionID string) *GetTransactionHistoryRequest {
	return &GetTransactionHistoryRequest{
		AbstractRequest: NewAbstractRequest(key, payload, queryParams, body),
		TransactionID:   transactionID,
	}
}

// GetHTTPMethod 获取HTTP方法
func (r *GetTransactionHistoryRequest) GetHTTPMethod() string {
	return HTTPMethodGET
}

// GetURLPattern 获取URL模式
func (r *GetTransactionHistoryRequest) GetURLPattern() (string, error) {
	return "{baseUrl}/v1/history/{transactionId}", nil
}
