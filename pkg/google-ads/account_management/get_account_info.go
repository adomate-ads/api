package account_management

import (
	"context"
	"fmt"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
	"os"
)

func GetAccountInformation() {
	conn := google_ads.GetGRPCConnection()
	defer conn.Close()
	ctx := examples.SetContext(context.Background(),
		examples.WithContext("authorization", "Bearer "+os.Getenv("ACCESS_TOKEN")),
		examples.WithContext("developer-token", os.Getenv("DEVELOPER_TOKEN")),
		examples.WithContext("login-customer-id", "283-459-0997"),
	)
	HelperGetAccountInformation(ctx, conn, "283-459-0997")
}

func HelperGetAccountInformation(ctx context.Context, conn *grpc.ClientConn, customerID string) {
	request := services.SearchGoogleAdsRequest{
		CustomerId: customerID,
		Query:      "SELECT customer.descriptive_name FROM customer WHERE customer.id = " + customerID,
	}
	search, err := services.NewGoogleAdsServiceClient(conn).Search(ctx, &request)
	if err != nil {
		return
	}

	if len(search.Results) == 0 {
		fmt.Println("Google Ads - No results found!")
		return
	}

	for _, resource := range search.Results {
		fmt.Println(resource.Customer.GetDescriptiveName())
	}
}
