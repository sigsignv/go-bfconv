package tojsonfeed

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestParser_Parse(t *testing.T) {
	file, err := os.Open("testdata/sample.rss")
	if err != nil {
		t.Fatalf("failed to open sample.rss: %v", err)
	}
	defer file.Close()

	p := NewParser()
	r, err := p.Parse(file)

	if err != nil {
		t.Fatalf("Parser.Parse failed: %v", err)
	}

	if !strings.HasPrefix(r.Link, "https://b.hatena.ne.jp/") {
		t.Errorf("feed link mismatch: got %q", r.Link)
	}

	if len(r.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(r.Items))
	}
}

func TestParser_Parse_Invalid(t *testing.T) {
	p := NewParser()
	_, err := p.Parse(bytes.NewReader([]byte("invalid feed data")))
	if err == nil {
		t.Fatalf("expected error for invalid RSS, got nil")
	}
}

func TestParser_Parse_RealFeeds(t *testing.T) {
	urls := []string{
		"https://b.hatena.ne.jp/entrylist.rss",
		"https://b.hatena.ne.jp/sigsign/rss",
	}

	// This test is skipped when running `go test -short`
	if testing.Short() {
		t.Skip("skipping live feed test in short mode")
	}

	client := &http.Client{Timeout: 15 * time.Second}
	p := NewParser()

	for i := range urls {
		u := urls[i]
		t.Run(u, func(t *testing.T) {
			resp, err := client.Get(u)
			if err != nil {
				// Skip on network errors so CI without network doesn't fail the test suite.
				t.Skipf("skipping fetch %s due to network error: %v", u, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Skipf("GET %s returned status %d", u, resp.StatusCode)
			}

			rss, err := p.Parse(resp.Body)
			if err != nil {
				t.Fatalf("Parse failed for %s: %v", u, err)
			}

			if !strings.HasPrefix(rss.Link, "https://b.hatena.ne.jp/") {
				t.Errorf("feed link mismatch for %s: got %q", u, rss.Link)
			}
		})
	}
}
