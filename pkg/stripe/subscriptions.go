package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/subscription"
)

type Subscription struct {
	ID            string `json:"id"`
	PaymentIntent string `json:"payment_intent"`
}

func CreateSubscription(customerId string, price string, description string) (*Subscription, error) {

	//Automatically save the payment method to the subscription when the first payment is successful.
	paymentSettings := &stripe.SubscriptionPaymentSettingsParams{
		SaveDefaultPaymentMethod: stripe.String("on_subscription"),
	}

	//Create the subscription
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(price), //THIS IS NOT THE PRICE; This is the ID that is associated with the price that is associated with the payment
			},
		},
		PaymentSettings: paymentSettings,
		PaymentBehavior: stripe.String("default_incomplete"),
		Description:     stripe.String(description),
	}
	params.AddExpand("latest_invoice.payment_intent") //Expand the payment intent so we can get the PI to send to frontend
	s, err := subscription.New(params)
	if err != nil {
		return nil, err
	}

	sub := Subscription{
		ID:            s.ID,
		PaymentIntent: s.LatestInvoice.PaymentIntent.ClientSecret,
	}

	return &sub, nil
}
