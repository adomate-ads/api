package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/price"
)

func getPriceID(name string) (string, error) {
	params := &stripe.PriceSearchParams{}
	params.Query = *stripe.String("name:" + name)
	iter := price.Search(params)
	for iter.Next() {
		return iter.Price().ID, nil
	}
	return "", iter.Err()
}
