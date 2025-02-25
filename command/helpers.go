package command

import (
	"bytes"
	"fmt"
	"os/exec"
)

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

// Get last command from history, excl. first entry (which is the current command)
//
// Use "<>@%/:" as a delimiter to separate command, status, and output
func getShellHistoryCommand(shell string) (string, error) {
	switch shell {
	case "fish":
		return `set cmd (history -2 | sed -n "2p")
set output (eval $cmd 2>&1)
echo "$cmd<>@%/:$status<>@%/:$output"`, nil
	case "bash":
		return `cmd="$(history 2 | sed -n '1p' | cut -c 8-)"
output="$(eval "$cmd" 2>&1)"
echo "$cmd<>@%/:$?<>@%/:$output"`, nil
	}

	return "", fmt.Errorf("unsupported shell: %s", shell)
}
