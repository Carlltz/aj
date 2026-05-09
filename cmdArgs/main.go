package cmdArgs

import (
	"fmt"
	"strings"

	"github.com/Carlltz/aj/config"
	"github.com/Carlltz/aj/utils"
)

// Skip the program name (os.Args[1:])
func GetCmdFlags(args []string) (Flags, error) {
	flags := Flags{}

	argsIndex := 0
outerLoop:
	for argsIndex < len(args)-1 {
		arg := args[argsIndex]

		var extraFields []string
		if numExtras, ok := FlagsWithExtraFields[arg]; ok {
			for range numExtras {
				argsIndex++
				if argsIndex >= len(args) {
					return flags, fmt.Errorf("missing value for flag: %s", arg)
				}
				extraFields = append(extraFields, args[argsIndex])
			}
		}

		switch arg {
		case "--shell", "-s":
			flags.Shell = Shells(extraFields[0])
			if err := flags.Shell.Validate(); err != nil {
				return flags, err
			}
		default:
			// No more flags
			break outerLoop
		}

		argsIndex++
	}

	if len(args) > 0 {
		flags.Cmd = CmdType(args[argsIndex])
		if err := flags.Cmd.Validate(); err != nil {
			// If no command identified default to correct if no instructions or generate if instructions
			if argsIndex == len(args) {
				flags.Cmd = CmdCorrect
			} else {
				flags.Cmd = CmdGenerate
			}
		} else {
			argsIndex++
		}
	} else {
		flags.Cmd = CmdCorrect
	}

	if flags.Shell == "" {
		cfg := config.GetConfig()
		if cfg.Shell != "" {
			flags.Shell = Shells(cfg.Shell)
			if err := flags.Shell.Validate(); err != nil {
				return flags, err
			}
		} else {
			flags.Shell = Shells(utils.GetShell())
			if err := flags.Shell.Validate(); err != nil {
				return flags, err
			}
		}
	}

	flags.Content = strings.Join(args[argsIndex:], " ")

	return flags, nil
}
