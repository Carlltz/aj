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

type CommandResult struct {
	Command   string
	Messages  []anthropic.MessageParam
	ToolUseID string
}

func callClaude(ctx context.Context, model anthropic.Model, messages []anthropic.MessageParam) (CommandResult, error) {
	message, err := Client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     model,
		MaxTokens: 1024,
		Tools: []anthropic.ToolUnionParam{
			toolDefinition,
		},
		ToolChoice: anthropic.ToolChoiceParamOfTool("response"),
		Messages:   messages,
	})
	if err != nil {
		return CommandResult{}, err
	}

	for _, block := range message.Content {
		if block.Type == "tool_use" {
			toolUse := block.AsToolUse()
			resp := response{}
			if err = json.Unmarshal(toolUse.Input, &resp); err != nil {
				return CommandResult{}, err
			}
			updated := append(messages, anthropic.NewAssistantMessage(
				anthropic.NewToolUseBlock(toolUse.ID, toolUse.Input, toolUse.Name),
			))
			return CommandResult{
				Command:   resp.NewCommand,
				Messages:  updated,
				ToolUseID: toolUse.ID,
			}, nil
		}
	}

	return CommandResult{}, fmt.Errorf("no tool use response found")
}

// refine appends a tool result + follow-up text and calls Claude again.
func refine(ctx context.Context, model anthropic.Model, prior CommandResult, followUp string) (CommandResult, error) {
	messages := append(prior.Messages,
		anthropic.NewUserMessage(
			anthropic.NewToolResultBlock(prior.ToolUseID, prior.Command, false),
			anthropic.NewTextBlock(followUp),
		),
	)
	return callClaude(ctx, model, messages)
}

// CorrectCommand corrects a command using Claude.
// Pass nil for priorMessages on the first call.
func CorrectCommand(ctx context.Context, cmd command.Command, priorMessages []anthropic.MessageParam) (CommandResult, error) {
	messages := priorMessages
	if len(messages) == 0 {
		question := fmt.Sprintf(`This command ran in %s shell on %s:
%s

Status:
%s

Output:
%s

Correct it so that it executes successfully without changing anything else. Use the response tool to provide the corrected command.`, utils.GetShell(), utils.GetOS(), cmd.Command, cmd.Status, cmd.Output)
		messages = []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(question)),
		}
	}
	return callClaude(ctx, anthropic.ModelClaudeHaiku4_5, messages)
}

// RefineCorrection continues a CorrectCommand conversation with user follow-up text.
func RefineCorrection(ctx context.Context, prior CommandResult, followUp string) (CommandResult, error) {
	return refine(ctx, anthropic.ModelClaudeHaiku4_5, prior, followUp)
}

// GenerateCommand generates a command using Claude.
// Pass nil for priorMessages on the first call.
func GenerateCommand(ctx context.Context, flags cmdArgs.Flags, priorMessages []anthropic.MessageParam) (CommandResult, error) {
	cfg := config.GetConfig()

	messages := priorMessages
	if len(messages) == 0 {
		question := fmt.Sprintf(`Please generate a command for %s shell on %s that achieves the following: %s. Use the response tool to provide the command.`, flags.Shell, cfg.Os, flags.Content)
		messages = []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(question)),
		}
	}
	return callClaude(ctx, anthropic.ModelClaudeSonnet4_6, messages)
}

// RefineCommand continues a GenerateCommand conversation with user follow-up text.
func RefineCommand(ctx context.Context, flags cmdArgs.Flags, prior CommandResult, followUp string) (CommandResult, error) {
	return refine(ctx, anthropic.ModelClaudeSonnet4_6, prior, followUp)
}
