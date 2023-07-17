package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/price"
)

func createPriceID(productID string, adBudget uint) (string, error) {
	params := &stripe.PriceParams{
		Product:    stripe.String(productID),
		UnitAmount: stripe.Int64(int64(adBudget)),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String("month"),
		},
	}

	p, err := price.New(params)
	if err != nil {
		return "", err
	}

	return p.ID, nil
}

//func getPriceID(name string) (string, error) {
//	params := &stripe.PriceSearchParams{}
//	params.Query = *stripe.String("name:" + name)
//	iter := price.Search(params)
//	for iter.Next() {
//		return iter.Price().ID, nil
//	}
//	return "", iter.Err()
//}
