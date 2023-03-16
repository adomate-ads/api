package google_ads_controller

import (
	"context"
	"encoding/json"
	"fmt"
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

type Message struct {
	Route string `json:"route" example:"/get_customers"`
	Body  string `json:"body" example:"{'customer_id': '1234567890'}"`
}

var RMQConfig RabbitMQConfig

func Setup() {
	RMQConfig = RabbitMQConfig{
		Host:     os.Getenv("RABBIT_HOST"),
		Port:     os.Getenv("RABBIT_PORT"),
		User:     os.Getenv("RABBIT_USER"),
		Password: os.Getenv("RABBIT_PASS"),
		Queue:    os.Getenv("RABBIT_GAC_QUEUE"),
	}
}

func SendToQueue(message Message) string {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		return ""
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel")
		return ""
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RMQConfig.Queue, // name
		false,           // durable
		false,           // delete when unused
		true,            // exclusive
		false,           // noWait
		nil,             // arguments
	)
	if err != nil {
		fmt.Println("Failed to declare a queue")
		return ""
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
		fmt.Println("Failed to register a consumer")
		return ""
	}

	corrId := uuid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Failed to marshal message")
		return ""
	}

	err = ch.PublishWithContext(ctx,
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          msg,
		})
	if err != nil {
		fmt.Println("Failed to publish a message")
		return ""
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			res := string(d.Body)
			fmt.Println(" [.] Got %s", res)
			return res
			break
		}
	}

	return ""
}
