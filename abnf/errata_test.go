package abnf_test

import (
	"github.com/0x51-dev/upeg/abnf/core"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

func TestErrata(t *testing.T) {
	t.Run("CRLF", func(t *testing.T) {
		for _, s := range []string{"\r\n", "\n"} {
			p, err := parser.New([]rune(s))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Parse(op.And{core.CRLF, op.EOF{}}); err != nil {
				t.Fatal(err)
			}
		}
	})
}
