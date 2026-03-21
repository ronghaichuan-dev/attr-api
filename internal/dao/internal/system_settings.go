// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemSettingsDao is the data access object for the table system_settings.
type SystemSettingsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SystemSettingsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SystemSettingsColumns defines and stores column names for the table system_settings.
type SystemSettingsColumns struct {
	Id        string // 主键ID
	Key       string // 设置键，唯一
	Value     string // 设置值
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间（软删除）
}

// systemSettingsColumns holds the columns for the table system_settings.
var systemSettingsColumns = SystemSettingsColumns{
	Id:        "id",
	Key:       "key",
	Value:     "value",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewSystemSettingsDao creates and returns a new DAO object for table data access.
func NewSystemSettingsDao(handlers ...gdb.ModelHandler) *SystemSettingsDao {
	return &SystemSettingsDao{
		group:    "default",
		table:    "system_settings",
		columns:  systemSettingsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemSettingsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemSettingsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemSettingsDao) Columns() SystemSettingsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemSettingsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemSettingsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemSettingsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
