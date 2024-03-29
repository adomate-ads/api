package gads

import (
	"errors"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/helpers"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"net/http"
	"strconv"
)

type Campaign struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ResourceName string `json:"resource_name"`
}

// GetCampaigns Google Ads godoc
// @Summary Get Google Ads Campaigns
// @Description Gets all Google Ads Campaigns
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Success 200 {object} []helpers.Campaign
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/campaigns/ [get]
func GetCampaigns(c *gin.Context) {
	// TODO - Refactor into helpers
	user := c.MustGet("x-user").(*models.User)
	if !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client ID is required."})
		return
	}

	request := services.SearchGoogleAdsRequest{
		CustomerId: google_ads.SuperUser,
		Query:      "SELECT customer_client.id FROM customer_client",
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

		customerClient := row.GetCustomerClient()
		if customerClient == nil {
			continue
		}

		clientId := *customerClient.Id

		request := services.SearchGoogleAdsRequest{
			CustomerId: strconv.Itoa(int(clientId)),
			Query:      "SELECT campaign.id, campaign.name, campaign.resource_name FROM campaign ORDER BY campaign.id",
		}

		resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

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
			campaign.ResourceName = campaignResp.ResourceName

			campaigns = append(campaigns, campaign)
		}
	}

	c.JSON(http.StatusOK, gin.H{"campaigns": campaigns})
}

// GetCampaignsInClient Google Ads godoc
// @Summary Get Google Ads Campaigns
// @Description Gets all Google Ads Campaigns for specific Client
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param clientId path int true "Client ID"
// @Success 200 {object} []helpers.Campaign
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/campaigns/{clientId} [get]
func GetCampaignsInClient(c *gin.Context) {
	clientId := c.Param("clientId")
	if clientId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client ID is required."})
		return
	}

	//TODO - Only allow the user to get the campaigns attached to their company

	campaigns := helpers.GetCampaigns(clientId)

	c.JSON(http.StatusOK, gin.H{"campaigns": campaigns})
}

// GetCampaign Google Ads godoc
// @Summary Get Google Ads Campaign
// @Description Gets all information about specific Google Ads Campaign
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param clientId path int true "Client ID"
// @Param campaignId path int true "Campaign ID"
// @Success 200 {object} helpers.Campaign
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/campaign/{clientId}/{campaignId} [get]
func GetCampaign(c *gin.Context) {
	clientId := c.Param("clientId")
	campaignId := c.Param("campaignId")
	if clientId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Client ID is required."})
		return
	}
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Campaign ID is required."})
		return
	}

	//TODO - Only allow the user to get the campaigns attached to their company

	campaign := helpers.GetCampaign(clientId, campaignId)

	c.JSON(http.StatusOK, gin.H{"campaign": campaign})
}
