package order

import (
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CreateRequest struct {
	Company string    `json:"company" binding:"required"`
	Amount  float64   `json:"amount" binding:"required"`
	Status  string    `json:"status" binding:"required"`
	Type    string    `json:"type" binding:"required"`
	StartAt time.Time `json:"start_at" binding:"required"`
	EndAt   time.Time `json:"end_at" binding:"required"`
}

// CreateOrder Order godoc
// @Summary Create Order
// @Description creates an order
// @Tags Order
// @Accept json
// @Produce json
// @Param create body CreateRequest true "Create Request"
// @Success 200 {object} []models.Industry
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order [post]
func CreateOrder(c *gin.Context) {
	fmt.Println("function create order called")
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(request.Company, " ") == "" || strings.Trim(request.Status, " ") == "" || strings.Trim(request.Type, " ") == "" || request.Amount < 0 || (request.Type != "base" && request.Type != "premium" && request.Type != "enterprise" && request.Type != "ads") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	if request.Status != "pending" && request.Status != "active" && request.Status != "cancelled" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Status input"})
		return
	}

	// Get company ID
	company, err := models.GetCompanyByName(request.Company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "That company does not exist"})
		return
	}
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != company.ID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get information about your company"})
		return
	}
	// Create industry
	order := models.Order{
		Company:   *company,
		CompanyID: company.ID,
		Amount:    request.Amount,
		Status:    request.Status,
		Type:      request.Type,
		StartAt:   request.StartAt,
		EndAt:     request.EndAt,
	}

	if err := order.CreateOrder(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created order"})
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get a slice of all orders
// @Tags Order
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Order
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order [get]
func GetOrders(c *gin.Context) {
	orders, err := models.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetOrder godoc
// @Summary Gets an order
// @Description Gets all information about specific order
// @Tags Order
// @Accept */*
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} []models.Order
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order/{id} [get]
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	order, err := models.GetOrder(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Only allow the user to get orders from their company
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != order.CompanyID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only get orders from your company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// UpdateOrder godoc
// @Summary Update Order
// @Description Update an order.
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param update body models.UpdateOrder true "Update Request"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order/{id} [patch]
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := models.GetOrder(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Only allow the user to update orders from their company
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != order.CompanyID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update orders from your company"})
		return
	}

	var request models.UpdateOrder
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate form input
	// TODO - Check the all fields are valid
	order, err = order.UpdateOrder(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// DeleteOrder godoc
// @Summary Delete Order
// @Description Delete an order.
// @Tags Order
// @Accept */*
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order/{id} [delete]
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := models.GetOrder(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != order.CompanyID && !auth.InGroup(user, "super-admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update orders from your company"})
		return
	}

	if err := order.DeleteOrder(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
