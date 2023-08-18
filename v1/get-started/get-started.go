package get_started

import (
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/email"
	google_ads_controller "github.com/adomate-ads/api/pkg/google-ads-controller"
	site_analyzer "github.com/adomate-ads/api/pkg/site-analyzer"
	stripe_pkg "github.com/adomate-ads/api/pkg/stripe"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	ipdata "github.com/ipdata/go"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type CreateAccountRequest struct {
	FirstName    string   `json:"first_name" binding:"required" example:"John"`
	LastName     string   `json:"last_name" binding:"required" example:"Doe"`
	Email        string   `json:"email" binding:"required" example:"johndoe@adomate.ai"`
	CompanyName  string   `json:"company_name" binding:"required" example:"Adomate"`
	Industry     string   `json:"industry" binding:"required" example:"Software"`
	Domain       string   `json:"domain" binding:"required" example:"adomate.ai"`
	Locations    []string `json:"locations" binding:"required" example:"[\"Houston, TX\"]"`
	Services     []string `json:"services" binding:"required" example:"[\"Google Ads\"]"`
	Headlines    []string `json:"headlines" binding:"required" example:"[\"Headline 1\", \"Headline 2\"]"`
	Descriptions []string `json:"descriptions" binding:"required" example:"[\"Description 1\", \"Description 2\"]"`
	Budget       uint     `json:"budget" binding:"required" example:"1000"`
	Price        string   `json:"price" binding:"required" example:"price_1MzQkOFzHmjFR1Qwa4QajKrY"`
	Ip           string   `json:"ip" binding:"required" example:"192.168.1.1"`
}

// CreateAccount godoc
// @Summary Create Account
// @Description Create account for user
// @Tags Getting Started
// @Accept json
// @Produce json
// @Param create body CreateAccountRequest true "Create Account Request"
// @Success 201 {object} stripe.Subscription
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

	// TODO should we allow this?
	// Check if company already exists by name
	//_, err = models.GetCompanyByName(request.CompanyName)
	//if err == nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "A company by that name already exists"})
	//	return
	//}

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

	// Create Locations
	for _, loc := range request.Locations {
		location := models.Location{
			Name:      loc,
			CompanyID: newCompany.ID,
			Company:   *newCompany,
		}
		if err := location.CreateLocation(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Create Services
	for _, serv := range request.Services {
		service := models.Service{
			Name:      serv,
			CompanyID: newCompany.ID,
			Company:   *newCompany,
		}
		if err := service.CreateService(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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
		Tax:   &stripe.CustomerTaxParams{IPAddress: stripe.String(request.Ip)},
	}
	params.AddExpand("tax")
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

	subscription, err := stripe_pkg.CreateSubscription(stripeCustomer.ID, request.Price, "Adomate - Initial Subscription", request.Budget)
	if err != nil {
		msg := fmt.Sprintf("Failed to create a stripe subscription for company %s", newCompany.Name)
		suggestion := fmt.Sprintf("Create Stripe Subscription, CustomerID:%s, PriceID:%s, CompanyID:%d", stripeCustomer.ID, request.Price, newCompany.ID)
		discord.SendMessage(discord.Error, msg, suggestion)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create Google Ads Account
	gadsCustomer, err := google_ads_controller.CreateCustomer(newCompany.Email)
	if err != nil {
		msg := fmt.Sprintf("Failed to create a google ads customer for company %s", newCompany.Email)
		suggestion := fmt.Sprintf("Create Google Ads Customer, Company Email:%s", request.Email)
		discord.SendMessage(discord.Error, msg, suggestion)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update company with Google Ads ID
	newCompany.GoogleAdsID = gadsCustomer.Id
	if _, err := newCompany.UpdateCompany(); err != nil {
		discord.SendMessage(discord.Error, "Failed to update company with Google Ads ID", fmt.Sprintf("Company ID: %d, should have gadsId: %d", newCompany.ID, gadsCustomer.Id))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create Google Ads Campaign
	campaignBody := google_ads_controller.Body{
		CustomerId:     gadsCustomer.Id,
		CampaignName:   "Initial Campaign",
		CampaignBudget: request.Budget,
	}
	gadsCampaign, err := google_ads_controller.CreateCampaign(campaignBody)
	if err != nil {
		msg := fmt.Sprintf("Failed to create a google ads campaign for company %s", newCompany.Email)
		suggestion := fmt.Sprintf("Create Google Ads Campaign, Company Email:%s", request.Email)
		discord.SendMessage(discord.Error, msg, suggestion)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create Google Ads Ad Group
	campaignId, err := google_ads_controller.GetCampaignID(gadsCampaign.ResourceName)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse google ads campaign id for company %s", newCompany.Email)
		suggestion := fmt.Sprintf("Get Google Ads Campaign ID, Company Email:%s", request.Email)
		discord.SendMessage(discord.Error, msg, suggestion)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	adGroupBody := google_ads_controller.Body{
		CustomerId:  gadsCustomer.Id,
		CampaignId:  campaignId,
		AdGroupName: "Initial Ad Group",
		MinCPCBid:   100000,
	}
	gadsAdGroup, err := google_ads_controller.CreateAdGroup(adGroupBody)
	if err != nil {
		msg := fmt.Sprintf("Failed to create a google ads ad group for company %s", newCompany.Email)
		suggestion := fmt.Sprintf("Create Google Ads Ad Group, Company Email:%s", request.Email)
		discord.SendMessage(discord.Error, msg, suggestion)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create Google Ads Ad
	adGroupId, err := google_ads_controller.GetAdGroupID(gadsAdGroup.ResourceName)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse google ads ad group id for company %s", newCompany.Email)
		suggestion := fmt.Sprintf("Get Google Ads Ad Group ID, Company Email:%s", request.Email)
		discord.SendMessage(discord.Error, msg, suggestion)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	adBody := google_ads_controller.Body{
		CustomerId:   gadsCustomer.Id,
		AdGroupId:    adGroupId,
		Headlines:    request.Headlines,
		Descriptions: request.Descriptions,
		FinalURL:     request.Domain,
	}
	_, err = google_ads_controller.CreateAdGroupAds(adBody)
	if err != nil {
		msg := fmt.Sprintf("Failed to create a google ads ad for company %s", newCompany.Email)
		suggestion := fmt.Sprintf("Create Google Ads Ad, Company Email:%s", request.Email)
		discord.SendMessage(discord.Error, msg, suggestion)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	discord.SendMessage(discord.Log, fmt.Sprintf("New Member Registered: %s %s - %s | %s", request.FirstName, request.LastName, request.Email, request.CompanyName), "")

	pwdResetToken, err := u.GeneratePasswordResetToken()
	if err != nil {
		discord.SendMessage(discord.Error, "Failed to generate password reset token", fmt.Sprintf("User ID: %d", u.ID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	variables := email.WelcomeData{
		Company:      request.CompanyName,
		Domain:       request.Domain,
		CreationLink: fmt.Sprintf("%s/new-user/%s", os.Getenv("FRONTEND_URL"), pwdResetToken),
	}

	variablesString, err := json.Marshal(variables)
	if err != nil {
		discord.SendMessage(discord.Error, "Failed to marshal welcome email variables", fmt.Sprintf("User ID: %d", u.ID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	emailBody := email.Email{
		To:        request.Email,
		Subject:   fmt.Sprintf("Welcome to Adomate, %s!", request.FirstName),
		Template:  "welcome email",
		Variables: string(variablesString),
	}

	email.SendEmail(emailBody)

	c.JSON(http.StatusCreated, gin.H{"message": subscription})
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
// @Param CF_Token body cloudflare.SiteVerifyRequest true "Cloudflare Token"
// @Param domain path string true "Domain URL"
// @Success 200 {object} []LocsAndSers
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /get-started/location-service/{domain} [get]
func GetLocationsAndServices(c *gin.Context) {
	domain := c.Param("domain")
	if domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "domain is required"})
		return
	}

	services, err := site_analyzer.GetServices(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	locations := []string{}

	locationsAndServices := LocsAndSers{
		Locations: locations,
		Services:  services,
	}

	c.JSON(http.StatusOK, locationsAndServices)
}

type AdContentRequest struct {
	Domain   string   `json:"domain" binding:"required" example:"adomate.ai"`
	Services []string `json:"services" binding:"required" example:"['Ad Generation', 'Ad Optimization', 'Ad Management']"`
}

type AdContentResponse struct {
	Headlines    []string `json:"headlines"`
	Descriptions []string `json:"descriptions"`
}

// GetAdContent godoc
// @Summary Get Ad Headlines and Descriptions
// @Description Get Headlines and Description for domain given services
// @Tags Getting Started
// @Accept json
// @Produce json
// @Param create body CreateAccountRequest true "Create Account Request"
// @Success 200 {object} []AdContentResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /get-started/ad-content [post]
func GetAdContent(c *gin.Context) {
	var request AdContentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	headlines, descriptions, err := site_analyzer.GetAdContent(request.Domain, request.Services)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content := AdContentResponse{
		Headlines:    headlines,
		Descriptions: descriptions,
	}

	c.JSON(http.StatusOK, content)
}

type IPInfo struct {
	IP     string `json:"ip"`
	City   string `json:"city"`
	Region string `json:"region"`
}

// GetIpInfo godoc
// @Summary Get IP Information
// @Description Get IP and related info
// @Tags Getting Started
// @Produce json
// @Success 200 {object} []IPInfo
// @Failure 500 {object} dto.ErrorResponse
// @Router /get-started/ip-info [post]
func GetIpInfo(c *gin.Context) {
	ipd, _ := ipdata.NewClient(os.Getenv("IPDATA_API_KEY"))

	data, err := ipd.Lookup(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := IPInfo{
		IP:     data.IP,
		City:   data.City,
		Region: data.Region,
	}
	c.JSON(http.StatusOK, resp)
}
