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

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created bill."})
}

func GetBillings(c *gin.Context) {
	billings, err := models.GetBillings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, billings)
}

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

func PayBill(c *gin.Context) {
	// Some Stripe magic here...

}

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
