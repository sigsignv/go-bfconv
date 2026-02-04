package bfconv

import (
	"fmt"
	"io"

	"github.com/mmcdole/gofeed/rss"
)

type Parser struct{}

func (p *Parser) Parse(data io.Reader) (*rss.Feed, error) {
	rssParser := &rss.Parser{}
	r, err := rssParser.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS: %w", err)
	}

	return r, nil
}
