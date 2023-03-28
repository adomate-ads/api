package campaign

import (
	"bytes"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/adomate-ads/api/pkg/email"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CreateRequest struct {
	Name            string `json:"name" binding:"required"`
	Company         string `json:"company" binding:"required"`
	BiddingStrategy string `json:"bidding_strategy" binding:"required"`
	Budget          uint   `json:"budget" binding:"required"`
}

// CreateCampaign godoc
// @Summary Create a campaign
// @Description creates a campaign for certain company/user
// @Tags Campaign
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "That bidding strategy does not exist"})
		return
	}

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

	data := email.NewCampaign{
		Company:   company.Name,
		Campaign:  campaign.Name,
		StartDate: time.Now().Format("2006-01-02"),
	}
	body := new(bytes.Buffer)
	if err := email.Templates["new-campaign"].Tmpl.Execute(body, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	email.SendEmail(company.Email, email.Templates["new-campaign"].Subject, body.String())

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
// @Param id path int true "Campaign ID"
// @Success 200 {object} []models.Campaign
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /campaign/company/{id} [get]
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

// GetCampaign godoc
// @Summary Gets a Campaign
// @Description Gets all information about a single campaign.
// @Tags Campaign
// @Accept */*
// @Produce json
// @Param id path int true "Campaign ID"
// @Success 200 {object} models.Campaign
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /campaign/{id} [get]
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

// DeleteCampaign godoc
// @Summary Delete Campaign
// @Description Delete a campaign.
// @Tags Campaign
// @Accept */*
// @Produce json
// @Param id path int true "Campaign ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /campaign/{id} [delete]
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

	data := email.DeleteCampaign{
		Company:  campaign.Company.Name,
		Campaign: campaign.Name,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
	}
	body := new(bytes.Buffer)
	if err := email.Templates["delete-campaign"].Tmpl.Execute(body, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	email.SendEmail(campaign.Company.Email, email.Templates["delete-campaign"].Subject, body.String())

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
