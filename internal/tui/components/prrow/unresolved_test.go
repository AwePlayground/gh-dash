package prrow

import (
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/dlvhdr/gh-dash/v4/internal/data"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/context"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/theme"
)

func TestRenderUnresolvedThreads(t *testing.T) {
	pr := &PullRequest{Ctx: &context.ProgramContext{Theme: *theme.DefaultTheme}, Data: &Data{Primary: &data.PullRequestData{ReviewThreads: data.ReviewThreads{Nodes: []data.ReviewThreadStatus{{IsResolved: true}, {IsResolved: false}}}}}}
	if got := ansi.Strip(pr.renderUnresolvedThreads()); got != "1" {
		t.Fatalf("got %q; want 1", got)
	}
	pr.Data.Primary.ReviewThreads.PageInfo.HasNextPage = true
	if got := ansi.Strip(pr.renderUnresolvedThreads()); got != "1+" {
		t.Fatalf("got %q; want 1+", got)
	}
}
