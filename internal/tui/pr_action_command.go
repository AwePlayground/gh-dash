package tui

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/atotto/clipboard"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/prrow"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/constants"
)

type prPromptCopiedMsg struct{}

func (m *Model) runPRAction(action string, pr *prrow.Data) tea.Cmd {
	repo := pr.GetRepoNameWithOwner()
	repoPath, ok := m.ctx.Config.RepoPaths[repo]
	if !ok {
		return func() tea.Msg { return constants.ErrMsg{Err: fmt.Errorf("configure repoPaths for %s", repo)} }
	}
	if action == "prompt" {
		return func() tea.Msg {
			launcher := filepath.Join(os.Getenv("HOME"), "bin", "pr-ai")
			prompt, err := exec.Command(launcher, "prompt", repo, fmt.Sprint(pr.Primary.Number), pr.Primary.HeadRefName, repoPath).Output()
			if err != nil {
				return constants.ErrMsg{Err: fmt.Errorf("generate fix prompt: %w", err)}
			}
			if err := clipboard.WriteAll(string(prompt)); err != nil {
				return constants.ErrMsg{Err: fmt.Errorf("copy fix prompt: %w", err)}
			}
			return prPromptCopiedMsg{}
		}
	}
	launcher := filepath.Join(os.Getenv("HOME"), "bin", "pr-ai")
	command := exec.Command(launcher, action, repo, fmt.Sprint(pr.Primary.Number), pr.Primary.HeadRefName, repoPath)
	if action == "shell" || action == "lazygit" {
		return tea.ExecProcess(command, func(err error) tea.Msg {
			if err != nil {
				return constants.ErrMsg{Err: err}
			}
			return execProcessFinishedMsg{}
		})
	}
	return func() tea.Msg {
		if err := command.Run(); err != nil {
			return constants.ErrMsg{Err: err}
		}
		return execProcessFinishedMsg{}
	}
}
