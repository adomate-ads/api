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

func Setup() {
	ctx := context.Background()
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
		option.WithTokenSource(oAuthConf.TokenSource(ctx, &oAuthToken)),
		option.WithEndpoint(GoogleAdsEndpoint),
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "developer-token", os.Getenv("GADS_DEVELOPER_TOKEN"))
	ctx = metadata.AppendToOutgoingContext(ctx, "login-customer-id", os.Getenv("GADS_LOGIN_ID")) // Manager Account ID

	getAndListCampaigns(ctx, "1644244393", opts...)
	getAndListClients(ctx, "1644244393", opts...)
	getAndListCustomers(ctx, opts...)
	getAndListAdGroups(ctx, "1644244393", opts...)
	getAndListAdGroupAds(ctx, "1644244393", opts...)
	getAndListKeywords(ctx, "1644244393", "1644244393", opts...)
}

func getAndListCustomers(ctx context.Context, opts ...option.ClientOption) {
	customerServiceClient, err := googleAdsApi.NewCustomerClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating NewCustomerClient:%v\n", err)
	}

	listAccessibleAccountsRequest := services.ListAccessibleCustomersRequest{}

	accessibleCustomersResponse, err := customerServiceClient.ListAccessibleCustomers(ctx, &listAccessibleAccountsRequest)
	if err != nil {
		log.Fatalf("Error occured when calling ListAccessibleCustomers:%v\n", err)
	}

	log.Printf("Accessible Accounts from Manager ID: %s", os.Getenv("GADS_LOGIN_ID"))
	for _, accountResource := range accessibleCustomersResponse.ResourceNames {
		log.Printf("Account Resource: %s\n", accountResource)
	}
}

func getAndListClients(ctx context.Context, customerId string, opts ...option.ClientOption) {
	fmt.Println("Getting and listing clients for customer ID " + customerId)
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
	fmt.Println("Getting and listing campaigns for customer ID " + customerId)
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
	fmt.Println("Getting and listing ad groups for customer ID " + customerId)
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
	fmt.Println("Getting and listing ad group ads for customer ID " + customerId)
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
	fmt.Println("Getting and listing keywords for adGroup ID " + adGroupId + "customer ID " + customerId)
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
