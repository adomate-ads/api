package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

func CreatePaymentIntent(customerEmail string, stripeId string, amount int64) (string, error) {
	params := &stripe.PaymentIntentParams{
		Amount:       stripe.Int64(amount),
		Customer:     stripe.String(stripeId),
		ReceiptEmail: stripe.String(customerEmail),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}

	return pi.ClientSecret, nil
}

func GetPaymentIntent(id string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{}
	params.AddExpand("customer")

	pi, err := paymentintent.Get(id, params)
	if err != nil {
		return nil, err
	}

	return pi, nil
}
