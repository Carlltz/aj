package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Carlltz/aj/command"
	"github.com/openai/openai-go"
)

type Response struct {
	NewCommand string `json:"new_command" jsonschema_description:"The new fixed command"`
}

var ResponseSchema = GenerateSchema[Response]()
var schemaParam = openai.ResponseFormatJSONSchemaJSONSchemaParam{
	Name:        openai.F("response"),
	Description: openai.F("The response with the new command"),
	Schema:      openai.F(ResponseSchema),
	Strict:      openai.Bool(true),
}

func CorrectCommand(command command.Command) (string, error) {
	ctx := context.Background()

	question := fmt.Sprintf(`This command ran in fish shell: %s
Gave the following output: %s

Could you correct it so that it will execute successfully, change as little as possible.`, command.Command, command.Output)

	chat, err := Client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		}),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		// Only certain models: https://platform.openai.com/docs/guides/structured-outputs#supported-models
		Model: openai.F(openai.ChatModelGPT4oMini),
	})
	if err != nil {
		return "", err
	}

	response := Response{}
	err = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &response)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(response.NewCommand), nil
}
