package email

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/pkg/discord"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"time"
)

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Queue    string
}

type Email struct {
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Template  string `json:"template"`
	Variables string `json:"variables"`
}

var RMQConfig RabbitMQConfig

func Setup() {
	RMQConfig = RabbitMQConfig{
		Host:     os.Getenv("RABBIT_HOST"),
		Port:     os.Getenv("RABBIT_PORT"),
		User:     os.Getenv("RABBIT_USER"),
		Password: os.Getenv("RABBIT_PASS"),
		Queue:    os.Getenv("RABBIT_MAIL_QUEUE"),
	}
}

func SendEmail(body Email) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("Failed to connect to RabbitMQ: %s", err), "Mail-Server")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("Failed to open a channel: %s", err), "Mail-Server")
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RMQConfig.Queue, // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("Failed to declare a queue: %s", err), "Mail-Server")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message, err := json.Marshal(body)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("Failed to marshal message: %s", err), "Mail-Server")
		return
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         message,
		})
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("Failed to publish a message: %s", err), "Mail-Server")
		return
	}
}
