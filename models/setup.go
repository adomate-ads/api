package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

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

func ConnectDatabase(dbConfig *DBConfig, clearDB bool) {
	//DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database.")
		log.Fatal("Connection Error:", err)
	} else {
		fmt.Println("Connected to database.")
	}

	//if clearDB {
	//	DB.DropTableIfExists(&User{})
	//	DB.DropTableIfExists(&Company{})
	//	DB.DropTableIfExists(&Industry{})
	//	DB.DropTableIfExists(&EmailTemplate{})
	//	DB.DropTableIfExists(&Email{})
	//	DB.DropTableIfExists(&Campaign{})
	//	DB.DropTableIfExists(&Order{})
	//	DB.DropTableIfExists(&Billing{})
	//	DB.DropTableIfExists(&PasswordReset{})
	//	DB.DropTableIfExists(&Location{})
	//	DB.DropTableIfExists(&Service{})
	//}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Company{})
	DB.AutoMigrate(&Industry{})
	DB.AutoMigrate(&EmailTemplate{})
	DB.AutoMigrate(&Email{})
	DB.AutoMigrate(&Campaign{})
	DB.AutoMigrate(&Order{})
	DB.AutoMigrate(&Billing{})
	DB.AutoMigrate(&PasswordReset{})
	DB.AutoMigrate(&Location{})
	DB.AutoMigrate(&Service{})
}
