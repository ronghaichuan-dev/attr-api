// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemAccountDao is the data access object for the table system_account.
type SystemAccountDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SystemAccountColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SystemAccountColumns defines and stores column names for the table system_account.
type SystemAccountColumns struct {
	Id          string // id
	Appid       string // 应用ID组
	AccountType string // 账号类型 1-AppStore 2-Google play 3-TikTok 4-ASA 5-Facebook
	CompanyId   string // 公司ID
	AccountInfo string // 账号信息
	Creator     string // 创建人
	Modifier    string // 修改人
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
}

// systemAccountColumns holds the columns for the table system_account.
var systemAccountColumns = SystemAccountColumns{
	Id:          "id",
	Appid:       "appid",
	AccountType: "account_type",
	CompanyId:   "company_id",
	AccountInfo: "account_info",
	Creator:     "creator",
	Modifier:    "modifier",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewSystemAccountDao creates and returns a new DAO object for table data access.
func NewSystemAccountDao(handlers ...gdb.ModelHandler) *SystemAccountDao {
	return &SystemAccountDao{
		group:    "default",
		table:    "system_account",
		columns:  systemAccountColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemAccountDao) Columns() SystemAccountColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemAccountDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
