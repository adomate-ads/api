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

func ListAccessibleCustomers() {
	conn := google_ads.GetGRPCConnection()
	defer conn.Close()
	ctx := examples.SetContext(context.Background(),
		examples.WithContext("authorization", "Bearer "+os.Getenv("ACCESS_TOKEN")),
		examples.WithContext("developer-token", os.Getenv("DEVELOPER_TOKEN")),
	)
	HelperListAccessibleCustomers(ctx, conn)
}

func HelperListAccessibleCustomers(ctx context.Context, conn *grpc.ClientConn) {
	customers, err := services.NewCustomerServiceClient(conn).
		ListAccessibleCustomers(ctx, &services.ListAccessibleCustomersRequest{})
	if err != nil {
		fmt.Printf("%+#v", err.Error())
		return
	}

	for _, customer := range customers.ResourceNames {
		fmt.Println("ResourceName: " + customer)
	}
}
