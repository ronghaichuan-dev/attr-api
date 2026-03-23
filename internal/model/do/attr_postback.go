// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrPostback is the golang structure of table attr_postback for DAO operations like Where/Data.
type AttrPostback struct {
	g.Meta                `orm:"table:attr_postback, do:true"`
	Id                    any // id
	AppId                 any // 应用ID
	PostbackType          any // 回传类型: install/event/reengagement
	Network               any // 渠道
	OriginalTransactionId any // 原始交易ID
	EventName             any // 事件名
	PostbackUrl           any // 回传URL
	ResponseCode          any // 响应码
	ResponseBody          any // 响应内容
	Status                any // 状态: 1-成功 2-失败 3-重试中
	RetryCount            any // 重试次数
	CreatedAt             any // 创建时间
}
