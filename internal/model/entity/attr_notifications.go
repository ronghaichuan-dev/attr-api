// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrNotifications is the golang structure for table attr_notifications.
type AttrNotifications struct {
	Id            int64       `json:"id"            orm:"id"              description:"id"`      // id
	Uuid          string      `json:"uuid"          orm:"uuid"            description:"用户ID"`    // 用户ID
	Token         string      `json:"token"         orm:"token"           description:"用户Token"` // 用户Token
	NoticeType    int         `json:"noticeType"    orm:"notice_type"     description:"通知类型"`    // 通知类型
	TxId          string      `json:"txId"          orm:"tx_id"           description:"事务ID"`    // 事务ID
	RenewalStatus int         `json:"renewalStatus" orm:"renewal _status" description:"续费状态"`    // 续费状态
	Sku           string      `json:"sku"           orm:"sku"             description:"商品sku"`   // 商品sku
	NoticeAt      *gtime.Time `json:"noticeAt"      orm:"notice_at"       description:"通知时间"`    // 通知时间
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`    // 创建时间
}
