// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrPostbackDao is the data access object for the table attr_postback.
type AttrPostbackDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  AttrPostbackColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// AttrPostbackColumns defines and stores column names for the table attr_postback.
type AttrPostbackColumns struct {
	Id                    string // id
	AppId                 string // 应用ID
	PostbackType          string // 回传类型: install/event/reengagement
	Network               string // 渠道
	OriginalTransactionId string // 原始交易ID
	EventName             string // 事件名
	PostbackUrl           string // 回传URL
	ResponseCode          string // 响应码
	ResponseBody          string // 响应内容
	Status                string // 状态: 1-成功 2-失败 3-重试中
	RetryCount            string // 重试次数
	CreatedAt             string // 创建时间
}

// attrPostbackColumns holds the columns for the table attr_postback.
var attrPostbackColumns = AttrPostbackColumns{
	Id:                    "id",
	AppId:                 "app_id",
	PostbackType:          "postback_type",
	Network:               "network",
	OriginalTransactionId: "original_transaction_id",
	EventName:             "event_name",
	PostbackUrl:           "postback_url",
	ResponseCode:          "response_code",
	ResponseBody:          "response_body",
	Status:                "status",
	RetryCount:            "retry_count",
	CreatedAt:             "created_at",
}

// NewAttrPostbackDao creates and returns a new DAO object for table data access.
func NewAttrPostbackDao(handlers ...gdb.ModelHandler) *AttrPostbackDao {
	return &AttrPostbackDao{
		group:    "default",
		table:    "attr_postback",
		columns:  attrPostbackColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrPostbackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrPostbackDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrPostbackDao) Columns() AttrPostbackColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrPostbackDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrPostbackDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrPostbackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
