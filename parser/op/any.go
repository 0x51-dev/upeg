package op

import (
	"github.com/0x51-dev/upeg/parser"
)

// Any matches any character.
type Any struct{}

func (a Any) Match(start parser.Cursor, p *parser.Parser) (parser.Cursor, error) {
	if p.Reader.Done() {
		return start, p.NewNoMatchError(a, start)
	}
	return p.Reader.Next().Cursor(), nil
}

func (a Any) String() string {
	return "."
}
