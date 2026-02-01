package tojsonfeed

import "io"

type Converter struct {
	Parser     *Parser
	Translator *Translator
}

func NewConverter() *Converter {
	return &Converter{
		Parser:     &Parser{},
		Translator: &Translator{},
	}
}

func (c *Converter) Convert(data io.Reader) (*Feed, error) {
	feed, err := c.Parser.Parse(data)
	if err != nil {
		return nil, err
	}

	return c.Translator.Translate(feed)
}
