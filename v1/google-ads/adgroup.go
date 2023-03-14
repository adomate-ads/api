package gads

import (
	"github.com/adomate-ads/api/pkg/google-ads/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
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

	adGroup := helpers.GetAdGroup(clientId, adGroupId)

	c.JSON(http.StatusOK, gin.H{"adgroup": adGroup})
}
