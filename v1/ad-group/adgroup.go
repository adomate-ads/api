package campaign

import (
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAdGroups godoc
// @Summary Get all adgroups
// @Description Get a slice of all adgroups
// @Tags AdGroup
// @Accept */*
// @Produce json
// @Success 200 {object} []models.AdGroup
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /adgroup [get]
func GetAdGroups(c *gin.Context) {
	adGroups, err := models.GetAdGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adGroups)
}

// GetAdGroupsForCompany godoc
// @Summary Get all adgroup for a company
// @Description get a slice of all adgroup for certain company
// @Tags AdGroup
// @Accept */*
// @Produce json
// @Param id path int true "AdGroup ID"
// @Success 200 {object} []models.AdGroup
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /adgroup/company/{id} [get]
func GetAdGroupsForCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Make sure that the user can only get information about campaigns from the company they're in.
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != uint(companyID) && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get adgroup for your company"})
		return
	}

	adgroups, err := models.GetAdGroupsByCompanyID(uint(companyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, adgroups)
}

// GetAdGroup godoc
// @Summary Gets a AdGroup
// @Description Gets all information about a single adgroup.
// @Tags AdGroup
// @Accept */*
// @Produce json
// @Param id path int true "AdGroup ID"
// @Success 200 {object} models.AdGroup
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /campaign/{id} [get]
func GetAdGroup(c *gin.Context) {
	id := c.Param("id")
	adGroupId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	adGroup, err := models.GetAdGroup(uint(adGroupId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Make sure that the user can only get information about a campaign from the company they're in.
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != adGroup.CompanyID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get a adgroups from your company"})
		return
	}

	c.JSON(http.StatusOK, adGroup)
}
