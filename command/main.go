package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Carlltz/aj/utils"
)

// GetLastCommand gets the last command and its output from the history
func GetLastCommand() (Command, error) {
	historyCommand, err := getShellHistoryCommand()
	if err != nil {
		return Command{}, err
	}

	out, errOut := runCommandReturnOut("history -2")
	fmt.Println(out)
	fmt.Println(errOut)
	os.Exit(0)

	out, errOut = runCommandReturnOut(historyCommand)
	if errOut != "" {
		return Command{}, fmt.Errorf("couldn't get last command: %s", errOut)
	}

	fmt.Println(out)
	os.Exit(0)

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

// RunCommandStdOut runs a command and outputs to os.Stdout and os.Stderr
func RunCommandStdOut(command string) error {
	cmd := exec.Command(utils.GetShell(), "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
