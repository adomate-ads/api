package service

import (
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CreateRequest struct {
	Name    string `json:"name" binding:"required"`
	Company string `json:"company" binding:"required"`
}

// CreateService godoc
// @Summary Create Service
// @Description Create a new service.
// @Tags Service
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /service [post]
func CreateService(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Company, " ") == "" || strings.Trim(request.Name, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Get company ID
	company, err := models.GetCompanyByName(request.Company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "That company does not exist"})
		return
	}

	s := models.Service{
		Name:      request.Name,
		CompanyID: company.ID,
		Company:   *company,
	}

	if err := s.CreateService(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created service"})
}

// GetServices godoc
// @Summary Get all service
// @Description Gets a slice of all service.
// @Tags Service
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Service
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /service [get]
func GetServices(c *gin.Context) {
	services, err := models.GetServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// GetServicesForCompany godoc
// @Summary Get all Services for a Company
// @Description Gets a slice of all the Services for a specific company.
// @Tags Service
// @Accept */*
// @Produce json
// @Param id path integer true "Company ID"
// @Success 200 {object} []models.Service
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /service/company/{id} [get]
func GetServicesForCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := c.MustGet("x-user").(*models.User)
	if auth.InGroup(user, "admin") {
		if user.CompanyID != uint(companyID) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only get service from your company"})
			return
		}
	}

	services, err := models.GetServicesByCompanyID(uint(companyID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// GetService godoc
// @Summary Gets a Service
// @Description Gets all information about a single Service.
// @Tags Service
// @Accept */*
// @Produce json
// @Param id path integer true "Service ID"
// @Success 200 {object} models.Service
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /service/{id} [get]
func GetService(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	service, err := models.GetService(uint(serviceID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != service.CompanyID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get service from your company"})
		return
	}

	c.JSON(http.StatusOK, service)
}

type UpdateRequest struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

// UpdateService godoc
// @Summary Update Service
// @Description Update information about a Service.
// @Tags Service
// @Accept  json
// @Produce json
// @Param update body CreateRequest true "Create Request"
// @Param id path integer true "Service ID"
// @Success 202 {object} models.Service
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /service/{id} [patch]
func UpdateService(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	service, err := models.GetService(uint(serviceID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var request UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get company ID
	if request.Company != "" {
		company, err := models.GetCompanyByName(request.Company)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "That company does not exist"})
			return
		}
		service.CompanyID = company.ID
		service.Company = *company
	}

	if request.Name != "" {
		service.Name = request.Name
	}

	updatedService, err := service.UpdateService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, updatedService)
}

// DeleteService godoc
// @Summary Delete Service
// @Description Delete a service.
// @Tags Service
// @Accept */*
// @Produce json
// @Param id path integer true "Service ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /service/{id} [delete]
func DeleteService(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	service, err := models.GetService(uint(serviceID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.DeleteService(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}
