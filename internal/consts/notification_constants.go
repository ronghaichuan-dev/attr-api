package consts

// 通知类型常量
const (
	// NotificationTypeConsumptionRequest 消费请求通知类型
	// 表示客户发起了对消耗性应用内购买的退款请求，App Store正在请求您提供消费数据。
	NotificationTypeConsumptionRequest = "CONSUMPTION_REQUEST"

	// NotificationTypeDidChangeRenewalPref 续订偏好更改通知类型
	// 表示用户更改了他们的订阅计划。
	NotificationTypeDidChangeRenewalPref = "DID_CHANGE_RENEWAL_PREF"

	// NotificationTypeDidChangeRenewalStatus 续订状态更改通知类型
	// 表示用户更改了订阅续订状态。
	NotificationTypeDidChangeRenewalStatus = "DID_CHANGE_RENEWAL_STATUS"

	// NotificationTypeDidFailToRenew 续订失败通知类型
	// 表示由于计费问题，订阅未能续订。
	NotificationTypeDidFailToRenew = "DID_FAIL_TO_RENEW"

	// NotificationTypeDidRenew 续订成功通知类型
	// 表示订阅已成功续订。
	NotificationTypeDidRenew = "DID_RENEW"

	// NotificationTypeExpired 过期通知类型
	// 表示订阅已过期。
	NotificationTypeExpired = "EXPIRED"

	// NotificationTypeGracePeriodExpired 宽限期过期通知类型
	// 表示计费宽限期已结束，订阅未续订，您可以关闭对服务或内容的访问。
	NotificationTypeGracePeriodExpired = "GRACE_PERIOD_EXPIRED"

	// NotificationTypeOfferRedeemed 优惠兑换通知类型
	// 表示用户兑换了促销优惠或优惠代码。
	NotificationTypeOfferRedeemed = "OFFER_REDEEMED"

	// NotificationTypePriceIncrease 价格上涨通知类型
	// 表示系统已通知用户自动续订订阅价格上涨。
	NotificationTypePriceIncrease = "PRICE_INCREASE"

	// NotificationTypeRefund 退款通知类型
	// 表示App Store已成功退款了消耗性应用内购买、非消耗性应用内购买、自动续订订阅或非续订订阅的交易。
	NotificationTypeRefund = "REFUND"

	// NotificationTypeRefundDeclined 退款拒绝通知类型
	// 表示App Store拒绝了应用开发者发起的退款请求。
	NotificationTypeRefundDeclined = "REFUND_DECLINED"

	// NotificationTypeRefundReversed 退款撤销通知类型
	// 表示App Store撤销了先前批准的退款，因为客户提出了争议。
	NotificationTypeRefundReversed = "REFUND_REVERSED"

	// NotificationTypeRenewalExtended 续订延长通知类型
	// 表示App Store延长了特定订阅的续订日期。
	NotificationTypeRenewalExtended = "RENEWAL_EXTENDED"

	// NotificationTypeRenewalExtension 续订延长请求通知类型
	// 表示App Store正在尝试延长您通过调用批量延长订阅续订日期请求的续订日期。
	NotificationTypeRenewalExtension = "RENEWAL_EXTENSION"

	// NotificationTypeRevoke 撤销通知类型
	// 表示用户通过家庭共享获得的应用内购买不再通过共享提供。
	NotificationTypeRevoke = "REVOKE"

	// NotificationTypeSubscribed 订阅通知类型
	// 表示用户订阅了产品。
	NotificationTypeSubscribed = "SUBSCRIBED"

	// NotificationTypeOneTimeCharge 一次性收费通知类型
	// 表示客户购买了消耗性、非消耗性或非续订订阅。
	NotificationTypeOneTimeCharge = "ONE_TIME_CHARGE"

	// NotificationTypeTest 测试通知类型
	// 表示这是一个测试通知，由您通过调用请求测试通知端点发起。
	NotificationTypeTest = "TEST"
)

// 子类型常量
const (
	// SubtypeAccepted 接受子类型
	// 适用于PRICE_INCREASE通知类型。表示用户接受了订阅价格上涨。
	SubtypeAccepted = "ACCEPTED"

	// SubtypeAutoRenewDisabled 自动续订禁用子类型
	// 适用于DID_CHANGE_RENEWAL_STATUS通知类型。表示用户禁用了订阅自动续订，或App Store在用户请求退款后禁用了订阅自动续订。
	SubtypeAutoRenewDisabled = "AUTO_RENEW_DISABLED"

	// SubtypeAutoRenewEnabled 自动续订启用子类型
	// 适用于DID_CHANGE_RENEWAL_STATUS通知类型。表示用户启用了订阅自动续订。
	SubtypeAutoRenewEnabled = "AUTO_RENEW_ENABLED"

	// SubtypeBillingRecovery 计费恢复子类型
	// 适用于DID_RENEW通知类型。表示之前续订失败的过期订阅已成功续订。
	SubtypeBillingRecovery = "BILLING_RECOVERY"

	// SubtypeBillingRetry 计费重试子类型
	// 适用于EXPIRED通知类型。表示由于订阅在计费重试期结束前未成功续订，订阅已过期。
	SubtypeBillingRetry = "BILLING_RETRY"

	// SubtypeDowngrade 降级子类型
	// 适用于DID_CHANGE_RENEWAL_PREF通知类型。表示用户降级了订阅或交叉降级到具有不同持续时间的订阅。降级在下次续订日期生效。
	SubtypeDowngrade = "DOWNGRADE"

	// SubtypeFailure 失败子类型
	// 适用于RENEWAL_EXTENSION通知类型。表示单个订阅的续订日期延长失败。
	SubtypeFailure = "FAILURE"

	// SubtypeGracePeriod 宽限期子类型
	// 适用于DID_FAIL_TO_RENEW通知类型。表示由于计费问题，订阅未能续订。在宽限期内继续提供对订阅的访问。
	SubtypeGracePeriod = "GRACE_PERIOD"

	// SubtypeInitialBuy 初始购买子类型
	// 适用于SUBSCRIBED通知类型。表示用户首次购买了订阅，或用户首次通过家庭共享获得了订阅的访问权限。
	SubtypeInitialBuy = "INITIAL_BUY"

	// SubtypePending 待定子类型
	// 适用于PRICE_INCREASE通知类型。表示系统已通知用户订阅价格上涨，但用户尚未接受。
	SubtypePending = "PENDING"

	// SubtypePriceIncrease 价格上涨子类型
	// 适用于EXPIRED通知类型。表示由于用户不同意价格上涨，订阅已过期。
	SubtypePriceIncrease = "PRICE_INCREASE"

	// SubtypeProductNotForSale 产品不可售子类型
	// 适用于EXPIRED通知类型。表示由于产品在订阅尝试续订时不可购买，订阅已过期。
	SubtypeProductNotForSale = "PRODUCT_NOT_FOR_SALE"

	// SubtypeResubscribe 重新订阅子类型
	// 适用于SUBSCRIBED通知类型。表示用户重新订阅或通过家庭共享获得了同一订阅或同一订阅组内的另一个订阅的访问权限。
	SubtypeResubscribe = "RESUBSCRIBE"

	// SubtypeSummary 摘要子类型
	// 适用于RENEWAL_EXTENSION通知类型。表示App Store服务器已完成您的请求，为所有符合条件的订阅者延长了订阅续订日期。
	SubtypeSummary = "SUMMARY"

	// SubtypeUpgrade 升级子类型
	// 适用于DID_CHANGE_RENEWAL_PREF通知类型。表示用户升级了订阅或交叉升级到具有相同持续时间的订阅。升级立即生效。
	SubtypeUpgrade = "UPGRADE"

	// SubtypeVoluntary 自愿子类型
	// 适用于EXPIRED通知类型。表示订阅在用户禁用订阅自动续订后过期。
	SubtypeVoluntary = "VOLUNTARY"
)
