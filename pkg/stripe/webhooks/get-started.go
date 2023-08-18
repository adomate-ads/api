package webhooks

import (
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/email"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v74"
	"os"
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
		UUID:   uuid.New().String(),
	}
	if err := pr.CreatePasswordReset(); err != nil {
		discord.SendMessage(discord.Error, "Stripe Webhook Payment Error", err.Error())
		return
	}

	// Send get-started welcome email
	variables := email.WelcomeData{
		Company:      user.Company.Name,
		Domain:       user.Company.Domain,
		CreationLink: fmt.Sprintf("%s/new-user/%s", os.Getenv("FRONTEND_URL"), pr.UUID),
	}

	variablesString, err := json.Marshal(variables)
	if err != nil {
		discord.SendMessage(discord.Error, "Failed to marshal welcome email variables", fmt.Sprintf("User ID: %d", user.ID))
		return
	}

	emailBody := email.Email{
		To:        user.Email,
		Subject:   fmt.Sprintf("Welcome to Adomate, %s!", user.FirstName),
		Template:  "welcome email",
		Variables: string(variablesString),
	}

	email.SendEmail(emailBody)

	discord.SendMessage(discord.Log, "Stripe Webhook - Payment Succeeded", "Sent get-started email to "+user.Email)
}
