package helpers

import (
	"errors"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	"google.golang.org/api/iterator"
)

type Client struct {
	Id          int64  `json:"id"`
	Name        string `json:"descriptive_name"`
	CurrentCode string `json:"current_code"`
	Timezone    string `json:"timezone"`
}

func GetClients() []Client {
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
			return clients
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

	return clients
}

func GetClient(clientId string) Client {
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

	var client Client

	for {
		row, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			return client
		}

		customerClient := row.GetCustomerClient()
		if customerClient == nil {
			continue
		}

		client.Id = *customerClient.Id
		if customerClient.DescriptiveName != nil {
			client.Name = *customerClient.DescriptiveName
		}
		if customerClient.CurrencyCode != nil {
			client.CurrentCode = *customerClient.CurrencyCode
		}
		if customerClient.TimeZone != nil {
			client.Timezone = *customerClient.TimeZone
		}

	}

	return client
}
