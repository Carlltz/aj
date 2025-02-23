package ai

import (
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var Client *openai.Client

// ConnectOpenAI connects to OpenAI
func ConnectOpenAI(envFile string) {
	openaiKey := strings.Trim(strings.SplitN(envFile, "=", 2)[1], "\n")
	Client = openai.NewClient(option.WithAPIKey(openaiKey))
}
