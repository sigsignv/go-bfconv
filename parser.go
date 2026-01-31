package tojsonfeed

import (
	"fmt"
	"io"

	"github.com/mmcdole/gofeed/rss"
)

type Parser struct {
	RSSParser *rss.Parser
}

func NewParser() *Parser {
	return &Parser{
		RSSParser: &rss.Parser{},
	}
}

func (p *Parser) Parse(feed io.Reader) (*rss.Feed, error) {
	r, err := p.RSSParser.Parse(feed)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS: %w", err)
	}

	return r, nil
}
