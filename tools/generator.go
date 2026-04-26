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

func GenerateCommand(ctx context.Context, flags cmdArgs.Flags) {
	err := claude.ConnectClaude()
	if err != nil {
		fmt.Printf("%s\n%s", red("Error connecting to Claude"), err)
		return
	}

	result, err := claude.GenerateCommand(ctx, flags, nil)
	if err != nil {
		fmt.Printf("%s\n%s", red("Error generating command"), err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s ", green("Command:"))
		color.Cyan(result.Command)
		fmt.Printf("%s %s %s\n", green("Enter to run,"), red("Ctrl+C to exit"), blue("or type a follow-up to refine:"))

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			fmt.Printf("%s\n", green("Output:"))
			command.RunCommandStdOut(result.Command)
			return
		}

		result, err = claude.RefineCommand(ctx, flags, result, line)
		if err != nil {
			fmt.Printf("%s\n%s", red("Error refining command"), err)
			return
		}
	}
}
