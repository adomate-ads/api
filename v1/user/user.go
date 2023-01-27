package user

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
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Company   string `json:"company" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

// CreateUser godoc
// @Summary Create User
// @Description Create a new user.
// @Tags User
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
// @Success 201 {object} []models.User
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user is an Admin of a company, if so make sure they're only adding the member to their company
	user := c.MustGet("x-user").(*models.User)
	if auth.InGroup(user, "admin") {
		if user.Company.Name != request.Company {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only add users to your company"})
			return
		}
	}

	// Validate form input
	if strings.Trim(request.FirstName, " ") == "" || strings.Trim(request.LastName, " ") == "" || strings.Trim(request.Email, " ") == "" || strings.Trim(request.Password, " ") == "" || strings.Trim(request.Company, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if user already exists
	_, err := models.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An account by that email already exists"})
		return
	}

	// Get company ID
	company, err := models.GetCompanyByName(request.Company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "That company does not exist"})
		return
	}

	// Create user
	u := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		CompanyID: company.ID,
		Company:   *company,
		Role:      request.Role,
	}
	if err := u.CreateUser(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	email.SendEmail(company.Email, email.Templates["new-user-notification"].Subject, email.Templates["new-user-notification"].Body)
	email.SendEmail(u.Email, email.Templates["new-user"].Subject, email.Templates["new-user"].Body)

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created user"})
}

// GetUsers godoc
// @Summary Get all Users
// @Description Gets a slice of all users.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} []models.User
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user [get]
func GetUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUsersByCompanyID godoc
// @Summary Get all Users for a Company
// @Description Gets a slice of all the users for a specific company.
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.User
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /user/company/{id} [get]
func GetUsersByCompanyID(c *gin.Context) {
	id := c.Param("id")
	CompanyID, err := strconv.ParseUint(id, 10, 64)

	// Check if the user is an Admin of a company, if so make sure they're only able to see the members in their company
	user := c.MustGet("x-user").(*models.User)
	if auth.InGroup(user, "admin") {
		if user.CompanyID != uint(CompanyID) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only get users in your company"})
			return
		}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := models.GetUsersByCompanyID(uint(CompanyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser godoc
// @Summary Gets a User
// @Description Gets all information about a single user.
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.GetUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the user is an Admin of a company, if so make sure they're only getting the members of their company
	xUser := c.MustGet("x-user").(*models.User)
	if auth.InGroup(xUser, "admin") || xUser.ID == user.ID {
		if xUser.CompanyID != user.CompanyID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only get users from your company"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser godoc
// @Summary Update User
// @Description Update information about a user.
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 202 {object} models.User
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user/{id} [patch]
func UpdateUser(c *gin.Context) {

}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete a user.
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the user is an Admin of a company, if so make sure they're only deleting the members of their company
	xUser := c.MustGet("x-user").(*models.User)
	if auth.InGroup(xUser, "admin") {
		if xUser.CompanyID != user.CompanyID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only get users from your company"})
			return
		}
	}
	// Check if the user is trying to delete themselves
	if xUser.ID == user.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't delete yourself"})
		return
	}

	if err := user.DeleteUser(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	email.SendEmail(user.Company.Email, email.Templates["delete-user-notification"].Subject, email.Templates["delete-user-notification"].Body)
	email.SendEmail(user.Email, email.Templates["delete-user"].Subject, email.Templates["delete-user"].Body)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
