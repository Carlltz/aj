package ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Carlltz/aj/command"
	"github.com/Carlltz/aj/utils"
	"github.com/openai/openai-go"
)

type response struct {
	NewCommand string `json:"new_command" jsonschema_description:"The new fixed command"`
}

var schemaParam = openai.ResponseFormatJSONSchemaJSONSchemaParam{
	Name:        openai.F("response"),
	Description: openai.F("The response with the new command"),
	Schema:      openai.F(generateSchema[response]()),
	Strict:      openai.Bool(true),
}

// CorrectCommand corrects a command using OpenAI
func CorrectCommand(ctx context.Context, command command.Command) (string, error) {
	// Question to ask the AI, fine-tuning needed!
	// No newline after output print, since it's already included in the output
	question := fmt.Sprintf(`This command ran in %s shell on %s:
%s

Status:
%s

Output:
%s
Correct it so that it executes successfully, change as little as possible.`, utils.GetShell(), utils.GetOS(), command.Command, command.Status, command.Output)

	// Ask the AI to correct the command
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

	response := response{}
	err = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &response)
	if err != nil {
		return "", err
	}

	return response.NewCommand, nil
}
