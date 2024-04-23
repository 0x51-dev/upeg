package op_test

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
	"unicode/utf8"
)

func Test0xFFFD(t *testing.T) {
	for _, r := range []rune{utf8.RuneError, 0x10FFFF} {
		p, err := parser.New([]rune{r})
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.ParseEOF(op.RuneRange{Min: 0x7D, Max: 0x10FFFF}); err != nil {
			t.Fatal(err)
		}
		if _, err := p.ParseEOF(op.ZeroOrMore{Value: op.RuneRange{Min: 0x7D, Max: 0x10FFFF}}); err != nil {
			t.Fatal(err)
		}
	}
}
