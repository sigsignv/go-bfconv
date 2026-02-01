package tojsonfeed

import (
	"fmt"
	"time"

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
		item.Image = "" // Todo: extract image if available
		item.DatePublished = t.translateItemDatePublished(r)
		item.Tags = []string{} // Todo: extract tags if available
		item.BookmarkExt = nil // Todo: extract bookmark extension

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
