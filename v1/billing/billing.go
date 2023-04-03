package billing

import (
	"bytes"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/adomate-ads/api/pkg/email"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Company  string    `json:"company" binding:"required"`
	Amount   float64   `json:"amount" binding:"required"`
	Status   string    `json:"status" binding:"required"`
	Comments string    `json:"comments" binding:"required"`
	DueAt    time.Time `json:"due_at" binding:"required"`
	IssuedAt time.Time `json:"issued_at" binding:"required"`
}

// CreateBilling godoc
// @Summary Create Bill
// @Description Create a new bill.
// @Tags Billing
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /billing [post]
func CreateBilling(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Company, " ") == "" || strings.Trim(request.Status, " ") == "" || strings.Trim(request.Comments, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}
	if request.Status != "paid" && request.Status != "unpaid" && request.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Status input"})
		return
	}

	// Get company ID
	company, err := models.GetCompanyByName(request.Company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "That company does not exist"})
		return
	}

	b := models.Billing{
		CompanyID: company.ID,
		Company:   *company,
		// TODO Request Company Address for email
		// CompanyAddressLine1: ,
		// CompanyAddressLine2: ,
		// CompanyAddressState: ,
		// CompanyAddressZip:   ,
		Amount:   request.Amount,
		Status:   request.Status,
		Comments: request.Comments,
		DueAt:    request.DueAt,
		IssuedAt: request.IssuedAt,
	}

	if err := b.CreateBilling(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//params := &stripe.CustomerSearchParams{}
	//params.Query = *stripe.String(fmt.Sprintf("metadata['company_id']:'%s'", company.ID))
	//// TODO - grab the customer ID from stripe
	//customerID := "12345"
	//iter := customer.Search(params)
	//for iter.Next() {
	//	result := iter.Current()
	//
	//}
	//billParams := &stripe.InvoiceParams{
	//	Customer:    stripe.String(customerID),
	//	Description: "10/12-12/10 something like that..",
	//}

	data := email.NewInvoice{
		InvoiceID:     b.ID,
		Company:       company.Name,
		InvoiceAmount: b.Amount,
		Status:        b.Status,
		DueAt:         b.DueAt.Format("2006-01-02"),
	}
	body := new(bytes.Buffer)
	if err := email.Templates["new-invoice"].Tmpl.Execute(body, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	email.SendEmail(company.Email, email.Templates["new-invoice"].Subject, body.String())

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created bill"})
}

// GetBillings godoc
// @Summary Get all Bills
// @Description Gets a slice of all bills.
// @Tags Billing
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Billing
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /billing [get]
func GetBillings(c *gin.Context) {
	billings, err := models.GetBillings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, billings)
}

// GetBillingsForCompany godoc
// @Summary Get all Bills for a Company
// @Description Gets a slice of all the bills for a specific company.
// @Tags Billing
// @Accept */*
// @Produce json
// @Param id path integer true "Company ID"
// @Success 200 {object} []models.Billing
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /billing/company/{id} [get]
func GetBillingsForCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Check if the user is an admin of a company, if so make sure they're getting bills of their company
	user := c.MustGet("x-user").(*models.User)
	if auth.InGroup(user, "admin") {
		if user.CompanyID != uint(companyID) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only get bills from your company"})
			return
		}
	}

	billings, err := models.GetBillingsByCompanyID(uint(companyID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, billings)
}

// GetBilling godoc
// @Summary Gets a Bill
// @Description Gets all information about a single bill.
// @Tags Billing
// @Accept */*
// @Produce json
// @Param id path integer true "Billing ID"
// @Success 200 {object} models.Billing
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /billing/{id} [get]
func GetBilling(c *gin.Context) {
	id := c.Param("id")
	billingID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	billing, err := models.GetBilling(uint(billingID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Only allow the user to get bills from their company
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != billing.CompanyID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get bills from your company"})
		return
	}

	c.JSON(http.StatusOK, billing)
}

type UpdateRequest struct {
	Company  string    `json:"company"`
	Amount   float64   `json:"amount"`
	Status   string    `json:"status"`
	Comments string    `json:"comments"`
	DueAt    time.Time `json:"due_at"`
	IssuedAt time.Time `json:"issued_at"`
}

// UpdateBilling godoc
// @Summary Update Bill
// @Description Update information about a bill.
// @Tags Billing
// @Accept  json
// @Produce json
// @Param update body CreateRequest true "Create Request"
// @Param id path integer true "Billing ID"
// @Success 202 {object} models.Billing
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /billing/{id} [patch]
func UpdateBilling(c *gin.Context) {
	id := c.Param("id")
	billingID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	billing, err := models.GetBilling(uint(billingID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var request UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get company ID
	if request.Company != "" {
		company, err := models.GetCompanyByName(request.Company)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "That company does not exist"})
			return
		}
		billing.CompanyID = company.ID
		billing.Company = *company
	}

	if request.Amount != 0 {
		billing.Amount = request.Amount
	}

	if request.Status != "" {
		if request.Status != "paid" && request.Status != "unpaid" && request.Status != "pending" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Status input"})
			return
		}
		billing.Status = request.Status
	}

	if request.Comments != "" {
		billing.Comments = request.Comments
	}

	if !request.DueAt.IsZero() {
		billing.DueAt = request.DueAt
	}

	if !request.IssuedAt.IsZero() {
		billing.IssuedAt = request.IssuedAt
	}

	updatedBilling, err := billing.UpdateBilling()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, updatedBilling)
}

// DeleteBilling godoc
// @Summary Delete Bill
// @Description Delete a bill.
// @Tags Billing
// @Accept */*
// @Produce json
// @Param id path integer true "Billing ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /billing/{id} [delete]
func DeleteBilling(c *gin.Context) {
	id := c.Param("id")
	billingID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	billing, err := models.GetBilling(uint(billingID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := billing.DeleteBilling(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bill deleted successfully"})
}
