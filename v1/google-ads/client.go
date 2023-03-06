package gads

import (
	"errors"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"net/http"
)

type Client struct {
	Id          int64  `json:"id"`
	Name        string `json:"descriptive_name"`
	CurrentCode string `json:"current_code"`
	Timezone    string `json:"timezone"`
}

// GetClients Google Ads godoc
// @Summary Get Google Ads Clients
// @Description Gets all Google Ads Clients
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Success 200 {object} []Client
// @Failure 500 {object} dto.ErrorResponse
// @Router /gads/client [get]
func GetClients(c *gin.Context) {
	request := services.SearchGoogleAdsRequest{
		CustomerId: google_ads.SuperUser,
		Query:      "SELECT customer_client.id, customer_client.descriptive_name FROM customer_client",
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	var clients []Client

	for {
		row, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		customerClient := row.GetCustomerClient()
		if customerClient == nil {
			continue
		}

		client := Client{}

		client.Id = *customerClient.Id
		if customerClient.DescriptiveName != nil {
			client.Name = *customerClient.DescriptiveName
		}

		clients = append(clients, client)
	}

	c.JSON(http.StatusOK, gin.H{"clients": clients})
}

// GetClient Google Ads godoc
// @Summary Get Google Ads Client
// @Description Gets all information about specific Google Ads Client
// @Tags Google Ads
// @Accept */*
// @Produce json
// @Param clientId path int true "Client ID"
// @Success 200 {object} Client
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

	request := services.SearchGoogleAdsRequest{
		CustomerId: clientId,
		Query: `	SELECT 
    						customer.id, 
							customer.descriptive_name, 
							customer.currency_code, 
							customer.time_zone 
						FROM 
						    customer 
						LIMIT 1`,
	}

	resp := google_ads.GADSClient.Search(google_ads.Ctx, &request)

	row, err := resp.Next()
	if errors.Is(err, iterator.Done) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Client not found."})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer := row.GetCustomer()
	client := Client{}
	client.Id = *customer.Id
	if customer.DescriptiveName != nil {
		client.Name = *customer.DescriptiveName
	}
	if customer.CurrencyCode != nil {
		client.CurrentCode = *customer.CurrencyCode
	}
	if customer.TimeZone != nil {
		client.Timezone = *customer.TimeZone
	}

	c.JSON(http.StatusOK, gin.H{"client": client})
}
