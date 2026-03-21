package app

import (
	g "github.com/gogf/gf/v2/frame/g"
)

// EventReportReq 移动端事件上报通知请求参数结构体
type EventReportReq struct {
	g.Meta      `path:"/event/report" method:"post" tags:"订阅事件" summary:"APP订阅事件上报通知"`
	Environment string     `json:"environment" dc:"环境"`
	AppVersion  string     `json:"app_version" dc:"app版本"`
	Payload     []*Payload `json:"payload"     dc:"请求负载"`
	Appid       string     `json:"appid"       dc:"应用ID，必填" binding:"required" ` // 应用ID，必填
	Country     string     `json:"country"     dc:"国家，可选"`                       // 国家，可选
	Region      string     `json:"region" dc:"州/省"`
	City        string     `json:"city" dc:"城市"`
	Rsid        string     `json:"rsid"        dc:"设备ID，可选" binding:"required" ` // 设备ID
	CreatedAt   int64      `json:"created_at"  dc:"事件时间"`
	EventUuid   string     `json:"event_uuid" dc:"事件唯一ID"`
	SentAt      int64      `json:"sent_at"     dc:"事件上报时间"`
}

type Payload struct {
	EventCode string `json:"event_code"`
	Params    Params `json:"params"`
	EventUuid string `json:"event_uuid"  dc:"事件唯一标识"`
}

type EventParam struct {
	Environment           string `json:"environment" dc:"环境"`
	AppVersion            string `json:"app_version" dc:"app版本"`
	Appid                 string `json:"appid"       dc:"应用ID，必填" binding:"required" ` // 应用ID，必填
	Country               string `json:"country"     dc:"国家，可选"`                       // 国家，可选
	EventCode             string `json:"event_code" dc:"事件编码"`
	EventUuid             string `json:"event_uuid"  dc:"事件唯一标识"`
	TransactionId         string `json:"transaction_id"`
	OriginalTransactionId string `json:"original_transaction_id"`
	Rsid                  string `json:"rsid"        dc:"设备ID，可选" binding:"required" ` // 设备ID，可选
	CreatedAt             int64  `json:"created_at"  dc:"事件时间"`
	SentAt                int64  `json:"sent_at"     dc:"事件上报时间"`
}

type Params struct {
	TransactionId         string `json:"transaction_id"`
	OriginalTransactionId string `json:"original_transaction_id"`
	Exts                  string `json:"exts"`
}

// EventReportRes 事件上报通知响应参数结构体
type EventReportRes struct {
	Success bool   `json:"success" dc:"处理是否成功"` // 处理是否成功
	Message string `json:"message" dc:"处理信息"`   // 处理信息
}
