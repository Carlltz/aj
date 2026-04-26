package tools

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Carlltz/aj/claude"
	"github.com/Carlltz/aj/cmdArgs"
	"github.com/Carlltz/aj/command"
	"github.com/fatih/color"
)

func CorrectCommand(ctx context.Context, cmdFlags cmdArgs.Flags) {
	// Connect to Claude
	err := claude.ConnectClaude()
	if err != nil {
		fmt.Printf("%s\n%s", red("Error connecting to Claude"), err)
		return
	}

	// Get the last command
	lastCommand, err := command.GetLastCommand(cmdFlags)
	if err != nil {
		fmt.Printf("%s\n%s", red("Error getting failed command"), err)
		return
	}

	// Print the last command
	fmt.Print("Last command: ")
	color.Blue(lastCommand.Command)

	// Correct the last command
	updatedCommand, err := claude.CorrectCommand(ctx, lastCommand)
	if err != nil {
		fmt.Printf("%s\n%s", red("Error correcting command"), err)
		return
	}

	// Print the corrected command
	fmt.Printf("%s: ", green("Enter to run"))
	color.Cyan(updatedCommand)
	fmt.Printf("%s\n", red("Ctrl+C to exit"))

	// Listen for Enter key press
	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n') // Wait for Enter
		fmt.Printf("%s\n", green("Output:"))
		command.RunCommandStdOut(updatedCommand)
		os.Exit(0)
	}()

	// Keep alive
	<-ctx.Done()
}
