// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrDailyStatsDao is the data access object for the table attr_daily_stats.
type AttrDailyStatsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  AttrDailyStatsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// AttrDailyStatsColumns defines and stores column names for the table attr_daily_stats.
type AttrDailyStatsColumns struct {
	Id                 string //
	StatDate           string // 统计日期 YYYY-MM-DD
	AppId              string // 应用ID
	Country            string // 国家
	TrackerNetwork     string // 归因渠道
	CampaignId         string // 推广活动ID
	InstallCount       string // 安装量
	TrialCount         string // 试用量
	SubscribeCount     string // 订阅量（付费）
	RenewCount         string // 续订量
	RefundCount        string // 退款量
	Revenue            string // 收入（分）
	RefundAmount       string // 退款金额（分）
	NetRevenue         string // 净收入（分）
	InstallToTrialRate string // 安装转试用率%
	TrialToPaidRate    string // 试用转付费率%
	CreatedAt          string //
	UpdatedAt          string //
}

// attrDailyStatsColumns holds the columns for the table attr_daily_stats.
var attrDailyStatsColumns = AttrDailyStatsColumns{
	Id:                 "id",
	StatDate:           "stat_date",
	AppId:              "app_id",
	Country:            "country",
	TrackerNetwork:     "tracker_network",
	CampaignId:         "campaign_id",
	InstallCount:       "install_count",
	TrialCount:         "trial_count",
	SubscribeCount:     "subscribe_count",
	RenewCount:         "renew_count",
	RefundCount:        "refund_count",
	Revenue:            "revenue",
	RefundAmount:       "refund_amount",
	NetRevenue:         "net_revenue",
	InstallToTrialRate: "install_to_trial_rate",
	TrialToPaidRate:    "trial_to_paid_rate",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewAttrDailyStatsDao creates and returns a new DAO object for table data access.
func NewAttrDailyStatsDao(handlers ...gdb.ModelHandler) *AttrDailyStatsDao {
	return &AttrDailyStatsDao{
		group:    "default",
		table:    "attr_daily_stats",
		columns:  attrDailyStatsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrDailyStatsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrDailyStatsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrDailyStatsDao) Columns() AttrDailyStatsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrDailyStatsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrDailyStatsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrDailyStatsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
