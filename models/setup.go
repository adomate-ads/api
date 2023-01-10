package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	var err error
	DB, err = gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Println("Cannot connect to database.")
		log.Fatal("Connection Error:", err)
	} else {
		fmt.Println("Connected to database.")
	}

	if clearDB {
		DB.DropTableIfExists(&User{})
		DB.DropTableIfExists(&Company{})
		DB.DropTableIfExists(&Industry{})
		DB.DropTableIfExists(&EmailTemplate{})
		DB.DropTableIfExists(&Email{})
		DB.DropTableIfExists(&Campaign{})
		DB.DropTableIfExists(&Billing{})
		DB.DropTableIfExists(&BiddingStrategy{})
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Company{})
	DB.AutoMigrate(&Industry{})
	DB.AutoMigrate(&EmailTemplate{})
	DB.AutoMigrate(&Email{})
	DB.AutoMigrate(&Campaign{})
	DB.AutoMigrate(&Billing{})
	DB.AutoMigrate(&BiddingStrategy{})
}
