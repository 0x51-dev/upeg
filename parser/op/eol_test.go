package op_test

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

var EOLTestCases = []EOLTestCase{
	{"\n"},
	{"\r"},
	{"\n\r"},
}

func TestEOL(t *testing.T) {
	eol := op.And{op.EndOfLine{}, op.EOF{}} // op.EOL{}
	t.Run("Match", func(t *testing.T) {
		for _, test := range EOLTestCases {
			p, err := parser.New([]rune(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Match(eol); err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Parse", func(t *testing.T) {
		for _, test := range EOLTestCases {
			p, err := parser.New([]rune(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Parse(eol); err != nil {
				t.Fatal(err)
			}
		}
	})
}

type EOLTestCase struct {
	input string
}
