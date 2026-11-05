# Local PR action menu

This fork adds a native `f` menu to the PR list. The menu stays inside the current terminal screen: Up/Down or j/k selects an action, Enter runs it, and Esc closes it. In the PR list, Enter focuses the preview, Up/Down or j/k scrolls its content without changing the selected PR, and Esc returns to the list.

The menu calls `~/bin/pr-ai` with the selected PR's repository, number, head branch, and matching `repoPaths` entry. Install that launcher separately and configure `repoPaths` for the repositories you work on. Options 4 and 5 open a shell and lazygit in the PR worktree. AI actions use the launcher's terminal-specific mode: Tern and tmux open a new tab/window; without a supported terminal, the selected tool runs inline and returns to the list when closed. Press `o` in the PR list to open the selected PR in the browser; this existing shortcut is not duplicated in the menu. The final menu option copies a fix prompt for the PR to the clipboard without launching a session, fetching branches, or creating a worktree; paste it into your existing session.

The PR table uses `` for all comments and review threads and `` for unresolved review threads. `unresolvedThreads.hidden: true` under `defaults.layout.prs` or a PR section's `layout` hides the unresolved column. GitHub returns at most 100 thread states per list result; a `+` suffix means more pages exist, so the displayed count is a lower bound rather than an exact total.

Build locally with `go build -o ~/bin/gh-dash-fork .` and run `~/bin/gh-dash-fork --config ~/.config/gh-dash/config.yml` (or `prs` if the launcher points to the fork binary).
