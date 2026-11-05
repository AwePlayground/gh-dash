package data

import "testing"

func TestReviewThreadsUnresolvedCount(t *testing.T) {
	threads := ReviewThreads{Nodes: []ReviewThreadStatus{{IsResolved: false}, {IsResolved: true}, {IsResolved: false}}}
	if n, incomplete := threads.UnresolvedCount(); n != 2 || incomplete {
		t.Fatalf("got %d, %v; want 2, false", n, incomplete)
	}
	threads.PageInfo.HasNextPage = true
	if n, incomplete := threads.UnresolvedCount(); n != 2 || !incomplete {
		t.Fatalf("got %d, %v; want 2, true", n, incomplete)
	}
}
