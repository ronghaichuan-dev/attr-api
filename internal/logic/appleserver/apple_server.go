package appleserver

import (
	"context"
	"encoding/json"
	"god-help-service/api/v1/api"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/net/proxy"
)

type sAppleServer struct {
}

func init() {
	service.RegisterAppleServer(NewAppleServer())
}

func NewAppleServer() *sAppleServer {
	return &sAppleServer{}
}

// AccountInfo 账户信息结构体
type AccountInfo struct {
	Socket5Proxy string `json:"socket5_proxy"`
	// 其他字段
}

func (s *sAppleServer) GetAttributionInfo(ctx context.Context, token string, proxyURL, username, password string) (*api.AppleAttributionInfoResponse, string, error) {
	logger.Infof("开始调用苹果归因接口 token:%s, proxyURL:%s, username:%s", token, proxyURL, username)
	url := "https://api-adservices.apple.com/api/v1/"

	var client *http.Client

	if proxyURL != "" {
		// 去掉代理URL中的前缀（如果有）
		proxyURL = strings.TrimPrefix(proxyURL, "socks5h://")
		proxyURL = strings.TrimPrefix(proxyURL, "socks5://")
		proxyURL = strings.TrimPrefix(proxyURL, "http://")
		proxyURL = strings.TrimPrefix(proxyURL, "https://")

		var auth *proxy.Auth
		if username != "" && password != "" {
			auth = &proxy.Auth{
				User:     username,
				Password: password,
			}
		}

		socks5Dialer, err := proxy.SOCKS5("tcp", proxyURL, auth, proxy.Direct)
		if err != nil {
			logger.Errorf("创建SOCKS5代理失败: %v", err)
			return nil, "", err
		}

		client = &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return socks5Dialer.Dial(network, addr)
				},
			},
			Timeout: 30 * time.Second,
		}
		logger.Infof("使用SOCKS5代理: %s", proxyURL)
	} else {
		client = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	payload := strings.NewReader(token)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		logger.Errorf("创建请求失败: %v", err)
		return nil, "", err
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("X-Apple-App-Store-Region", "US")

	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("发送请求失败: %v", err)
		return nil, "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("读取响应失败: %v", err)
		return nil, "", err
	}

	logger.Infof("苹果归因接口响应 statusCode:%d headers:%v", resp.StatusCode, resp.Header)

	if resp.StatusCode != 200 {
		logger.Errorf("苹果归因接口返回错误 statusCode:%d response:%s", resp.StatusCode, string(body))
		return nil, "", gerror.Newf("苹果归因接口错误 statusCode:%d response:%s", resp.StatusCode, string(body))
	}

	logger.Infof("苹果归因接口调用成功: %s", string(body))
	result := &api.AppleAttributionInfoResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, "", err
	}
	return result, string(body), nil
}
