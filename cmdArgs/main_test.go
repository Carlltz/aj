package cmdArgs_test

import (
	"reflect"
	"testing"

	"github.com/Carlltz/aj/cmdArgs"
)

func TestGetCmdFlags(t *testing.T) {
	tests := []struct {
		testId  string
		name    string
		args    []string
		want    cmdArgs.Flags
		wantErr bool
	}{
		{
			testId: "nothing provided",
			name:   "no flags, default to correct",
			args:   []string{},
			want: cmdArgs.Flags{
				Cmd: cmdArgs.CmdCorrect,
			},
		},
		{
			testId: "only instructions provided",
			name:   "no flags with instructions, default to generate",
			args:   []string{"some instructions"},
			want: cmdArgs.Flags{
				Cmd:     cmdArgs.CmdGenerate,
				Content: "some instructions",
			},
		},
		{
			testId: "explicit correct command",
			name:   "explicit correct command",
			args:   []string{"fix"},
			want: cmdArgs.Flags{
				Cmd: cmdArgs.CmdCorrect,
			},
		},
		{
			testId: "explicit generate command",
			name:   "explicit generate command with instructions",
			args:   []string{"gen", "some instructions"},
			want: cmdArgs.Flags{
				Cmd:     cmdArgs.CmdGenerate,
				Content: "some instructions",
			},
		},
		{
			testId: "explicit generate command with shell flag",
			name:   "explicit generate command with shell flag",
			args:   []string{"--shell", "bash", "gen", "some instructions"},
			want: cmdArgs.Flags{
				Shell:   cmdArgs.Bash,
				Cmd:     cmdArgs.CmdGenerate,
				Content: "some instructions",
			},
		},
		{
			testId:  "invalid shell flag",
			name:    "invalid shell flag",
			args:    []string{"--shell", "invalid", "generate", "some instructions"},
			wantErr: true,
		},
		{
			testId:  "missing value for shell flag",
			name:    "missing value for shell flag",
			args:    []string{"--shell", "generate", "some instructions"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmdArgs.GetCmdFlags(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCmdFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want.Shell == "" {
				tt.want.Shell = got.Shell // Set shell to got.Shell for comparison since it may come from config or env
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: GetCmdFlags() = %v, want %v", tt.testId, got, tt.want)
			}
		})
	}
}
