package helpers

import (
	"errors"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"google.golang.org/api/iterator"
)

type Campaign struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ResourceName string `json:"resource_name"`
}

func GetCampaigns(ClientId string) []Campaign {
	request := services.SearchGoogleAdsRequest{
		CustomerId: ClientId,
		Query:      "SELECT campaign.id, campaign.name, campaign.resource_name FROM campaign",
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	var campaigns []Campaign

	for {
		row, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			return campaigns
		}

		campaign := row.GetCampaign()
		if campaign == nil {
			continue
		}

		c := Campaign{}

		c.Id = *campaign.Id
		if campaign.Name != nil {
			c.Name = *campaign.Name
		}
		c.ResourceName = campaign.ResourceName

		campaigns = append(campaigns, c)
	}

	return campaigns
}

func GetCampaign(ClientId string, CampaignId string) Campaign {
	request := services.SearchGoogleAdsRequest{
		CustomerId: ClientId,
		Query: `	SELECT 
    					campaign.id, 
						campaign.name 
					FROM 
						campaign
					WHERE
					    campaign.id = ` + CampaignId + `
					LIMIT 1`,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	row, err := resp.Next()
	if errors.Is(err, iterator.Done) {
		return Campaign{}
	} else if err != nil {
		return Campaign{}
	}

	campaignResp := row.GetCampaign()
	campaign := Campaign{}
	campaign.Id = *campaignResp.Id
	if campaignResp.Name != nil {
		campaign.Name = *campaignResp.Name
	}
	campaign.ResourceName = campaignResp.ResourceName

	return campaign
}
