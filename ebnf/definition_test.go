package ebnf

import (
	_ "embed"
	"github.com/0x51-dev/upeg/ebnf/ir"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

var (
	//go:embed definition.ebnf
	DefinitionEBNF string

	//go:embed syntax.ebnf
	SyntaxEBNF string

	//go:embed table2.ebnf
	Table2EBNF string
)

func TestAlternation(t *testing.T) {
	for _, test := range []string{
		"Alt=SP|SP;",
		"Alt = SP\n\t| SP\n\t;",
		"Alt = SP | SP ;",
		"Alt = SP | SP | SP ;",
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{Definition, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestComment(t *testing.T) {
	for _, comment := range []string{
		"(**)",
		"(* Comment *)",
		"(* Comment\n * Line2\n *)",
		"/**/",
		"/* Comment */",
		"/* Comment\n * Line2\n */",
		"//\n",
		"// Comment\n",
	} {
		p, err := NewParser([]rune(comment))
		p.SetIgnoreList([]any{' ', op.EndOfLine{}})
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{Comment, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestComplex(t *testing.T) {
	for _, test := range []string{
		"X=('');",
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{Grammar, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestConcatenation(t *testing.T) {
	for _, test := range []string{
		"Concat=SP;",
		"Concat = SP ;",
		"Concat = SP, SP;",
		"Concat = SP , SP , SP ;",
		"Concat = (A,B,C);",
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{Definition, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestDefinition(t *testing.T) {
	for _, test := range []string{
		"Def=SP|SP|SP,SP|SP,SP,SP|SP;",
		"Def = SP | SP , SP , SP | SP ;",
		"D e f\n\t= SP ;",
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		p.SetIgnoreList([]any{' ', '\t', '\n', '\r'})
		if _, err := p.Parse(op.And{Definition, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestEBNF(t *testing.T) {
	for _, test := range []string{
		DefinitionEBNF,
		SyntaxEBNF,
		Table2EBNF,
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		n, err := p.Parse(op.And{Grammar, op.EOF{}})
		if err != nil {
			t.Fatal(err)
		}
		if _, err := ir.ParseGrammar(n); err != nil {
			t.Fatal(err)
		}
	}
}

func TestGrammar(t *testing.T) {
	for _, test := range []string{
		"(* Comment *)DEF=SP;",
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{Grammar, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestIdentifier(t *testing.T) {
	for _, test := range []string{
		"Id",
		"I D",
		"Id1",
		"Id_1",
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{Identifier, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestOptional(t *testing.T) {
	for _, test := range []string{
		"Opt=[SP];",
		"Opt = [ SP ] ;",
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{Definition, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestString(t *testing.T) {
	for _, test := range []string{
		"''",
		"\"\"",
		"'test'",
		"\"test\"",
		"'\\'",
	} {
		p, err := NewParser([]rune(test))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{String, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}
