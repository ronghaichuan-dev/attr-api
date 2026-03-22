// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrAppSubscriptionsDao is the data access object for the table attr_app_subscriptions.
type AttrAppSubscriptionsDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  AttrAppSubscriptionsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// AttrAppSubscriptionsColumns defines and stores column names for the table attr_app_subscriptions.
type AttrAppSubscriptionsColumns struct {
	Id                    string // id
	Environment           string // 环境
	OrignialTransactionId string // 原始交易ID
	Rsid                  string // 用户设备ID
	Appid                 string // appID
	ProductId             string // 产品ID
	Status                string // 订阅状态 1-自动续订服务已激活 2-自动续订服务已过期自动  3-自动续订服务目前处于计费重试期 4-自动续订服务目前处于账单宽限期 5-自动续订订阅已取消
	AutoRenewStatus       string // 自动续费状态 1-启用 2-禁用
	IsTrial               string // 是否试订 1-是 2-否
	IsPaid                string // 是否付费 1-是 2-否
	LastEventAt           string // 上次事件时间
	ExpiresReason         string // 过期原因 1-无 2-订阅在计费重试期结束后过期 3-订阅因价格上涨过期 4-订阅因产品不可售过期 5-用户自愿取消订阅导致过期
	ExpiresAt             string // 过期时间
	CreatedAt             string // 创建时间
	UpdatedAt             string // 更新时间
}

// attrAppSubscriptionsColumns holds the columns for the table attr_app_subscriptions.
var attrAppSubscriptionsColumns = AttrAppSubscriptionsColumns{
	Id:                    "id",
	Environment:           "environment",
	OrignialTransactionId: "orignial_transaction_id",
	Rsid:                  "rsid",
	Appid:                 "appid",
	ProductId:             "product_id",
	Status:                "status",
	AutoRenewStatus:       "auto_renew_status",
	IsTrial:               "is_trial",
	IsPaid:                "is_paid",
	LastEventAt:           "last_event_at",
	ExpiresReason:         "expires_reason",
	ExpiresAt:             "expires_at",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewAttrAppSubscriptionsDao creates and returns a new DAO object for table data access.
func NewAttrAppSubscriptionsDao(handlers ...gdb.ModelHandler) *AttrAppSubscriptionsDao {
	return &AttrAppSubscriptionsDao{
		group:    "default",
		table:    "attr_app_subscriptions",
		columns:  attrAppSubscriptionsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrAppSubscriptionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrAppSubscriptionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrAppSubscriptionsDao) Columns() AttrAppSubscriptionsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrAppSubscriptionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrAppSubscriptionsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrAppSubscriptionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
