package webhooks

import (
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/email"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v74"
)

func PaymentSucceeded(paymentIntent stripe.PaymentIntent) {
	company, err := models.GetCompanyByStripeID(paymentIntent.Customer.ID)
	if err != nil {
		discord.SendMessage(discord.Error, "Stripe Webhook Payment Error", "Attempted to search for company by stripe id, but no company found: "+paymentIntent.Customer.ID)
		return
	}

	user, err := models.GetUserByEmail(company.Email)
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
	email.Templates["get-started"].Execute(data)

	discord.SendMessage(discord.Log, "Stripe Webhook - Payment Succeeded", "Sent get-started email to "+user.Email)
}
