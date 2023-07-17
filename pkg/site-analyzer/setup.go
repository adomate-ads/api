package site_analyzer

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

// TODO - get rid of this fail on error stuff as it doesnt return the function so it can lead to unintended consequences

func failOnError(err error, msg string) {
	discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", msg, err.Error()), "API - SA Fix")
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Queue    string
}

type Request struct {
	Route string      `json:"route"`
	Body  RequestBody `json:"body"`
}

type RequestBody struct {
	URL      string   `json:"url"`
	Services []string `json:"services"`
}

type Response struct {
	Body  ResponseBody `json:"body"`
	Error string       `json:"error"`
}

type ResponseBody struct {
	Services     []string `json:"services"`
	Headlines    []string `json:"headlines"`
	Descriptions []string `json:"descriptions"`
}

var RMQConfig RabbitMQConfig

func Setup() {
	RMQConfig = RabbitMQConfig{
		Host:     os.Getenv("RABBIT_HOST"),
		Port:     os.Getenv("RABBIT_PORT"),
		User:     os.Getenv("RABBIT_USER"),
		Password: os.Getenv("RABBIT_PASS"),
		Queue:    os.Getenv("RABBIT_SA_QUEUE"),
	}
}

func GetServices(url string) ([]string, error) {
	// establish a connection
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to connect to RabbitMQ", err.Error()), "API - Site_Analyzer Fix")
	}

	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			discord.SendMessage(discord.Error, "Failed to close RabbitMQ connection", err.Error())
		}
	}(conn)

	// create a channel
	ch, err := conn.Channel()
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to open channel", err.Error()), "API - Site_Analyzer Fix")
	}

	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			discord.SendMessage(discord.Error, "Failed to close RabbitMQ channel", err.Error())
		}
	}(ch)

	// declare a queue
	q, err := ch.QueueDeclare(
		RMQConfig.Queue, // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to declare a queue", err.Error()), "API - Site_Analyzer Fix")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// generate a random string
	corrId := uuid.New().String()

	// Create a request
	req := Request{
		Route: "GuessServices",
		Body: RequestBody{
			URL: url,
		},
	}

	message, err := json.Marshal(req)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to marshal the request", err.Error()), "API - Site_Analyzer Fix")
	}

	// publish the message
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(message),
		})
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to publish a message", err.Error()), "API - Site_Analyzer Fix")
	}

	// consume the response
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
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to register a consumer", err.Error()), "API - Site_Analyzer Fix")
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			fmt.Println("Got a response: ", d.CorrelationId)
			fmt.Println(string(d.Body))
			var res Response
			err := json.Unmarshal(d.Body, &res)
			if err != nil {
				discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to unmarshal the response", err.Error()), "API - Site_Analyzer Fix")
			}
			return res.Body.Services, nil
		}
	}

	return nil, errors.New("failed to get services")
}

func GetAdContent(url string, services []string) ([]string, []string, error) {
	// establish a connection
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			discord.SendMessage(discord.Error, "Failed to close RabbitMQ connection", err.Error())
		}
	}(conn)

	// create a channel
	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			discord.SendMessage(discord.Error, "Failed to close RabbitMQ connection", err.Error())
		}
	}(ch)

	// declare a queue
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

	// consume the response
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
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// generate a random string
	corrId := uuid.New().String()

	// Create a request
	req := Request{
		Route: "GenerateAdContent",
		Body: RequestBody{
			URL:      url,
			Services: services,
		},
	}

	message, err := json.Marshal(req)
	if err != nil {
		failOnError(err, "Failed to marshal the request")
	}

	// publish the message
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(message),
		})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			var res Response
			err := json.Unmarshal(d.Body, &res)
			if err != nil {
				failOnError(err, "Failed to unmarshal the response")
			}
			return res.Body.Headlines, res.Body.Descriptions, nil
		}
	}

	return nil, nil, errors.New("failed to get services")
}
