package openai

import (
	"context"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/sashabaranov/go-openai"
)

func GPT35Turbo(prompt string) string {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		discord.SendMessage(discord.Error, "OpenAI", err.Error())
	}
	return resp.Choices[0].Message.Content
}
