// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemUsersDao is the data access object for the table system_users.
type SystemUsersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SystemUsersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SystemUsersColumns defines and stores column names for the table system_users.
type SystemUsersColumns struct {
	Id        string //
	Username  string //
	Password  string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string // 删除时间（软删除）
	RoleId    string // 角色ID，关联角色表
}

// systemUsersColumns holds the columns for the table system_users.
var systemUsersColumns = SystemUsersColumns{
	Id:        "id",
	Username:  "username",
	Password:  "password",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	RoleId:    "role_id",
}

// NewSystemUsersDao creates and returns a new DAO object for table data access.
func NewSystemUsersDao(handlers ...gdb.ModelHandler) *SystemUsersDao {
	return &SystemUsersDao{
		group:    "default",
		table:    "system_users",
		columns:  systemUsersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemUsersDao) Columns() SystemUsersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemUsersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
