package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/product"
)

func getProductID(name string) (string, error) {
	params := &stripe.ProductSearchParams{}
	params.Query = *stripe.String("name:" + name)
	iter := product.Search(params)
	for iter.Next() {
		return iter.Product().ID, nil
	}
	return "", iter.Err()
}
