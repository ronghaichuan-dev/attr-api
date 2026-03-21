package appleapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// ResponseBodyV2 苹果服务器通知V2响应体
type ResponseBodyV2 struct {
	NotificationType string              `json:"notificationType"`
	Subtype          string              `json:"subtype,omitempty"`
	NotificationUUID string              `json:"notificationUUID"`
	AppMetadata      *AppMetadata        `json:"appMetadata"`
	Version          string              `json:"version"`
	SignedDate       int64               `json:"signedDate"`
	Data             *ResponseBodyV2Data `json:"data,omitempty"`
}

// ResponseBodyV2Data 响应体数据
type ResponseBodyV2Data struct {
	AppMetadata           *AppMetadata           `json:"appMetadata,omitempty"`
	Environment           string                 `json:"environment"`
	RenewalInfo           *RenewalInfo           `json:"renewalInfo,omitempty"`
	TransactionInfo       *TransactionInfo       `json:"transactionInfo,omitempty"`
	SignedRenewalInfo     string                 `json:"signedRenewalInfo,omitempty"`
	SignedTransactionInfo string                 `json:"signedTransactionInfo,omitempty"`
	Summary               *ResponseBodyV2Summary `json:"summary,omitempty"`
}

// ResponseBodyV2Summary 响应体摘要
type ResponseBodyV2Summary struct {
	RequestIdentifier      string   `json:"requestIdentifier"`
	StorefrontCountryCodes []string `json:"storefrontCountryCodes"`
	OriginalExpirationDate int64    `json:"originalExpirationDate"`
	ExtendedExpirationDate int64    `json:"extendedExpirationDate"`
	SuccessfulCount        int      `json:"successfulCount"`
	FailedCount            int      `json:"failedCount"`
}

// AppMetadata 应用元数据
type AppMetadata struct {
	BundleID               string   `json:"bundleId"`
	BundleVersion          string   `json:"bundleVersion"`
	Environment            string   `json:"environment"`
	AppAppleID             int64    `json:"appAppleId,omitempty"`
	AppStoreVersion        string   `json:"appStoreVersion,omitempty"`
	SellerName             string   `json:"sellerName,omitempty"`
	StorefrontCountryCodes []string `json:"storefrontCountryCodes,omitempty"`
	Status                 int      `json:"status"`
}

// RenewalInfo 订阅续订信息
type RenewalInfo struct {
	AutoRenewProductID            string `json:"autoRenewProductId"`
	AutoRenewStatus               int    `json:"autoRenewStatus"`
	BillingGracePeriodExpiresDate int64  `json:"billingGracePeriodExpiresDate,omitempty"`
	ExpirationIntent              int    `json:"expirationIntent,omitempty"`
	GracePeriodExpiresDate        int64  `json:"gracePeriodExpiresDate,omitempty"`
	IsInBillingRetryPeriod        bool   `json:"isInBillingRetryPeriod,omitempty"`
	OfferCodeRefName              string `json:"offerCodeRefName,omitempty"`
	OriginalTransactionID         string `json:"originalTransactionId"`
	PriceIncreaseStatus           int    `json:"priceIncreaseStatus,omitempty"`
	ProductID                     string `json:"productId"`
	RecentSubscriptionStartDate   int64  `json:"recentSubscriptionStartDate,omitempty"`
	RenewalDate                   int64  `json:"renewalDate,omitempty"`
	SignedDate                    int64  `json:"signedDate"`
	SubscriptionGroupID           string `json:"subscriptionGroupIdentifier,omitempty"`
	WebOrderLineItemID            string `json:"webOrderLineItemId,omitempty"`
}

// TransactionInfo 交易信息
type TransactionInfo struct {
	AppAccountToken       string `json:"appAccountToken,omitempty"`
	BundleID              string `json:"bundleId"`
	Environment           string `json:"environment"`
	ExpiresDate           int64  `json:"expiresDate,omitempty"`
	InAppOwnershipType    string `json:"inAppOwnershipType"`
	IsUpgraded            bool   `json:"isUpgraded,omitempty"`
	OfferDiscountType     string `json:"offerDiscountType,omitempty"`
	OfferID               string `json:"offerId,omitempty"`
	OriginalPurchaseDate  int64  `json:"originalPurchaseDate"`
	OriginalTransactionID string `json:"originalTransactionId"`
	ProductID             string `json:"productId"`
	PurchaseDate          int64  `json:"purchaseDate"`
	Quantity              int    `json:"quantity"`
	Price                 int64  `json:"price"`
	Currency              string `json:"currency"`
	RevocationDate        int64  `json:"revocationDate,omitempty"`
	RevocationReason      int    `json:"revocationReason,omitempty"`
	SubscriptionGroupID   string `json:"subscriptionGroupIdentifier,omitempty"`
	TransactionID         string `json:"transactionId"`
	Type                  string `json:"type"`
	WebOrderLineItemID    string `json:"webOrderLineItemId"`
}

// CreateFromRawNotification 从原始通知创建ResponseBodyV2实例
func CreateFromRawNotification(signedPayload string) (*ResponseBodyV2, error) {

	// 解析JWT
	payloadMap, err := ParseJWT(signedPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	// 解析data字段
	data, ok := payloadMap["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid data format in JWT payload")
	}

	// 解析嵌套的JWT
	renewalInfo, transactionInfo, err := parseNestedJWTs(data)
	if err != nil {
		return nil, err
	}

	// 创建AppMetadata
	appMetadata, err := createAppMetadata(data)
	if err != nil {
		return nil, err
	}

	// 解析notificationType、subtype、notificationUUID、version、signedDate
	notificationType, _ := payloadMap["notificationType"].(string)
	subtype, _ := payloadMap["subtype"].(string)
	notificationUUID, _ := payloadMap["notificationUUID"].(string)
	version, _ := payloadMap["version"].(string)

	// 转换signedDate为int64
	signedDateFloat, ok := payloadMap["signedDate"].(float64)
	if !ok {
		return nil, errors.New("invalid signedDate format")
	}
	signedDate := int64(signedDateFloat)

	// 创建并返回ResponseBodyV2实例
	responseBodyV2 := &ResponseBodyV2{
		NotificationType: notificationType,
		NotificationUUID: notificationUUID,
		AppMetadata:      appMetadata,
		Version:          version,
		SignedDate:       signedDate,
	}

	// 设置subtype
	if subtype != "" {
		responseBodyV2.Subtype = subtype
	}

	// 设置数据
	responseBodyV2.Data = &ResponseBodyV2Data{
		AppMetadata:     appMetadata,
		RenewalInfo:     renewalInfo,
		TransactionInfo: transactionInfo,
	}

	return responseBodyV2, nil
}

// parseNestedJWTs 解析嵌套的JWT
func parseNestedJWTs(data map[string]interface{}) (*RenewalInfo, *TransactionInfo, error) {
	var renewalInfo *RenewalInfo
	var transactionInfo *TransactionInfo

	// 解析SignedRenewalInfo
	if signedRenewalInfo, ok := data["signedRenewalInfo"].(string); ok && signedRenewalInfo != "" {
		renewalMap, err := ParseJWT(signedRenewalInfo)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse signedRenewalInfo: %w", err)
		}

		renewalInfoBytes, err := json.Marshal(renewalMap)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal renewalInfo: %w", err)
		}

		if err = json.Unmarshal(renewalInfoBytes, &renewalInfo); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal renewalInfo: %w", err)
		}
	}

	// 解析SignedTransactionInfo
	if signedTransactionInfo, ok := data["signedTransactionInfo"].(string); ok && signedTransactionInfo != "" {
		transactionMap, err := ParseJWT(signedTransactionInfo)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse signedTransactionInfo: %w", err)
		}

		transactionInfoBytes, err := json.Marshal(transactionMap)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal transactionInfo: %w", err)
		}

		if err := json.Unmarshal(transactionInfoBytes, &transactionInfo); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal transactionInfo: %w", err)
		}
	}

	return renewalInfo, transactionInfo, nil
}

// createAppMetadata 创建AppMetadata
func createAppMetadata(data map[string]interface{}) (*AppMetadata, error) {
	// 从data字段获取appMetadata
	if data["appMetaData"] != nil {
		appMetadataMap, ok := data["appMetadata"].(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid appMetadata format")
		}
		appMetadataBytes, err := json.Marshal(appMetadataMap)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal appMetadata: %w", err)
		}
		var appMetadata AppMetadata
		if err := json.Unmarshal(appMetadataBytes, &appMetadata); err != nil {
			return nil, fmt.Errorf("failed to unmarshal appMetadata: %w", err)
		}

		return &appMetadata, nil
	}

	return nil, nil

}

// GetNotificationType 获取通知类型
func (r *ResponseBodyV2) GetNotificationType() string {
	return r.NotificationType
}

// GetSubtype 获取子类型
func (r *ResponseBodyV2) GetSubtype() string {
	return r.Subtype
}

// GetNotificationUUID 获取通知UUID
func (r *ResponseBodyV2) GetNotificationUUID() string {
	return r.NotificationUUID
}

// GetAppMetadata 获取应用元数据
func (r *ResponseBodyV2) GetAppMetadata() *AppMetadata {
	return r.AppMetadata
}

// GetVersion 获取版本
func (r *ResponseBodyV2) GetVersion() string {
	return r.Version
}

// GetSignedDate 获取签名日期
func (r *ResponseBodyV2) GetSignedDate() int64 {
	return r.SignedDate
}

// GetSignedDateAsTime 获取签名日期作为time.Time
func (r *ResponseBodyV2) GetSignedDateAsTime() time.Time {
	return time.Unix(r.SignedDate/1000, (r.SignedDate%1000)*1000000)
}
