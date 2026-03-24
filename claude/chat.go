package claude

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Carlltz/aj/cmdArgs"
	"github.com/Carlltz/aj/command"
	"github.com/Carlltz/aj/config"
	"github.com/Carlltz/aj/utils"
	"github.com/anthropics/anthropic-sdk-go"
)

type response struct {
	NewCommand string `json:"new_command" jsonschema_description:"The new working command"`
}

var toolInputSchema = anthropic.ToolInputSchemaParam{
	Properties: map[string]any{
		"new_command": map[string]any{
			"type":        "string",
			"description": "The new working command",
		},
	},
	Required: []string{"new_command"},
}

var toolDefinition = anthropic.ToolUnionParam{
	OfTool: &anthropic.ToolParam{
		Name:        "response",
		Description: anthropic.String("The response with the new working command"),
		InputSchema: toolInputSchema,
	},
}

// CorrectCommand corrects a command using Claude
func CorrectCommand(ctx context.Context, cmd command.Command) (string, error) {
	// Question to ask the AI, fine-tuning needed!
	// No newline after output print, since it's already included in the output
	question := fmt.Sprintf(`This command ran in %s shell on %s:
%s

Status:
%s

Output:
%s

Correct it so that it executes successfully without changing anything else. Use the response tool to provide the corrected command.`, utils.GetShell(), utils.GetOS(), cmd.Command, cmd.Status, cmd.Output)

	// Ask the AI to correct the command
	message, err := Client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeHaiku4_5,
		MaxTokens: 1024,
		Tools: []anthropic.ToolUnionParam{
			toolDefinition,
		},
		ToolChoice: anthropic.ToolChoiceParamOfTool("response"),
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(question)),
		},
	})
	if err != nil {
		return "", err
	}

	// Extract the tool use response
	for _, block := range message.Content {
		if block.Type == "tool_use" {
			toolUse := block.AsToolUse()
			resp := response{}
			err = json.Unmarshal(toolUse.Input, &resp)
			if err != nil {
				return "", err
			}
			return resp.NewCommand, nil
		}
	}

	return "", fmt.Errorf("no tool use response found")
}

// GenerateCommand generates a command using Claude
func GenerateCommand(ctx context.Context, flags cmdArgs.Flags) (string, error) {
	cfg := config.GetConfig()
	// Question to ask the AI, fine-tuning needed!
	question := fmt.Sprintf(`Please generate a command for %s shell on %s that achieves the following: %s. Use the response tool to provide the command.`, flags.Shell, cfg.Os, flags.Content)

	// Ask the AI to generate the command
	message, err := Client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeHaiku4_5,
		MaxTokens: 1024,
		Tools: []anthropic.ToolUnionParam{
			toolDefinition,
		},
		ToolChoice: anthropic.ToolChoiceParamOfTool("response"),
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(question)),
		},
	})
	if err != nil {
		return "", err
	}

	// Extract the tool use response
	for _, block := range message.Content {
		if block.Type == "tool_use" {
			toolUse := block.AsToolUse()
			resp := response{}
			err = json.Unmarshal(toolUse.Input, &resp)
			if err != nil {
				return "", err
			}
			return resp.NewCommand, nil
		}
	}

	return "", fmt.Errorf("no tool use response found")
}
