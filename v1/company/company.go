package company

import (
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/adomate-ads/api/pkg/email"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Industry string `json:"industry" binding:"required"`
	Domain   string `json:"domain" binding:"required"`
}

// CreateCompany godoc
// @Summary Create Company
// @Description creates a company that can start campaigns, etc
// @Tags Company
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
// @Success 200 {object} []models.Company
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /company [post]
func CreateCompany(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Name, " ") == "" || strings.Trim(request.Email, " ") == "" || strings.Trim(request.Industry, " ") == "" || strings.Trim(request.Domain, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if company already exists
	_, err := models.GetCompanyByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An company by that email already exists"})
		return
	}

	// Get Industry ID
	industry, err := models.GetIndustryByName(request.Industry)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An industry by that name does not exist"})
		return
	}

	// Create company
	company := models.Company{
		Name:       request.Name,
		Email:      request.Email,
		IndustryID: industry.ID,
		Industry:   *industry,
		Domain:     request.Domain,
	}

	if err := company.CreateCompany(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	email.SendEmail(company.Email, email.Templates["registration"].Subject, email.Templates["registration"].Body)

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully registered company"})
}

// GetCompanies godoc
// @Summary Get all companies
// @Description Get a slice of all companies
// @Tags Company
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Company
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /company [get]
func GetCompanies(c *gin.Context) {
	companies, err := models.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

// GetCompany godoc
// @Summary Get Company
// @Description Get a company.
// @Tags Company
// @Accept */*
// @Produce json
// @Param id path int true "Company ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /company/{id} [get]
func GetCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Make sure that the user can only get information about the company that they're from.
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != uint(companyID) && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get information about your company"})
		return
	}

	company, err := models.GetCompany(uint(companyID))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company doesn't exist"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		}
		return
	}

	c.JSON(http.StatusOK, company)
}

// DeleteCompany godoc
// @Summary Delete Company
// @Description Delete a company.
// @Tags Company
// @Accept */*
// @Produce json
// @Param id path int true "Company ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /company/{id} [delete]
func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	company, err := models.GetCompany(uint(companyID))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company doesn't exist"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		}
		return
	}

	if err := company.DeleteCompany(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	email.SendEmail(company.Email, email.Templates["delete-company"].Subject, email.Templates["delete-company"].Body)

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
