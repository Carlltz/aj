package command

import (
	"bytes"
	"os/exec"
	"strings"
)

// ExecuteCommand executes a command and returns its stdout and stderr
func ExecuteCommand(command string) (string, string) {
	cmd := exec.Command("fish", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", strings.TrimSpace(stderr.String())
	}
	return strings.TrimSpace(out.String()), ""
}
