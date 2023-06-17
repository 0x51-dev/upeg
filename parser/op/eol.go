package op

import "github.com/0x51-dev/upeg/parser"

// EndOfLine matches the end of a line.
type EndOfLine struct{}

func (e EndOfLine) Match(_ parser.Cursor, p *parser.Parser) (parser.Cursor, error) {
	return p.Match(Or{"\n\r", '\n', '\r'})
}

func (e EndOfLine) String() string {
	return `\n`
}
