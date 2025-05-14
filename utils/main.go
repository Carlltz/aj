package utils

import (
	"os"
	"runtime"
	"strings"
)

var shell string

// /opt/homebrew/bin/fish
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

func GetOS() string {
	os := runtime.GOOS
	switch os {
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default:
		return os
	}
}
