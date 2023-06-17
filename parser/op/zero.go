package op

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
)

// ZeroOrMore matches the given expression zero or more times.
type ZeroOrMore struct {
	Value any
}

func (zero ZeroOrMore) Match(start parser.Cursor, p *parser.Parser) (parser.Cursor, error) {
	end := start
	for {
		var err error
		if end, err = p.Match(zero.Value); err != nil {
			break
		}
	}
	return end, nil
}

func (zero ZeroOrMore) Parse(p *parser.Parser) (*parser.Node, error) {
	var nodes []*parser.Node
	for {
		node, err := p.Parse(zero.Value)
		if err != nil {
			break
		}
		if node != nil {
			nodes = append(nodes, node)
		}
	}
	return nil, nil
}

func (zero ZeroOrMore) String() string {
	return fmt.Sprintf("%v*", StringAny(zero.Value))
}
