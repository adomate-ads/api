package main

import (
	"bytes"
	"encoding/json"
	"github.com/adomate-ads/api/v1/billing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestOnlineCheck(t *testing.T) {
	mockResponse := `{"message":"Adomate Ads API Online."}`
	r := SetUpRouter()
	r.GET("/", OnlineCheck)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBillingHandler(t *testing.T) {
	mockResponse := `{"message":"Successfully created bill."}`
	r := SetUpRouter()
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
	assert.Equal(t, http.StatusOK, w.Code)
}
