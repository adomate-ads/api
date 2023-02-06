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
	"strings"
	"testing"
)

var r *gin.Engine = SetUpRouter()
var authCookie string = ""

func SetUpRouter() *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	models.ConnectDatabase(models.Config(), true)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(sessions.Sessions("adomate", sessions.NewCookieStore([]byte("testing"))))
	r.Use(auth.Auth)

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

func RequestTesting(method string, url string, body *bytes.Buffer, expectedResponse string, expectedStatus int, t *testing.T, cookies ...*http.Cookie) {
	if body == nil {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		assert.Equal(t, expectedResponse, string(responseData))
		assert.Equal(t, expectedStatus, w.Code)
	} else {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		assert.Equal(t, expectedResponse, string(responseData))
		assert.Equal(t, expectedStatus, w.Code)
	}
}

func TestOnlineCheck(t *testing.T) {
	RequestTesting("GET", "/v1/", nil, `{"message":"Adomate Ads API Online."}`, http.StatusOK, t)
}

func TestRegisterHandler(t *testing.T) {
	industry := models.Industry{
		Industry: "software",
	}
	if err := industry.CreateIndustry(); err != nil {
		t.Fatal(err)
	}

	mockResponse := `{"message":"Successfully created user and company"}`
	emptyResponse := `{"error":"Parameters can't be empty"}`
	duplicateUserResposne := `{"error":"An account by that email already exists"}`
	sameCompanyResposne := `{"error":"A company by that name already exists"}`
	industryDoesntExist := `{"error":"An industry by that name does not exist"}`

	userIndustryExist := user.RegisterRequest{
		FirstName:   "Raaj",
		LastName:    "Patel",
		Email:       "the2@raajpatel.dev",
		Password:    "Password123",
		CompanyName: "Raaj LLC2.",
		Industry:    "ads",
		Domain:      "raajpatel.dev",
		Budget:      1000,
	}
	userWithNoBusiness := user.RegisterRequest{
		FirstName:   "Raaj",
		LastName:    "Patel",
		Email:       "the@raajpatel.dev",
		Password:    "Password123",
		CompanyName: " ",
		Industry:    "software",
		Domain:      "raajpatel.dev",
		Budget:      1000,
	}
	userWithNoName := user.RegisterRequest{
		FirstName:   " ",
		LastName:    "Patel",
		Email:       "the@raajpatel.dev",
		Password:    "Password123",
		CompanyName: "Raaj LLC.",
		Industry:    "software",
		Domain:      "raajpatel.dev",
		Budget:      1000,
	}
	userWithDupilicateCompany := user.RegisterRequest{
		FirstName:   "Raaj",
		LastName:    "Patel",
		Email:       "the2@raajpatel2.dev",
		Password:    "Password123",
		CompanyName: "Raaj LLC.",
		Industry:    "software",
		Domain:      "raajpatel.dev",
		Budget:      1000,
	}
	user := user.RegisterRequest{
		FirstName:   "Raaj",
		LastName:    "Patel",
		Email:       "the@raajpatel.dev",
		Password:    "Password123",
		CompanyName: "Raaj LLC.",
		Industry:    "software",
		Domain:      "raajpatel.dev",
		Budget:      1000,
	}

	jsonValue, _ := json.Marshal(user)
	jsonValueIndustryDoesntexit, _ := json.Marshal(userIndustryExist)
	jsonValueSameCompany, _ := json.Marshal(userWithDupilicateCompany)
	jsonValueNoname, _ := json.Marshal(userWithNoName)
	jsonValueBusiness, _ := json.Marshal(userWithNoBusiness)

	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValue), mockResponse, http.StatusCreated, t)
	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValueNoname), emptyResponse, http.StatusBadRequest, t)
	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValueBusiness), emptyResponse, http.StatusBadRequest, t)
	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValue), duplicateUserResposne, http.StatusBadRequest, t)
	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValue), duplicateUserResposne, http.StatusBadRequest, t)
	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValueSameCompany), sameCompanyResposne, http.StatusBadRequest, t)
	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValueIndustryDoesntexit), industryDoesntExist, http.StatusBadRequest, t)

}

func TestLoginHandler(t *testing.T) {
	mockResponse := `{"message":"Successfully authenticated user"}`
	mockResponseEmptyName := `{"error":"Parameters can't be empty"}`
	mockResponseNoEmail := `{"error":"An account by that email does not exist"}`
	mockResponseWrongPassword := `{"error":"Incorrect password"}`
	userWrongPassword := user.LoginRequest{
		Email:    "the@raajpatel.dev",
		Password: "123",
	}
	userDoesntExist := user.LoginRequest{
		Email:    "Wyatt@raajpatel.dev",
		Password: "Password123",
	}
	userEmptyName := user.LoginRequest{
		Email:    " ",
		Password: "Password123",
	}
	user := user.LoginRequest{
		Email:    "the@raajpatel.dev",
		Password: "Password123",
	}

	jsonValue, _ := json.Marshal(user)
	jsonValue2, _ := json.Marshal(userEmptyName)
	jsonValue3, _ := json.Marshal(userDoesntExist)
	jsonValue4, _ := json.Marshal(userWrongPassword)
	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	authCookie = (w.Header().Get("Set-Cookie"))[8:strings.Index(w.Header().Get("Set-Cookie"), ";")]

	//responseData, _ := io.ReadAll(w.Body)
	//assert.Equal(t, mockResponse, string(responseData))
	//assert.Equal(t, http.StatusOK, w.Code)
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}
	RequestTesting("POST", "/v1/login", bytes.NewBuffer(jsonValue), mockResponse, http.StatusOK, t, cookie)
	RequestTesting("POST", "/v1/login", bytes.NewBuffer(jsonValue2), mockResponseEmptyName, http.StatusBadRequest, t, cookie)
	RequestTesting("POST", "/v1/login", bytes.NewBuffer(jsonValue3), mockResponseNoEmail, http.StatusBadRequest, t, cookie)
	RequestTesting("POST", "/v1/login", bytes.NewBuffer(jsonValue4), mockResponseWrongPassword, http.StatusBadRequest, t, cookie)
	RequestTesting("POST", "/v1/login", bytes.NewBuffer(jsonValue4), mockResponseWrongPassword, http.StatusBadRequest, t, cookie)

}

func TestMeHandler(t *testing.T) {
	user, _ := models.GetUser(1)
	userString, _ := json.Marshal(user)
	mockResponse := fmt.Sprintf(`{"user":%s}`, userString)
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}
	RequestTesting("GET", "/v1/me", nil, mockResponse, http.StatusOK, t, cookie)
}

func TestIndustryHandler(t *testing.T) {

}
