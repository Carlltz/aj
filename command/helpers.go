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
//
// Fish:
// function record_last_command --on-event fish_postexec
//
//	echo $argv[1] > ~/.last_fish_command
//
// end
//
// Note: bash requires `history -a` to be run before aj to ensure the history file is up to date
// better: PROMPT_COMMAND="history 1 | sed -n '1p' > ~/.last_bash_command"
//
// Note: zsh should have:
// # Before each prompt, store the last command in an env var
//
//	precmd() {
//	  fc -ln -1 > ~/.last_zsh_command
//	}
func getShellHistoryCommand() (string, error) {
	shell := utils.GetShell()
	switch shell {
	case "fish":
		return `set LASTCMD (cat ~/.last_fish_command)
set OUTPUT (eval $LASTCMD 2>&1)
echo "$LASTCMD<>@%/:$status<>@%/:$OUTPUT"`, nil
	case "bash":
		return `LASTCMD="$(cat ~/.last_bash_command)"
OUTPUT="$(eval "$LASTCMD" 2>&1)"
echo "$LASTCMD<>@%/:$?<>@%/:$OUTPUT"`, nil
	case "zsh":
		return `LASTCMD="$(cat ~/.last_zsh_command)"
OUTPUT="$(eval "$LASTCMD" 2>&1)"
echo "$LASTCMD<>@%/:$?<>@%/:$OUTPUT"`, nil
	}

	return "", fmt.Errorf("unsupported shell: %s", shell)
}
