// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApps is the golang structure of table system_apps for DAO operations like Where/Data.
type SystemApps struct {
	g.Meta          `orm:"table:system_apps, do:true"`
	Id              any         // 主键ID
	Appid           any         // 应用ID，主键
	AppToken        any         // 应用Token
	CompanyId       any         // 公司ID
	AppName         any         // 应用名称
	Icon            any         // 应用图标
	SubscriptionFee any         // 订阅费用
	Creator         any         // 创建人
	Modifier        any         // 修改人
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 删除时间（软删除）
}
