package tools

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/Carlltz/aj/config"
	"github.com/Carlltz/aj/tools/views"
)

var choices = []string{ /* "Set LLM Type", */ "Set API Key", "Set Default Shell"}

type screen int

const (
	screenMainMenu screen = iota
	screenSetShell
	screenSetAPIKey
	screenConfirm
)

type confirmDoneMsg struct{}

func waitThenReturn() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second)
		return confirmDoneMsg{}
	}
}

type menuModel struct {
	cursor int
	choice string
}

func (m menuModel) Init() tea.Cmd { return nil }

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "enter":
			m.choice = choices[m.cursor]
			return m, nil
		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}
	return m, nil
}

func (m menuModel) View() tea.View {
	s := strings.Builder{}
	s.WriteString("Currently configured:\n")
	cfg := config.GetConfig()
	s.WriteString(fmt.Sprintf("Default Shell: %s\n", cfg.Shell))
	s.WriteString("\n")

	s.WriteString("What would you like to configure?\n\n")
	for i := range choices {
		if m.cursor == i {
			s.WriteString("(X) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")
	return tea.NewView(s.String())
}

type rootModel struct {
	screen     screen
	menu       menuModel
	shell      views.ShellModel
	apiKey     views.APIKeyModel
	confirmMsg string
}

func (r rootModel) Init() tea.Cmd { return nil }

func (r rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch r.screen {
	case screenMainMenu:
		updated, cmd := r.menu.Update(msg)
		r.menu = updated.(menuModel)
		switch r.menu.choice {
		case "Set Default Shell":
			r.screen = screenSetShell
			r.shell = views.ShellModel{}
			r.menu.choice = ""
			return r, nil
		case "Set API Key":
			r.screen = screenSetAPIKey
			r.apiKey = views.APIKeyModel{}
			r.menu.choice = ""
			return r, nil
		}
		return r, cmd

	case screenSetShell:
		updated, cmd := r.shell.Update(msg)
		r.shell = updated.(views.ShellModel)
		if r.shell.Back {
			r.screen = screenMainMenu
			r.shell = views.ShellModel{}
			return r, nil
		}
		if r.shell.Choice != "" {
			cfg := config.GetConfig()
			cfg.Shell = r.shell.Choice
			config.SetConfig(cfg)
			r.shell = views.ShellModel{}
			r.confirmMsg = "Default shell set to: " + cfg.Shell
			r.screen = screenConfirm
			return r, waitThenReturn()
		}
		return r, cmd

	case screenSetAPIKey:
		updated, cmd := r.apiKey.Update(msg)
		r.apiKey = updated.(views.APIKeyModel)
		if r.apiKey.Back {
			r.screen = screenMainMenu
			r.apiKey = views.APIKeyModel{}
			return r, nil
		}
		if r.apiKey.Key != "" {
			if err := config.SetAPIKey(r.apiKey.Key); err != nil {
				r.confirmMsg = "Error saving API key: " + err.Error()
			} else {
				r.confirmMsg = "API key saved."
			}
			r.apiKey = views.APIKeyModel{}
			r.screen = screenConfirm
			return r, waitThenReturn()
		}
		return r, cmd

	case screenConfirm:
		if _, ok := msg.(confirmDoneMsg); ok {
			r.screen = screenMainMenu
			r.confirmMsg = ""
		}
		return r, nil
	}
	return r, nil
}

func (r rootModel) View() tea.View {
	switch r.screen {
	case screenMainMenu:
		return r.menu.View()
	case screenSetShell:
		return r.shell.View()
	case screenSetAPIKey:
		return r.apiKey.View()
	case screenConfirm:
		return tea.NewView(r.confirmMsg + "\n")
	}
	return tea.NewView("")
}

func ConfigCommand(ctx context.Context) {
	p := tea.NewProgram(rootModel{})
	_, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
