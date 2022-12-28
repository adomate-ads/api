package main

import (
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/v1/company"
	"github.com/adomate-ads/api/v1/industry"
	"github.com/adomate-ads/api/v1/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	models.ConnectDatabase(models.Config())

	r := engine()
	r.Use(gin.Logger())

	// TODO - At some point we should break down the router into smaller files

	// Add router group for v1
	v1 := r.Group("/v1")
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
		auth.POST("/company", company.Register)
		auth.GET("/company", company.GetCompanies)
		auth.GET("/company/:id", company.GetCompany)

		// Industry Routes
		auth.POST("/industry", industry.Register)
		auth.GET("/industry", industry.GetIndustries)
		auth.GET("/industry/:industry", industry.GetIndustry)

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
