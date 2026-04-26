package views

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type APIKeyModel struct {
	input []rune
	Key   string
	Back  bool
}

func (m APIKeyModel) Init() tea.Cmd { return nil }

func (m APIKeyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.PasteMsg:
		m.input = append(m.input, []rune(msg.Content)...)
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.Back = true
			return m, nil
		case "enter":
			m.Key = string(m.input)
			return m, nil
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			r := []rune(msg.String())
			if len(r) == 1 {
				m.input = append(m.input, r[0])
			}
		}
	}
	return m, nil
}

func (m APIKeyModel) View() tea.View {
	s := strings.Builder{}
	s.WriteString("Enter API key:\n\n")
	s.WriteString(strings.Repeat("*", len(m.input)))
	s.WriteString("█\n")
	s.WriteString("\n(enter to confirm, esc to go back)\n")
	return tea.NewView(s.String())
}
