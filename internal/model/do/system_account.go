// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemAccount is the golang structure of table system_account for DAO operations like Where/Data.
type SystemAccount struct {
	g.Meta      `orm:"table:system_account, do:true"`
	Id          any         // id
	Appid       any         // 应用ID组
	AccountType any         // 账号类型 1-AppStore 2-Google play 3-TikTok 4-ASA 5-Facebook
	CompanyId   any         // 公司ID
	AccountInfo any         // 账号信息
	Creator     any         // 创建人
	Modifier    any         // 修改人
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
