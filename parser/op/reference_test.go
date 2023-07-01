package op_test

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

func TestReference(t *testing.T) {
	circular := op.Or{op.And{'0', op.Reference{Name: "circular"}}, "1"}
	t.Run("Parse", func(t *testing.T) {
		for _, test := range []string{
			"01",
			"001",
			"0001",
			// etc.
		} {
			p, err := parser.New([]rune(test))
			if err != nil {
				t.Fatal(err)
			}
			p.Rules["circular"] = circular
			n, err := p.Parse(op.Capture{Value: circular})
			if err != nil {
				t.Fatal(err)
			}
			if n.Value() != test {
				t.Errorf("expected %s, got %s", test, n.Value())
			}
		}
	})
}
