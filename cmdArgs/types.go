package cmdArgs

import "fmt"

type CmdType string

const (
	CmdCorrect  CmdType = "fix"
	CmdGenerate CmdType = "gen"
	CmdConfig   CmdType = "config"
)

func (c CmdType) Validate() error {
	if c != CmdCorrect && c != CmdGenerate && c != CmdConfig {
		return fmt.Errorf("invalid command: %s", c)
	}
	return nil
}

type Shells string

const (
	Fish Shells = "fish"
	Bash Shells = "bash"
)

func (s Shells) Validate() error {
	if s != Fish && s != Bash {
		return fmt.Errorf("invalid shell: %s", s)
	}
	return nil
}

var FlagsWithExtraFields = map[string]int{
	"--shell": 1,
	"-s":      1,
}

type Flags struct {
	Cmd     CmdType
	Shell   Shells
	Content string
}
