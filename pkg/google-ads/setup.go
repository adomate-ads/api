package google_ads

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"

	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	googleAdsApi "github.com/adomate-ads/api/pkg/google-ads/v12"
)

const GoogleAdsEndpoint string = "googleads.googleapis.com:443"

var GADSClient *googleAdsApi.Client
var SuperUser string
var Ctx context.Context

func Setup() {
	Ctx = context.Background()
	oAuthToken := oauth2.Token{
		RefreshToken: os.Getenv("GADS_REFRESH_TOKEN"),
		TokenType:    "Bearer",
	}

	oAuthConf := oauth2.Config{
		ClientID:     os.Getenv("GADS_CLIENT_ID"),
		ClientSecret: os.Getenv("GADS_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://127.0.0.1",
		Scopes:       []string{"https://www.googleapis.com/auth/adwords"},
	}

	opts := []option.ClientOption{
		option.WithTokenSource(oAuthConf.TokenSource(Ctx, &oAuthToken)),
		option.WithEndpoint(GoogleAdsEndpoint),
	}

	Ctx = metadata.AppendToOutgoingContext(Ctx, "developer-token", os.Getenv("GADS_DEVELOPER_TOKEN"))
	Ctx = metadata.AppendToOutgoingContext(Ctx, "login-customer-id", os.Getenv("GADS_LOGIN_ID")) // Manager Account ID

	var err error
	GADSClient, err = googleAdsApi.NewClient(Ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating googleAdsClient:%v\n", err)
	}

	SuperUser = os.Getenv("GADS_LOGIN_ID")

	//getAndListCampaigns(ctx, "1644244393", opts...)
	//getAndListClients(ctx, "1644244393", opts...)
	//getAndListAdGroups(ctx, "1787212549", opts...)
	//getAndListAdGroupAds(ctx, "1787212549", opts...)
	//getAndListKeywords(ctx, "1787212549", "149197460347", opts...)
	//GetClientInfo(ctx, "1787212549", opts...)
}

func getAndListClients(ctx context.Context, customerId string, opts ...option.ClientOption) {
	fmt.Println("\n\nGetting and listing clients for customer ID " + customerId)
	googleAdsClient, err := googleAdsApi.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating googleAdsClient:%v\n", err)
	}

	request := services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      "SELECT customer_client.id, customer_client.descriptive_name FROM customer_client",
	}

	resp := googleAdsClient.Search(ctx, &request)

	for {
		row, err := resp.Next()
		customerClient := row.GetCustomerClient()
		//Structure of customerClient:
		//customerClient.id
		//customerClient.descriptive_name

		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			log.Fatalf("An error occured: %v", err)
		}

		log.Printf("Client: %v \n", customerClient)
	}
}

func getAndListCampaigns(ctx context.Context, customerId string, opts ...option.ClientOption) {
	fmt.Println("\n\nGetting and listing campaigns for customer ID " + customerId)
	googleAdsClient, err := googleAdsApi.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating googleAdsClient:%v\n", err)
	}

	request := services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      "SELECT campaign.id, campaign.name FROM campaign ORDER BY campaign.id",
	}

	resp := googleAdsClient.Search(ctx, &request)

	for {
		row, err := resp.Next()
		campaign := row.GetCampaign()

		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			log.Fatalf("An error occured: %v", err)
		}

		log.Printf("Campaign: %v", campaign)
	}

}

func getAndListAdGroups(ctx context.Context, customerId string, opts ...option.ClientOption) {
	fmt.Println("\n\nGetting and listing ad groups for customer ID " + customerId)
	googleAdsClient, err := googleAdsApi.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating googleAdsClient:%v\n", err)
	}

	request := services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      "SELECT ad_group.id, ad_group.name FROM ad_group ORDER BY ad_group.id",
	}

	resp := googleAdsClient.Search(ctx, &request)

	for {
		row, err := resp.Next()
		adGroup := row.GetAdGroup()

		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			log.Fatalf("An error occured: %v", err)
		}

		log.Printf("Ad Group: %v", adGroup)
	}

}

func getAndListAdGroupAds(ctx context.Context, customerId string, opts ...option.ClientOption) {
	fmt.Println("\n\nGetting and listing ad group ads for customer ID " + customerId)
	googleAdsClient, err := googleAdsApi.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating googleAdsClient:%v\n", err)
	}

	request := services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      "SELECT ad_group_ad.ad.id, ad_group_ad.ad.final_urls FROM ad_group_ad ORDER BY ad_group_ad.ad.id",
	}

	resp := googleAdsClient.Search(ctx, &request)

	for {
		row, err := resp.Next()
		adGroupAd := row.GetAdGroupAd()

		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			log.Fatalf("An error occured: %v", err)
		}

		log.Printf("Ad Group Ad: %v", adGroupAd)
	}

}

func getAndListKeywords(ctx context.Context, customerId string, adGroupId string, opts ...option.ClientOption) {
	fmt.Println("\n\nGetting and listing keywords for adGroup ID " + adGroupId + " customer ID " + customerId)
	googleAdsClient, err := googleAdsApi.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating googleAdsClient:%v\n", err)
	}

	request := services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      "SELECT ad_group.id, ad_group_criterion.type, ad_group_criterion.criterion_id, ad_group_criterion.keyword.text, ad_group_criterion.keyword.match_type FROM ad_group_criterion WHERE ad_group_criterion.type = 'KEYWORD' AND ad_group.id = " + adGroupId,
	}

	resp := googleAdsClient.Search(ctx, &request)

	for {
		row, err := resp.Next()
		adGroupCriterion := row.GetAdGroupCriterion()

		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			log.Fatalf("An error occured: %v", err)
		}

		log.Printf("Ad Group Criterion: %v", adGroupCriterion)
	}

}

// GetAllCustomers - No params
// GetCustomer - Customer ID
func GetClientInfo(ctx context.Context, customerId string, opts ...option.ClientOption) {
	fmt.Println("\n\nGetting client info for customer ID " + customerId)
	googleAdsClient, err := googleAdsApi.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating googleAdsClient:%v\n", err)
	}

	request := services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      "SELECT customer.id, customer.descriptive_name, customer.currency_code, customer.time_zone FROM customer LIMIT 1",
	}

	resp := googleAdsClient.Search(ctx, &request)

	for {
		row, err := resp.Next()
		customer := row.GetCustomer()

		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			log.Fatalf("An error occured: %v", err)
		}

		log.Printf("Customer: %v", customer)
	}
}

//GetAllCampaignsForCustomer - CustomerID
//GetCampaign - Campaign ID
//GetAllAdGroups - CustomerID, CampaignID
//GetAdGroup - AdGroupID
//GetAllAds - CustomerID, CampaignID, AdGroupID
//GetAd - AdID
