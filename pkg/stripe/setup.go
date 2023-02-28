package stripe

import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/product"
	"github.com/stripe/stripe-go/v74/subscription"
	"log"
	"os"
)

func Setup() {
	stripe.Key = os.Getenv("STRIPE_KEY")
}

func SetupProducts() bool {
	// TODO - Add product URLs once front end is ready

	// Check if plans already have been created
	params := &stripe.ProductSearchParams{}
	params.Query = "name~'Base Plan' OR name~'Premium Plan' OR name~'Enterprise Plan'"
	iter := product.Search(params)
	for iter.Next() {
		fmt.Printf("Product %s already exists", iter.Product().Name)
		return false
	}

	//Base Plan
	paramsBP := &stripe.ProductParams{
		Name:        stripe.String("Base Plan"),
		Description: stripe.String("Great for small to mid sized businesses spending under $500/mo on ads."),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String("USD"),
			UnitAmount: stripe.Int64(20 * 100), // $20.00
			Recurring: &stripe.ProductDefaultPriceDataRecurringParams{
				Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
			},
			TaxBehavior: stripe.String(string(stripe.PriceCurrencyOptionsTaxBehaviorExclusive)),
		},
	}
	if _, err := product.New(paramsBP); err != nil {
		log.Fatal(err)
	}

	//Premium Plan
	paramsPP := &stripe.ProductParams{
		Name:        stripe.String("Premium Plan"),
		Description: stripe.String("Great for mid to larger sized businesses spending $501-$2500/mo on ads."),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String("USD"),
			UnitAmount: stripe.Int64(50 * 100), // $50.00
			Recurring: &stripe.ProductDefaultPriceDataRecurringParams{
				Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
			},
			TaxBehavior: stripe.String(string(stripe.PriceCurrencyOptionsTaxBehaviorExclusive)),
		},
	}
	if _, err := product.New(paramsPP); err != nil {
		log.Fatal(err)
	}

	//Enterprise Plan
	paramsEP := &stripe.ProductParams{
		Name:        stripe.String("Enterprise Plan"),
		Description: stripe.String("Great for large+ size businesses spending over $2501/mo on ads."),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String("USD"),
			UnitAmount: stripe.Int64(150 * 100), // $150.00
			Recurring: &stripe.ProductDefaultPriceDataRecurringParams{
				Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
			},
			TaxBehavior: stripe.String(string(stripe.PriceCurrencyOptionsTaxBehaviorExclusive)),
		},
	}
	if _, err := product.New(paramsEP); err != nil {
		log.Fatal(err)
	}
	return true
}

func GetSubscriptions() {
	fmt.Println("Getting subscriptions...")
	params := &stripe.SubscriptionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := subscription.List(params)
	for i.Next() {
		s := i.Subscription()
		sub, _ := json.Marshal(s)
		fmt.Printf("Subscription: %s", sub)
	}
}
