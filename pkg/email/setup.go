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

func failOnError(err error, msg string) {
	discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", msg, err), "API - Email Fix")
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Queue    string
}

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
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

func SendEmail(to, subject, body string) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
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
		failOnError(err, "Failed to declare a queue")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	email := Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	message, err := json.Marshal(email)
	if err != nil {
		failOnError(err, "Failed to marshal email")
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message),
		})
	if err != nil {
		failOnError(err, "Failed to publish a message")
	}
}
