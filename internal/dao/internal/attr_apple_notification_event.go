// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrAppleNotificationEventDao is the data access object for the table attr_apple_notification_event.
type AttrAppleNotificationEventDao struct {
	table    string                            // table is the underlying table name of the DAO.
	group    string                            // group is the database configuration group name of the current DAO.
	columns  AttrAppleNotificationEventColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler                // handlers for customized model modification.
}

// AttrAppleNotificationEventColumns defines and stores column names for the table attr_apple_notification_event.
type AttrAppleNotificationEventColumns struct {
	Id                    string // id
	Envirment             string // 环境
	Version               string // 应用版本
	NotificationUuid      string // 通知唯一ID
	SignedPayload         string // 加密数据
	NotificationType      string // 通知类型
	Subtype               string // 通知子类型
	OriginalTransactionId string // 原始交易ID
	TransactionId         string // 交易ID
	ResponseText          string // 解密后的数据
	ReceivedAt            string // 通知接收时间
	ProcessedAt           string // 解密处理时间
}

// attrAppleNotificationEventColumns holds the columns for the table attr_apple_notification_event.
var attrAppleNotificationEventColumns = AttrAppleNotificationEventColumns{
	Id:                    "id",
	Envirment:             "envirment",
	Version:               "version",
	NotificationUuid:      "notification_uuid",
	SignedPayload:         "signed_payload",
	NotificationType:      "notification_type",
	Subtype:               "subtype",
	OriginalTransactionId: "original_transaction_id",
	TransactionId:         "transaction_id",
	ResponseText:          "response_text",
	ReceivedAt:            "received_at",
	ProcessedAt:           "processed_at",
}

// NewAttrAppleNotificationEventDao creates and returns a new DAO object for table data access.
func NewAttrAppleNotificationEventDao(handlers ...gdb.ModelHandler) *AttrAppleNotificationEventDao {
	return &AttrAppleNotificationEventDao{
		group:    "default",
		table:    "attr_apple_notification_event",
		columns:  attrAppleNotificationEventColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrAppleNotificationEventDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrAppleNotificationEventDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrAppleNotificationEventDao) Columns() AttrAppleNotificationEventColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrAppleNotificationEventDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrAppleNotificationEventDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrAppleNotificationEventDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
