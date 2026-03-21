// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrNotificationsDao is the data access object for the table attr_notifications.
type AttrNotificationsDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  AttrNotificationsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// AttrNotificationsColumns defines and stores column names for the table attr_notifications.
type AttrNotificationsColumns struct {
	Id            string // id
	Uuid          string // 用户ID
	Token         string // 用户Token
	NoticeType    string // 通知类型
	TxId          string // 事务ID
	RenewalStatus string // 续费状态
	Sku           string // 商品sku
	NoticeAt      string // 通知时间
	CreatedAt     string // 创建时间
}

// attrNotificationsColumns holds the columns for the table attr_notifications.
var attrNotificationsColumns = AttrNotificationsColumns{
	Id:            "id",
	Uuid:          "uuid",
	Token:         "token",
	NoticeType:    "notice_type",
	TxId:          "tx_id",
	RenewalStatus: "renewal _status",
	Sku:           "sku",
	NoticeAt:      "notice_at",
	CreatedAt:     "created_at",
}

// NewAttrNotificationsDao creates and returns a new DAO object for table data access.
func NewAttrNotificationsDao(handlers ...gdb.ModelHandler) *AttrNotificationsDao {
	return &AttrNotificationsDao{
		group:    "default",
		table:    "attr_notifications",
		columns:  attrNotificationsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrNotificationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrNotificationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrNotificationsDao) Columns() AttrNotificationsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrNotificationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrNotificationsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrNotificationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
