package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/v1/billing"
	"github.com/adomate-ads/api/v1/campaign"
	"github.com/adomate-ads/api/v1/company"
	"github.com/adomate-ads/api/v1/industry"
	"github.com/adomate-ads/api/v1/user"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var r *gin.Engine = SetUpRouter()

func SetUpRouter() *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	models.ConnectDatabase(models.Config(), false)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(sessions.Sessions("adomate", sessions.NewCookieStore([]byte(""))))

	v1 := r.Group("v1")

	v1.GET("/", OnlineCheck)

	v1.POST("/login", user.Login)
	v1.POST("/register", user.Register)
	v1.GET("/logout", auth.NotGuest, user.Logout)

	// Protected routes, requires authentication
	v1.GET("/me", auth.NotGuest, user.Me)
	v1.GET("/status", auth.NotGuest, user.Status)

	company.Routes(v1)
	industry.Routes(v1)
	billing.Routes(v1)
	campaign.Routes(v1)
	return r
}

func TestOnlineCheck(t *testing.T) {
	mockResponse := `{"message":"Adomate Ads API Online."}`
	req, _ := http.NewRequest("GET", "/v1/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginHandler(t *testing.T) {
	mockResponse := `{"message":"Successfully authenticated user"}`
	mockUserId := "1"
	user := user.LoginRequest{
		Email:    "the@raajpatel.dev",
		Password: "Password123",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseHeader := w.Result().Header
	fmt.Println(responseHeader["Set-Cookie"])
	assert.Equal(t, mockUserId, responseHeader.Get("user-id"))

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

//
//func TestCreateIndustryHandler(t *testing.T) {
//	mockResponse := `{"message":"Successfully created industry"}`
//
//	industry := industry.CreateRequest{
//		Industry: "Software",
//	}
//
//	jsonValue, _ := json.Marshal(industry)
//	req, _ := http.NewRequest("POST", "/v1/industry", bytes.NewBuffer(jsonValue))
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//
//	responseData, _ := io.ReadAll(w.Body)
//	assert.Equal(t, mockResponse, string(responseData))
//	assert.Equal(t, http.StatusOK, w.Code)
//
//	mockResponse = `{"error":"An industry by that name already exists"}`
//	req, _ = http.NewRequest("POST", "/v1/industry", bytes.NewBuffer(jsonValue))
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//
//	responseData, _ = io.ReadAll(w.Body)
//	assert.Equal(t, mockResponse, string(responseData))
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//}
//
//func TestCreateCompanyHandler(t *testing.T) {
//	mockResponse := `{"message":"Successfully registered company"}`
//
//	company := company.CreateRequest{
//		Name:     "Raaj Inc.",
//		Email:    "the@raajpatel.dev",
//		Industry: "Software",
//		Domain:   "https://raajpatel.dev",
//		Budget:   10,
//	}
//
//	jsonValue, _ := json.Marshal(company)
//	req, _ := http.NewRequest("POST", "/v1/company", bytes.NewBuffer(jsonValue))
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//
//	responseData, _ := io.ReadAll(w.Body)
//	assert.Equal(t, mockResponse, string(responseData))
//	assert.Equal(t, http.StatusCreated, w.Code)
//
//	mockResponse = `{"error":"An company by that email already exists"}`
//	req, _ = http.NewRequest("POST", "/v1/company", bytes.NewBuffer(jsonValue))
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//
//	responseData, _ = io.ReadAll(w.Body)
//	assert.Equal(t, mockResponse, string(responseData))
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//}
//
//func TestCreateBillingHandler(t *testing.T) {
//	mockResponse := `{"error":"That company does not exist"}`
//
//	billing := billing.CreateRequest{
//		Company:  "Raaj123 Inc.",
//		Amount:   100,
//		Status:   "unpaid",
//		Comments: "This is a test",
//		DueAt:    time.Now(),
//		IssuedAt: time.Now(),
//	}
//
//	jsonValue, _ := json.Marshal(billing)
//	req, _ := http.NewRequest("POST", "/v1/billing", bytes.NewBuffer(jsonValue))
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//
//	responseData, _ := io.ReadAll(w.Body)
//	assert.Equal(t, mockResponse, string(responseData))
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//
//	mockResponse = `{"message":"Successfully created bill"}`
//	billing.Company = "Raaj Inc."
//
//	jsonValue, _ = json.Marshal(billing)
//	req, _ = http.NewRequest("POST", "/v1/billing", bytes.NewBuffer(jsonValue))
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//
//	responseData, _ = io.ReadAll(w.Body)
//	assert.Equal(t, mockResponse, string(responseData))
//	assert.Equal(t, http.StatusCreated, w.Code)
//}
