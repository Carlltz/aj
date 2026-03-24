package claude

import (
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

var Client *anthropic.Client

// ConnectClaude connects to Claude API
func ConnectClaude(envFile string) {
	claudeKey := strings.Trim(strings.SplitN(envFile, "=", 2)[1], "\n")
	client := anthropic.NewClient(option.WithAPIKey(claudeKey))
	Client = &client
}
