package google_ads_controller

import (
	"context"
	"encoding/json"
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

type Message struct {
	Route string `json:"route" example:"/get_customers"`
	Body  Body   `json:"body,omitempty" example:"{'customer_name': 'Test Customer'}"`
}

type Body struct {
	Id           uint   `json:"id,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	//Campaign
	CustomerId     uint   `json:"customer_id,omitempty"`
	CampaignName   string `json:"campaign_name,omitempty"`
	CampaignBudget uint   `json:"campaign_budget,omitempty"`
	//Ad Group Ads
	AdGroupId    uint     `json:"ad_group_id,omitempty"`
	Headlines    []string `json:"headlines,omitempty"`
	Descriptions []string `json:"descriptions,omitempty"`
	FinalURL     string   `json:"final_url,omitempty"`
	//Ad Group
	CampaignId  uint   `json:"campaign_id,omitempty"`
	AdGroupName string `json:"ad_group_name,omitempty"`
	MinCPCBid   uint   `json:"min_cpc_bid,omitempty"`
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
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to connect to RabbitMQ", err.Error()), "API - Google_Ads_Controller Fix")
		return ""
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to open a channel", err.Error()), "API - Google_Ads_Controller Fix")
		return ""
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RMQConfig.Queue, // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // noWait
		nil,             // arguments
	)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to declare a queue", err.Error()), "API - Google_Ads_Controller Fix")
		return ""
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
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to declare a reply queue", err.Error()), "API - Google_Ads_Controller Fix")
		return ""
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
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to register a consumer", err.Error()), "API - Google_Ads_Controller Fix")
		return ""
	}

	corrId := uuid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg, err := json.Marshal(message)
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to marshal message", err.Error()), "API - Google_Ads_Controller Fix")
		return ""
	}

	err = ch.PublishWithContext(ctx,
		"",          // exchange
		"gac_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       replyQ.Name,
			Body:          msg,
		})
	if err != nil {
		discord.SendMessage(discord.Error, fmt.Sprintf("%s: %s", "Failed to publish a message", err.Error()), "API - Google_Ads_Controller Fix")
		return ""
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			res := string(d.Body)
			return res
		}
	}

	return ""
}
