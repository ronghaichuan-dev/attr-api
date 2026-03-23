// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrClick is the golang structure for table attr_click.
type AttrClick struct {
	Id           int64  `json:"id"           orm:"id"            description:"id"`                   // id
	ClickUuid    string `json:"clickUuid"    orm:"click_uuid"    description:"点击唯一ID"`               // 点击唯一ID
	AppId        string `json:"appId"        orm:"app_id"        description:"应用ID"`                 // 应用ID
	ClickType    string `json:"clickType"    orm:"click_type"    description:"类型: click/impression"` // 类型: click/impression
	Idfa         string `json:"idfa"         orm:"idfa"          description:"IDFA"`                 // IDFA
	Idfv         string `json:"idfv"         orm:"idfv"          description:"IDFV"`                 // IDFV
	GpsAdid      string `json:"gpsAdid"      orm:"gps_adid"      description:"GAID"`                 // GAID
	Ip           string `json:"ip"           orm:"ip"            description:"IP地址"`                 // IP地址
	UserAgent    string `json:"userAgent"    orm:"user_agent"    description:"UA"`                   // UA
	Network      string `json:"network"      orm:"network"       description:"渠道名称"`                 // 渠道名称
	CampaignId   string `json:"campaignId"   orm:"campaign_id"   description:"推广活动ID"`               // 推广活动ID
	CampaignName string `json:"campaignName" orm:"campaign_name" description:"推广活动名称"`               // 推广活动名称
	AdgroupId    string `json:"adgroupId"    orm:"adgroup_id"    description:"广告组ID"`                // 广告组ID
	AdId         string `json:"adId"         orm:"ad_id"         description:"广告ID"`                 // 广告ID
	KeywordId    string `json:"keywordId"    orm:"keyword_id"    description:"关键词ID"`                // 关键词ID
	Creative     string `json:"creative"     orm:"creative"      description:"素材"`                   // 素材
	ClickUrl     string `json:"clickUrl"     orm:"click_url"     description:"点击链接"`                 // 点击链接
	RedirectUrl  string `json:"redirectUrl"  orm:"redirect_url"  description:"跳转地址"`                 // 跳转地址
	ClickAt      int64  `json:"clickAt"      orm:"click_at"      description:"点击时间"`                 // 点击时间
	CreatedAt    int64  `json:"createdAt"    orm:"created_at"    description:"创建时间"`                 // 创建时间
}
