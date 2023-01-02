package main

import (
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/v1/billing"
	"github.com/adomate-ads/api/v1/campaign"
	"github.com/adomate-ads/api/v1/company"
	"github.com/adomate-ads/api/v1/industry"
	"github.com/adomate-ads/api/v1/role"
	"github.com/adomate-ads/api/v1/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func OnlineCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Adomate Ads API Online."})
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	models.ConnectDatabase(models.Config(), false)

	r := engine()
	r.Use(gin.Logger())

	// TODO - At some point we should break down the router into smaller files

	// Add router group for v1
	v1 := r.Group("/v1")

	// Online Handler - Primarily for testing purposes
	v1.GET("/", OnlineCheck)

	v1.POST("/login", user.Login)
	v1.GET("/logout", user.Logout)

	// Protected routes, requires authentication
	// TODO - I need to change the auth to a middleware function that way we can determine user access level
	auth := v1.Group("/auth")
	auth.Use(user.AuthRequired)
	{
		// Some debug user routes
		auth.GET("/me", user.Me)
		auth.GET("/status", user.Status)

		// Company Routes
		auth.POST("/company", company.CreateCompany)
		auth.GET("/company", company.GetCompanies)
		auth.GET("/company/:id", company.GetCompany)
		auth.DELETE("/company/:id", company.DeleteCompany)

		auth.GET("/company/billing/:id", billing.GetBillingsForCompany)
		auth.GET("/company/campaign/:id", campaign.GetCampaignsForCompany)

		// Industry Routes
		auth.POST("/industry", industry.CreateIndustry)
		auth.GET("/industry", industry.GetIndustries)
		auth.GET("/industry/:industry", industry.GetIndustry)
		auth.DELETE("/industry/:id", industry.DeleteIndustry)

		// Billing Routes
		auth.POST("/billing", billing.CreateBilling)
		auth.GET("/billing", billing.GetBillings)
		auth.GET("/billing/:id", billing.GetBilling)
		// Just test this one patch request for now... we can add more later once we know this one works
		auth.PATCH("/billing/:id", billing.UpdateBilling)
		auth.DELETE("/billing/:id", billing.DeleteBilling)

		// Role Routes
		auth.POST("/role", role.CreateRole)
		auth.GET("/role", role.GetRoles)
		auth.GET("/role/:role", role.GetRole)
		auth.DELETE("/role/:id", role.DeleteRole)

		// Campaign Routes
		auth.POST("/campaign", campaign.CreateCampaign)
		auth.GET("/campaign", campaign.GetCampaigns)
		auth.GET("/campaign/:id", campaign.GetCampaign)
		auth.DELETE("/campaign/:id", campaign.DeleteCampaign)
	}

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()

	// Setup CORS and only allow origin from APP URL
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{os.Getenv("APP_URL")}
	corsConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(corsConfig))

	// Set up the cookie store for session management
	r.Use(sessions.Sessions("adomate", sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	return r
}
