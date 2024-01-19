// Package abnf is autogenerated by https://github.com/0x51-dev/upeg. DO NOT EDIT.
package abnf

import (
	. "github.com/0x51-dev/upeg/abnf/core"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
)

var (
	Rulelist      = op.Capture{Name: "Rulelist", Value: op.OneOrMore{Value: op.Or{Rule, op.And{op.ZeroOrMore{Value: WSP}, CNl}}}}
	Rule          = op.Capture{Name: "Rule", Value: op.And{Rulename, DefinedAs, Elements, CNl}}
	Rulename      = op.Capture{Name: "Rulename", Value: op.And{ALPHA, op.ZeroOrMore{Value: op.Or{ALPHA, DIGIT, '-'}}}}
	DefinedAs     = op.And{op.ZeroOrMore{Value: CWsp}, op.Or{'=', "=/"}, op.ZeroOrMore{Value: CWsp}}
	Elements      = op.And{Alternation, op.ZeroOrMore{Value: WSP}}
	CWsp          = op.Or{WSP, op.And{CNl, WSP}}
	CNl           = op.Or{Comment, CRLF}
	Comment       = op.Capture{Name: "Comment", Value: op.And{';', op.ZeroOrMore{Value: op.Or{WSP, VCHAR}}, CRLF}}
	Alternation   = op.Capture{Name: "Alternation", Value: op.And{Concatenation, op.ZeroOrMore{Value: op.And{op.ZeroOrMore{Value: CWsp}, '/', op.ZeroOrMore{Value: CWsp}, Concatenation}}}}
	Concatenation = op.Capture{Name: "Concatenation", Value: op.And{Repetition, op.ZeroOrMore{Value: op.And{op.OneOrMore{Value: CWsp}, Repetition}}}}
	Repetition    = op.Capture{Name: "Repetition", Value: op.And{op.Optional{Value: Repeat}, Element}}
	Repeat        = op.Capture{Name: "Repeat", Value: op.Or{op.And{op.ZeroOrMore{Value: DIGIT}, '*', op.ZeroOrMore{Value: DIGIT}}, op.OneOrMore{Value: DIGIT}}}
	Element       = op.Or{Rulename, Group, Option, CharVal, NumVal, ProseVal}
	Group         = op.And{'(', op.ZeroOrMore{Value: CWsp}, op.Reference{Name: "Alternation"}, op.ZeroOrMore{Value: CWsp}, ')'}
	Option        = op.Capture{Name: "Option", Value: op.And{'[', op.ZeroOrMore{Value: CWsp}, op.Reference{Name: "Alternation"}, op.ZeroOrMore{Value: CWsp}, ']'}}
	CharVal       = op.Capture{Name: "CharVal", Value: op.And{DQUOTE, op.ZeroOrMore{Value: op.Or{op.RuneRange{Min: 0x20, Max: 0x21}, op.RuneRange{Min: 0x23, Max: 0x7E}}}, DQUOTE}}
	NumVal        = op.Capture{Name: "NumVal", Value: op.And{'%', op.Or{BinVal, DecVal, HexVal}}}
	BinVal        = op.Capture{Name: "BinVal", Value: op.And{'b', op.OneOrMore{Value: BIT}, op.Optional{Value: op.Or{op.OneOrMore{Value: op.And{'.', op.OneOrMore{Value: BIT}}}, op.And{'-', op.OneOrMore{Value: BIT}}}}}}
	DecVal        = op.Capture{Name: "DecVal", Value: op.And{'d', op.OneOrMore{Value: DIGIT}, op.Optional{Value: op.Or{op.OneOrMore{Value: op.And{'.', op.OneOrMore{Value: DIGIT}}}, op.And{'-', op.OneOrMore{Value: DIGIT}}}}}}
	HexVal        = op.Capture{Name: "HexVal", Value: op.And{'x', op.OneOrMore{Value: HEXDIG}, op.Optional{Value: op.Or{op.OneOrMore{Value: op.And{'.', op.OneOrMore{Value: HEXDIG}}}, op.And{'-', op.OneOrMore{Value: HEXDIG}}}}}}
	ProseVal      = op.Capture{Name: "ProseVal", Value: op.And{'<', op.ZeroOrMore{Value: op.Or{op.RuneRange{Min: 0x20, Max: 0x3D}, op.RuneRange{Min: 0x3F, Max: 0x7E}}}, '>'}}
)

func NewParser(input []rune) (*parser.Parser, error) {
	p, err := parser.New(input)
	if err != nil {
		return nil, err
	}
	p.Rules["Alternation"] = Alternation
	return p, nil
}
