package appleapi

import (
	"time"
)

// Payload 表示JWT的payload部分
type Payload struct {
	Issuer   string `json:"iss"`
	Audience string `json:"aud"`
	Expires  int64  `json:"exp"`
	IssuedAt int64  `json:"iat"`
	BundleID string `json:"bid"`
}

// NewPayload 创建一个新的Payload实例
func NewPayload(issuerID, bundleID string) *Payload {
	now := time.Now().Unix()
	return &Payload{
		Issuer:   issuerID,
		Audience: "appstoreconnect-v1",
		Expires:  now + 1800, // 30分钟过期
		IssuedAt: now,
		BundleID: bundleID,
	}
}

// Refresh 更新过期时间
func (p *Payload) Refresh() {
	now := time.Now().Unix()
	p.Expires = now + 1800
	p.IssuedAt = now
}
