package main

import (
	"bytes"
	"encoding/json"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/v1/billing"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var r *gin.Engine = SetUpRouter()

func SetUpRouter() *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	models.ConnectDatabase(models.Config())

	r := gin.Default()
	return r
}

func TestOnlineCheck(t *testing.T) {
	mockResponse := `{"message":"Adomate Ads API Online."}`
	r.GET("/", OnlineCheck)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBillingHandler(t *testing.T) {
	mockResponse := `{"error":"That company does not exist"}`
	r.POST("/billing", billing.CreateBilling)

	billing := billing.CreateRequest{
		Company:  "Raaj Inc.",
		Amount:   100,
		Status:   "unpaid",
		Comments: "This is a test",
		DueAt:    time.Now(),
		IssuedAt: time.Now(),
	}

	jsonValue, _ := json.Marshal(billing)

	req, _ := http.NewRequest("POST", "/billing", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
