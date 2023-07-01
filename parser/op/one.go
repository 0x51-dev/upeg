package op

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
)

type OneOrMore struct {
	Value any
}

func (one OneOrMore) Match(start parser.Cursor, p *parser.Parser) (parser.Cursor, error) {
	var err error
	if _, err = p.Match(one.Value); err != nil {
		p.Reader.Jump(start)
		return start, err
	}
	for {
		if _, err = p.Match(one.Value); err != nil {
			break
		}
	}
	return start, nil
}

func (one OneOrMore) Parse(p *parser.Parser) (*parser.Node, error) {
	start := p.Reader.Cursor()
	var nodes []*parser.Node
	node, err := p.Parse(one.Value)
	if err != nil {
		p.Reader.Jump(start)
		return nil, err
	}
	if node != nil {
		nodes = append(nodes, node)
	}
	for {
		node, err := p.Parse(one.Value)
		if err != nil {
			break
		}
		if node != nil {
			nodes = append(nodes, node)
		}
	}
	if len(nodes) == 0 {
		return nil, nil
	}
	if len(nodes) == 1 {
		return nodes[0], nil
	}
	return parser.NewParentNode("", nodes), nil
}

func (one OneOrMore) String() string {
	return fmt.Sprintf("%v+", StringAny(one.Value))
}
