package email

import (
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/rabbitmq"
)

type Email struct {
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Template  string `json:"template"`
	Variables string `json:"variables"`
}

func SendEmail(body Email) {
	msgBody, err := json.Marshal(body)
	if err != nil {
		discord.SendMessage(discord.Error, "[Email] Failed to marshal message", fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	if err := rabbitmq.SendMessage(msgBody, rabbitmq.RMQConfig.MailQueue); err != nil {
		discord.SendMessage(discord.Error, "[Email] Failed to send message", fmt.Sprintf("Error: %s", err.Error()))
		return
	}
}
