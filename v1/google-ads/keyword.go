package gads

import (
	"errors"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"net/http"
)

type Keyword struct {
	Id      int64  `json:"id"`
	Keyword string `json:"keyword"`
}

// GetKeywords Google Ads godoc
// @Summary Get Google Ads Keyword
// @Description Gets all information about specific Google Keyword
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param id path int true "Client ID"
// @Param id path int true "Campaign ID"
// @Param id path int true "AdGroup ID"
// @Success 200 {object} Keyword
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/keyword/{clientId}/{campaignId}/{adgroupId} [get]
func GetKeywords(c *gin.Context) {
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

	request := services.SearchGoogleAdsRequest{
		CustomerId: clientId,
		Query: `	SELECT 
    					ad_group.id, 
						ad_group_criterion.criterion_id,
						ad_group_criterion.keyword.text
					FROM 
						ad_group_criterion
					WHERE
					    ad_group.id = ` + adGroupId,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	var keywords []Keyword

	for {
		row, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		keywordResp := row.GetAdGroupCriterion()
		keyword := Keyword{}
		keyword.Id = *keywordResp.CriterionId
		if keywordResp.GetKeyword().Text != nil {
			keyword.Keyword = *keywordResp.GetKeyword().Text
		}

		keywords = append(keywords, keyword)
	}

	c.JSON(http.StatusOK, gin.H{"keywords": keywords})
}
