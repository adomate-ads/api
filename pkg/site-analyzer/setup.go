package site_analyzer

import (
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/rabbitmq"
)

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

func GetServices(url string) ([]string, error) {
	msg := Request{
		Route: "GuessServices",
		Body: RequestBody{
			URL: fmt.Sprintf("https://%s", url),
		},
	}

	msgBody, err := json.Marshal(msg)
	if err != nil {
		discord.SendMessage(discord.Error, "[SA] Failed to marshal message", fmt.Sprintf("Error: %s", err.Error()))
		return nil, err
	}

	resp, err := rabbitmq.SendMessageWithResponse(msgBody, rabbitmq.RMQConfig.SAQueue)
	if err != nil {
		discord.SendMessage(discord.Error, "[SA] Failed to send message", fmt.Sprintf("Error: %s", err.Error()))
		return nil, err
	}

	var res Response
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		discord.SendMessage(discord.Error, "[SA] Failed to unmarshal response", fmt.Sprintf("Error: %s", err.Error()))
		return nil, err
	}

	return res.Body.Services, nil
}

func GetAdContent(url string, services []string) ([]string, []string, error) {
	msg := Request{
		Route: "GenerateAdContent",
		Body: RequestBody{
			URL:      fmt.Sprintf("https://%s", url),
			Services: services,
		},
	}

	msgBody, err := json.Marshal(msg)
	if err != nil {
		discord.SendMessage(discord.Error, "[SA] Failed to marshal message", fmt.Sprintf("Error: %s", err.Error()))
		return nil, nil, err
	}

	resp, err := rabbitmq.SendMessageWithResponse(msgBody, rabbitmq.RMQConfig.SAQueue)
	if err != nil {
		discord.SendMessage(discord.Error, "[SA] Failed to send message", fmt.Sprintf("Error: %s", err.Error()))
		return nil, nil, err
	}

	var res Response
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		discord.SendMessage(discord.Error, "[SA] Failed to unmarshal response", fmt.Sprintf("Error: %s", err.Error()))
		return nil, nil, err
	}

	return res.Body.Headlines, res.Body.Descriptions, nil
}
