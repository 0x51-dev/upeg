package ir_test

import (
	"testing"

	"github.com/0x51-dev/upeg/abnf"
	"github.com/0x51-dev/upeg/bnf"
	"github.com/0x51-dev/upeg/ir"
	"github.com/0x51-dev/upeg/parser"
)

func TestBNF_Rulelist(t *testing.T) {
	for _, test := range []struct {
		raw      string
		expected string
	}{
		{
			raw:      "<X> ::= <Y>",
			expected: "X = Y",
		},
		{
			raw:      "<X> ::= <Y> <Z>",
			expected: "X = ( Y Z )",
		},
		{
			raw:      "<X> ::= <Y> | <Z>",
			expected: "X = ( Y / Z )",
		},
		{
			raw:      "<X> ::= <Y>+",
			expected: "X = 1*Y",
		},
		{
			raw:      "<X> ::= <Y>*",
			expected: "X = *Y",
		},
	} {
		p, err := parser.New([]rune(test.raw + "\n"))
		if err != nil {
			t.Fatal(err)
		}
		n, err := p.ParseEOF(bnf.Rulelist)
		if err != nil {
			t.Fatal(err)
		}
		l, err := ir.ParseRulelist(n)
		if err != nil {
			t.Fatal(err)
		}
		if l.String() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, l.String())
		}
	}
}

func TestParseAlternation(t *testing.T) {
	for _, test := range []struct {
		raw      string
		expected string
	}{
		{
			raw:      "X",
			expected: "X",
		},
		{
			raw:      "X / Y",
			expected: "( X / Y )",
		},
		{
			raw:      "X / Y / Z",
			expected: "( X / Y / Z )",
		},
	} {
		p, err := abnf.NewParser([]rune(test.raw))
		if err != nil {
			t.Fatal(err)
		}
		n, err := p.Parse(abnf.Alternation)
		if err != nil {
			t.Fatal(err)
		}
		r, err := ir.ParseAlternation(n)
		if err != nil {
			t.Error(err)
		}
		if r.String() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, r.String())
		}
	}
}

func TestParseOption(t *testing.T) {
	for _, test := range []struct {
		raw      string
		expected string
	}{
		{
			raw:      "[ X ]",
			expected: "[ X ]",
		},
		{
			raw:      "[ ( ( X ) ) ]",
			expected: "[ X ]",
		},
	} {
		p, err := abnf.NewParser([]rune(test.raw))
		if err != nil {
			t.Fatal(err)
		}
		n, err := p.ParseEOF(abnf.Option)
		if err != nil {
			t.Fatal(err)
		}
		r, err := ir.ParseOption(n)
		if err != nil {
			t.Error(err)
		}
		if r.String() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, r.String())
		}
	}
}

func TestParseRepetition(t *testing.T) {
	for _, test := range []struct {
		raw      string
		expected string
	}{
		{
			raw:      "X",
			expected: "X",
		},
		{
			raw:      "*X",
			expected: "*X",
		},
		{
			raw:      "1*X",
			expected: "1*X",
		},
		{
			raw:      "*1X",
			expected: "*1X",
		},
		{
			raw:      "1*1X",
			expected: "1X",
		},
		{
			raw:      "1X",
			expected: "1X",
		},
		{
			raw:      "( X )",
			expected: "X",
		},
		{
			raw:      "[ X ]",
			expected: "[ X ]",
		},
		{
			raw:      "%x00",
			expected: "x00",
		},
		{
			raw:      "\"X\"",
			expected: "\"X\"",
		},
		{
			raw:      "<X>",
			expected: "<X>",
		},
	} {
		p, err := abnf.NewParser([]rune(test.raw))
		if err != nil {
			t.Fatal(err)
		}
		n, err := p.ParseEOF(abnf.Repetition)
		if err != nil {
			t.Fatal(err)
		}
		r, err := ir.ParseRepetition(n)
		if err != nil {
			t.Error(err)
		}
		if r.String() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, r.String())
		}
	}
}

func TestParseRulename(t *testing.T) {
	rulename := "Rulename"
	p, err := abnf.NewParser([]rune(rulename))
	if err != nil {
		t.Fatal(err)
	}
	n, err := p.ParseEOF(abnf.Rulename)
	if err != nil {
		t.Fatal(err)
	}
	name, err := ir.ParseRulename(n)
	if err != nil {
		t.Error(err)
	}
	if name != rulename {
		t.Errorf("expected %s, got %s", rulename, name)
	}

	// invalid
	if _, err := ir.ParseRulename(nil); err == nil {
		t.Error("expected error")
	}
	if _, err := ir.ParseRulename(parser.NewNode("invalid", "")); err == nil {
		t.Error("expected error")
	}
}
