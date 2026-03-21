// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApps is the golang structure for table system_apps.
type SystemApps struct {
	Id              int64       `json:"id"              orm:"id"               description:"主键ID"`      // 主键ID
	Appid           string      `json:"appid"           orm:"appid"            description:"应用ID，主键"`   // 应用ID，主键
	AppToken        string      `json:"appToken"        orm:"app_token"        description:"应用Token"`   // 应用Token
	CompanyId       int         `json:"companyId"       orm:"company_id"       description:"公司ID"`      // 公司ID
	AppName         string      `json:"appName"         orm:"app_name"         description:"应用名称"`      // 应用名称
	Icon            string      `json:"icon"            orm:"icon"             description:"应用图标"`      // 应用图标
	SubscriptionFee float64     `json:"subscriptionFee" orm:"subscription_fee" description:"订阅费用"`      // 订阅费用
	Creator         int         `json:"creator"         orm:"creator"          description:"创建人"`       // 创建人
	Modifier        int         `json:"modifier"        orm:"modifier"         description:"修改人"`       // 修改人
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`      // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`      // 更新时间
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"       description:"删除时间（软删除）"` // 删除时间（软删除）
}
