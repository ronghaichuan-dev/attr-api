// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrRecordDao is the data access object for the table attr_record.
type AttrRecordDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AttrRecordColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AttrRecordColumns defines and stores column names for the table attr_record.
type AttrRecordColumns struct {
	Id                string // id
	AttrUuid          string // 归因事件唯一ID
	Environment       string // 环境
	AppId             string // 应用ID
	AppToken          string // 应用token
	AppVersion        string // 应用版本
	Uuid              string // 用户ID
	Idfa              string // 广告标识符
	Idfv              string // 应用开发商标识符
	GpsAdid           string // 谷歌市场广告ID
	AndroidId         string // 安卓ID
	OsName            string // 苹果系统名称
	OsVersion         string // 苹果系统版本
	Language          string // 语言
	Country           string // 国家
	Tracker           string // 设备当前归因链接来源 adServices|adjust|branch|appsFlyer
	TrackerToken      string // 设备当前归因链接的识别码
	TrackerUid        string // 归因唯一ID
	TrackerVersion    string // 归因版本
	TrackerNetwork    string // 归因设备当前归因渠道的名称 tiktok|facebook|google|twitter
	TrackerChannel    string // 归因渠道
	TrackerCampaignId string // 归因设备当前归因推广活动ID
	TrackerAdgroupId  string // 归因广告组ID
	TrackerAdId       string // 归因广告ID
	TrackerKeywordId  string // 归因关键词ID
	TrackerAgency     string // 归因渠道代理
	Network           string // 渠道的名称
	Channel           string // 归因渠道
	CampaignId        string // 设备当前归因推广活动ID
	AdgroupId         string // 广告组ID
	AdId              string // 广告ID
	InstallAt         string // 安装时间戳
	SentAt            string // 发送时间
	IsHandleToken     string // token是否已调用苹果归因接口 1-已调用 2-未调用
	AdServicesToken   string // 苹果归因token
}

// attrRecordColumns holds the columns for the table attr_record.
var attrRecordColumns = AttrRecordColumns{
	Id:                "id",
	AttrUuid:          "attr_uuid",
	Environment:       "environment",
	AppId:             "app_id",
	AppToken:          "app_token",
	AppVersion:        "app_version",
	Uuid:              "uuid",
	Idfa:              "idfa",
	Idfv:              "idfv",
	GpsAdid:           "gps_adid",
	AndroidId:         "android_id",
	OsName:            "os_name",
	OsVersion:         "os_version",
	Language:          "language",
	Country:           "country",
	Tracker:           "tracker",
	TrackerToken:      "tracker_token",
	TrackerUid:        "tracker_uid",
	TrackerVersion:    "tracker_version",
	TrackerNetwork:    "tracker_network",
	TrackerChannel:    "tracker_channel",
	TrackerCampaignId: "tracker_campaign_id",
	TrackerAdgroupId:  "tracker_adgroup_id",
	TrackerAdId:       "tracker_ad_id",
	TrackerKeywordId:  "tracker_keyword_id",
	TrackerAgency:     "tracker_agency",
	Network:           "network",
	Channel:           "channel",
	CampaignId:        "campaign_id",
	AdgroupId:         "adgroup_id",
	AdId:              "ad_id",
	InstallAt:         "install_at",
	SentAt:            "sent_at",
	IsHandleToken:     "is_handle_token",
	AdServicesToken:   "ad_services_token",
}

// NewAttrRecordDao creates and returns a new DAO object for table data access.
func NewAttrRecordDao(handlers ...gdb.ModelHandler) *AttrRecordDao {
	return &AttrRecordDao{
		group:    "default",
		table:    "attr_record",
		columns:  attrRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrRecordDao) Columns() AttrRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrRecordDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AttrRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
