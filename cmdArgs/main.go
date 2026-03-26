package cmdArgs

import (
	"fmt"
	"os"
	"strings"

	"github.com/Carlltz/aj/config"
	"github.com/Carlltz/aj/utils"
)

func GetCmdFlags() (Flags, error) {
	args := os.Args[1:]
	flags := Flags{}

	if len(args) == 0 {
		return flags, nil
	}

	argsIndex := 0
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
			break
		}

		argsIndex++
	}

	flags.Cmd = CmdType(args[argsIndex])
	if err := flags.Cmd.Validate(); err != nil {
		// If no command identified default to correct if no instructions or generate if instructions
		if argsIndex == len(args)-1 {
			flags.Cmd = CmdCorrect
		} else {
			flags.Cmd = CmdGenerate
		}
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

	flags.Content = strings.Join(args[argsIndex+1:], " ")

	return flags, nil
}
