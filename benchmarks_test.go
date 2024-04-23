package upeg

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"strings"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	in := strings.Repeat("a;", 1000)
	p, err := parser.New([]rune(in[:len(in)-1]))
	if err != nil {
		b.Fatal(err)
	}
	b.Run("ZeroOrMore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := p.ParseEOF(op.And{'a', op.ZeroOrMore{Value: op.And{';', 'a'}}}); err != nil {
				b.Fatal(err)
			}
			p.Reset()
		}
	})
	b.Run("Or", func(b *testing.B) {
		as := op.Or{
			op.And{'a', ';', op.Reference{Name: "as"}},
			'a',
		}
		p.Rules["as"] = as
		for i := 0; i < b.N; i++ {
			if _, err := p.ParseEOF(as); err != nil {
				b.Fatal(err)
			}
			p.Reset()
		}
	})
}
