// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemSettings is the golang structure for table system_settings.
type SystemSettings struct {
	Id        int         `json:"id"        orm:"id"         description:"主键ID"`      // 主键ID
	Key       string      `json:"key"       orm:"key"        description:"设置键，唯一"`    // 设置键，唯一
	Value     string      `json:"value"     orm:"value"      description:"设置值"`       // 设置值
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`      // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`      // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间（软删除）"` // 删除时间（软删除）
}
