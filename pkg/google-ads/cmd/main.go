package main

import (
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	google_ads.Setup()

}
