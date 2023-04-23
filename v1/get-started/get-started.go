package get_started

import (
	"bytes"
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/email"
	stripe_pkg "github.com/adomate-ads/api/pkg/stripe"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
	"net/http"
	"strconv"
	"strings"
)

type CreateAccountRequest struct {
	FirstName   string `json:"first_name" binding:"required" example:"John"`
	LastName    string `json:"last_name" binding:"required" example:"Doe"`
	Email       string `json:"email" binding:"required" example:"johndoe@adomate.ai"`
	CompanyName string `json:"company_name" binding:"required" example:"Adomate"`
	Industry    string `json:"industry" binding:"required" example:"Software"`
	Domain      string `json:"domain" binding:"required" example:"adomate.ai"`
}

// CreateAccount godoc
// @Summary Create Account
// @Description Create account for user
// @Tags Getting Started
// @Accept json
// @Produce json
// @Param create body CreateAccountRequest true "Create Account Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /get-started [post]
func CreateAccount(c *gin.Context) {
	var request CreateAccountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate User form input
	if strings.Trim(request.FirstName, " ") == "" || strings.Trim(request.LastName, " ") == "" || strings.Trim(request.Email, " ") == "" {
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

	// Check if company already exists by name - TODO should we allow this?
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
		Password:  "",
		CompanyID: newCompany.ID,
		Company:   *newCompany,
		Role:      "owner",
	}

	if err := u.CreateUser(); err != nil {
		err := newCompany.DeleteCompany()
		if err != nil {
			msg := fmt.Sprintf("Failed to delete company %s after failed user creation", newCompany.Name)
			suggestion := fmt.Sprintf("Delete company %s manually and email %s.", newCompany.Name, u.Email)
			discord.SendMessage(discord.Error, msg, suggestion)

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("user-id", u.ID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	params := &stripe.CustomerParams{
		Name:  stripe.String(request.CompanyName),
		Email: stripe.String(request.Email),
	}
	params.AddMetadata("company_id", strconv.Itoa(int(newCompany.ID)))

	stripeCustomer, err := customer.New(params)
	if err != nil {
		msg := fmt.Sprintf("Failed to create a stripe customer for company %s", newCompany.Name)
		suggestion := fmt.Sprintf("Create Stripe Customer, Name:%s, Email:%s, CompanyID:%d", request.CompanyName, request.Email, newCompany.ID)
		discord.SendMessage(discord.Error, msg, suggestion)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newCompany.StripeID = stripeCustomer.ID
	if _, err := newCompany.UpdateCompany(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send welcome Email
	data := email.WelcomeData{
		FirstName: u.FirstName,
		Company:   u.Company.Name,
		Domain:    u.Company.Domain,
	}
	body := new(bytes.Buffer)
	if err := email.Templates["register"].Tmpl.Execute(body, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	email.SendEmail(u.Email, email.Templates["register"].Subject, body.String())

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created user and company"})
	// TODO - In the future, we should send an email to the user with a link to verify their email address
	// TODO - In the future, we should possibly send a session token back to the user
}

type CreatePaymentIntentRequest struct {
	Email  string `json:"email" binding:"required"`
	Amount int64  `json:"amount" binding:"required"`
}

func CreatePaymentIntent(c *gin.Context) {
	var request CreatePaymentIntentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Email, " ") == "" || request.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check if company exists
	company, err := models.GetCompanyByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A company by that email does not exist"})
		return
	}

	clientSecret, err := stripe_pkg.CreatePaymentIntent(request.Email, company.StripeID, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating payment."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"client_secret": clientSecret})
}

type LocsAndSers struct {
	Locations []string `json:"locations"`
	Services  []string `json:"services"`
}

// GetLocationsAndServices godoc
// @Summary Get Locations and Services
// @Description Get Locations and Services for domain
// @Tags Getting Started
// @Accept json
// @Produce json
// @Param domain path string true "Domain URL"
// @Success 201 {object} []LocsAndSers
// @Router /get-started/location-services/{domain} [get]
func GetLocationsAndServices(c *gin.Context) {
	// TODO - Uncomment this when microservice is done
	//domain := c.Param("domain")
	//
	//locations, services, err := website_parse.GetLocAndSer(domain)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	locations := []string{"Houston, TX", "Dallas, TX"}
	services := []string{"Ad Bot Blocking", "Automatic Google Ads"}

	locationsAndServices := LocsAndSers{
		Locations: locations,
		Services:  services,
	}

	c.JSON(http.StatusOK, locationsAndServices)
}
