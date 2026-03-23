// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrClickDao is the data access object for the table attr_click.
type AttrClickDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AttrClickColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AttrClickColumns defines and stores column names for the table attr_click.
type AttrClickColumns struct {
	Id           string // id
	ClickUuid    string // 点击唯一ID
	AppId        string // 应用ID
	ClickType    string // 类型: click/impression
	Idfa         string // IDFA
	Idfv         string // IDFV
	GpsAdid      string // GAID
	Ip           string // IP地址
	UserAgent    string // UA
	Network      string // 渠道名称
	CampaignId   string // 推广活动ID
	CampaignName string // 推广活动名称
	AdgroupId    string // 广告组ID
	AdId         string // 广告ID
	KeywordId    string // 关键词ID
	Creative     string // 素材
	ClickUrl     string // 点击链接
	RedirectUrl  string // 跳转地址
	ClickAt      string // 点击时间
	CreatedAt    string // 创建时间
}

// attrClickColumns holds the columns for the table attr_click.
var attrClickColumns = AttrClickColumns{
	Id:           "id",
	ClickUuid:    "click_uuid",
	AppId:        "app_id",
	ClickType:    "click_type",
	Idfa:         "idfa",
	Idfv:         "idfv",
	GpsAdid:      "gps_adid",
	Ip:           "ip",
	UserAgent:    "user_agent",
	Network:      "network",
	CampaignId:   "campaign_id",
	CampaignName: "campaign_name",
	AdgroupId:    "adgroup_id",
	AdId:         "ad_id",
	KeywordId:    "keyword_id",
	Creative:     "creative",
	ClickUrl:     "click_url",
	RedirectUrl:  "redirect_url",
	ClickAt:      "click_at",
	CreatedAt:    "created_at",
}

// NewAttrClickDao creates and returns a new DAO object for table data access.
func NewAttrClickDao(handlers ...gdb.ModelHandler) *AttrClickDao {
	return &AttrClickDao{
		group:    "default",
		table:    "attr_click",
		columns:  attrClickColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrClickDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrClickDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrClickDao) Columns() AttrClickColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrClickDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrClickDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrClickDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
