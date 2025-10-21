package ai

import (
	"strings"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

var Client *openai.Client

// ConnectOpenAI connects to OpenAI
func ConnectOpenAI(envFile string) {
	openaiKey := strings.Trim(strings.SplitN(envFile, "=", 2)[1], "\n")
	client := openai.NewClient(option.WithAPIKey(openaiKey))
	Client = &client
}
