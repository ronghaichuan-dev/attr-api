// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrEventDao is the data access object for the table attr_event.
type AttrEventDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AttrEventColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AttrEventColumns defines and stores column names for the table attr_event.
type AttrEventColumns struct {
	Id        string // id
	EventName string // 事件名称
	EventCode string // 事件代码
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Status    string // 状态 1-启用 2-禁用
}

// attrEventColumns holds the columns for the table attr_event.
var attrEventColumns = AttrEventColumns{
	Id:        "id",
	EventName: "event_name",
	EventCode: "event_code",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Status:    "status",
}

// NewAttrEventDao creates and returns a new DAO object for table data access.
func NewAttrEventDao(handlers ...gdb.ModelHandler) *AttrEventDao {
	return &AttrEventDao{
		group:    "default",
		table:    "attr_event",
		columns:  attrEventColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrEventDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrEventDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrEventDao) Columns() AttrEventColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrEventDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrEventDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrEventDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
