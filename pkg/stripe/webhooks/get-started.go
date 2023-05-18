package webhooks

import (
	"bytes"
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/email"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

func PaymentSucceeded(paymentIntent stripe.PaymentIntent) {
	discord.SendMessage(discord.Log, "Stripe Webhook - Payment Succeeded", "")

	fmt.Println(paymentIntent.Customer.Email)
	fmt.Println(paymentIntent.Customer)

	fmt.Println("Break...")
	params := &stripe.PaymentIntentParams{}
	params.AddExpand("customer")

	pi, err := paymentintent.Get(paymentIntent.ID, params)
	if err != nil {
		discord.SendMessage(discord.Error, "Stripe Webhook Payment Error", "Attempted to get payment intent, but failed: "+err.Error())
	}

	user, err := models.GetUserByEmail(pi.Customer.Email)
	if err != nil {
		discord.SendMessage(discord.Error, "Stripe Webhook Payment Error", "Attempted to search for user by email, but no user found: "+paymentIntent.Customer.Email)
		return
	}

	// Generate PW Reset Token
	pr := models.PasswordReset{
		UserID: user.ID,
		User:   *user,
	}
	pr.UUID = uuid.New().String()
	if err := pr.CreatePasswordReset(); err != nil {
		discord.SendMessage(discord.Error, "Stripe Webhook Payment Error", err.Error())
		return
	}

	// Send get-started welcome email
	data := email.GetStartedData{
		URL: "https://app.adomate.com/setup/" + pr.UUID,
	}
	body := new(bytes.Buffer)
	if err := email.Templates["get-started"].Tmpl.Execute(body, data); err != nil {
		discord.SendMessage(discord.Error, "Email Error - Sending Payment Success", err.Error())
		return
	}
	email.SendEmail(paymentIntent.Customer.Email, email.Templates["get-started"].Subject, body.String())
	discord.SendMessage(discord.Log, "Stripe Webhook - Payment Succeeded", "Sent get-started email to "+paymentIntent.Customer.Email)
}
