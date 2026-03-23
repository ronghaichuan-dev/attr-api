package consts

const (
	IsHandleTokenYes = iota + 1
	IsHandleTokenNo
)

// 归因匹配方式
const (
	MatchTypeDeviceId      = "device_id"      // 设备ID匹配（IDFA/GAID）
	MatchTypeReferrer      = "referrer"        // Referrer匹配
	MatchTypeProbabilistic = "probabilistic"   // 概率匹配
	MatchTypeTracker       = "tracker"         // 第三方Tracker匹配
	MatchTypeAdServices    = "ad_services"     // Apple Ad Services匹配
)

// 归因匹配置信度
const (
	MatchConfidenceHigh   = "high"
	MatchConfidenceMedium = "medium"
	MatchConfidenceLow    = "low"
)

// 归因窗口（秒）
const (
	ClickAttributionWindow      = 7 * 24 * 3600  // 点击归因窗口 7 天
	ImpressionAttributionWindow = 24 * 3600       // 展示归因窗口 24 小时
)

// 点击类型
const (
	ClickTypeClick      = "click"
	ClickTypeImpression = "impression"
)

// 回传状态
const (
	PostbackStatusSuccess = iota + 1
	PostbackStatusFailed
	PostbackStatusRetrying
)

// 回传类型
const (
	PostbackTypeInstall      = "install"
	PostbackTypeEvent        = "event"
	PostbackTypeReengagement = "reengagement"
)

// 订阅状态
const (
	SubscriptionStatusActive       = iota + 1 // 1-自动续订服务已激活
	SubscriptionStatusExpired                  // 2-自动续订服务已过期
	SubscriptionStatusBillingRetry             // 3-计费重试期
	SubscriptionStatusGracePeriod              // 4-账单宽限期
	SubscriptionStatusCanceled                 // 5-已取消
	SubscriptionStatusRevoked                  // 6-已撤销
)

const (
	IsTrialFreeYes = iota + 1
	IsTrialFreeNo
)

const (
	IsPaidYes = iota + 1
	IsPaidNo
)

const (
	AutoRenewStatusYes = iota + 1
	AutoRenewStatusNo
)

const (
	EventCodeAttribution  = "Attribution"  //归因事件
	EventCodeTrialFree    = "TrialFree"    //试订事件
	EventCodeInstall      = "Install"      //安装事件
	EventCodeScreenUnlock = "ScreenUnlock" //开屏事件
	EventCodeSubscribe    = "Subscribe"    //订阅事件
	EventCodeSubscribeFix = "SubscribeFix" //订阅修复事件
	EventCodeStartUp      = "StartUp"      //启动事件
)

const (
	ExpiresReasonDefault             = iota + 1
	ExpiresReasonRetryFinished       //2-订阅在计费重试期结束后过期
	ExpiresReasonPriceUpgrade        //3-订阅因价格上涨过期
	ExpiresReasonUnavailableProduct  //4-订阅因产品不可售过期
	ExpiresReasonUserCancelSubscribe //5-用户自愿取消订阅导致过期
)
