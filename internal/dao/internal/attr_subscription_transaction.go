// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrSubscriptionTransactionDao is the data access object for the table attr_subscription_transaction.
type AttrSubscriptionTransactionDao struct {
	table    string                             // table is the underlying table name of the DAO.
	group    string                             // group is the database configuration group name of the current DAO.
	columns  AttrSubscriptionTransactionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler                 // handlers for customized model modification.
}

// AttrSubscriptionTransactionColumns defines and stores column names for the table attr_subscription_transaction.
type AttrSubscriptionTransactionColumns struct {
	Id                    string // id
	TransactionType       string // 交易类型 RENEW / REFUND / TRIAL
	Envirment             string // 环境
	AppVersion            string // 应用版本
	Appid                 string // 应用ID
	OriginalTransactionId string // 原始交易ID
	TransactionId         string // 子交易ID
	InAppOwnership        string // 是否为用户购买 PURCHASED-购买 FAMILY_SHARED-家庭分享
	Uuid                  string // 用户ID
	ProductId             string // sku
	Price                 string // 订阅金额
	Currency              string // 币种
	SubscribeStatus       string // 订阅状态
	PurchaseAt            string // 购买时间
	CreatedAt             string // 创建时间
}

// attrSubscriptionTransactionColumns holds the columns for the table attr_subscription_transaction.
var attrSubscriptionTransactionColumns = AttrSubscriptionTransactionColumns{
	Id:                    "id",
	TransactionType:       "transaction_type",
	Envirment:             "envirment",
	AppVersion:            "app_version",
	Appid:                 "appid",
	OriginalTransactionId: "original_transaction_id",
	TransactionId:         "transaction_id",
	InAppOwnership:        "in_app_ownership",
	Uuid:                  "uuid",
	ProductId:             "product_id",
	Price:                 "price",
	Currency:              "currency",
	SubscribeStatus:       "subscribe_status",
	PurchaseAt:            "purchase_at",
	CreatedAt:             "created_at",
}

// NewAttrSubscriptionTransactionDao creates and returns a new DAO object for table data access.
func NewAttrSubscriptionTransactionDao(handlers ...gdb.ModelHandler) *AttrSubscriptionTransactionDao {
	return &AttrSubscriptionTransactionDao{
		group:    "default",
		table:    "attr_subscription_transaction",
		columns:  attrSubscriptionTransactionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrSubscriptionTransactionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrSubscriptionTransactionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrSubscriptionTransactionDao) Columns() AttrSubscriptionTransactionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrSubscriptionTransactionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrSubscriptionTransactionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrSubscriptionTransactionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
