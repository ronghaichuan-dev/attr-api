// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrAppleNotificationEventsDao is the data access object for the table attr_apple_notification_events.
type AttrAppleNotificationEventsDao struct {
	table    string                             // table is the underlying table name of the DAO.
	group    string                             // group is the database configuration group name of the current DAO.
	columns  AttrAppleNotificationEventsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler                 // handlers for customized model modification.
}

// AttrAppleNotificationEventsColumns defines and stores column names for the table attr_apple_notification_events.
type AttrAppleNotificationEventsColumns struct {
	Id                    string //
	NotificationUuid      string //
	SignedPayload         string //
	NotificationType      string //
	Subtype               string //
	OriginalTransactionId string //
	TransactionId         string //
	ResponseText          string //
	ReceivedAt            string //
	ProcessedAt           string //
}

// attrAppleNotificationEventsColumns holds the columns for the table attr_apple_notification_events.
var attrAppleNotificationEventsColumns = AttrAppleNotificationEventsColumns{
	Id:                    "id",
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

// NewAttrAppleNotificationEventsDao creates and returns a new DAO object for table data access.
func NewAttrAppleNotificationEventsDao(handlers ...gdb.ModelHandler) *AttrAppleNotificationEventsDao {
	return &AttrAppleNotificationEventsDao{
		group:    "default",
		table:    "attr_apple_notification_events",
		columns:  attrAppleNotificationEventsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrAppleNotificationEventsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrAppleNotificationEventsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrAppleNotificationEventsDao) Columns() AttrAppleNotificationEventsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrAppleNotificationEventsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrAppleNotificationEventsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrAppleNotificationEventsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
