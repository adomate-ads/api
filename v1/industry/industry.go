package industry

import (
	"github.com/adomate-ads/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RegisterRequest struct {
	Industry string `json:"industry" binding:"required"`
}

func Register(c *gin.Context) {
	var request RegisterRequest
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created industry"})
}

func GetIndustries(c *gin.Context) {
	industries, err := models.GetIndustries()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"industries": industries})
}

// Im not sure if we should do this by name or ID
func GetIndustry(c *gin.Context) {
	industry, err := models.GetIndustryByName(c.Param("industry"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"industry": industry})
}
