package tools

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Carlltz/aj/claude"
	"github.com/Carlltz/aj/cmdArgs"
	"github.com/Carlltz/aj/command"
	"github.com/fatih/color"
)

func CorrectCommand(ctx context.Context, cmdFlags cmdArgs.Flags) {
	err := claude.ConnectClaude()
	if err != nil {
		fmt.Printf("%s\n%s", red("Error connecting to Claude"), err)
		return
	}

	lastCommand, err := command.GetLastCommand(cmdFlags)
	if err != nil {
		fmt.Printf("%s\n%s", red("Error getting failed command"), err)
		return
	}

	fmt.Print("Last command: ")
	color.Blue(lastCommand.Command)

	result, err := claude.CorrectCommand(ctx, lastCommand, nil)
	if err != nil {
		fmt.Printf("%s\n%s", red("Error correcting command"), err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s ", green("Command:"))
		color.Cyan(result.Command)
		fmt.Printf("%s %s\n", green("Enter to run"), blue("or type a follow-up to refine:"))

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			fmt.Printf("%s\n", green("Output:"))
			command.RunCommandStdOut(result.Command)
			return
		}

		result, err = claude.RefineCorrection(ctx, result, line)
		if err != nil {
			fmt.Printf("%s\n%s", red("Error refining command"), err)
			return
		}
	}
}
