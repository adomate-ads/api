package email

import (
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/rabbitmq"
	"os"
)

type Email struct {
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Template  string `json:"template"`
	Variables string `json:"variables"`
}

var Queue string = os.Getenv("RABBIT_MAIL_QUEUE")

func SendEmail(body Email) {
	msgBody, err := json.Marshal(body)
	if err != nil {
		discord.SendMessage(discord.Error, "[Email] Failed to marshal message", fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	if err := rabbitmq.SendMessage(msgBody, Queue); err != nil {
		discord.SendMessage(discord.Error, "[Email] Failed to send message", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}
