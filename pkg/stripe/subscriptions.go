package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/subscription"
)

type Subscription struct {
	ID            string `json:"id"`
	PaymentIntent string `json:"payment_intent"`
	Total         int64  `json:"total"`
	Tax           int64  `json:"tax"`
	Items         []Item `json:"items"`
}

type Item struct {
	ID    string `json:"id"`
	Price int64  `json:"price"`
}

func CreateSubscription(customerId string, price string, description string, adBudget uint) (*Subscription, error) {

	//Automatically save the payment method to the subscription when the first payment is successful.
	paymentSettings := &stripe.SubscriptionPaymentSettingsParams{
		SaveDefaultPaymentMethod: stripe.String("on_subscription"),
	}

	p, err := createPriceID("prod_Ns9ONV54VtPZFF", adBudget)
	if err != nil {
		return nil, err
	}

	//Create the subscription
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(price), //THIS IS NOT THE PRICE; This is the ID that is associated with the price that is associated with the payment
			},
			{
				Price: stripe.String(p),
			},
		},
		PaymentSettings: paymentSettings,
		PaymentBehavior: stripe.String("default_incomplete"),
		Description:     stripe.String(description),
		AutomaticTax:    &stripe.SubscriptionAutomaticTaxParams{Enabled: stripe.Bool(true)},
	}
	params.AddExpand("latest_invoice.payment_intent") //Expand the payment intent so we can get the PI to send to frontend
	params.AddExpand("latest_invoice")

	s, err := subscription.New(params)
	if err != nil {
		return nil, err
	}

	var items []Item
	for _, item := range s.Items.Data {
		items = append(items, Item{
			ID:    item.Price.ID,
			Price: item.Price.UnitAmount,
		})
	}

	sub := Subscription{
		ID:            s.ID,
		PaymentIntent: s.LatestInvoice.PaymentIntent.ClientSecret,
		Total:         s.LatestInvoice.Total,
		Tax:           s.LatestInvoice.Tax,
		Items:         items,
	}

	return &sub, nil
}
