// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrInstall is the golang structure for table attr_install.
type AttrInstall struct {
	Id                int64  `json:"id"                orm:"id"                  description:"id"`                                                         // id
	AttrUuid          string `json:"attrUuid"          orm:"attr_uuid"           description:"归因事件唯一ID"`                                                   // 归因事件唯一ID
	Environment       string `json:"environment"       orm:"environment"         description:"环境"`                                                         // 环境
	AppId             string `json:"appId"             orm:"app_id"              description:"应用ID"`                                                       // 应用ID
	AppToken          string `json:"appToken"          orm:"app_token"           description:"应用token"`                                                    // 应用token
	AppVersion        string `json:"appVersion"        orm:"app_version"         description:"应用版本"`                                                       // 应用版本
	Rsid              string `json:"rsid"              orm:"rsid"                description:"设备ID"`                                                       // 设备ID
	Idfa              string `json:"idfa"              orm:"idfa"                description:"广告标识符"`                                                      // 广告标识符
	Idfv              string `json:"idfv"              orm:"idfv"                description:"应用开发商标识符"`                                                   // 应用开发商标识符
	GpsAdid           string `json:"gpsAdid"           orm:"gps_adid"            description:"谷歌市场广告ID"`                                                   // 谷歌市场广告ID
	AndroidId         string `json:"androidId"         orm:"android_id"          description:"安卓ID"`                                                       // 安卓ID
	OsName            string `json:"osName"            orm:"os_name"             description:"苹果系统名称"`                                                     // 苹果系统名称
	OsVersion         string `json:"osVersion"         orm:"os_version"          description:"苹果系统版本"`                                                     // 苹果系统版本
	Language          string `json:"language"          orm:"language"            description:"语言"`                                                         // 语言
	Country           string `json:"country"           orm:"country"             description:"国家"`                                                         // 国家
	Region            string `json:"region"            orm:"region"              description:"州/省"`                                                        // 州/省
	City              string `json:"city"              orm:"city"                description:"城市"`                                                         // 城市
	Tracker           string `json:"tracker"           orm:"tracker"             description:"设备当前归因链接来源 adServices|adjust|branch|appsFlyer"`              // 设备当前归因链接来源 adServices|adjust|branch|appsFlyer
	TrackerToken      string `json:"trackerToken"      orm:"tracker_token"       description:"设备当前归因链接的识别码"`                                               // 设备当前归因链接的识别码
	TrackerUid        string `json:"trackerUid"        orm:"tracker_uid"         description:"归因唯一ID"`                                                     // 归因唯一ID
	TrackerVersion    string `json:"trackerVersion"    orm:"tracker_version"     description:"归因版本"`                                                       // 归因版本
	TrackerNetwork    string `json:"trackerNetwork"    orm:"tracker_network"     description:"归因设备当前归因渠道的名称 tiktok|facebook|google|twitter"`               // 归因设备当前归因渠道的名称 tiktok|facebook|google|twitter
	TrackerChannel    string `json:"trackerChannel"    orm:"tracker_channel"     description:"归因渠道"`                                                       // 归因渠道
	TrackerCampaignId string `json:"trackerCampaignId" orm:"tracker_campaign_id" description:"归因设备当前归因推广活动ID"`                                             // 归因设备当前归因推广活动ID
	TrackerAdgroupId  string `json:"trackerAdgroupId"  orm:"tracker_adgroup_id"  description:"归因广告组ID"`                                                    // 归因广告组ID
	TrackerAdId       string `json:"trackerAdId"       orm:"tracker_ad_id"       description:"归因广告ID"`                                                     // 归因广告ID
	TrackerKeywordId  string `json:"trackerKeywordId"  orm:"tracker_keyword_id"  description:"归因关键词ID"`                                                    // 归因关键词ID
	TrackerAgency     string `json:"trackerAgency"     orm:"tracker_agency"      description:"归因渠道代理"`                                                     // 归因渠道代理
	Network           string `json:"network"           orm:"network"             description:"渠道的名称"`                                                      // 渠道的名称
	Channel           string `json:"channel"           orm:"channel"             description:"归因渠道"`                                                       // 归因渠道
	CampaignId        string `json:"campaignId"        orm:"campaign_id"         description:"设备当前归因推广活动ID"`                                               // 设备当前归因推广活动ID
	AdgroupId         string `json:"adgroupId"         orm:"adgroup_id"          description:"广告组ID"`                                                      // 广告组ID
	AdId              string `json:"adId"              orm:"ad_id"               description:"广告ID"`                                                       // 广告ID
	InstallAt         int64  `json:"installAt"         orm:"install_at"          description:"安装时间戳"`                                                      // 安装时间戳
	SentAt            int64  `json:"sentAt"            orm:"sent_at"             description:"发送时间"`                                                       // 发送时间
	IsHandleToken     int    `json:"isHandleToken"     orm:"is_handle_token"     description:"token是否已调用苹果归因接口 1-已调用 2-未调用"`                               // token是否已调用苹果归因接口 1-已调用 2-未调用
	AdServicesToken   string `json:"adServicesToken"   orm:"ad_services_token"   description:"苹果归因token"`                                                  // 苹果归因token
	IsFirstInstall    int    `json:"isFirstInstall"    orm:"is_first_install"    description:"是否首次安装"`                                                     // 是否首次安装
	TokenResponseText string `json:"tokenResponseText" orm:"token_response_text" description:"解析token原始数据"`                                                // 解析token原始数据
	MatchType         string `json:"matchType"         orm:"match_type"          description:"匹配方式: device_id/referrer/probabilistic/tracker/ad_services"` // 匹配方式: device_id/referrer/probabilistic/tracker/ad_services
	MatchConfidence   string `json:"matchConfidence"   orm:"match_confidence"    description:"匹配置信度: high/medium/low"`                                     // 匹配置信度: high/medium/low
	ClickId           int64  `json:"clickId"           orm:"click_id"            description:"关联的点击记录ID"`                                                  // 关联的点击记录ID
	ClickToInstall    int64  `json:"clickToInstall"    orm:"click_to_install"    description:"点击到安装的时间间隔（秒）"`                                              // 点击到安装的时间间隔（秒）
	Ip                string `json:"ip"                orm:"ip"                  description:"安装时IP"`                                                      // 安装时IP
	UserAgent         string `json:"userAgent"         orm:"user_agent"          description:"安装时UA"`                                                      // 安装时UA
}
