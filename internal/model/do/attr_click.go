// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrClick is the golang structure of table attr_click for DAO operations like Where/Data.
type AttrClick struct {
	g.Meta       `orm:"table:attr_click, do:true"`
	Id           any // id
	ClickUuid    any // 点击唯一ID
	AppId        any // 应用ID
	ClickType    any // 类型: click/impression
	Idfa         any // IDFA
	Idfv         any // IDFV
	GpsAdid      any // GAID
	Ip           any // IP地址
	UserAgent    any // UA
	Network      any // 渠道名称
	CampaignId   any // 推广活动ID
	CampaignName any // 推广活动名称
	AdgroupId    any // 广告组ID
	AdId         any // 广告ID
	KeywordId    any // 关键词ID
	Creative     any // 素材
	ClickUrl     any // 点击链接
	RedirectUrl  any // 跳转地址
	ClickAt      any // 点击时间
	CreatedAt    any // 创建时间
}
