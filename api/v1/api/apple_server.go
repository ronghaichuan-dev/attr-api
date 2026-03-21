package api

type AppleAttributionInfoResponse struct {
	Attribution     bool   `json:"attribution"`
	OrgId           int    `json:"orgId"`
	CampaignId      int    `json:"campaignId"`
	ConversionType  string `json:"conversionType"`
	ImpressionDate  string `json:"impressionDate"`
	ClickDate       string `json:"clickDate"`
	ClaimType       string `json:"claimType"`
	AdGroupId       int    `json:"adGroupId"`
	CountryOrRegion string `json:"countryOrRegion"`
	KeywordId       int    `json:"keywordId"`
	AdId            int    `json:"adId"`
}
