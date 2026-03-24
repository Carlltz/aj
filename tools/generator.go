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

func GenerateCommand(ctx context.Context, flags cmdArgs.Flags) {
	// Connect to Claude
	claude.ConnectClaude("envFile") // TODO

	// Generate a new command
	generatedCommand, err := claude.GenerateCommand(ctx, flags)
	if err != nil {
		fmt.Printf("%s\n%s", red("Error generating command"), err)
		return
	}

	// Print the generated command
	fmt.Printf("%s: ", green("Enter to run"))
	color.Cyan(generatedCommand)
	fmt.Printf("%s\n", red("Ctrl+C to exit"))

	// Listen for Enter key press
	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n') // Wait for Enter
		fmt.Printf("%s\n", green("Output:"))
		command.RunCommandStdOut(generatedCommand)
		os.Exit(0)
	}()

	// Keep alive
	<-ctx.Done()
}
