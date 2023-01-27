package user

import (
	"github.com/adomate-ads/api/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Login User
// @Description Login using user credentials.
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	session := sessions.Default(c)
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Email, " ") == "" || strings.Trim(request.Password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match
	u, err := models.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An account by that email does not exist"})
		return
	}

	if err := models.VerifyPassword(request.Password, u.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	// Save the ID in the session
	session.Set("user-id", u.ID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

type RegisterRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Industry    string `json:"industry" binding:"required"`
	Domain      string `json:"domain" binding:"required"`
	Budget      uint   `json:"budget" binding:"required"`
}

// Register godoc
// @Summary Register New User
// @Description Registers a new user.
// @Tags Auth
// @Accept json
// @Param register body RegisterRequest true "Register Request"
// @Produce json
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate User form input
	if strings.Trim(request.FirstName, " ") == "" || strings.Trim(request.LastName, " ") == "" || strings.Trim(request.Email, " ") == "" || strings.Trim(request.Password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}
	// Validate Company form input
	if strings.Trim(request.CompanyName, " ") == "" || strings.Trim(request.Industry, " ") == "" || strings.Trim(request.Domain, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if user already exists
	_, err := models.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An account by that email already exists"})
		return
	}

	// Check if company already exists by name
	_, err = models.GetCompanyByName(request.CompanyName)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A company by that name already exists"})
		return
	}

	// Check if company already exists by email
	_, err = models.GetCompanyByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A company by that email already exists"})
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
		Name:       request.CompanyName,
		Email:      request.Email,
		IndustryID: industry.ID,
		Industry:   *industry,
		Domain:     request.Domain,
		Budget:     request.Budget,
	}

	if err := company.CreateCompany(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newCompany, err := models.GetCompanyByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create user
	u := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		CompanyID: newCompany.ID,
		Company:   *newCompany,
		Role:      "owner",
	}

	//TODO - If we run into this error, that means it created the company but not the user, so we actually need to delete the company now and return an error.
	if err := u.CreateUser(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//params := &stripe.CustomerParams{
	//	Name:  stripe.String(request.CompanyName),
	//	Email: stripe.String(request.Email),
	//}
	//params.AddMetadata("company_id", strconv.Itoa(int(newCompany.ID)))
	//
	////TODO - If we run into this error, that means that we created the user and company, but not the stripe customer, so we actually need to delete the user and company now and return an error.
	//if _, err := customer.New(params); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created user and company"})
	// TODO - In the future, we should send an email to the user with a link to verify their email address
	// TODO - In the future, we should possibly send a session token back to the user
}

// Logout godoc
// @Summary Logout User
// @Description Logout of a user.
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /logout [post]
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user-id")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete("user-id")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func Me(c *gin.Context) {
	//session := sessions.Default(c)
	//user := session.Get("user-id")
	user, _ := c.Get("x-user")
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
