package role

import (
	"github.com/adomate-ads/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CreateRequest struct {
	Role string `json:"role" binding:"required"`
}

func CreateRole(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Role, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if role already exists
	_, err := models.GetRoleByName(request.Role)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A role by that name already exists"})
		return
	}

	// Create industry
	role := models.Role{
		Role: request.Role,
	}

	if err := role.CreateRole(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created role"})
}

func GetRoles(c *gin.Context) {
	roles, err := models.GetRoles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

// Im not sure if we should do this by name or ID
func GetRole(c *gin.Context) {
	role, err := models.GetRoleByName(c.Param("role"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	role, err := models.GetRole(uint(roleID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := role.DeleteRole(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
