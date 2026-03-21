// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotificationEventsDao is the data access object for the table notification_events.
type NotificationEventsDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  NotificationEventsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// NotificationEventsColumns defines and stores column names for the table notification_events.
type NotificationEventsColumns struct {
	Id             string //
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
	NotificationId string //
	EventType      string //
	EventData      string //
	Status         string //
}

// notificationEventsColumns holds the columns for the table notification_events.
var notificationEventsColumns = NotificationEventsColumns{
	Id:             "id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	NotificationId: "notification_id",
	EventType:      "event_type",
	EventData:      "event_data",
	Status:         "status",
}

// NewNotificationEventsDao creates and returns a new DAO object for table data access.
func NewNotificationEventsDao(handlers ...gdb.ModelHandler) *NotificationEventsDao {
	return &NotificationEventsDao{
		group:    "default",
		table:    "notification_events",
		columns:  notificationEventsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NotificationEventsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NotificationEventsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NotificationEventsDao) Columns() NotificationEventsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NotificationEventsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NotificationEventsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *NotificationEventsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
