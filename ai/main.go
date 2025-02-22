package ai

import (
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/openai/openai-go"
)

var Client *openai.Client

// ConnectOpenAI connects to OpenAI and returns a WaitGroup to wait for the connection
func ConnectOpenAI() *sync.WaitGroup {
	wg := sync.WaitGroup{}
	wg.Add(1)
	if Client != nil {
		wg.Done()
		return &wg
	}

	go func() {
		defer wg.Done()
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			color.Red("OPENAI_API_KEY is not set!")
			os.Exit(1)
		}

		Client = openai.NewClient() // defaults to os.LookupEnv("OPENAI_API_KEY")
	}()

	return &wg
}
