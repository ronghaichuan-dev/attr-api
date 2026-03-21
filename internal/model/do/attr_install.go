// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AttrInstall is the golang structure of table attr_install for DAO operations like Where/Data.
type AttrInstall struct {
	g.Meta            `orm:"table:attr_install, do:true"`
	Id                any // id
	AttrUuid          any // 归因事件唯一ID
	Environment       any // 环境
	AppId             any // 应用ID
	AppToken          any // 应用token
	AppVersion        any // 应用版本
	Rsid              any // 设备ID
	Idfa              any // 广告标识符
	Idfv              any // 应用开发商标识符
	GpsAdid           any // 谷歌市场广告ID
	AndroidId         any // 安卓ID
	OsName            any // 苹果系统名称
	OsVersion         any // 苹果系统版本
	Language          any // 语言
	Country           any // 国家
	Region            any // 州/省
	City              any // 城市
	Tracker           any // 设备当前归因链接来源 adServices|adjust|branch|appsFlyer
	TrackerToken      any // 设备当前归因链接的识别码
	TrackerUid        any // 归因唯一ID
	TrackerVersion    any // 归因版本
	TrackerNetwork    any // 归因设备当前归因渠道的名称 tiktok|facebook|google|twitter
	TrackerChannel    any // 归因渠道
	TrackerCampaignId any // 归因设备当前归因推广活动ID
	TrackerAdgroupId  any // 归因广告组ID
	TrackerAdId       any // 归因广告ID
	TrackerKeywordId  any // 归因关键词ID
	TrackerAgency     any // 归因渠道代理
	Network           any // 渠道的名称
	Channel           any // 归因渠道
	CampaignId        any // 设备当前归因推广活动ID
	AdgroupId         any // 广告组ID
	AdId              any // 广告ID
	InstallAt         any // 安装时间戳
	SentAt            any // 发送时间
	IsHandleToken     any // token是否已调用苹果归因接口 1-已调用 2-未调用
	AdServicesToken   any // 苹果归因token
	IsFirstInstall    any // 是否首次安装
	TokenResponseText any // 解析token原始数据
}
