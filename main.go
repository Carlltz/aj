package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/Carlltz/aj/ai"
	"github.com/Carlltz/aj/command"
	"github.com/fatih/color"
)

//go:embed .env
var envFile string

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func main() {
	// Load environment variables from embedded .env file
	lines := strings.Split(envFile, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			_ = os.Setenv(parts[0], parts[1])
		}
	}

	// Connect to OpenAI
	wg := ai.ConnectOpenAI()

	// Get the last command
	lastCommand, err := command.GetLastCommand()
	if err != nil {
		fmt.Printf("%s\n%s", red("Error getting failed command"), err)
		return
	}

	// Print the last command
	fmt.Print("Last command: ")
	color.Blue(lastCommand.Command)

	// Wait for OpenAI to connect
	wg.Wait()

	// Correct the last command
	updatedCommand, err := ai.CorrectCommand(lastCommand)
	if err != nil {
		fmt.Printf("%s\n%s", red("Error correcting command"), err)
		return
	}

	// Print the corrected command
	fmt.Printf("%s or %s: ", red("Ctrl+C to exit"), green("Enter to run"))
	color.Cyan(updatedCommand)

	// Wait for user input
	var input string
	fmt.Scanln(&input)

	// Run the updated command on enter
	if input == "" {
		err := command.RunCommandStdOut(updatedCommand)
		if err != nil {
			fmt.Printf("%s %s", red("Error running command:"), err)
		}
	}
}
