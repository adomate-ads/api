package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"time"
)

const (
	InitialBackOff = 5 * time.Second
	MaxBackOff     = 1 * time.Minute
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

func connectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RMQConfig.User, RMQConfig.Password, RMQConfig.Host, RMQConfig.Port))
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

func handleReconnection() (*amqp.Connection, *amqp.Channel) {
	backoff := InitialBackOff

	for {
		conn, ch, err := connectRabbitMQ()
		if err == nil {
			return conn, ch
		}

		fmt.Printf("[RabbitMQ] Failed to connect to RabbitMQ. Retrying in %v... Error: %v", backoff, err)
		time.Sleep(backoff)

		if backoff < MaxBackOff {
			backoff *= 2
		}
	}
}
