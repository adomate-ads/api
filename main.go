package main

import (
	"fmt"
	"github.com/adomate-ads/api/models"
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

	r.POST("/login", user.Login)
	r.GET("/logout", user.Logout)

	// Protected routes, requires authentication
	auth := r.Group("/auth")
	auth.Use(user.AuthRequired)
	{
		auth.GET("/me", user.Me)
		auth.GET("/status", user.Status)
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
