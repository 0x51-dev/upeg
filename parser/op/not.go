package op

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
)

// Not matches if the given expression does not match.
type Not struct {
	Value any
}

func (n Not) Match(start parser.Cursor, p *parser.Parser) (parser.Cursor, error) {
	if end, err := p.Match(n.Value); err == nil {
		p.Reader.Jump(start)
		return start, parser.NewErrorStack(p.NewNoMatchError(n, end), err)
	}
	return start, nil
}

func (n Not) String() string {
	return fmt.Sprintf("!%v", StringAny(n.Value))
}
