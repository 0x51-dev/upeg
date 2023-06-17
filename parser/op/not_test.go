package op_test

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

var NotTestCases = []NotTestCase{
	{"a", op.Not{Value: "b"}},
	{"a", op.Not{Value: "ab"}},
}

func TestNot(t *testing.T) {
	t.Run("Match", func(t *testing.T) {
		for _, test := range NotTestCases {
			p, err := parser.New([]rune(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Match(op.And{test.consumer, test.input, op.EOF{}}); err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Parse", func(t *testing.T) {
		for _, test := range NotTestCases {
			p, err := parser.New([]rune(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Parse(op.And{test.consumer, test.input, op.EOF{}}); err != nil {
				t.Fatal(err)
			}
		}
	})
}

type NotTestCase struct {
	input    string
	consumer op.Not
}
