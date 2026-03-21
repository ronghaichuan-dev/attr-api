// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AttrNotifications is the golang structure of table attr_notifications for DAO operations like Where/Data.
type AttrNotifications struct {
	g.Meta        `orm:"table:attr_notifications, do:true"`
	Id            any         // id
	Uuid          any         // 用户ID
	Token         any         // 用户Token
	NoticeType    any         // 通知类型
	TxId          any         // 事务ID
	RenewalStatus any         // 续费状态
	Sku           any         // 商品sku
	NoticeAt      *gtime.Time // 通知时间
	CreatedAt     *gtime.Time // 创建时间
}
