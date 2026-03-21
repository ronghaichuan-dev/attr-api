// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemRolesDao is the data access object for the table system_roles.
type SystemRolesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SystemRolesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SystemRolesColumns defines and stores column names for the table system_roles.
type SystemRolesColumns struct {
	Id        string // 角色ID，主键，自增
	RoleName  string // 角色名称，最多100字符，不能为空
	RoleCode  string // 角色代码，最多100字符，不能为空
	RoleDesc  string // 角色描述，最多500字符，可选
	Status    string // 状态，0：禁用，1：启用，默认为1
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间（软删除）
}

// systemRolesColumns holds the columns for the table system_roles.
var systemRolesColumns = SystemRolesColumns{
	Id:        "id",
	RoleName:  "role_name",
	RoleCode:  "role_code",
	RoleDesc:  "role_desc",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewSystemRolesDao creates and returns a new DAO object for table data access.
func NewSystemRolesDao(handlers ...gdb.ModelHandler) *SystemRolesDao {
	return &SystemRolesDao{
		group:    "default",
		table:    "system_roles",
		columns:  systemRolesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemRolesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemRolesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemRolesDao) Columns() SystemRolesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemRolesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemRolesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemRolesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
