package gads

import (
	"errors"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"net/http"
	"strconv"
)

type AdGroupAd struct {
	Id           int64    `json:"id"`
	Name         string   `json:"name"`
	ResourceName string   `json:"resource_name"`
	FinalURL     []string `json:"final_url"`
}

// GetAdGroupAds Google Ads godoc
// @Summary Get Google Ads AdGroupAd
// @Description Gets all information about specific Google AdGroupAd
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param clientId path int true "Client ID"
// @Param campaignId path int true "Campaign ID"
// @Param adgroupId path int true "AdGroup ID"
// @Success 200 {object} AdGroupAd
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/adgroupad/{clientId}/{campaignId}/{adgroupId} [get]
func GetAdGroupAds(c *gin.Context) {
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
						ad_group_ad.ad.id,
						ad_group_ad.ad.name,
						ad_group_ad.ad.final_urls,
						ad_group_ad.resource_name
					FROM 
						ad_group_ad
					WHERE
					    ad_group.id = ` + adGroupId,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	var adGroupAds []AdGroupAd

	for {
		row, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		adGroupAdResp := row.GetAdGroupAd()
		adGroupAd := AdGroupAd{}
		adGroupAd.Id = *adGroupAdResp.Ad.Id
		if adGroupAdResp.Ad.Name != nil {
			adGroupAd.Name = *adGroupAdResp.Ad.Name
		}
		adGroupAd.ResourceName = adGroupAdResp.ResourceName
		adGroupAd.FinalURL = adGroupAdResp.Ad.FinalUrls

		adGroupAds = append(adGroupAds, adGroupAd)
	}

	c.JSON(http.StatusOK, gin.H{"adgroupads": adGroupAds})
}
