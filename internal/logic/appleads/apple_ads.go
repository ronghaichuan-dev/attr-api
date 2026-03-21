package appleads

import (
	"context"
	"encoding/json"
	"god-help-service/api/v1/api"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service"
	"time"

	"github.com/gungoren/apple-search-ads-go/asa"
)

type sAppleAds struct {
}

func init() {
	service.RegisterAppleAds(NewAppleAds())
}

func NewAppleAds() *sAppleAds {
	return &sAppleAds{}
}

// GetAllCampaigns 获取所有广告系列
func (s *sAppleAds) GetAllCampaigns(ctx context.Context, companyId int) ([]*asa.Campaign, error) {
	var info entity.SystemAccount
	err := dao.SystemAccount.Ctx(ctx).Fields("account_info").Where(dao.SystemAccount.Columns().CompanyId, companyId).Scan(&info)
	if err != nil {
		return nil, err
	}
	appleAdsInfo := api.AppleAdsInfo{}
	err = json.Unmarshal([]byte(info.AccountInfo), &appleAdsInfo)
	if err != nil {
		return nil, err
	}
	// Organization ID in Apple Search Ads
	orgID := appleAdsInfo.OrgId
	// Key ID for the given private key, described in Apple Search Ads
	keyID := appleAdsInfo.KeyId
	// Team ID for the given private key for the Apple Search Ads
	teamID := appleAdsInfo.TeamId
	// ClientID ID for the given private key for the Apple Search Ads
	clientID := appleAdsInfo.ClientId
	// A duration value for the lifetime of a token. Apple Search Ads does not accept a token with a lifetime of longer than 20 minutes
	expiryDuration := 20 * time.Minute
	// The bytes of the private key created you have uploaded to it Apple Search Ads.
	auth, err := asa.NewTokenConfig(orgID, keyID, teamID, clientID, expiryDuration, []byte(appleAdsInfo.PrivateKey))
	if err != nil {
		return nil, err
	}
	client := asa.NewClient(auth.Client())
	params := &asa.GetAllCampaignQuery{
		Limit:  100,
		Offset: 1,
	}
	apps, _, err := client.Campaigns.GetAllCampaigns(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return apps.Campaigns, nil
}
