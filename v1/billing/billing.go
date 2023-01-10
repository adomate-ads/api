package billing

import (
	"github.com/adomate-ads/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
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
// @Accept */*
// @Produce json
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
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
		Amount:    request.Amount,
		Status:    request.Status,
		Comments:  request.Comments,
		DueAt:     request.DueAt,
		IssuedAt:  request.IssuedAt,
	}

	if err := b.CreateBilling(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created bill"})
}

// GetBillings godoc
// @Summary Get all Bills
// @Description Gets a slice of all bills.
// @Tags Billing
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Billing
// @Failure 400 {object} dto.ErrorResponse
// @Router /billing [get]
func GetBillings(c *gin.Context) {
	billings, err := models.GetBillings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
// @Success 200 {object} []models.Billing
// @Failure 400 {object} dto.ErrorResponse
// @Router /billing/company/:id [get]
func GetBillingsForCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	billings, err := models.GetBillingsByCompanyID(uint(companyID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
// @Success 200 {object} models.Billing
// @Failure 400 {object} dto.ErrorResponse
// @Router /billing/:id [get]
func GetBilling(c *gin.Context) {
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
	c.JSON(http.StatusOK, billing)
}

// This will need to be tested with unit tests or something cause honestly i've never handled a patch request before
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
// @Accept */*
// @Produce json
// @Success 202 {object} models.Billing
// @Failure 400 {object} dto.ErrorResponse
// @Router /billing/:id [patch]
func UpdateBilling(c *gin.Context) {
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

	var request UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	// TODO - This needs to be verified, we might need to add another flag for if it is not passed there for not changed.
	if request.Status != "paid" && request.Status != "unpaid" && request.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Status input"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, updatedBilling)
}

func PayBill(c *gin.Context) {
	// TODO - Some Stripe magic here...

}

// DeleteBilling godoc
// @Summary Delete Bill
// @Description Delete a bill.
// @Tags Billing
// @Accept */*
// @Produce json
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /billing/:id [delete]
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bill deleted successfully"})
}
