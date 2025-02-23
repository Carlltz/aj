package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetLastCommand gets the last command and its output from the history
func GetLastCommand() (Command, error) {
	// Get last command from history, excl. first entry (which is the current command)
	// Use "<>@%/:" as a delimiter to separate command, status, and output
	out, err := runCommandReturnOut(`
	set cmd (history -2 | sed -n "2p")
    set output (eval $cmd 2>&1)
	echo "$cmd<>@%/:$status<>@%/:$output"
	`)
	if err != "" {
		return Command{}, fmt.Errorf("couldn't get last command: %s", err)
	}

	parts := strings.Split(out, "<>@%/:")
	if len(parts) != 3 {
		return Command{}, fmt.Errorf("couldn't parse last command output: %s", out)
	}

	var command Command
	command.Command = parts[0]
	command.Status = parts[1]
	command.Output = parts[2]

	return command, nil
}

// runCommandReturnOut executes a command and returns its stdout and stderr
func runCommandReturnOut(command string) (string, string) {
	cmd := exec.Command("fish", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", stderr.String()
	}
	return out.String(), ""
}

// RunCommandStdOut runs a command and outputs to os.Stdout and os.Stderr
func RunCommandStdOut(command string) error {
	cmd := exec.Command("fish", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
