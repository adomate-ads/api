package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func SendMessage(body []byte, queue string) error {
	conn, ch := handleReconnection()
	defer conn.Close()
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
	conn, ch := handleReconnection()
	defer conn.Close()
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
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
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
