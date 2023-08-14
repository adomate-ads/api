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
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	// declare a reply queue
	replyQ, err := ch.QueueDeclare(
		fmt.Sprintf("reply_%s", q.Name), // name
		true,                            // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to declare a reply queue", err.Error()), "API - Site_Analyzer Fix")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// generate a random string
	corrId := uuid.New().String()

	// Create a request
	req := Request{
		Route: "GuessServices",
		Body: RequestBody{
			URL: fmt.Sprintf("https://%s", url),
		},
	}

	message, err := json.Marshal(req)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to marshal the request", err.Error()), "API - Site_Analyzer Fix")
		return nil, err
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
			ReplyTo:       replyQ.Name,
			Body:          message,
		})
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to publish a message", err.Error()), "API - Site_Analyzer Fix")
		return nil, err
	}

	timeout := time.After(90 * time.Second)

	// consume the response
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
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to register a consumer", err.Error()), "API - Site_Analyzer Fix")
		return nil, err
	}

	for {
		select {
		case d, ok := <-msgs:
			if !ok {
				return nil, errors.New("failed to get services: messages channel closed")
			}
			if corrId == d.CorrelationId {
				var res Response
				err := json.Unmarshal(d.Body, &res)
				if err != nil {
					discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to unmarshal the response", err.Error()), "API - Site_Analyzer Fix")
					return nil, err
				}
				return res.Body.Services, nil
			}
		case <-timeout:
			return nil, errors.New("failed to get services: timed out after 90 seconds")
		}
	}
}

func GetAdContent(url string, services []string) ([]string, []string, error) {
	// establish a connection
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to connect to RabbitMQ", err.Error()), "API - Site_Analyzer Fix")
		return nil, nil, err
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
		return nil, nil, err
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
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to declare a queue", err.Error()), "API - Site_Analyzer Fix")
		return nil, nil, err
	}

	// declare a reply queue
	replyQ, err := ch.QueueDeclare(
		fmt.Sprintf("reply_%s", q.Name), // name
		true,                            // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to declare a reply queue", err.Error()), "API - Site_Analyzer Fix")
		return nil, nil, err
	}

	// consume the response
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
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to register a consumer", err.Error()), "API - Site_Analyzer Fix")
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// generate a random string
	corrId := uuid.New().String()

	// Create a request
	req := Request{
		Route: "GenerateAdContent",
		Body: RequestBody{
			URL:      fmt.Sprintf("https://%s", url),
			Services: services,
		},
	}

	message, err := json.Marshal(req)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to marshal the request", err.Error()), "API - Site_Analyzer Fix")
		return nil, nil, err
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
			ReplyTo:       replyQ.Name,
			Body:          message,
		})
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to publish a message", err.Error()), "API - Site_Analyzer Fix")
		return nil, nil, err
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			var res Response
			err := json.Unmarshal(d.Body, &res)
			if err != nil {
				discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to unmarshal the response", err.Error()), "API - Site_Analyzer Fix")
				return nil, nil, err
			}
			return res.Body.Headlines, res.Body.Descriptions, nil
		}
	}

	return nil, nil, errors.New("failed to get services")
}
