package bfconv

import (
	"fmt"
	"io"

	"github.com/mmcdole/gofeed/rss"
)

type Parser struct{}

func (p *Parser) Parse(data io.Reader) (*rss.Feed, error) {
	r := &rss.Parser{}

	rss, err := r.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS: %w", err)
	}

	return rss, nil
}
