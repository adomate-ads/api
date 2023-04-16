package openai

import (
	"github.com/sashabaranov/go-openai"
	"os"
)

var client *openai.Client

func Setup() {
	client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}
