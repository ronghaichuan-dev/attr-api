package admin

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttrEventLog 操作`attr_event_log`表的DAO结构
type AttrEventLog struct {
	*gdb.Model
}

// NewAppEventLog 创建并返回一个操作`attr_event_log`表的DAO实例
func NewAppEventLog(ctx context.Context, option ...interface{}) *AttrEventLog {
	return &AttrEventLog{
		Model: g.DB().Model("attr_event_log").Safe(),
	}
}
