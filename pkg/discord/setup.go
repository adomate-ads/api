package discord

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
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

type Message struct {
	Type       string    `json:"type" example:"error/warning/log"`
	Message    string    `json:"message"`
	Suggestion string    `json:"suggestion,omitempty"`
	Time       time.Time `json:"time,omitempty"`
	Origin     string    `json:"origin,omitempty"`
}

var RMQConfig RabbitMQConfig

func Setup() {
	RMQConfig = RabbitMQConfig{
		Host:     os.Getenv("RABBIT_HOST"),
		Port:     os.Getenv("RABBIT_PORT"),
		User:     os.Getenv("RABBIT_USER"),
		Password: os.Getenv("RABBIT_PASS"),
		Queue:    os.Getenv("RABBIT_DISCORD_QUEUE"),
	}
}

const Error string = "Error"
const Warn string = "Warn"
const Log string = "Log"
const General string = "General"

func SendMessage(level string, message string, suggestion string) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %s", err)
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
		log.Panicf("Failed to declare a queue: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg := &Message{
		Type:       level,
		Message:    message,
		Suggestion: suggestion,
		Time:       time.Now(),
		Origin:     "api",
	}

	msgString, err := json.Marshal(msg)
	if err != nil {
		log.Panicf("Failed to marshal message: %s", err)
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(msgString),
		})
	if err != nil {
		log.Panicf("Failed to publish message: %s", err)
	}
}
