package bfconv

import (
	"fmt"
	"strconv"
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
		item.Image = t.extractExtensionValue(r, "imageurl")
		item.DatePublished = t.translateItemDatePublished(r)
		item.Tags = t.translateItemTags(r)

		ext := &BookmarkExtension{}
		ext.Count = t.translateBookmarkCount(r)
		ext.CommentListPageURL = t.extractExtensionValue(r, "bookmarkCommentListPageUrl")
		ext.SiteEntriesListURL = t.extractExtensionValue(r, "bookmarkSiteEntriesListUrl")

		item.BookmarkExt = ext

		items = append(items, item)
	}
	feed.Items = items

	return &feed, nil
}

func (t *Translator) translateItemDatePublished(item *rss.Item) string {
	if item == nil || item.DublinCoreExt == nil || len(item.DublinCoreExt.Date) == 0 {
		return ""
	}

	date := item.DublinCoreExt.Date[0]
	if _, err := time.Parse(time.RFC3339, date); err == nil {
		return date
	}

	return ""
}

func (t *Translator) translateItemTags(item *rss.Item) []string {
	if item == nil {
		return []string{}
	}

	elems := item.Extensions["dc"]
	if len(elems) == 0 {
		return []string{}
	}

	values := elems["subject"]
	if len(values) == 0 {
		return []string{}
	}

	tags := make([]string, 0, len(values))
	for _, v := range values {
		tags = append(tags, v.Value)
	}

	return tags
}

func (t *Translator) translateBookmarkCount(item *rss.Item) int {
	c := t.extractExtensionValue(item, "bookmarkcount")
	n, err := strconv.Atoi(c)
	if err != nil {
		return 0
	}

	return n
}

func (t *Translator) extractExtensionValue(item *rss.Item, name string) string {
	if item == nil {
		return ""
	}

	elems := item.Extensions["hatena"]
	if len(elems) == 0 {
		return ""
	}

	values := elems[name]
	if len(values) == 0 {
		return ""
	}

	return values[0].Value
}
