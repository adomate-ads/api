package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"os"
)

func Setup() {
	stripe.Key = os.Getenv("STRIPE_KEY")
}
