package campaign

import (
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CreateRequest struct {
	Name            string `json:"name" binding:"required"`
	Company         string `json:"company" binding:"required"`
	BiddingStrategy string `json:"bidding_strategy" binding:"required"`
	Budget          uint   `json:"budget" binding:"required"`
}

func CreateCampaign(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Make sure that the user can only create a campaign in their company.
	user := c.MustGet("x-user").(*models.User)
	if user.Company.Name != request.Company && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only create a campaign for your company"})
		return
	}

	// Validate form input
	if strings.Trim(request.Name, " ") == "" || strings.Trim(request.Company, " ") == "" || strings.Trim(request.BiddingStrategy, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Get company ID
	company, err := models.GetCompanyByName(request.Company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "That company does not exist"})
		return
	}

	// Get bidding strategy ID
	biddingStrategy, err := models.GetBiddingStrategyByName(request.BiddingStrategy)

	campaign := models.Campaign{
		Name:              request.Name,
		CompanyID:         company.ID,
		Company:           *company,
		Budget:            request.Budget,
		BiddingStrategyID: biddingStrategy.ID,
		BiddingStrategy:   *biddingStrategy,
		Keywords:          []models.Keyword{},
	}

	// TODO - Fetch and fill keywords

	if err := campaign.CreateCampaign(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully registered campaign"})
}

func GetCampaigns(c *gin.Context) {
	campaigns, err := models.GetCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

func GetCampaignsForCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Make sure that the user can only get information about campaigns from the company they're in.
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != uint(companyID) && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get campaigns for your company"})
		return
	}

	campaigns, err := models.GetCampaignsByCompanyID(uint(companyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

func GetCampaign(c *gin.Context) {
	id := c.Param("id")
	campaignID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	campaign, err := models.GetCampaign(uint(campaignID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Make sure that the user can only get information about a campaign from the company they're in.
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != campaign.CompanyID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get a campaign from your company"})
		return
	}

	c.JSON(http.StatusOK, campaign)
}

func DeleteCampaign(c *gin.Context) {
	id := c.Param("id")
	campaignID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	campaign, err := models.GetCampaign(uint(campaignID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Make sure that the user can only delete a campaign in their company.
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != campaign.CompanyID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete a campaign from your company"})
		return
	}

	if err := campaign.DeleteCampaign(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
