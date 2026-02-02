package bfconv

import (
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	p := &Parser{}

	t.Run("parses minimal RSS", func(t *testing.T) {
		file, err := os.Open("testdata/minimal.rss")
		if err != nil {
			t.Fatalf("failed to open minimal.rss: %v", err)
		}
		defer file.Close()

		rss, err := p.Parse(file)
		if err != nil {
			t.Fatalf("Parse failed: %v", err)
		}

		if rss.Link != "https://example.com/alice/bookmark" {
			t.Errorf("feed link mismatch: got %q", rss.Link)
		}
	})

	t.Run("parses sample RSS", func(t *testing.T) {
		file, err := os.Open("testdata/sample.rss")
		if err != nil {
			t.Fatalf("failed to open sample.rss: %v", err)
		}
		defer file.Close()

		rss, err := p.Parse(file)
		if err != nil {
			t.Fatalf("Parse failed: %v", err)
		}

		if !strings.HasPrefix(rss.Link, "https://b.hatena.ne.jp/") {
			t.Errorf("feed link mismatch: got %q", rss.Link)
		}

		if len(rss.Items) != 3 {
			t.Errorf("expected 3 items, got %d", len(rss.Items))
		}
	})

	t.Run("fails on invalid RSS", func(t *testing.T) {
		file, err := os.Open("testdata/invalid.rss")
		if err != nil {
			t.Fatalf("failed to open invalid.rss: %v", err)
		}
		defer file.Close()

		_, err = p.Parse(file)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("parses live RSS feeds", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping live RSS feed tests")
		}

		urls := []string{
			"https://b.hatena.ne.jp/entrylist.rss",
			"https://b.hatena.ne.jp/sigsign/rss",
		}

		client := &http.Client{Timeout: 15 * time.Second}

		for i := range urls {
			u := urls[i]
			t.Run(u, func(t *testing.T) {
				resp, err := client.Get(u)
				if err != nil {
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
	})
}
