package api

import "github.com/gogf/gf/v2/frame/g"

// TrackingClickReq 点击追踪请求
type TrackingClickReq struct {
	g.Meta       `path:"/tracking/click" method:"get" tags:"点击追踪" summary:"广告点击追踪"`
	AppId        string `json:"app_id" form:"app_id" binding:"required"`
	Network      string `json:"network" form:"network"`
	CampaignId   string `json:"campaign_id" form:"campaign_id"`
	CampaignName string `json:"campaign_name" form:"campaign_name"`
	AdgroupId    string `json:"adgroup_id" form:"adgroup_id"`
	AdId         string `json:"ad_id" form:"ad_id"`
	KeywordId    string `json:"keyword_id" form:"keyword_id"`
	Creative     string `json:"creative" form:"creative"`
	Idfa         string `json:"idfa" form:"idfa"`
	Idfv         string `json:"idfv" form:"idfv"`
	GpsAdid      string `json:"gps_adid" form:"gps_adid"`
	RedirectUrl  string `json:"redirect_url" form:"redirect_url" binding:"required"`
}

// TrackingClickRes 点击追踪响应
type TrackingClickRes struct {
}

// TrackingImpressionReq 展示追踪请求
type TrackingImpressionReq struct {
	g.Meta       `path:"/tracking/impression" method:"get" tags:"展示追踪" summary:"广告展示追踪"`
	AppId        string `json:"app_id" form:"app_id" binding:"required"`
	Network      string `json:"network" form:"network"`
	CampaignId   string `json:"campaign_id" form:"campaign_id"`
	CampaignName string `json:"campaign_name" form:"campaign_name"`
	AdgroupId    string `json:"adgroup_id" form:"adgroup_id"`
	AdId         string `json:"ad_id" form:"ad_id"`
	KeywordId    string `json:"keyword_id" form:"keyword_id"`
	Creative     string `json:"creative" form:"creative"`
	Idfa         string `json:"idfa" form:"idfa"`
	GpsAdid      string `json:"gps_adid" form:"gps_adid"`
}

// TrackingImpressionRes 展示追踪响应
type TrackingImpressionRes struct {
}
