// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrEventLogDao is the data access object for the table attr_event_log.
type AttrEventLogDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  AttrEventLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// AttrEventLogColumns defines and stores column names for the table attr_event_log.
type AttrEventLogColumns struct {
	Id           string // id
	Country      string // 国家
	City         string // 城市
	Region       string // 州/省
	EventUuid    string // 事件唯一ID
	Appid        string // APP ID
	EventCode    string // 事件ID
	Rsid         string // 设备ID
	ResponseText string // 事件内容
	SentAt       string // 发送时间
	CreatedAt    string // 创建时间
}

// attrEventLogColumns holds the columns for the table attr_event_log.
var attrEventLogColumns = AttrEventLogColumns{
	Id:           "id",
	Country:      "country",
	City:         "city",
	Region:       "region",
	EventUuid:    "event_uuid",
	Appid:        "appid",
	EventCode:    "event_code",
	Rsid:         "rsid",
	ResponseText: "response_text",
	SentAt:       "sent_at",
	CreatedAt:    "created_at",
}

// NewAttrEventLogDao creates and returns a new DAO object for table data access.
func NewAttrEventLogDao(handlers ...gdb.ModelHandler) *AttrEventLogDao {
	return &AttrEventLogDao{
		group:    "default",
		table:    "attr_event_log",
		columns:  attrEventLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttrEventLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttrEventLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttrEventLogDao) Columns() AttrEventLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttrEventLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttrEventLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AttrEventLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
