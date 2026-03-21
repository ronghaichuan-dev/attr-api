// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemAccount is the golang structure for table system_account.
type SystemAccount struct {
	Id          int64       `json:"id"          orm:"id"           description:"id"`                                                      // id
	Appid       string      `json:"appid"       orm:"appid"        description:"应用ID组"`                                                   // 应用ID组
	AccountType int         `json:"accountType" orm:"account_type" description:"账号类型 1-AppStore 2-Google play 3-TikTok 4-ASA 5-Facebook"` // 账号类型 1-AppStore 2-Google play 3-TikTok 4-ASA 5-Facebook
	CompanyId   int         `json:"companyId"   orm:"company_id"   description:"公司ID"`                                                    // 公司ID
	AccountInfo string      `json:"accountInfo" orm:"account_info" description:"账号信息"`                                                    // 账号信息
	Creator     int         `json:"creator"     orm:"creator"      description:"创建人"`                                                     // 创建人
	Modifier    int         `json:"modifier"    orm:"modifier"     description:"修改人"`                                                     // 修改人
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`                                                    // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`                                                    // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`                                                    // 删除时间
}
