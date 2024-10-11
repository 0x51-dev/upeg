// Package bnf is autogenerated by https://github.com/0x51-dev/upeg. DO NOT EDIT.
package bnf

import (
	. "github.com/0x51-dev/upeg/abnf/core"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
)

var (
	Rulelist      = op.Capture{Name: "Rulelist", Value: op.OneOrMore{Value: op.Or{Rule, op.And{op.ZeroOrMore{Value: WSP}, CNl}}}}
	Rule          = op.Capture{Name: "Rule", Value: op.And{RulenameBr, DefinedAs, Elements, CNl}}
	RulenameBr    = op.And{'<', Rulename, '>'}
	Rulename      = op.Capture{Name: "Rulename", Value: op.And{ALPHA, op.ZeroOrMore{Value: op.Or{ALPHA, DIGIT, '-'}}}}
	DefinedAs     = op.And{op.ZeroOrMore{Value: CWsp}, "::=", op.ZeroOrMore{Value: CWsp}}
	Elements      = op.And{Alternation, op.ZeroOrMore{Value: WSP}}
	CWsp          = op.Or{WSP, op.And{CNl, WSP}}
	CNl           = op.Or{Comment, CRLF}
	Comment       = op.Capture{Name: "Comment", Value: op.And{';', op.ZeroOrMore{Value: op.Or{WSP, VCHAR}}, CRLF}}
	Alternation   = op.Capture{Name: "Alternation", Value: op.And{Concatenation, op.ZeroOrMore{Value: op.And{op.ZeroOrMore{Value: CWsp}, '|', op.ZeroOrMore{Value: CWsp}, Concatenation}}}}
	Concatenation = op.Capture{Name: "Concatenation", Value: op.And{Repetition, op.ZeroOrMore{Value: op.And{op.OneOrMore{Value: CWsp}, Repetition}}}}
	Repetition    = op.Capture{Name: "Repetition", Value: op.And{Element, op.Optional{Value: Repeat}}}
	Repeat        = op.Capture{Name: "Repeat", Value: op.Or{'*', '+'}}
	Element       = op.Or{RulenameBr, Group, Option, CharVal}
	Group         = op.And{'(', op.ZeroOrMore{Value: CWsp}, op.Reference{Name: "Alternation"}, op.ZeroOrMore{Value: CWsp}, ')'}
	Option        = op.Capture{Name: "Option", Value: op.And{'[', op.ZeroOrMore{Value: CWsp}, op.Reference{Name: "Alternation"}, op.ZeroOrMore{Value: CWsp}, ']'}}
	CharVal       = op.Capture{Name: "CharVal", Value: op.Or{LiteralDouble, LiteralSingle}}
	LiteralDouble = op.And{rune(0x22), op.ZeroOrMore{Value: op.Or{op.RuneRange{Min: 0x20, Max: 0x21}, op.RuneRange{Min: 0x23, Max: 0x7E}}}, rune(0x22)}
	LiteralSingle = op.And{rune(0x27), op.ZeroOrMore{Value: op.Or{op.RuneRange{Min: 0x20, Max: 0x26}, op.RuneRange{Min: 0x28, Max: 0x7E}}}, rune(0x27)}
)

func NewParser(input []rune) (*parser.Parser, error) {
	p, err := parser.New(input)
	if err != nil {
		return nil, err
	}
	p.Rules["Alternation"] = Alternation
	return p, nil
}
