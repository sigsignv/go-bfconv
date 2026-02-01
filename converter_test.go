package tojsonfeed

import (
	"bytes"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	c := NewConverter()

	t.Run("sample rss", func(t *testing.T) {
		file, err := os.Open("testdata/sample.rss")
		if err != nil {
			t.Skipf("failed to open sample.rss: %v", err)
		}
		defer file.Close()

		f, err := c.Convert(file)
		if err != nil {
			t.Fatalf("Convert failed: %v", err)
		}

		if f.Title != "Sample Hatena Bookmark Feed" {
			t.Errorf("feed title mismatch: got %s", f.Title)
		}
		if f.HomePageURL != "https://b.hatena.ne.jp/entrylist/all" {
			t.Errorf("feed home_page_url mismatch: got %s", f.HomePageURL)
		}
		if f.Description != "Sample entries for testing" {
			t.Errorf("feed description mismatch: got %s", f.Description)
		}
		if len(f.Items) != 3 {
			t.Fatalf("items: expected 3, got %d", len(f.Items))
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		if _, err := c.Convert(bytes.NewReader([]byte("invalid RSS data"))); err == nil {
			t.Fatalf("expected error for invalid RSS, got nil")
		}
	})

	t.Run("real feeds", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping live feed test in short mode")
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
					// Skip on network errors so CI without network doesn't fail the test suite.
					t.Skipf("skipping fetch %s due to network error: %v", u, err)
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					t.Skipf("GET %s returned status %d", u, resp.StatusCode)
				}

				f, err := c.Convert(resp.Body)
				if err != nil {
					t.Fatalf("Convert failed for %s: %v", u, err)
				}

				if f.HomePageURL == "" {
					t.Errorf("feed home page url is empty for %s", u)
				}
			})
		}
	})
}
