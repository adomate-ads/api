package preregistration

import (
	"github.com/adomate-ads/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type CreateRequest struct {
	Domain string `json:"domain" binding:"required"`
}

// CreatePreRegistration Preregistration godoc
// @Summary Create Preregistration
// @Description preregisters a company
// @Tags Preregister
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
// @Success 201 {object} []models.PreRegistration
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /preregistration [post]
func CreatePreRegistration(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Domain, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if domain is already registered
	if _, err := models.GetPreRegistrationByDomain(request.Domain); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Domain already exists"})
		return
	}

	pr := models.PreRegistration{
		Domain: request.Domain,
	}

	if err := pr.CreatePreRegistration(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created pre-registration."})

	// TODO: - Start fetching the locations & services
}

// GetPreRegistrations godoc
// @Summary Get all preregistration
// @Description Get a slice of all preregistration
// @Tags Preregister
// @Accept */*
// @Produce json
// @Success 200 {object} []models.PreRegistration
// @Failure 500 {object} dto.ErrorResponse
// @Router /preregistration [get]
func GetPreRegistrations(c *gin.Context) {
	pr, err := models.GetPreRegistrations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pr)
}

// GetPreRegistration godoc
// @Summary Get Preregistration
// @Description Get a preregistration domain.
// @Tags Preregister
// @Accept */*
// @Produce json
// @Param domain path string true "Preregistered Domain"
// @Success 200 {object} []models.PreRegistration
// @Failure 404 {object} dto.ErrorResponse
// @Router /preregistration/{domain} [get]
func GetPreRegistration(c *gin.Context) {
	domain := c.Param("domain")

	pr, err := models.GetPreRegistrationByDomain(domain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pre-registration not found"})
		return
	}

	c.JSON(http.StatusOK, pr)
}

type LocationRequest struct {
	Domain   string   `json:"domain" binding:"required"`
	Location []string `json:"location" binding:"required"`
}

// AddLocations godoc
// @Summary Add Location
// @Description Add a location to preregistration domain
// @Tags Preregister
// @Accept json
// @Produce json
// @Param create body LocationRequest true "Location Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /preregistration/locations [post]
func AddLocations(c *gin.Context) {
	var request LocationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if len(request.Location) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if domain exists
	pr, err := models.GetPreRegistrationByDomain(request.Domain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Domain not found"})
		return
	}

	// Check if locations already exist if not create them
	for _, location := range request.Location {
		exists := false
		for _, l := range pr.Locations {
			if l.Location == location {
				exists = true
				break
			}
		}
		if !exists {
			prl := models.PreRegLocation{
				PreRegistrationID: pr.ID,
				Location:          location,
			}
			if err := prl.CreatePreRegLocation(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created pre-registration locations."})
}

// GetLocations godoc
// @Summary Get Locations
// @Description Get locations for preregistration domain
// @Tags Preregister
// @Accept */*
// @Produce json
// @Param string path string true "Domain Name"
// @Success 200 {object} dto.MessageResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /preregistration/locations [get]
func GetLocations(c *gin.Context) {
	domain := c.Param("domain")

	pr, err := models.GetPreRegistrationByDomain(domain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pre-registration not found"})
		return
	}

	c.JSON(http.StatusOK, pr.Locations)
}

// DeleteLocations godoc
// @Summary Delete Locations
// @Description Delete locations from domain.
// @Tags Preregister
// @Accept json
// @Produce json
// @Param create body LocationRequest true "Location Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /preregistration/locations [delete]
func DeleteLocations(c *gin.Context) {
	var request LocationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if len(request.Location) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if domain exists
	pr, err := models.GetPreRegistrationByDomain(request.Domain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Domain not found"})
		return
	}

	// Check if locations already exist if not create them
	for _, location := range request.Location {
		for _, l := range pr.Locations {
			if l.Location == location {
				if err := l.DeletePreRegLocation(); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully deleted pre-registration locations."})
}

type ServiceRequest struct {
	Domain  string   `json:"domain" binding:"required"`
	Service []string `json:"service" binding:"required"`
}

// AddServices godoc
// @Summary Add Services
// @Description Add services to preregistration domain
// @Tags Preregister
// @Accept json
// @Produce json
// @Param create body ServiceRequest true "Service Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /preregistration/services [post]
func AddServices(c *gin.Context) {
	var request ServiceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if len(request.Service) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if domain exists
	pr, err := models.GetPreRegistrationByDomain(request.Domain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Domain not found"})
		return
	}

	// Check if services already exist if not create them
	for _, service := range request.Service {
		exists := false
		for _, s := range pr.Services {
			if s.Service == service {
				exists = true
				break
			}
		}
		if !exists {
			prs := models.PreRegService{
				PreRegistrationID: pr.ID,
				Service:           service,
			}
			if err := prs.CreatePreRegService(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created pre-registration services."})
}

// GetServices godoc
// @Summary Get Services
// @Description Get services for preregistration domain
// @Tags Preregister
// @Accept */*
// @Produce json
// @Param string path string true "Domain Name"
// @Success 200 {object} dto.MessageResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /preregistration/services [get]
func GetServices(c *gin.Context) {
	domain := c.Param("domain")

	pr, err := models.GetPreRegistrationByDomain(domain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pre-registration not found"})
		return
	}

	c.JSON(http.StatusOK, pr.Services)
}

// DeleteServices godoc
// @Summary Delete Services
// @Description Delete Services from domain.
// @Tags Preregister
// @Accept json
// @Produce json
// @Param create body ServiceRequest true "Service Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /preregistration/services [delete]
func DeleteServices(c *gin.Context) {
	var request ServiceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if len(request.Service) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if domain exists
	pr, err := models.GetPreRegistrationByDomain(request.Domain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Domain not found"})
		return
	}

	// Check if services already exist if not create them
	for _, service := range request.Service {
		for _, s := range pr.Services {
			if s.Service == service {
				if err := s.DeletePreRegService(); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully deleted pre-registration services."})
}

type BudgetRequest struct {
	Domain string `json:"domain" binding:"required"`
	Budget uint   `json:"budget" binding:"required"`
}

// SetBudget godoc
// @Summary Set Budget
// @Description Set Budget of a domain.
// @Tags Preregister
// @Accept json
// @Produce json
// @Param create body BudgetRequest true "Budget Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /preregistration/budget [post]
func SetBudget(c *gin.Context) {
	var request BudgetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if request.Budget == 0 || request.Domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if domain exists
	pr, err := models.GetPreRegistrationByDomain(c.Param("domain"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Domain not found"})
		return
	}

	pr.Budget = request.Budget

	if err := pr.UpdatePreRegistration(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully set budget."})
}
