package campaign

import (
	"github.com/adomate-ads/api/models"
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

// CreateCampaign godoc
// @Summary Create add campaign
// @Description creates a campaign for certain company/user
// @Tags Campaign
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Campaign
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /campaign [post]
func CreateCampaign(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

// GetCampaigns godoc
// @Summary Get all campaigns
// @Description Get a slice of all campaigns
// @Tags Campaign
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Campaign
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /campaign [get]
func GetCampaigns(c *gin.Context) {
	campaigns, err := models.GetCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

// GetCampaignsForCompany godoc
// @Summary Get all campaigns for a company
// @Description get a slice of all campaigns for certain company
// @Tags Campaign
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Campaign
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /campaign/company/:id [get]
func GetCampaignsForCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	campaigns, err := models.GetCampaignsByCompanyID(uint(companyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

// GetCampaign godoc
// @Summary Gets a campaign 
// @Description Gets all information about specific campaign
// @Tags Campaign
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Campaign
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /campaign/:id [get]
func GetCampaign(c *gin.Context) {
	id := c.Param("id")
	campaignID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	campaign, err := models.GetCompany(uint(campaignID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

// DeleteCampaign godoc
// @Summary Delete Campaign
// @Description Delete a campaign.
// @Tags Campaign
// @Accept */*
// @Produce json
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /campaign/:id [delete]
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

	if err := campaign.DeleteCampaign(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
