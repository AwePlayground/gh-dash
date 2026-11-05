package tui

import (
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/dlvhdr/gh-dash/v4/internal/config"
	"github.com/dlvhdr/gh-dash/v4/internal/data"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/branchsidebar"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/footer"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/issueview"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/notificationview"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/prrow"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/prssection"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/prview"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/section"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/sidebar"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/components/tabs"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/context"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/keys"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/theme"
)

func TestEnterFocusesPRPreviewAndArrowKeysScrollIt(t *testing.T) {
	cfg, err := config.ParseConfig(config.Location{ConfigFlag: "../config/testdata/test-config.yml", SkipGlobalConfig: true})
	if err != nil {
		t.Fatal(err)
	}
	ctx := &context.ProgramContext{Config: &cfg, View: config.PRsView, PreviewPosition: "bottom", DynamicPreviewWidth: 80, DynamicPreviewHeight: 8, MainContentHeight: 8}
	ctx.Theme = theme.ParseTheme(ctx.Config)
	ctx.Styles = context.InitStyles(ctx.Theme)
	sec := prssection.NewModel(0, ctx, config.PrsSectionConfig{}, time.Now(), time.Now())
	sec.Prs = []prrow.Data{{Primary: &data.PullRequestData{Title: "first"}}, {Primary: &data.PullRequestData{Title: "second"}}, {Primary: &data.PullRequestData{Title: "third"}}}
	sec.Table.SetRows(sec.BuildRows())
	preview := sidebar.NewModel()
	preview.IsOpen = true
	preview.UpdateProgramContext(ctx)
	preview.SetContent(strings.Repeat("line\n", 60))
	m := Model{ctx: ctx, keys: keys.Keys, prs: []section.Section{&sec}, sidebar: preview, prView: prview.NewModel(ctx), issueSidebar: issueview.NewModel(ctx), branchSidebar: branchsidebar.NewModel(ctx), notificationView: notificationview.NewModel(ctx), footer: footer.NewModel(ctx), tabs: tabs.NewModel(ctx)}
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if !m.prPreviewFocused {
		t.Fatal("Enter should focus PR preview")
	}
	row := sec.CurrRow()
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	if sec.CurrRow() != row || m.sidebar.YOffset() == 0 {
		t.Fatalf("down moved PR %d->%d or failed to scroll preview %d", row, sec.CurrRow(), m.sidebar.YOffset())
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	if m.prPreviewFocused {
		t.Fatal("Esc should return to PR list")
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	if sec.CurrRow() == row {
		t.Fatal("down in PR list should select the next PR")
	}
}
