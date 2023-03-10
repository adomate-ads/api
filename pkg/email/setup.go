package email

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Queue    string
}

type Email struct {
	To      string
	Subject string
	Body    string
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

func SendEmail(to string, subject string, body string) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RMQConfig.Queue, // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	email := &Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	message, err := json.Marshal(email)
	failOnError(err, "Failed to marshal email")

	log.Printf("So far so good...")

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			// TODO: Is this the right content type? Maybe application/json or text/html not sure...?
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf("Email sent!")
}
