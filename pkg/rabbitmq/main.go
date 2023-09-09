package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"time"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string

	DiscordQueue string
	MailQueue    string
	GacQueue     string
	SAQueue      string
}

var RMQConfig Config

func Setup() {
	RMQConfig = Config{
		Host:     os.Getenv("RABBIT_HOST"),
		Port:     os.Getenv("RABBIT_PORT"),
		User:     os.Getenv("RABBIT_USER"),
		Password: os.Getenv("RABBIT_PASS"),

		DiscordQueue: os.Getenv("RABBIT_DISCORD_QUEUE"),
		MailQueue:    os.Getenv("RABBIT_MAIL_QUEUE"),
		GacQueue:     os.Getenv("RABBIT_GAC_QUEUE"),
		SAQueue:      os.Getenv("RABBIT_SA_QUEUE")}
}

func SendMessage(body []byte, queue string) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		fmt.Printf("[RabbitMQ] Failed to connect to RabbitMQ: %s", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("[RabbitMQ] Failed to open a channel: %s", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		fmt.Printf("[RabbitMQ] Failed to declare a queue: %s", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})
	if err != nil {
		fmt.Printf("[RabbitMQ] Failed to publish a message: %s", err)
		return err
	}
	return nil
}

func SendMessageWithResponse(body []byte, queue string) (string, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		fmt.Printf("failed to connect to RabbitMQ: %s", err)
		return "", err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("failed to open a channel: %s", err)
		return "", err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		fmt.Printf("failed to declare a queue: %s", err)
		return "", err
	}

	replyQ, err := ch.QueueDeclare(
		fmt.Sprintf("reply_%s", q.Name),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("failed to declare a queue: %s", err)
		return "", err
	}

	msgs, err := ch.Consume(
		replyQ.Name, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		fmt.Printf("failed to register a consumer: %s", err)
		return "", err
	}

	corrId := uuid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",          // exchange
		"gac_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       replyQ.Name,
			Body:          body,
		})
	if err != nil {
		fmt.Printf("failed to publish a message: %s", err)
		return "", err
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			res := string(d.Body)
			return res, nil
		}
	}

	return "", errors.New("failed to get response from RabbitMQ")
}
