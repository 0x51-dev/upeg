package ebnf_test

import (
	_ "embed"
	"github.com/0x51-dev/upeg/ebnf"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

var (
	//go:embed definition.ebnf
	definitionSpec string
)

func TestSpecifications(t *testing.T) {
	for _, spec := range []string{definitionSpec} {
		p, err := parser.New([]rune(spec))
		if err != nil {
			t.Fatal(err)
		}
		p.Rules["Alternation"] = ebnf.Alternation
		if _, err := p.Parse(op.And{ebnf.Grammar, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}