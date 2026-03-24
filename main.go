package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Carlltz/aj/cmdArgs"
	"github.com/Carlltz/aj/tools"
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

	cmdFlags, err := cmdArgs.GetCmdFlags()
	if err != nil {
		fmt.Printf("%s", red(fmt.Sprintf("Error parsing command flags: %v", err)))
		os.Exit(1)
	}

	switch cmdFlags.Cmd {
	case cmdArgs.CmdCorrect:
		tools.CorrectCommand(ctx)
	case cmdArgs.CmdGenerate:
		// Handle CmdGenerate
	}
}
