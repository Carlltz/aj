package views

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

var ShellChoices = []string{"fish", "bash"}

type ShellModel struct {
	Cursor int
	Choice string
	Back   bool
}

func (m ShellModel) Init() tea.Cmd { return nil }

func (m ShellModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			m.Back = true
			return m, nil
		case "enter":
			m.Choice = ShellChoices[m.Cursor]
			return m, nil
		case "down", "j":
			m.Cursor++
			if m.Cursor >= len(ShellChoices) {
				m.Cursor = 0
			}
		case "up", "k":
			m.Cursor--
			if m.Cursor < 0 {
				m.Cursor = len(ShellChoices) - 1
			}
		}
	}
	return m, nil
}

func (m ShellModel) View() tea.View {
	s := strings.Builder{}
	s.WriteString("Select default shell:\n\n")
	for i, shell := range ShellChoices {
		if m.Cursor == i {
			s.WriteString("(X) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(shell)
		s.WriteString("\n")
	}
	s.WriteString("\n(press esc to go back)\n")
	return tea.NewView(s.String())
}
