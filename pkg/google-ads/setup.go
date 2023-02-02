package google_ads

import (
	"context"
	"errors"
	"log"

	"github.com/joeshaw/envdecode"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"

	"github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	google_ads_api "github.com/adomate-ads/api/pkg/google-ads/v12"
)

var env struct {
	ClientID        string `env:"GADS_CLIENT_ID,required"`
	ClientSecret    string `env:"GADS_CLIENT_SECRET,required"`
	RefreshToken    string `env:"GADS_REFRESH_TOKEN,required"`
	DeveloperToken  string `env:"GADS_DEVELOPER_TOKEN,required"`
	LoginCustomerID string `env:"GADS_LOGIN_CUSTOMER_ID,required"`
}

const GoogleAdsEndpoint string = "googleads.googleapis.com:443"

func Setup() {
	_ = envdecode.Decode(&env)

	ctx := context.Background()
	oAuthToken := oauth2.Token{
		RefreshToken: env.RefreshToken,
		TokenType:    "Bearer",
	}
	oAuthConf := oauth2.Config{
		ClientID:     env.ClientID,
		ClientSecret: env.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://127.0.0.1",
		Scopes:       []string{"https://www.googleapis.com/auth/adwords"},
	}

	opts := []option.ClientOption{
		option.WithTokenSource(oAuthConf.TokenSource(ctx, &oAuthToken)),
		option.WithEndpoint(GoogleAdsEndpoint),
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "developer-token", env.DeveloperToken)
	ctx = metadata.AppendToOutgoingContext(ctx, "login-customer-id", env.LoginCustomerID)

	googleAdsClient, err := google_ads_api.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Error occured when creating googleAdsClient:%v\n", err)
	}

	getAndListCampaigns(ctx, googleAdsClient, "1162019409")
}

func getAndListCampaigns(ctx context.Context, client *google_ads_api.Client, customerId string) {
	request := services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      "SELECT campaign.id, campaign.name FROM campaign ORDER BY campaign.id",
	}

	resp := client.Search(ctx, &request)

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
