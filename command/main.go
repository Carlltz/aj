package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var cyan = color.New(color.FgCyan).SprintFunc()

// GetLastCommand gets the last command and its output from the history
func GetLastCommand() (Command, error) {
	var command Command

	// Get last command from history, excl. first entry (which is the current command)
	out, err := ExecuteCommand("history --max 2")
	if err != "" {
		return Command{}, fmt.Errorf("couldn't get last command: %s", err)
	}
	command.Command = strings.TrimSpace(strings.Split(out, "\n")[1])

	// Execute the last command to get its error output
	out, command.Output = ExecuteCommand(command.Command)
	if out != "" {
		return Command{}, fmt.Errorf("command %s executed successfully:\n%s", cyan(command.Command), out)
	}

	return command, nil
}

// RunCommandStdOut runs a command and outputs to os.Stdout and os.Stderr
func RunCommandStdOut(command string) error {
	cmd := exec.Command("fish", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
