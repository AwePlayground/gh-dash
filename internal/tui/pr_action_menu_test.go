package tui

import (
	tea "charm.land/bubbletea/v2"
	"testing"
)

func TestPRActionMenuNavigation(t *testing.T) {
	m := NewPRActionMenu("owner/repo #42")
	if m.Selected() != "omp" {
		t.Fatalf("default = %q", m.Selected())
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	if m.Selected() != "claude" {
		t.Fatalf("down = %q", m.Selected())
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	if m.Selected() != "omp" {
		t.Fatalf("up = %q", m.Selected())
	}
}

func TestPRActionMenuCanSelectCopyPrompt(t *testing.T) {
	m := NewPRActionMenu("owner/repo #42")
	for range 5 {
		m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	}
	if m.Selected() != "prompt" {
		t.Fatalf("last action = %q; want prompt", m.Selected())
	}
}
