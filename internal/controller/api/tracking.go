package api

import (
	"god-help-service/internal/consts"
	"god-help-service/internal/service"
	"god-help-service/internal/util/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TrackingController 点击追踪控制器
type TrackingController struct{}

// Click 处理广告点击追踪
func (c *TrackingController) Click(ctx *gin.Context) {
	appId := ctx.Query("app_id")
	redirectUrl := ctx.Query("redirect_url")

	if appId == "" || redirectUrl == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "app_id and redirect_url are required"})
		return
	}

	// 异步记录点击
	go func() {
		err := service.Attr().RecordClick(ctx.Request.Context(), &service.ClickRecord{
			AppId:        appId,
			ClickType:    consts.ClickTypeClick,
			Network:      ctx.Query("network"),
			CampaignId:   ctx.Query("campaign_id"),
			CampaignName: ctx.Query("campaign_name"),
			AdgroupId:    ctx.Query("adgroup_id"),
			AdId:         ctx.Query("ad_id"),
			KeywordId:    ctx.Query("keyword_id"),
			Creative:     ctx.Query("creative"),
			Idfa:         ctx.Query("idfa"),
			Idfv:         ctx.Query("idfv"),
			GpsAdid:      ctx.Query("gps_adid"),
			Ip:           ctx.ClientIP(),
			UserAgent:    ctx.Request.UserAgent(),
			RedirectUrl:  redirectUrl,
		})
		if err != nil {
			logger.Errorf("记录点击失败: %v", err)
		}
	}()

	// 302 重定向
	ctx.Redirect(http.StatusFound, redirectUrl)
}

// Impression 处理广告展示追踪
func (c *TrackingController) Impression(ctx *gin.Context) {
	appId := ctx.Query("app_id")
	if appId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "app_id is required"})
		return
	}

	// 异步记录展示
	go func() {
		err := service.Attr().RecordClick(ctx.Request.Context(), &service.ClickRecord{
			AppId:        appId,
			ClickType:    consts.ClickTypeImpression,
			Network:      ctx.Query("network"),
			CampaignId:   ctx.Query("campaign_id"),
			CampaignName: ctx.Query("campaign_name"),
			AdgroupId:    ctx.Query("adgroup_id"),
			AdId:         ctx.Query("ad_id"),
			KeywordId:    ctx.Query("keyword_id"),
			Creative:     ctx.Query("creative"),
			Idfa:         ctx.Query("idfa"),
			GpsAdid:      ctx.Query("gps_adid"),
			Ip:           ctx.ClientIP(),
			UserAgent:    ctx.Request.UserAgent(),
		})
		if err != nil {
			logger.Errorf("记录展示失败: %v", err)
		}
	}()

	// 返回 1x1 透明像素
	ctx.Data(http.StatusOK, "image/gif", transparentPixel)
}

// 1x1 透明 GIF 像素
var transparentPixel = []byte{
	0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x01, 0x00,
	0x01, 0x00, 0x80, 0x00, 0x00, 0xff, 0xff, 0xff,
	0x00, 0x00, 0x00, 0x21, 0xf9, 0x04, 0x01, 0x00,
	0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x00,
	0x01, 0x00, 0x01, 0x00, 0x00, 0x02, 0x02, 0x44,
	0x01, 0x00, 0x3b,
}
