package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/email"
	"github.com/adomate-ads/api/pkg/stripe"
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
	"time"
)

var r *gin.Engine = SetUpRouter()
var authCookie string = ""
var otherAuthCookie string = ""

func SetUpRouter() *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	email.Setup()
	stripe.Setup()
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

		for _, cookie := range cookies {
			req.AddCookie(cookie)
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
	user2 := user.RegisterRequest{
		FirstName:   "wyatt",
		LastName:    "griffin",
		Email:       "test@gmail.com",
		Password:    "Password123",
		CompanyName: "LLC.",
		Industry:    "software",
		Domain:      "raajpatel.dev",
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
	}
	userWithNoBusiness := user.RegisterRequest{
		FirstName:   "Raaj",
		LastName:    "Patel",
		Email:       "the@raajpatel.dev",
		Password:    "Password123",
		CompanyName: " ",
		Industry:    "software",
		Domain:      "raajpatel.dev",
	}
	userWithNoName := user.RegisterRequest{
		FirstName:   " ",
		LastName:    "Patel",
		Email:       "the@raajpatel.dev",
		Password:    "Password123",
		CompanyName: "Raaj LLC.",
		Industry:    "software",
		Domain:      "raajpatel.dev",
	}
	userWithDupilicateCompany := user.RegisterRequest{
		FirstName:   "Raaj",
		LastName:    "Patel",
		Email:       "the2@raajpatel2.dev",
		Password:    "Password123",
		CompanyName: "Raaj LLC.",
		Industry:    "software",
		Domain:      "raajpatel.dev",
	}
	user := user.RegisterRequest{
		FirstName:   "Raaj",
		LastName:    "Patel",
		Email:       "the@raajpatel.dev",
		Password:    "Password123",
		CompanyName: "Raaj LLC.",
		Industry:    "software",
		Domain:      "raajpatel.dev",
	}

	jsonValue, _ := json.Marshal(user)
	jsonValue2, _ := json.Marshal(user2)
	jsonValueIndustryDoesntexit, _ := json.Marshal(userIndustryExist)
	jsonValueSameCompany, _ := json.Marshal(userWithDupilicateCompany)
	jsonValueNoname, _ := json.Marshal(userWithNoName)
	jsonValueBusiness, _ := json.Marshal(userWithNoBusiness)
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}
	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValue), mockResponse, http.StatusCreated, t, cookie)
	RequestTesting("POST", "/v1/register", bytes.NewBuffer(jsonValue2), mockResponse, http.StatusCreated, t)
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
	user22 := user.LoginRequest{
		Email:    "test@gmail.com",
		Password: "Password123",
	}
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
	jsonValue22, _ := json.Marshal(user22)
	jsonValue2, _ := json.Marshal(userEmptyName)
	jsonValue3, _ := json.Marshal(userDoesntExist)
	jsonValue4, _ := json.Marshal(userWrongPassword)
	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req2, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonValue22))
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	authCookie = (w.Header().Get("Set-Cookie"))[8:strings.Index(w.Header().Get("Set-Cookie"), ";")]
	otherAuthCookie = (w2.Header().Get("Set-Cookie"))[8:strings.Index(w2.Header().Get("Set-Cookie"), ";")]
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

func TestGetIndustry(t *testing.T) {
	userAdmin, _ := models.GetUser(1)
	userAdmin.Role = "super-admin"
	userAdmin.UpdateUser()
	userIndustry, _ := models.GetIndustry(1)
	jsonValue, _ := json.Marshal(userIndustry)
	mockResponse := fmt.Sprintf(`{"industries":[%s]}`, jsonValue)
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}
	RequestTesting("GET", "/v1/industry", nil, mockResponse, http.StatusOK, t, cookie)
}

func TestCreateIndustryHandler(t *testing.T) {

	userAdmin, _ := models.GetUser(1)
	userAdmin.Role = "super-admin"
	userAdmin.UpdateUser()

	newIndustry := industry.CreateRequest{
		Industry: "advertisment",
	}
	EmptyIndustry := industry.CreateRequest{
		Industry: " ",
	}
	jsonValue, _ := json.Marshal(newIndustry)
	jsonValue2, _ := json.Marshal(EmptyIndustry)
	mockResponse := `{"message":"Successfully created industry"}`
	mockResponseEmptyParam := `{"error":"Parameters can't be empty"}`
	mockResponse2 := `{"error":"An industry by that name already exists"}`
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}

	RequestTesting("POST", "/v1/industry", bytes.NewBuffer(jsonValue), mockResponse, http.StatusCreated, t, cookie)
	RequestTesting("POST", "/v1/industry", bytes.NewBuffer(jsonValue2), mockResponseEmptyParam, http.StatusBadRequest, t, cookie)
	RequestTesting("POST", "/v1/industry", bytes.NewBuffer(jsonValue), mockResponse2, http.StatusBadRequest, t, cookie)

}

func TestGetSpecificIndustryHandler(t *testing.T) {
	userIndustry, _ := models.GetIndustry(1)
	jsonValue, _ := json.Marshal(userIndustry)
	mockResponse := fmt.Sprintf(`{"industry":%s}`, jsonValue)
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}

	RequestTesting("GET", "/v1/industry/software", nil, mockResponse, http.StatusOK, t, cookie)

}

func TestDeleteIndustryHandler(t *testing.T) {}

func TestCreateCompanyHandler(t *testing.T) {
	mockResponse := `{"message":"Successfully registered company"}`
	mockResponseSameEmail := `{"error":"An company by that email already exists"}`
	mockResponseEmptyParam := `{"error":"Parameters can't be empty"}`
	mockResponseDNE := `{"error":"An industry by that name does not exist"}`
	companyDNE := company.CreateRequest{
		Name:     "twitter",
		Email:    "neil@othee.com",
		Industry: "chemical Engineering",
		Domain:   "domain",
	}
	companyEmptyName := company.CreateRequest{
		Name:     " ",
		Email:    "neil@othee.com",
		Industry: "software",
		Domain:   "domain",
	}
	company := company.CreateRequest{
		Name:     "twitter",
		Email:    "wy@othee.com",
		Industry: "software",
		Domain:   "domain",
	}

	jsonValue, _ := json.Marshal(company)
	jsonValue2, _ := json.Marshal(companyEmptyName)
	jsonValue3, _ := json.Marshal(companyDNE)
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}

	RequestTesting("POST", "/v1/company", bytes.NewBuffer(jsonValue), mockResponse, http.StatusCreated, t, cookie)
	RequestTesting("POST", "/v1/company", bytes.NewBuffer(jsonValue2), mockResponseEmptyParam, http.StatusBadRequest, t, cookie)
	RequestTesting("POST", "/v1/company", bytes.NewBuffer(jsonValue), mockResponseSameEmail, http.StatusBadRequest, t, cookie)
	RequestTesting("POST", "/v1/company", bytes.NewBuffer(jsonValue3), mockResponseDNE, http.StatusBadRequest, t, cookie)
}

func TestGetCompaniesHandler(t *testing.T) {
	//user, _ := models.GetUserByEmail("test@gmail.com")
	//user.Role = "super-admin"
	//user.UpdateUser()

	company, _ := models.GetCompanies()
	jsonValue, _ := json.Marshal(company)
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}
	mockResponse := fmt.Sprintf(`%s`, jsonValue)
	RequestTesting("GET", "/v1/company", nil, mockResponse, http.StatusOK, t, cookie)
}

func TestGetCompanyHandler(t *testing.T) {
	company1, _ := models.GetCompany(1)
	jsonValue, _ := json.Marshal(company1)
	mockResponse := fmt.Sprintf(`%s`, jsonValue)
	mockResponseNotAuthorizedUser := `{"error":"You can only get information about your company"}`
	mockResponseCompanyDNE := `{"error":"Company doesn't exist"}`
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}
	cookie2 := &http.Cookie{
		Name:   "adomate",
		Value:  otherAuthCookie,
		MaxAge: 300,
	}
	RequestTesting("GET", "/v1/company/1", nil, mockResponse, http.StatusOK, t, cookie)
	RequestTesting("GET", "/v1/company/1", nil, mockResponseNotAuthorizedUser, http.StatusForbidden, t, cookie2)
	RequestTesting("GET", "/v1/company/33", nil, mockResponseCompanyDNE, http.StatusNotFound, t, cookie)

}

func TestCreateBilling(t *testing.T) {

	company1, _ := models.GetCompany(1)
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}

	now := time.Now()
	later := now.Add(time.Hour)
	billRequest := billing.CreateRequest{
		Company:  company1.Name,
		Amount:   34.12,
		Status:   "unpaid",
		Comments: "something about something",
		IssuedAt: time.Now(),
		DueAt:    later,
	}
	billRequest2 := billing.CreateRequest{
		Company:  company1.Name + "nonSense",
		Amount:   34.12,
		Status:   "unpaid",
		Comments: "something about something",
		IssuedAt: time.Now(),
		DueAt:    later,
	}
	billRequest3 := billing.CreateRequest{
		Company:  company1.Name,
		Amount:   34.12,
		Status:   "not valid status",
		Comments: "something about something",
		IssuedAt: time.Now(),
		DueAt:    later,
	}
	billRequest4 := billing.CreateRequest{
		Company:  company1.Name,
		Amount:   34.12,
		Status:   " ",
		Comments: "something about something",
		IssuedAt: time.Now(),
		DueAt:    later,
	}
	jsonValue, _ := json.Marshal(billRequest)
	jsonValue2, _ := json.Marshal(billRequest2)
	jsonValue3, _ := json.Marshal(billRequest3)
	jsonValue4, _ := json.Marshal(billRequest4)

	mockResponse := `{"message":"Successfully created bill"}`
	mockCompNoExist := `{"error":"That company does not exist"}`
	mockResponseBadInput := `{"error":"Invalid Status input"}`
	mockResponseEmpty := `{"error":"Parameters can't be empty"}`
	RequestTesting("POST", "/v1/billing", bytes.NewBuffer(jsonValue), mockResponse, http.StatusCreated, t, cookie)
	RequestTesting("POST", "/v1/billing", bytes.NewBuffer(jsonValue2), mockCompNoExist, http.StatusBadRequest, t, cookie)
	RequestTesting("POST", "/v1/billing", bytes.NewBuffer(jsonValue3), mockResponseBadInput, http.StatusBadRequest, t, cookie)
	RequestTesting("POST", "/v1/billing", bytes.NewBuffer(jsonValue4), mockResponseEmpty, http.StatusBadRequest, t, cookie)

}

func TestGetBilling(t *testing.T) {
	//industry := models.Industry{
	//	Industry: "software",
	//}
	//newC := models.Company{
	//	ID:         2,
	//	Name:       "wyattomate",
	//	Email:      "theman#gmail.com",
	//	IndustryID: 1,
	//	Industry:   industry,
	//	Domain:     "fourlokodev.com",
	//	Budget:     10000,
	//	AdsBalance: 342,
	//	CreatedAt:  time.Now(),
	//	UpdatedAt:  time.Now(),
	//}
	//newC.CreateCompany()
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}
	cookie2 := &http.Cookie{
		Name:   "adomate",
		Value:  otherAuthCookie,
		MaxAge: 300,
	}
	company1, _ := models.GetCompany(1)

	billing, _ := models.GetBilling(uint(company1.ID))
	jsonValue, _ := json.Marshal(billing)
	mockResponse := fmt.Sprintf(`%s`, jsonValue)
	mockResponse2 := `{"error":"You can only get bills from your company"}`
	RequestTesting("GET", "/v1/billing/1", nil, mockResponse, http.StatusOK, t, cookie)
	RequestTesting("GET", "/v1/billing/1", nil, mockResponse2, http.StatusForbidden, t, cookie2)
}

//skip testupdate

//func TestUpdateBilling(t *testing.T) {
//	cookie := &http.Cookie{
//		Name:   "adomate",
//		Value:  authCookie,
//		MaxAge: 300,
//	}
//	company, _ := models.GetCompany(1)
//	//billings, _ := models.GetBillings()
//	//var bill Billing
//	//json.Unmarshal(billings[0], &bill)
//	//json.Unmarshal([]byte(billings[0]), &bill)
//	updateRequest := billing.UpdateRequest{
//		Company: company.Name,
//		Status:  "paid",
//		//DueAt:    billings[0].DueAt,
//		//IssuedAt: billings[0].IssuedAt,
//	}
//
//	//vals, _ := json.Marshal(billings[0])
//	jsonValue, _ := json.Marshal(updateRequest)
//	mockResponse := fmt.Sprintf(`%s`, jsonValue)
//
//	RequestTesting("PATCH", "/v1/billing/1", bytes.NewBuffer(jsonValue), mockResponse, http.StatusAccepted, t, cookie)
//
//}

func TestDeleteBilling(t *testing.T) {
	cookie := &http.Cookie{
		Name:   "adomate",
		Value:  authCookie,
		MaxAge: 300,
	}
	mockResponse := `{"message":"Bill deleted successfully"}`
	RequestTesting("DELETE", "/v1/billing/1", nil, mockResponse, http.StatusOK, t, cookie)
}

//skipped update user and delete user
//skipped most company and compaing routes
//get billing and don't do campaign
