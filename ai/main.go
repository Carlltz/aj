package ai

import (
	"log"
	"os"
	"sync"

	"github.com/openai/openai-go"
)

var Client *openai.Client

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
			log.Fatal("OPENAI_API_KEY is not set")
		}

		Client = openai.NewClient() // defaults to os.LookupEnv("OPENAI_API_KEY")
		// log.Println("Connected to OpenAI")
	}()

	return &wg
}
