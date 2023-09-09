package discord

import (
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/pkg/rabbitmq"
	"time"
)

type Message struct {
	Type    string    `json:"type" example:"error/warning/log"`
	Title   string    `json:"title"`
	Message string    `json:"message,omitempty"`
	Time    time.Time `json:"time,omitempty"`
	Origin  string    `json:"origin"`
}

const Error string = "Error"
const Warn string = "Warning"
const Log string = "Log"

func SendMessage(level string, title string, message string) {
	msg := &Message{
		Type:    level,
		Title:   title,
		Message: message,
		Time:    time.Now(),
		Origin:  "api",
	}

	msgBody, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("[Discord] Failed to marshal message body")
		return
	}

	if err := rabbitmq.SendMessage(msgBody, rabbitmq.RMQConfig.DiscordQueue); err != nil {
		fmt.Println("[Discord] Failed to send message to queue")
		return
	}
}
