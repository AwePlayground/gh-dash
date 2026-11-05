package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var prActions = []struct{ key, label string }{
	{"omp", "Fix comments / CI with omp"},
	{"claude", "Fix comments / CI with claude"},
	{"codex", "Fix comments / CI with codex"},
	{"shell", "Open shell in worktree"},
	{"lazygit", "Open lazygit in worktree"},
	{"prompt", "Copy fix prompt to clipboard"},
}

type PRActionMenu struct {
	title    string
	selected int
}

func NewPRActionMenu(title string) *PRActionMenu { return &PRActionMenu{title: title} }
func (m *PRActionMenu) Selected() string         { return prActions[m.selected].key }
func (m *PRActionMenu) Update(msg tea.KeyPressMsg) {
	switch msg.Code {
	case tea.KeyDown, 'j':
		if m.selected < len(prActions)-1 {
			m.selected++
		}
	case tea.KeyUp, 'k':
		if m.selected > 0 {
			m.selected--
		}
	}
}
func (m *PRActionMenu) View() string {
	var s strings.Builder
	s.WriteString(m.title)
	s.WriteString("\n\n")
	for i, a := range prActions {
		prefix := "  "
		if i == m.selected {
			prefix = "› "
		}
		s.WriteString(prefix + a.label + "\n")
	}
	s.WriteString("\nEnter: select   Esc: cancel")
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).Render(s.String())
}
