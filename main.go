package main

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Carlltz/aj/ai"
	"github.com/Carlltz/aj/command"
	"github.com/fatih/color"
)

//go:embed .env
var envFile string

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func main() {
	// Create a context that listens for Ctrl+C (SIGINT) and SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Listen for context cancellation
	go func() {
		<-ctx.Done()
		fmt.Printf("%s", red("\nExiting..."))
		os.Exit(0)
	}()

	// Connect to OpenAI
	ai.ConnectOpenAI(envFile)

	// Get the last command
	lastCommand, err := command.GetLastCommand()
	if err != nil {
		fmt.Printf("%s\n%s", red("Error getting failed command"), err)
		return
	}

	// Print the last command
	fmt.Print("Last command: ")
	color.Blue(lastCommand.Command)

	// Correct the last command
	updatedCommand, err := ai.CorrectCommand(ctx, lastCommand)
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
