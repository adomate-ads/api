package gads

import (
	"errors"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"net/http"
)

type Campaign struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ResourceName string `json:"resource_name"`
}

// GetCampaigns Google Ads godoc
// @Summary Get Google Ads Campaigns
// @Description Gets all Google Ads Campaigns for specific Client
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param id path int true "Campaign ID"
// @Success 200 {object} []Campaign
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/campaigns/{clientId} [get]
func GetCampaigns(c *gin.Context) {
	clientId := c.Param("clientId")

	request := services.SearchGoogleAdsRequest{
		CustomerId: clientId,
		Query:      "SELECT campaign.id, campaign.name FROM campaign ORDER BY campaign.id",
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	var campaigns []Campaign

	for {
		row, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		campaignResp := row.GetCampaign()
		campaign := Campaign{}
		campaign.Id = *campaignResp.Id
		if campaignResp.Name != nil {
			campaign.Name = *campaignResp.Name
		}

		campaigns = append(campaigns, campaign)
	}

	c.JSON(http.StatusOK, gin.H{"campaigns": campaigns})
}

// GetCampaign Google Ads godoc
// @Summary Get Google Ads Campaign
// @Description Gets all information about specific Google Ads Campaign
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param id path int true "Campaign ID"
// @Success 200 {object} Campaign
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/campaign/{campaignId} [get]
func GetCampaign(c *gin.Context) {
	clientId := c.Param("clientId")
	campaignId := c.Param("campaignId")

	//TODO - Only allow the user to get the client attached to their company

	request := services.SearchGoogleAdsRequest{
		CustomerId: clientId,
		Query: `	SELECT 
    					campaign.id, 
						campaign.name 
					FROM 
						campaign
					WHERE
					    campaign.id = ` + campaignId + `
					LIMIT 1`,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	row, err := resp.Next()
	if errors.Is(err, iterator.Done) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Campaign not found."})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	campaignResp := row.GetCampaign()
	campaign := Campaign{}
	campaign.Id = *campaignResp.Id
	if campaignResp.Name != nil {
		campaign.Name = *campaignResp.Name
	}
	campaign.ResourceName = campaignResp.ResourceName

	c.JSON(http.StatusOK, gin.H{"campaign": campaign})
}
