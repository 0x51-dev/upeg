package ebnf

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
)

var (
	Grammar = op.Capture{
		Name:  "Grammar",
		Value: op.OneOrMore{Value: Definition},
	}
	Comment = op.Capture{
		Name: "Comment",
		Value: op.Or{
			op.Ignore{Value: op.And{"(*", op.ZeroOrMore{Value: op.AnyBut{Value: "*)"}}, "*)"}},
			op.Ignore{Value: op.And{"/*", op.ZeroOrMore{Value: op.AnyBut{Value: "*/"}}, "*/"}},
			op.Ignore{Value: op.And{"//", op.ZeroOrMore{Value: op.AnyBut{Value: op.EndOfLine{}}}, op.EndOfLine{}}},
		},
	}
	Definition = op.Capture{
		Name: "Definition",
		Value: op.And{
			Identifier,
			'=',
			op.Or{
				Alternation,
				Empty,
			},
			op.Or{';', '.'},
		},
	}
	Letter = op.Or{
		op.RuneRange{Min: 'a', Max: 'z'},
		op.RuneRange{Min: 'A', Max: 'Z'},
	}
	Digit  = op.RuneRange{Min: '0', Max: '9'}
	String = op.Capture{
		Name: "String",
		Value: op.Or{
			op.Ignore{Value: op.And{'\'', op.ZeroOrMore{Value: op.AnyBut{Value: '\''}}, '\''}},
			op.Ignore{Value: op.And{'"', op.ZeroOrMore{Value: op.AnyBut{Value: '"'}}, '"'}},
			op.Ignore{Value: op.And{'’', op.ZeroOrMore{Value: op.AnyBut{Value: '’'}}, '’'}},
		},
	}
	Identifier = op.Capture{
		Name: "Identifier",
		Value: op.Ignore{Value: op.And{
			Letter,
			op.ZeroOrMore{Value: op.Or{
				Letter, Digit, "_",
				op.And{" ", op.Peek{Value: op.Or{Letter, Digit, "_"}}},
			}},
		}},
	}
	Alternation = op.Capture{
		Name:  "Alternation",
		Value: op.And{Concatenation, op.ZeroOrMore{Value: op.And{op.Or{'|', '/', '!'}, Concatenation}}},
	}
	Concatenation = op.Capture{
		Name:  "Concatenation",
		Value: op.And{Factor, op.ZeroOrMore{Value: op.And{',', Factor}}},
	}
	Optional = op.Capture{
		Name: "Optional",
		Value: op.Or{
			op.And{'[', op.Reference{Name: "Alternation"}, ']'},
			op.And{"(/", op.Reference{Name: "Alternation"}, "/)"},
		},
	}
	Repetition = op.Capture{
		Name: "Repetition",
		Value: op.Or{
			op.And{'{', op.Reference{Name: "Alternation"}, '}'},
			op.And{"(:", op.Reference{Name: "Alternation"}, ":)"},
		},
	}
	Grouping = op.And{'(', op.Reference{Name: "Alternation"}, ')'}
	Factor   = op.Capture{
		Name: "Factor",
		Value: op.And{
			op.Optional{Value: op.And{
				op.Capture{
					Name:  "Amount",
					Value: op.And{op.RuneRange{Min: '1', Max: '9'}, op.ZeroOrMore{Value: Digit}},
				},
				'*',
			}},
			Term,
			op.Optional{Value: op.Or{
				op.Capture{
					Name:  "Optional",
					Value: '?',
				},
				op.Capture{
					Name:  "ZeroOrMore",
					Value: '*',
				},
				op.Capture{
					Name:  "Except",
					Value: op.And{'-', op.Optional{Value: Term}},
				},
			}},
		},
	}
	Term = op.Capture{
		Name: "Term",
		Value: op.Or{
			Identifier,
			String,
			Grouping,
			Repetition,
			Optional,
			SpecialSequence,
		},
	}
	Empty           = op.Capture{Name: "Empty", Value: ""}
	SpecialSequence = op.And{'?', op.Capture{
		Name:  "SpecialSequence",
		Value: op.ZeroOrMore{Value: op.AnyBut{Value: '?'}},
	}, '?'}
)

func NewParser(input []rune) (*parser.Parser, error) {
	p, err := parser.New(input)
	if err != nil {
		return nil, err
	}
	p.Rules["Alternation"] = Alternation
	p.SetIgnoreList([]any{' ', '\t', op.EndOfLine{}, Comment})
	return p, nil
}
