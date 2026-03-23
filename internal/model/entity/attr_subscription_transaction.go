// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AttrSubscriptionTransaction is the golang structure for table attr_subscription_transaction.
type AttrSubscriptionTransaction struct {
	Id                    int64  `json:"id"                    orm:"id"                      description:"id"`                                      // id
	TransactionType       string `json:"transactionType"       orm:"transaction_type"        description:"交易类型 RENEW / REFUND / TRIAL"`             // 交易类型 RENEW / REFUND / TRIAL
	Envirment             string `json:"envirment"             orm:"envirment"               description:"环境"`                                      // 环境
	AppVersion            string `json:"appVersion"            orm:"app_version"             description:"应用版本"`                                    // 应用版本
	Appid                 string `json:"appid"                 orm:"appid"                   description:"应用ID"`                                    // 应用ID
	OriginalTransactionId string `json:"originalTransactionId" orm:"original_transaction_id" description:"原始交易ID"`                                  // 原始交易ID
	TransactionId         string `json:"transactionId"         orm:"transaction_id"          description:"子交易ID"`                                   // 子交易ID
	InAppOwnership        string `json:"inAppOwnership"        orm:"in_app_ownership"        description:"是否为用户购买 PURCHASED-购买 FAMILY_SHARED-家庭分享"` // 是否为用户购买 PURCHASED-购买 FAMILY_SHARED-家庭分享
	Rsid                  string `json:"rsid"                  orm:"rsid"                    description:"用户设备ID"`                                  // 用户设备ID
	ProductId             string `json:"productId"             orm:"product_id"              description:"sku"`                                     // sku
	Price                 int64  `json:"price"                 orm:"price"                   description:"订阅金额"`                                    // 订阅金额
	Currency              string `json:"currency"              orm:"currency"                description:"币种"`                                      // 币种
	SubscribeStatus       string `json:"subscribeStatus"       orm:"subscribe_status"        description:"订阅状态"`                                    // 订阅状态
	PurchaseAt            int64  `json:"purchaseAt"            orm:"purchase_at"             description:"购买时间"`                                    // 购买时间
	CreatedAt             int64  `json:"createdAt"             orm:"created_at"              description:"创建时间"`                                    // 创建时间
	Country               string `json:"country"               orm:"country"                 description:"国家"`                                      // 国家
	TrackerNetwork        string `json:"trackerNetwork"        orm:"tracker_network"         description:"归因渠道"`                                    // 归因渠道
	CampaignId            string `json:"campaignId"            orm:"campaign_id"             description:"推广活动ID"`                                  // 推广活动ID
	AdgroupId             string `json:"adgroupId"             orm:"adgroup_id"              description:"广告组ID"`                                   // 广告组ID
	AdId                  string `json:"adId"                  orm:"ad_id"                   description:"广告ID"`                                    // 广告ID
}
