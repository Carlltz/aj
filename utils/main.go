package utils

import (
	"os"
	"strings"
)

var shell string

func GetShell() string {
	if shell != "" {
		return shell
	}

	shell = os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}
	shell = shell[strings.LastIndex(shell, "/")+1:]
	return shell
}
