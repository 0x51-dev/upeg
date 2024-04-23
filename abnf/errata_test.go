package abnf_test

import (
	"github.com/0x51-dev/upeg/abnf/core"
	"github.com/0x51-dev/upeg/parser"
	"testing"
)

func TestErrata(t *testing.T) {
	t.Run("CRLF", func(t *testing.T) {
		for _, s := range []string{"\r\n", "\n"} {
			p, err := parser.New([]rune(s))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.ParseEOF(core.CRLF); err != nil {
				t.Fatal(err)
			}
		}
	})
}
