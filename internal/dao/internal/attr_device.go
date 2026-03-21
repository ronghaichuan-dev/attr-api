// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrDeviceDao is the data access object for the table attr_device.
type AttrDeviceDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AttrDeviceColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AttrDeviceColumns defines and stores column names for the table attr_device.
type AttrDeviceColumns struct {
	Id                 string // id
	Uuid               string // 设备ID
	Appid              string // 应用ID
	AttrSubscriptionId string // 归因订阅ID
	Country            string // 国家
	IsRefund           string // 是否退款 1-是 2-否
	IsRenew            string // 是否续订
	RenewCount         string // 续订次数
	DeductionCount     string // 扣费次数
	CreatedAt          string // 创建时间
	LastInstallAt      string // 最后安装时间
	LastTrialAt        string // 最后试用时间
	LastSubscribeAt    string // 最后订阅时间
	LastRenewAt        string // 最后续费时间
	LastRefundAt       string // 最后退款时间
}

// attrDeviceColumns holds the columns for the table attr_device.
var attrDeviceColumns = AttrDeviceColumns{
	Id:                 "id",
	Uuid:               "uuid",
	Appid:              "appid",
	AttrSubscriptionId: "attr_subscription_id",
	Country:            "country",
	IsRefund:           "is_refund",
	IsRenew:            "is_renew",
	RenewCount:         "renew_count",
	DeductionCount:     "deduction_count",
	CreatedAt:          "created_at",
	LastInstallAt:      "last_install_at",
	LastTrialAt:        "last_trial_at",
	LastSubscribeAt:    "last_subscribe_at",
	LastRenewAt:        "last_renew_at",
	LastRefundAt:       "last_refund_at",
}

// NewAttrDeviceDao creates and returns a new DAO object for table data access.
func NewAttrDeviceDao(handlers ...gdb.ModelHandler) *AttrDeviceDao {
	return &AttrDeviceDao{
		group:    "default",
		table:    "attr_device",
		columns:  attrDeviceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrDeviceDao) Columns() AttrDeviceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrDeviceDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
