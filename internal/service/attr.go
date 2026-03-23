package service

import (
	"context"
	adminApi "god-help-service/api/v1/admin"
	"god-help-service/api/v1/app"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service/appleapi"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClickRecord 点击记录参数
type ClickRecord struct {
	AppId        string
	ClickType    string
	Network      string
	CampaignId   string
	CampaignName string
	AdgroupId    string
	AdId         string
	KeywordId    string
	Creative     string
	Idfa         string
	Idfv         string
	GpsAdid      string
	Ip           string
	UserAgent    string
	ClickUrl     string
	RedirectUrl  string
}

type (
	IAttr interface {
		Test()
		// HandleAttribution 处理归因信息
		HandleAttribution(ctx context.Context, attr *app.Attribution) error
		HandleTrialFreeSubscribe(ctx context.Context, environment string, appid string, uuid string, payload *app.EventParam, resp *appleapi.ResponseBodyV2) error
		UpdateAppSubscriptionFieldByOriginalTransactionId(ctx context.Context, tx gdb.TX, originalTransactionId string, updateData g.Map) (int64, error)
		HandleSubscribe(ctx context.Context, environment string, appid string, uuid string, payload *app.Payload) error
		SaveNotificationEvent(ctx context.Context, payload string, body *appleapi.ResponseBodyV2) error
		UpdateAttrDeviceField(ctx context.Context, tx gdb.TX, uuid string, updateData g.Map) error
		GetAttrDevice(ctx context.Context, uuid []string) (map[string]struct{}, error)
		CreateAttrDeviceOrUpdate(ctx context.Context, appid string, uuid string, country string) error
		// CreateAttrDeviceOrUpdateWithAttribution 创建或更新设备归因记录（含归因信息）
		CreateAttrDeviceOrUpdateWithAttribution(ctx context.Context, attr *app.Attribution, installId int64) error
		// UpdateAttrDeviceSubscription 更新设备订阅相关字段
		UpdateAttrDeviceSubscription(ctx context.Context, rsid string, appid string, updateData g.Map) error
		// RecordClick 记录广告点击/展示
		RecordClick(ctx context.Context, record *ClickRecord) error
		// SendPostback 发送归因回传
		SendPostback(ctx context.Context, postback *entity.AttrPostback) error
		// GetEventByCode 根据事件编码获取事件详情
		GetEventByCode(ctx context.Context, eventCode string) (*entity.AttrEvent, error)
		// GetEventDropdownList 获取事件下拉选项列表
		GetEventDropdownList(ctx context.Context) (*adminApi.EventDropdownRes, error)
		// GetEventList 获取事件列表
		GetEventList(ctx context.Context, req *adminApi.EventListReq) (*adminApi.EventListRes, error)
		// GetEventDetailById 根据ID获取事件详情
		GetEventDetailById(ctx context.Context, eventId int64) (*adminApi.EventDetailItem, error)
		// CreateEvent 创建事件
		CreateEvent(ctx context.Context, req *adminApi.EventCreateReq) (int64, error)
		// UpdateEvent 更新事件
		UpdateEvent(ctx context.Context, req *adminApi.EventUpdateReq) error
		// DeleteEvent 删除事件（软删除）
		DeleteEvent(ctx context.Context, eventId int64) error
		// GetAppEventLogList 获取事件日志列表
		GetAppEventLogList(ctx context.Context, req *adminApi.AppEventLogListReq) (*adminApi.AppEventLogListRes, error)
		// GetAppEventLogById 根据ID获取事件日志详情
		GetAppEventLogById(ctx context.Context, logId int64) (*entity.AppEventLogCustom, error)
		SaveAttrSubscriptionTransaction(ctx context.Context, notificationType string, v2 *appleapi.ResponseBodyV2) error
		UpdateAttrSubscriptionTransaction(ctx context.Context, tx gdb.TX, originalTxId string, updateData gdb.Map) (int64, error)
	}
)

var (
	localAttr IAttr
)

func Attr() IAttr {
	if localAttr == nil {
		panic("implement not found for interface IAttr, forgot register?")
	}
	return localAttr
}

func RegisterAttr(i IAttr) {
	localAttr = i
}
