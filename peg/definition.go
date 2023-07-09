package peg

import "github.com/0x51-dev/upeg/parser/op"

var (
	Grammar = op.And{
		Spacing,
		op.OneOrMore{Value: Definition},
		EndOfFile,
	}
	Definition = op.And{
		Identifier,
		LEFTARROW,
		Expression,
	}
	Expression = op.And{
		Sequence,
		op.ZeroOrMore{
			Value: op.And{
				SLASH,
				Sequence,
			},
		},
	}
	Sequence = op.ZeroOrMore{
		Value: Prefix,
	}
	Prefix = op.And{
		op.Optional{
			Value: op.Or{AND, NOT},
		},
		Suffix,
	}
	Suffix = op.And{
		Primary,
		op.Optional{
			Value: op.Or{
				QUESTION,
				STAR,
				PLUS,
			},
		},
	}
	Primary = op.Or{
		op.And{Identifier, op.Not{Value: LEFTARROW}},
		op.And{OPEN, op.Reference{Name: "Expression"}, CLOSE},
		Literal, Class, DOT,
	}
	Identifier = op.And{
		IdentStart,
		op.ZeroOrMore{Value: IdentCont},
		Spacing,
	}
	IdentStart = op.Or{
		op.RuneRange{
			Min: 'a',
			Max: 'z',
		},
		op.RuneRange{
			Min: 'A',
			Max: 'Z',
		},
		'_',
	}
	IdentCont = op.Or{
		IdentStart,
		op.RuneRange{
			Min: '0',
			Max: '9',
		},
	}
	Literal = op.Or{
		op.And{'\'', op.ZeroOrMore{Value: op.And{op.Not{Value: '\''}, Char}}, '\'', Spacing},
		op.And{'"', op.ZeroOrMore{Value: op.And{op.Not{Value: '"'}, Char}}, '"', Spacing},
	}
	Class = op.And{
		'[', op.ZeroOrMore{Value: op.And{op.Not{Value: ']'}, Range}}, ']', Spacing,
	}
	Range = op.Or{
		op.And{Char, '-', Char},
		Char,
	}
	Char = op.Or{
		op.And{'\\', op.Or{'n', 'r', 't', '\'', '"', '[', ']', '\\'}},
		op.And{'\\', op.RuneRange{Min: '0', Max: '2'}, op.RuneRange{Min: '0', Max: '7'}, op.RuneRange{Min: '0', Max: '7'}},
		op.And{'\\', op.RuneRange{Min: '0', Max: '7'}, op.Optional{Value: op.RuneRange{Min: '0', Max: '7'}}},
		op.And{op.Not{Value: '\\'}, op.Any{}},
	}

	LEFTARROW = op.And{"<-", Spacing}
	SLASH     = op.And{"/", Spacing}
	AND       = op.And{"&", Spacing}
	NOT       = op.And{"!", Spacing}
	QUESTION  = op.And{"?", Spacing}
	STAR      = op.And{"*", Spacing}
	PLUS      = op.And{"+", Spacing}
	OPEN      = op.And{"(", Spacing}
	CLOSE     = op.And{")", Spacing}
	DOT       = op.And{".", Spacing}
	Spacing   = op.ZeroOrMore{Value: op.Or{Space, Comment}}
	Comment   = op.And{
		'#',
		op.ZeroOrMore{Value: op.And{op.Not{Value: EndOfLine}, op.Any{}}},
		EndOfLine,
	}
	Space     = op.Or{' ', '\t', EndOfLine}
	EndOfLine = op.Or{"\r\n", '\r', '\n'}
	EndOfFile = op.EOF{}
)
