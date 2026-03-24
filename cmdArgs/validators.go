package cmdArgs

import "slices"

var validShells = []string{"bash", "zsh", "fish"}

func validateShell(shell string) bool {
	return slices.Contains(validShells, shell)
}
