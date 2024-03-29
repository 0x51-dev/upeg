package op

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
)

type Peek struct {
	Value any
}

func (n Peek) Match(start parser.Cursor, p *parser.Parser) (parser.Cursor, error) {
	if _, err := p.Match(n.Value); err != nil {
		return start, err
	}
	p.Reader.Jump(start)
	return start, nil
}

func (n Peek) String() string {
	return fmt.Sprintf("&%v", StringAny(n.Value))
}
