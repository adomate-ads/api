package client_management

import (
	"context"
	"fmt"
	google_ads "github.com/adomate-ads/api/pkg/google-ads"
	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
	"os"
)

func CreateClient(customerName string, email string) {
	conn := google_ads.GetGRPCConnection()
	defer conn.Close()
	ctx := examples.SetContext(context.Background(),
		examples.WithContext("authorization", "Bearer "+os.Getenv("ACCESS_TOKEN")),
		examples.WithContext("developer-token", os.Getenv("DEVELOPER_TOKEN")),
		examples.WithContext("login-customer-id", "113-199-6258"),
	)
	HelperCreateClient(ctx, conn, customerName, email)
}

func HelperCreateClient(ctx context.Context, conn *grpc.ClientConn, customerName string, email string) {
	CurrencyCode := "USD"
	TimeZone := "America/Chicago"
	TrackingURLTemplate := "{lpurl}?device={device}"
	FinalURLSuffix := "keyword={keyword}&matchtype={matchtype}&adgroupid={adgroupid}"

	request := services.CreateCustomerClientRequest{
		CustomerId: "1131996258",
		CustomerClient: &resources.Customer{
			DescriptiveName:     &customerName,
			CurrencyCode:        &CurrencyCode,
			TimeZone:            &TimeZone,
			TrackingUrlTemplate: &TrackingURLTemplate,
			FinalUrlSuffix:      &FinalURLSuffix,
		},
		EmailAddress: &email,
		ValidateOnly: true,
	}
	response, err := services.NewCustomerServiceClient(conn).CreateCustomerClient(ctx, &request)
	if err != nil {
		fmt.Printf("%+#v", err.Error())
		return
	}

	fmt.Println(response)
}
