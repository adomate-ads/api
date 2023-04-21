package website_parse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"time"
)

func failOnError(err error, msg string) {
	discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", msg, err), "API - Website Parse Fix")
}

type WebsiteParseRequest struct {
	Domain string `json:"domain" binding:"required" example:"adomate.ai"`
}

type WebsiteParseResponse struct {
	Locations []string `json:"locations"`
	Services  []string `json:"services"`
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Queue    string
}

var RMQConfig RabbitMQConfig

func Setup() {
	RMQConfig = RabbitMQConfig{
		Host:     os.Getenv("RABBIT_HOST"),
		Port:     os.Getenv("RABBIT_PORT"),
		User:     os.Getenv("RABBIT_USER"),
		Password: os.Getenv("RABBIT_PASS"),
		Queue:    os.Getenv("RABBIT_WP_QUEUE"),
	}
}

func GetLocAndSer(domain string) ([]string, []string, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return nil, nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return nil, nil, err
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
		return nil, nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		failOnError(err, "Failed to register a consumer")
		return nil, nil, err
	}

	corrId := uuid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := WebsiteParseRequest{
		Domain: domain,
	}

	message, err := json.Marshal(req)
	if err != nil {
		failOnError(err, "Failed to marshal email")
		return nil, nil, err
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode:  amqp.Persistent,
			ContentType:   "text/plain",
			CorrelationId: corrId,
			Body:          []byte(message),
		})
	if err != nil {
		failOnError(err, "Failed to publish a message")
		return nil, nil, err
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			var resp WebsiteParseResponse
			err = json.Unmarshal(d.Body, &resp)
			if err != nil {
				failOnError(err, "Failed to unmarshal response")
				return nil, nil, err
			}
			return resp.Locations, resp.Services, nil
		}
	}

	return nil, nil, errors.New("unknown error occurred")
}
