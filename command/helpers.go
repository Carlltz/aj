package command

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/Carlltz/aj/utils"
)

// runCommandReturnOut executes a command and returns its stdout and stderr
func runCommandReturnOut(command string) (string, string) {
	cmd := exec.Command(utils.GetShell(), "-c", command)
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

// Get last command from history, excl. first entry (which is the current command)
//
// Use "<>@%/:" as a delimiter to separate command, status, and output
func getShellHistoryCommand() (string, error) {
	shell := utils.GetShell()
	switch shell {
	case "fish":
		return `set cmd (history -2 | sed -n "2p")
set output (eval $cmd 2>&1)
echo "$cmd<>@%/:$status<>@%/:$output"`, nil
	case "bash":
		return `cmd="$(tail -n 2 ~/.bash_history | sed -n '1p')"
output="$(eval "$cmd" 2>&1)"
echo "$cmd<>@%/:$?<>@%/:$output"`, nil
	}

	return "", fmt.Errorf("unsupported shell: %s", shell)
}
