package tojsonfeed

import (
	"fmt"
	"time"

	ext "github.com/mmcdole/gofeed/extensions"
	"github.com/mmcdole/gofeed/rss"
)

type Translator struct{}

func (t *Translator) Translate(rss *rss.Feed) (*Feed, error) {
	if rss == nil {
		return nil, fmt.Errorf("rss feed is nil")
	}

	feed := Feed{}

	feed.Version = "https://jsonfeed.org/version/1.1"
	feed.Title = rss.Title
	feed.HomePageURL = rss.Link
	feed.Description = rss.Description

	items := make([]Item, 0, len(rss.Items))
	for _, r := range rss.Items {
		if r == nil {
			continue
		}

		item := Item{}

		item.ID = r.Link
		item.URL = r.Link
		item.Title = r.Title
		item.ContentText = r.Description
		item.Image = t.translateItemImage(r)
		item.DatePublished = t.translateItemDatePublished(r)
		item.Tags = t.translateItemTags(r)

		bookmarkExt := &BookmarkExtension{}
		bookmarkExt.Count = t.translateBookmarkCount(r)
		bookmarkExt.CommentListPageURL = t.translateBookmarkCommentListPageURL(r)
		bookmarkExt.SiteEntriesListURL = t.translateBookmarkSiteEntriesListURL(r)

		item.BookmarkExt = bookmarkExt

		items = append(items, item)
	}
	feed.Items = items

	return &feed, nil
}

func (t *Translator) translateItemDatePublished(rssItem *rss.Item) string {
	if rssItem == nil || rssItem.DublinCoreExt == nil || len(rssItem.DublinCoreExt.Date) == 0 {
		return ""
	}

	date := rssItem.DublinCoreExt.Date[0]
	if _, err := time.Parse(time.RFC3339, date); err == nil {
		return date
	}

	return ""
}

func (t *Translator) translateItemImage(rssItem *rss.Item) string {
	imageURLs := t.extractExtension(rssItem, "hatena", "imageurl")
	if len(imageURLs) == 0 {
		return ""
	}

	image := imageURLs[0].Value
	return image
}

func (t *Translator) translateItemTags(rssItem *rss.Item) []string {
	subjects := t.extractExtension(rssItem, "dc", "subject")
	if len(subjects) == 0 {
		return []string{}
	}

	tags := make([]string, 0, len(subjects))
	for _, s := range subjects {
		tags = append(tags, s.Value)
	}

	return tags
}

func (t *Translator) translateBookmarkCount(rssItem *rss.Item) int {
	counts := t.extractExtension(rssItem, "hatena", "bookmarkcount")
	if len(counts) == 0 {
		return 0
	}

	var n int
	fmt.Sscanf(counts[0].Value, "%d", &n)
	return n
}

func (t *Translator) translateBookmarkCommentListPageURL(rssItem *rss.Item) string {
	URLs := t.extractExtension(rssItem, "hatena", "bookmarkCommentListPageUrl")
	if len(URLs) == 0 {
		return ""
	}

	url := URLs[0].Value
	return url
}

func (t *Translator) translateBookmarkSiteEntriesListURL(rssItem *rss.Item) string {
	URLs := t.extractExtension(rssItem, "hatena", "bookmarkSiteEntriesListUrl")
	if len(URLs) == 0 {
		return ""
	}

	url := URLs[0].Value
	return url
}

func (t *Translator) extractExtension(rssItem *rss.Item, ns string, name string) []ext.Extension {
	if rssItem == nil || len(rssItem.Extensions) == 0 {
		return nil
	}

	elms, ok := rssItem.Extensions[ns]
	if !ok || len(elms) == 0 {
		return nil
	}

	items, ok := elms[name]
	if !ok || len(items) == 0 {
		return nil
	}

	return items
}
