package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type NotificationReq struct {
	g.Meta        `path:"/notification/callback" method:"post" tags:"服务器通知回调" summary:"苹果服务器通知回调"`
	SignedPayload string `json:"signedPayload" dc:"苹果服务器通知的签名有效负载"`
}

type NotificationRes struct {
	Success bool `json:"success"`
}

type NotificationInfo struct {
	NotificationType string `json:"notificationType"`
	Subtype          string `json:"subtype"`
	NotificationUUID string `json:"notificationUUID"`
	Data             Data   `json:"data"`
	Version          string `json:"version"`
	SignedDate       int64  `json:"signedDate"`
}

type Data struct {
	AppAppleId            int64  `json:"appAppleId"`
	BundleId              string `json:"bundleId"`
	BundleVersion         string `json:"bundleVersion"`
	Environment           string `json:"environment"`
	SignedTransactionInfo string `json:"signedTransactionInfo"`
	SignedRenewalInfo     string `json:"signedRenewalInfo"`
	Status                int    `json:"status"`
}

// EmptyReq 空请求
type EmptyReq struct {
	g.Meta `path:"/notification/memory" method:"get" tags:"内存监控" summary:"获取内存监控数据"`
}
