package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Carlltz/aj/ai"
	"github.com/Carlltz/aj/command"
	"github.com/fatih/color"
)

//go:embed .env
var envFile string

func main() {
	// Load environment variables from embedded .env file
	lines := strings.Split(envFile, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			_ = os.Setenv(parts[0], parts[1])
		}
	}

	wg := ai.ConnectOpenAI()

	lastCommand, err := command.GetLastCommand()
	if err != nil {
		log.Println("Error getting last command:", err)
		return
	}

	fmt.Print("\n")
	fmt.Print("Last command: ")
	color.Blue(lastCommand.Command)
	fmt.Print("\n")

	wg.Wait()

	updatedCommand, err := ai.CorrectCommand(lastCommand)
	if err != nil {
		log.Println("Error correcting command:", err)
		return
	}

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s or %s: ", red("Ctrl+C to exit"), green("Enter to run"))
	color.Cyan(updatedCommand)

	// Take user input
	var input string
	fmt.Scanln(&input)

	// Run the updated command
	if input == "" {
		err := command.RunCommand(updatedCommand)
		if err != nil {
			fmt.Printf("%s %s", red("Error running command:"), err)
		}
	}
}
