package command

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetLastCommand() (Command, error) {
	var command Command

	// Get last command from history, excl. first entry (which is the current command)
	cmd := exec.Command("fish", "-c", "history --max 2")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return Command{}, err
	}
	command.Command = strings.TrimSpace(strings.Split(out.String(), "\n")[1])

	// Execute the last command to get its output
	cmd = exec.Command("fish", "-c", command.Command)
	out.Reset()
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err == nil {
		log.Println("Command executed successfully")
		// Return a new error
		return command, fmt.Errorf("command executed successfully: %s", out.String())
	}
	command.Output = strings.TrimSpace(stderr.String())

	return command, nil
}

func RunCommand(command string) error {
	cmd := exec.Command("fish", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
