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
	ResourceName string `json:"resource_name" binding:"required"`
	Company      string `json:"company" binding:"required"`
}

// CreateCampaign godoc
// @Summary Create a campaign
// @Description creates a campaign for certain company/user
// @Tags Campaign
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
// @Success 201 {object} []models.Campaign
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
	if strings.Trim(request.ResourceName, " ") == "" || strings.Trim(request.Company, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Get company ID
	company, err := models.GetCompanyByName(request.Company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "That company does not exist"})
		return
	}

	campaign := models.Campaign{
		ResourceName: request.ResourceName,
		CompanyID:    company.ID,
		Company:      *company,
	}

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

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
