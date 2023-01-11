package main

import (
	"fmt"
	"github.com/adomate-ads/api/docs"
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/v1/billing"
	"github.com/adomate-ads/api/v1/campaign"
	"github.com/adomate-ads/api/v1/company"
	"github.com/adomate-ads/api/v1/industry"
	"github.com/adomate-ads/api/v1/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
)

// OnlineCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags General
// @Accept */*
// @Produce json
// @Success 200 {object} dto.MessageResponse
// @Router / [get]
func OnlineCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Adomate Ads API Online."})
}

// @title Adomate API
// @version 1.0
// @description Adomate Monolithic API

// @contact.name Adomate API Support
// @contact.url https://adomate.com/support
// @contact.email support@adomate.com

// @host localhost:3000
// @BasePath /
// @schemes http https
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	models.ConnectDatabase(models.Config(), false)

	r := engine()
	r.Use(gin.Logger())

	r.Use(auth.Auth)

	// Add router group for v1
	v1 := r.Group("/v1")

	// Online Handler - Primarily for testing purposes
	v1.GET("/", OnlineCheck)

	v1.POST("/login", user.Login)
	v1.POST("/register", user.Register)
	v1.GET("/logout", auth.NotGuest, user.Logout)

	// Protected routes, requires authentication
	v1.GET("/me", auth.NotGuest, user.Me)
	v1.GET("/status", auth.NotGuest, user.Status)

	// Test Groups
	group := r.Group("/test-group")
	group.GET("/super-admin", auth.NotGuest, auth.InGroup("super-admin"), user.Me)
	group.GET("/support", auth.NotGuest, auth.InGroup("support"), user.Me)
	group.GET("/admin", auth.NotGuest, auth.InGroup("admin"), user.Me)
	group.GET("/user", auth.NotGuest, auth.InGroup("user"), user.Me)
	// Test Roles
	roles := r.Group("/test-roles")
	roles.GET("/super-admin", auth.NotGuest, auth.HasRole("super-admin"), user.Me)
	roles.GET("/support-billing", auth.NotGuest, auth.HasRole("support-billing"), user.Me)
	roles.GET("/support-ticket", auth.NotGuest, auth.HasRole("support-ticket"), user.Me)
	roles.GET("/owner", auth.NotGuest, auth.HasRole("owner"), user.Me)
	roles.GET("/admin", auth.NotGuest, auth.HasRole("admin"), user.Me)
	roles.GET("/user", auth.NotGuest, auth.HasRole("user"), user.Me)

	user.Routes(v1)
	company.Routes(v1)
	industry.Routes(v1)
	billing.Routes(v1)
	campaign.Routes(v1)

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

	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
