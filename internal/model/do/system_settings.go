// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemSettings is the golang structure of table system_settings for DAO operations like Where/Data.
type SystemSettings struct {
	g.Meta    `orm:"table:system_settings, do:true"`
	Id        any         // 主键ID
	Key       any         // 设置键，唯一
	Value     any         // 设置值
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间（软删除）
}
