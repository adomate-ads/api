package gads

import (
	"errors"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"github.com/adomate-ads/api/pkg/google-ads/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdGroup struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ResourceName string `json:"resource_name"`
}

// GetAdGroupsInCampaign Google Ads godoc
// @Summary Get Google Ads AdGroup
// @Description Gets all Google Ads Groups for specific Campaign
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param clientId path int true "Client ID"
// @Param campaignId path int true "Campaign ID"
// @Success 200 {object} []AdGroup
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/adgroup/{clientId}/{campaignId} [get]
func GetAdGroupsInCampaign(c *gin.Context) {
	clientId := c.Param("clientId")
	campaignId := c.Param("campaignId")
	if clientId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client ID is required."})
		return
	}
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Campaign ID is required."})
		return
	}

	//TODO - Only allow the user to get the campaigns attached to their company
	campaignNumber, err := strconv.ParseUint(campaignId, 10, 64)
	campaign, _ := models.GetCampaign(uint(campaignNumber))
	companyId := campaign.CompanyID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != companyId && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get information about your company"})
		return
	}
	request := services.SearchGoogleAdsRequest{
		CustomerId: clientId,
		Query:      `SELECT campaign.id, ad_group.id, ad_group.name, ad_group.resource_name FROM ad_group WHERE campaign.id = ` + campaignId + ` ORDER BY campaign.id`,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	var adGroups []AdGroup

	for {
		row, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
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


	adGroups := helpers.GetAdGroups(clientId, campaignId)

	c.JSON(http.StatusOK, gin.H{"adgroups": adGroups})
}

// GetAdGroup Google Ads godoc
// @Summary Get Google Ads AdGroup
// @Description Gets all information about specific Google AdGroup
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param clientId path int true "Client ID"
// @Param campaignId path int true "Campaign ID"
// @Param adgroupId path int true "AdGroup ID"
// @Success 200 {object} AdGroup
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/adgroup/{clientId}/{campaignId}/{adgroupId} [get]
func GetAdGroup(c *gin.Context) {
	clientId := c.Param("clientId")
	campaignId := c.Param("campaignId")
	adGroupId := c.Param("adgroupId")
	if clientId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Client ID is required."})
		return
	}
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Campaign ID is required."})
		return
	}
	if adGroupId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "AdGroup ID is required."})
		return
	}

	//TODO - Only allow the user to get the campaigns attached to their company
	campaignNumber, err := strconv.ParseUint(campaignId, 10, 64)
	campaign, _ := models.GetCampaign(uint(campaignNumber))
	companyId := campaign.CompanyID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != companyId && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get information about your company"})
		return
	}
	request := services.SearchGoogleAdsRequest{
		CustomerId: clientId,
		Query: `	SELECT 
    					ad_group.id, 
						ad_group.name,
						ad_group.resource_name
					FROM 
						ad_group
					WHERE
					    ad_group.id = ` + adGroupId + `
					LIMIT 1`,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	row, err := resp.Next()
	if errors.Is(err, iterator.Done) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "AdGroup not found."})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	adGroupResp := row.GetAdGroup()
	adGroup := AdGroup{}
	adGroup.Id = *adGroupResp.Id
	if adGroupResp.Name != nil {
		adGroup.Name = *adGroupResp.Name
	}
	adGroup.ResourceName = adGroupResp.ResourceName


	adGroup := helpers.GetAdGroup(clientId, adGroupId)

	c.JSON(http.StatusOK, gin.H{"adgroup": adGroup})
}
