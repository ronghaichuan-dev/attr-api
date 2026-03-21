// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrTransactionOrderDao is the data access object for the table attr_transaction_order.
type AttrTransactionOrderDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  AttrTransactionOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// AttrTransactionOrderColumns defines and stores column names for the table attr_transaction_order.
type AttrTransactionOrderColumns struct {
	Id               string // id
	AppId            string // 应用ID
	TransactionId    string // 交易ID
	SubTransactionId string // 子交易ID
	Uuid             string // 用户ID
	SkuId            string // sku
	Amount           string // 订阅金额
	SubscribeStatus  string // 订阅状态
	CreatedAt        string // 创建时间
}

// attrTransactionOrderColumns holds the columns for the table attr_transaction_order.
var attrTransactionOrderColumns = AttrTransactionOrderColumns{
	Id:               "id",
	AppId:            "app_id",
	TransactionId:    "transaction_id",
	SubTransactionId: "sub_transaction_id",
	Uuid:             "uuid",
	SkuId:            "sku_id",
	Amount:           "amount",
	SubscribeStatus:  " subscribe_status",
	CreatedAt:        "created_at",
}

// NewAttrTransactionOrderDao creates and returns a new DAO object for table data access.
func NewAttrTransactionOrderDao(handlers ...gdb.ModelHandler) *AttrTransactionOrderDao {
	return &AttrTransactionOrderDao{
		group:    "default",
		table:    "attr_transaction_order",
		columns:  attrTransactionOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrTransactionOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrTransactionOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrTransactionOrderDao) Columns() AttrTransactionOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrTransactionOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrTransactionOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrTransactionOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
