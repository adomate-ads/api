package cmd

import (
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Config() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil && os.Getenv("GIN_MODE") != "release" {
		log.Fatalf("Error loading .env file.")
	}

	dbConfig := Config()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database.")
		log.Fatal("Connection Error:", err)
	} else {
		fmt.Println("Connected to database.")
	}

	err = DB.AutoMigrate(&models.User{}, &models.Company{}, &models.EmailTemplate{}, &models.Email{}, &models.Campaign{}, &models.Order{}, &models.Billing{}, &models.PasswordReset{}, &models.Location{}, &models.Service{})
	if err != nil {
		log.Fatal("Cannot auto-migrate db:", err)
	}
}
