package op_test

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

var AndTestCases = []AndTestCase{
	{"abc", op.And{'a', "bc"}},
	{"abc", op.And{op.Any{}, "bc"}},
	{"abc", op.And{op.Any{}, op.Any{}, 'c'}},
	{"abc", op.And{op.Or{'a', 'b'}, 'b', 'c'}},
	{"abc", op.And{op.Or{'b', 'a'}, 'b', 'c'}},
	{"abbbc", op.And{'a', op.OneOrMore{Value: op.And{op.Not{Value: op.And{'a', 'c'}}, 'b'}}, 'c'}},
}

func TestAnd(t *testing.T) {
	t.Run("Match", func(t *testing.T) {
		for _, test := range AndTestCases {
			p, err := parser.New([]rune(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Match(append(test.consumer, op.EOF{})); err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Parse", func(t *testing.T) {
		for _, test := range AndTestCases {
			p, err := parser.New([]rune(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Parse(append(test.consumer, op.EOF{})); err != nil {
				t.Fatal(err)
			}
		}
	})
}

type AndTestCase struct {
	input    string
	consumer op.And
}
