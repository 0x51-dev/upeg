package op

import (
	"github.com/0x51-dev/upeg/parser"
)

// EOF matches the end of the input.
type EOF struct{}

func (eof EOF) Match(start parser.Cursor, p *parser.Parser) (parser.Cursor, error) {
	if p.Reader.Done() {
		return start, nil
	}
	c := p.Reader.Cursor()
	return start, p.NewNoMatchError(eof, c)
}

func (eof EOF) String() string {
	return "!."
}
