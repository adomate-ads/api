package industry

import (
	"github.com/adomate-ads/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CreateRequest struct {
	Industry string `json:"industry" binding:"required"`
}

// Create Industry godoc
// @Summary Create Industry
// @Description creates an industry category
// @Tags Industry
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
// @Success 200 {object} []models.Industry
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /industry [post]
func CreateIndustry(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Industry, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if industry already exists
	_, err := models.GetIndustryByName(request.Industry)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An industry by that name already exists"})
		return
	}

	// Create industry
	industry := models.Industry{
		Industry: request.Industry,
	}

	if err := industry.CreateIndustry(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created industry"})
}

// GetIndustries godoc
// @Summary Get all industries
// @Description Get a slice of all industries
// @Tags Industry
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Industry
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /industry [get]
func GetIndustries(c *gin.Context) {
	industries, err := models.GetIndustries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"industries": industries})
}

// GetIndustry godoc
// @Summary Gets a industry
// @Summary Gets an industry by name
// @Description Gets all information about specific industry
// @Tags Industry
// @Accept */*
// @Produce json
// @Param industry path string true "Industry Name"
// @Success 200 {object} []models.Industry
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /industry/{industry} [get]
func GetIndustry(c *gin.Context) {
	industry, err := models.GetIndustryByName(c.Param("industry"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"industry": industry})
}

// DeleteIndustry godoc
// @Summary Delete Industry
// @Description Delete an industry.
// @Tags Industry
// @Accept */*
// @Produce json
// @Param id path int true "Industry ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /industry/{id} [delete]
func DeleteIndustry(c *gin.Context) {
	id := c.Param("id")
	industryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	industry, err := models.GetIndustry(uint(industryID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := industry.DeleteIndustry(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Industry deleted successfully"})
}
