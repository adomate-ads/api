package stripe_v1

import (
	"encoding/json"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/stripe/webhooks"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/webhook"
	"io"
	"net/http"
	"os"
)

func handleWebhook(c *gin.Context) {
	endpointSecret := os.Getenv("WEBHOOK_SECRET")
	signatureHeader := c.GetHeader("Stripe-Signature")

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		discord.SendMessage(discord.Error, "Stripe Webhook Error", err.Error())
		return
	}

	var event stripe.Event
	event, err = webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		discord.SendMessage(discord.Error, "Stripe Webhook Error - Constructing Event", err.Error())
		return
	}

	switch event.Type {
	case "payment_intent.created":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			discord.SendMessage(discord.Error, "Stripe Webhook - PI Created", err.Error())
			return
		}

		discord.SendMessage(discord.Log, "Stripe Webhook", "Payment Intent Created: "+paymentIntent.ID)
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			discord.SendMessage(discord.Error, "Stripe Webhook - PI Succeeded", err.Error())
			return
		}

		webhooks.PaymentSucceeded(paymentIntent)
		c.Status(http.StatusOK)

	default:
		discord.SendMessage(discord.Warn, "Unknown Stripe Webhook", event.Type)
	}
}
