package gads

import (
	"github.com/adomate-ads/api/pkg/google-ads/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetClients Google Ads godoc
// @Summary Get Google Ads Clients
// @Description Gets all Google Ads Clients
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Success 200 {object} []helpers.Client
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/client [get]
func GetClients(c *gin.Context) {
	clients := helpers.GetClients()
	c.JSON(http.StatusOK, gin.H{"clients": clients})
}

// GetClient Google Ads godoc
// @Summary Get Google Ads Client
// @Description Gets all information about specific Google Ads Client
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param clientId path int true "Client ID"
// @Success 200 {object} helpers.Client
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/client/{clientId} [get]
func GetClient(c *gin.Context) {
	clientId := c.Param("clientId")
	if clientId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Client ID is required."})
		return
	}

	//TODO - Only allow the user to get the client attached to their company

	client := helpers.GetClient(clientId)

	c.JSON(http.StatusOK, gin.H{"client": client})
}
