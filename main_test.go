package main

import (
	"bytes"
	"encoding/json"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/v1/billing"
	"github.com/adomate-ads/api/v1/campaign"
	"github.com/adomate-ads/api/v1/company"
	"github.com/adomate-ads/api/v1/industry"
	"github.com/adomate-ads/api/v1/role"
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

	models.ConnectDatabase(models.Config(), true)

	r := gin.Default()

	r.GET("/", OnlineCheck)

	// Company Routes
	r.POST("/company", company.CreateCompany)
	r.GET("/company", company.GetCompanies)
	r.GET("/company/:id", company.GetCompany)
	r.DELETE("/company/:id", company.DeleteCompany)

	r.GET("/company/billing/:id", billing.GetBillingsForCompany)
	r.GET("/company/campaign/:id", campaign.GetCampaignsForCompany)

	// Industry Routes
	r.POST("/industry", industry.CreateIndustry)
	r.GET("/industry", industry.GetIndustries)
	r.GET("/industry/:industry", industry.GetIndustry)
	r.DELETE("/industry/:id", industry.DeleteIndustry)

	// Billing Routes
	r.POST("/billing", billing.CreateBilling)
	r.GET("/billing", billing.GetBillings)
	r.GET("/billing/:id", billing.GetBilling)
	// Just test this one patch request for now... we can add more later once we know this one works
	r.PATCH("/billing/:id", billing.UpdateBilling)
	r.DELETE("/billing/:id", billing.DeleteBilling)

	// Role Routes
	r.POST("/role", role.CreateRole)
	r.GET("/role", role.GetRoles)
	r.GET("/role/:role", role.GetRole)
	r.DELETE("/role/:id", role.DeleteRole)

	// Campaign Routes
	r.POST("/campaign", campaign.CreateCampaign)
	r.GET("/campaign", campaign.GetCampaigns)
	r.GET("/campaign/:id", campaign.GetCampaign)
	r.DELETE("/campaign/:id", campaign.DeleteCampaign)

	return r
}

func TestOnlineCheck(t *testing.T) {
	mockResponse := `{"message":"Adomate Ads API Online."}`
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateIndustryHandler(t *testing.T) {
	mockResponse := `{"message":"Successfully created industry"}`

	industry := industry.CreateRequest{
		Industry: "Software",
	}

	jsonValue, _ := json.Marshal(industry)
	req, _ := http.NewRequest("POST", "/industry", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

	mockResponse = `{"error":"An industry by that name already exists"}`
	req, _ = http.NewRequest("POST", "/industry", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ = ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateCompanyHandler(t *testing.T) {
	mockResponse := `{"message":"Successfully registered company"}`

	company := company.CreateRequest{
		Name:     "Raaj Inc.",
		Email:    "the@raajpatel.dev",
		Industry: "Software",
		Domain:   "https://raajpatel.dev",
		Budget:   10,
	}

	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/company", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

	mockResponse = `{"error":"An company by that email already exists"}`
	req, _ = http.NewRequest("POST", "/company", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ = ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateBillingHandler(t *testing.T) {
	mockResponse := `{"error":"That company does not exist"}`

	billing := billing.CreateRequest{
		Company:  "Raaj123 Inc.",
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

	mockResponse = `{"message":"Successfully created bill"}`
	billing.Company = "Raaj Inc."

	jsonValue, _ = json.Marshal(billing)
	req, _ = http.NewRequest("POST", "/billing", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ = ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
