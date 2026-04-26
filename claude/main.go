package claude

import (
	"github.com/Carlltz/aj/config"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

var Client *anthropic.Client

// ConnectClaude connects to Claude API
func ConnectClaude() error {
	claudeKey, err := config.GetAPIKey()
	if err != nil {
		return err
	}
	client := anthropic.NewClient(option.WithAPIKey(claudeKey))
	Client = &client
	return nil
}
