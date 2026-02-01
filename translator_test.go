package tojsonfeed

import (
	"os"
	"testing"

	"github.com/mmcdole/gofeed/rss"
)

func TestTranslate(t *testing.T) {
	tr := &Translator{}

	t.Run("translate sample RSS feed", func(t *testing.T) {
		file, err := os.Open("testdata/sample.rss")
		if err != nil {
			t.Skipf("failed to open sample.rss: %v", err)
		}
		defer file.Close()

		p := NewParser()
		r, err := p.Parse(file)
		if err != nil {
			t.Skipf("Parse failed: %v", err)
		}

		feed, err := tr.Translate(r)
		if err != nil {
			t.Fatalf("Translate failed: %v", err)
		}

		if feed.Title != "Sample Hatena Bookmark Feed" {
			t.Errorf("feed title mismatch: got %s", feed.Title)
		}

		if feed.HomePageURL != "https://b.hatena.ne.jp/entrylist/all" {
			t.Errorf("feed home_page_url mismatch: got %s", feed.HomePageURL)
		}

		if feed.Description != "Sample entries for testing" {
			t.Errorf("feed description mismatch: got %s", feed.Description)
		}

		if len(feed.Items) != 3 {
			t.Fatalf("items: expected 3, got %d", len(feed.Items))
		}

		item0 := feed.Items[0]
		if item0.Title != "Example Article Title One" {
			t.Errorf("item[0].title mismatch: got %s", item0.Title)
		}
		if item0.URL != "https://example.com/article/1" {
			t.Errorf("item[0].url mismatch: got %s", item0.URL)
		}
		if item0.ContentText != "This is the summary of the first article about technology." {
			t.Errorf("item[0].content_text mismatch: got %s", item0.ContentText)
		}
		if item0.Image != "https://example.com/image1.jpg" {
			t.Errorf("item[0].image mismatch: got %s", item0.Image)
		}
		if item0.DatePublished != "2026-01-30T10:15:00Z" {
			t.Errorf("item[0].date_published mismatch: got %s", item0.DatePublished)
		}
		if len(item0.Tags) != 3 || item0.Tags[0] != "Technology" || item0.Tags[1] != "Go" || item0.Tags[2] != "Programming" {
			t.Errorf("item[0].tags mismatch: got %v", item0.Tags)
		}
		if item0.BookmarkExt.Count != 42 {
			t.Errorf("item[0].BookmarkExt.Count mismatch: got %d", item0.BookmarkExt.Count)
		}
		if item0.BookmarkExt.CommentListPageURL != "https://b.hatena.ne.jp/entry/s/example.com/article/1" {
			t.Errorf("item[0].BookmarkExt.CommentPageURL mismatch: got %s", item0.BookmarkExt.CommentListPageURL)
		}
		if item0.BookmarkExt.SiteEntriesListURL != "https://b.hatena.ne.jp/site/example.com/" {
			t.Errorf("item[0].BookmarkExt.SiteEntriesURL mismatch: got %s", item0.BookmarkExt.SiteEntriesListURL)
		}

		item1 := feed.Items[1]
		if item1.Title != "Example Article Title Two" {
			t.Errorf("item[1].title mismatch: got %s", item1.Title)
		}
		if item1.URL != "https://example.com/article/2" {
			t.Errorf("item[1].url mismatch: got %s", item1.URL)
		}
		if item1.ContentText != "" {
			t.Errorf("item[1].content_text mismatch: got %s", item1.ContentText)
		}
		if item1.Image != "https://example.com/image2.jpg" {
			t.Errorf("item[1].image mismatch: got %s", item1.Image)
		}
		if item1.DatePublished != "2026-01-29T15:30:00Z" {
			t.Errorf("item[1].date_published mismatch: got %s", item1.DatePublished)
		}
		if len(item1.Tags) != 0 {
			t.Errorf("item[1].tags mismatch: got %v", item1.Tags)
		}
		if item1.BookmarkExt.Count != 15 {
			t.Errorf("item[1].BookmarkExt.Count mismatch: got %d", item1.BookmarkExt.Count)
		}
		if item1.BookmarkExt.CommentListPageURL != "https://b.hatena.ne.jp/entry/s/example.com/article/2" {
			t.Errorf("item[1].BookmarkExt.CommentPageURL mismatch: got %s", item1.BookmarkExt.CommentListPageURL)
		}
		if item1.BookmarkExt.SiteEntriesListURL != "https://b.hatena.ne.jp/site/example.com/" {
			t.Errorf("item[1].BookmarkExt.SiteEntriesURL mismatch: got %s", item1.BookmarkExt.SiteEntriesListURL)
		}

		item2 := feed.Items[2]
		if item2.Title != "Article Without Image" {
			t.Errorf("item[2].title mismatch: got %s", item2.Title)
		}
		if item2.URL != "https://example.com/article/3" {
			t.Errorf("item[2].url mismatch: got %s", item2.URL)
		}
		if item2.ContentText != "This article has no image URL." {
			t.Errorf("item[2].content_text mismatch: got %s", item2.ContentText)
		}
		if item2.Image != "" {
			t.Errorf("item[2].image mismatch: got %s", item2.Image)
		}
		if item2.DatePublished != "2026-01-28T08:00:00Z" {
			t.Errorf("item[2].date_published mismatch: got %s", item2.DatePublished)
		}
		if len(item2.Tags) != 1 || item2.Tags[0] != "Test" {
			t.Errorf("item[2].tags mismatch: got %v", item2.Tags)
		}
		if item2.BookmarkExt.Count != 3 {
			t.Errorf("item[2].BookmarkExt.Count mismatch: got %d", item2.BookmarkExt.Count)
		}
		if item2.BookmarkExt.CommentListPageURL != "https://b.hatena.ne.jp/entry/s/example.com/article/3" {
			t.Errorf("item[2].BookmarkExt.CommentPageURL mismatch: got %s", item2.BookmarkExt.CommentListPageURL)
		}
		if item2.BookmarkExt.SiteEntriesListURL != "https://b.hatena.ne.jp/site/example.com/" {
			t.Errorf("item[2].BookmarkExt.SiteEntriesURL mismatch: got %s", item2.BookmarkExt.SiteEntriesListURL)
		}
	})

	t.Run("should error if feed is nil", func(t *testing.T) {
		if _, err := tr.Translate(nil); err == nil {
			t.Fatalf("expected error for nil feed")
		}
	})

	t.Run("should skip if item is nil", func(t *testing.T) {
		rss := &rss.Feed{
			Title: "Test Feed",
			Items: []*rss.Item{nil},
		}
		feed, err := tr.Translate(rss)
		if err != nil {
			t.Fatalf("Translate failed: %v", err)
		}
		if len(feed.Items) != 0 {
			t.Fatalf("expected 0 items, got %d", len(feed.Items))
		}
	})
}
