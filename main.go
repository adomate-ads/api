package main

import (
	"fmt"
	"github.com/adomate-ads/api/docs"
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/discord"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/rabbitmq"
	"github.com/adomate-ads/api/pkg/stripe"
	"github.com/adomate-ads/api/v1/billing"
	"github.com/adomate-ads/api/v1/campaign"
	"github.com/adomate-ads/api/v1/company"
	get_started "github.com/adomate-ads/api/v1/get-started"
	gads "github.com/adomate-ads/api/v1/google-ads"
	"github.com/adomate-ads/api/v1/order"
	"github.com/adomate-ads/api/v1/service"
	stripe_v1 "github.com/adomate-ads/api/v1/stripe"
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
// @contact.url https://adomate.ai/support
// @contact.email support@adomate.ai

// @host api.adomate.ai
// @BasePath /
// @schemes https
func main() {
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	err := godotenv.Load(".env")
	if err != nil && os.Getenv("GIN_MODE") != "release" {
		log.Fatalf("Error loading .env file.")
	}

	rabbitmq.Setup()
	google_ads.Setup()
	stripe.Setup()

	//stripe.SetupProducts()
	//stripe.GetSubscriptions()

	models.ConnectDatabase(models.Config())

	r := engine()
	r.Use(gin.Logger())

	r.Use(auth.Auth)

	r.Static("/static", "./docs/static")
	r.LoadHTMLGlob("docs/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Add router group for v1
	v1 := r.Group("/v1")

	// Online Handler - Primarily for testing purposes
	v1.GET("/", OnlineCheck)

	v1.POST("/login", user.Login)
	v1.POST("/register", user.Register)
	v1.GET("/logout", auth.NotGuest, user.Logout)
	v1.POST("/forgot", user.ForgotPassword)
	v1.POST("/reset/:resetToken", user.ResetPassword)

	// Protected routes, requires authentication
	v1.GET("/me", auth.NotGuest, user.Me)
	v1.GET("/status", auth.NotGuest, user.Status)

	user.Routes(v1)
	company.Routes(v1)
	billing.Routes(v1)
	campaign.Routes(v1)
	order.Routes(v1)
	service.Routes(v1)
	gads.Routes(v1)
	get_started.Routes(v1)
	stripe_v1.Routes(v1)

	// Static files, such as images, css, and js
	v1.StaticFS("/storage", gin.Dir("./storage", false))

	discord.SendMessage(discord.Log, "[API] Starting...", "An API Instance is starting.")

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		discord.SendMessage(discord.Error, "An API Instance has failed to start.", err.Error())
		log.Fatal("Unable to start server:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()

	// Setup CORS and only allow origin from APP URL
	corsConfig := cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL"), os.Getenv("DASHBOARD_URL"), "http://app.adomate.local"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	// Set up the cookie store for session management
	r.Use(sessions.Sessions("adomate", sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
