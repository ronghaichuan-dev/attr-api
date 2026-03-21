// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gungoren/apple-search-ads-go/asa"
)

type (
	IAppleAds interface {
		// GetAllCampaigns 获取所有广告系列
		GetAllCampaigns(ctx context.Context, companyId int) ([]*asa.Campaign, error)
	}
)

var (
	localAppleAds IAppleAds
)

func AppleAds() IAppleAds {
	if localAppleAds == nil {
		panic("implement not found for interface IAppleAds, forgot register?")
	}
	return localAppleAds
}

func RegisterAppleAds(i IAppleAds) {
	localAppleAds = i
}
