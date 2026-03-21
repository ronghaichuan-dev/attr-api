package consts

const (
	IsHandleTokenYes = iota + 1
	IsHandleTokenNo
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
