package app

import "github.com/gogf/gf/v2/frame/g"

type AttributionReportReq struct {
	g.Meta `path:"/attribution/report" method:"post" tags:"归因事件" summary:"APP归因事件上报"`
	*Attribution
}

type Attribution struct {
	Environment       string `json:"environment" binding:"required"`
	AppId             string `json:"app_id" binding:"required"`
	AppToken          string `json:"app_token" binding:"required"`
	AppVersion        string `json:"app_version"`
	Rsid              string `json:"rsid" binding:"required"`
	Idfa              string `json:"idfa"`
	Idfv              string `json:"idfv"`
	GpsAdid           string `json:"gps_adid"`
	AndroidId         string `json:"android_id"`
	OsName            string `json:"os_name"`
	OsVersion         string `json:"os_version"`
	Language          string `json:"language"`
	Country           string `json:"country"`
	Region            string `json:"region"`
	City              string `json:"city"`
	AdServicesToken   string `json:"ad_services_token"`
	Tracker           string `json:"tracker"`
	TrackerToken      string `json:"tracker_token"`
	TrackerUid        string `json:"tracker_uid"`
	TrackerVersion    string `json:"tracker_version"`
	TrackerNetwork    string `json:"tracker_network"`
	TrackerChannel    string `json:"tracker_channel"`
	TrackerCampaignId string `json:"tracker_campaign_id"`
	TrackerAdgroupId  string `json:"tracker_adgroup_id"`
	TrackerAdId       string `json:"tracker_ad_id"`
	TrackerKeywordId  string `json:"tracker_keyword_id"`
	TrackerAgency     string `json:"tracker_agency"`
	InstallAt         int    `json:"install_at"`
	SentAt            int    `json:"sent_at"`
	IsHandleToken     int    `json:"is_handle_token"`
	AttrUuid          string `json:"attr_uuid" binding:"required"`
	IsFirstInstall    int    `json:"is_first_install"`
}

type AttributionReportRes struct {
}

type AccountInfo struct {
	Socks5   string `json:"socks5" dc:"socks5"`
	Username string `json:"username"`
	Password string `json:"password"`
}
