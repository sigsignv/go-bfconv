package tojsonfeed

// See also: https://jsonfeed.org/version/1.1/
type Feed struct {
	Version     string `json:"version"`
	Title       string `json:"title"`
	HomePageURL string `json:"home_page_url,omitempty"`
	FeedURL     string `json:"feed_url,omitempty"`
	Description string `json:"description,omitempty"`
	Items       []Item `json:"items"`
}

type Item struct {
	ID            string             `json:"id"`
	URL           string             `json:"url,omitempty"`
	Title         string             `json:"title,omitempty"`
	ContentText   string             `json:"content_text,omitempty"`
	Image         string             `json:"image,omitempty"`
	DatePublished string             `json:"date_published,omitempty"`
	Tags          []string           `json:"tags,omitempty"`
	BookmarkExt   *BookmarkExtension `json:"_bookmark,omitempty"`
}

// BookmarkExtension holds Hatena Bookmark specific extension data.
type BookmarkExtension struct {
	Count              int    `json:"count,omitempty"`
	CommentListPageURL string `json:"comment_list_page_url,omitempty"`
	SiteEntriesListURL string `json:"site_entries_list_url,omitempty"`
}
