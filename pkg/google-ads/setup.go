package google_ads

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"

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
}
