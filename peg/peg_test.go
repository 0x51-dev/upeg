package peg_test

import (
	_ "embed"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"github.com/0x51-dev/upeg/peg"
	"testing"
)

var (
	//go:embed definition.peg
	definitionSpec string
)

func TestSpacing(t *testing.T) {
	for _, test := range []string{
		" ",
		"\t",
		"\r",
		"\n",
		" \t\r\n",
		"# Comment\n\n",
	} {
		p, err := parser.New([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{peg.Spacing, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSpecifications(t *testing.T) {
	for _, spec := range []string{definitionSpec} {
		p, err := parser.New([]rune(spec))
		if err != nil {
			t.Fatal(err)
		}
		p.Rules["Expression"] = peg.Expression
		if _, err := p.Parse(op.And{peg.Grammar, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}
