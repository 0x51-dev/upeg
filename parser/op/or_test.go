package op_test

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

var OrTestCases = []OrTestCase{
	{"a", op.Or{'a'}},
	{"a", op.Or{'a', 'b'}},
	{"a", op.Or{'b', 'a'}},
	{"a", op.Or{"ab", 'a'}},
	{"ab", op.Or{"ab", 'a'}},
}

func TestOr(t *testing.T) {
	t.Run("Match", func(t *testing.T) {
		for _, test := range OrTestCases {
			p, err := parser.New([]rune(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Match(op.And{test.consumer, op.EOF{}}); err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Parse", func(t *testing.T) {
		for _, test := range OrTestCases {
			p, err := parser.New([]rune(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Parse(op.And{test.consumer, op.EOF{}}); err != nil {
				t.Fatal(err)
			}
		}
	})
}

type OrTestCase struct {
	input    string
	consumer op.Or
}
