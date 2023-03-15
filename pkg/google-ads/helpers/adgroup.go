package helpers

import (
	"errors"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"google.golang.org/api/iterator"
)

type AdGroup struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ResourceName string `json:"resource_name"`
}

func GetAdGroups(ClientId string, CampaignId string) []AdGroup {
	request := services.SearchGoogleAdsRequest{
		CustomerId: ClientId,
		Query: `SELECT campaign.id, ad_group.id, ad_group.name, 
       			ad_group.resource_name FROM ad_group WHERE campaign.id = 
				` + CampaignId + ` ORDER BY campaign.id`,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	var adGroups []AdGroup

	for {
		row, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			return adGroups
		}

		adGroupResp := row.GetAdGroup()
		adGroup := AdGroup{}
		adGroup.Id = *adGroupResp.Id
		if adGroupResp.Name != nil {
			adGroup.Name = *adGroupResp.Name
		}
		adGroup.ResourceName = adGroupResp.ResourceName

		adGroups = append(adGroups, adGroup)
	}

	return adGroups
}

func GetAdGroup(ClientId string, AdGroupId string) AdGroup {
	request := services.SearchGoogleAdsRequest{
		CustomerId: ClientId,
		Query: `	SELECT 
    					ad_group.id, 
						ad_group_ad.ad.id,
						ad_group_ad.ad.name,
						ad_group_ad.ad.final_urls,
						ad_group_ad.resource_name
					FROM 
						ad_group_ad
					WHERE
					    ad_group.id = ` + AdGroupId,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	row, err := resp.Next()
	if errors.Is(err, iterator.Done) {
		return AdGroup{}
	} else if err != nil {
		return AdGroup{}
	}

	adGroupResp := row.GetAdGroup()
	adGroup := AdGroup{}
	adGroup.Id = *adGroupResp.Id
	if adGroupResp.Name != nil {
		adGroup.Name = *adGroupResp.Name
	}
	adGroup.ResourceName = adGroupResp.ResourceName

	return adGroup
}
