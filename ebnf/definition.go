package ebnf

import "github.com/0x51-dev/upeg/parser/op"

var (
	Letter = op.Or{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
	Digit = op.Or{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	Symbol = op.Or{
		'[', ']', '{', '}', '(', ')', '<', '>',
		'\'', '"', '=', '|', '.', ',', ';', '-',
		'+', '*', '?', "\\n", "\\t", "\\r", "\\f", "\\b",
	}
	Character  = op.Or{Letter, Digit, Symbol, '_', ' '}
	Identifier = op.And{Letter, op.ZeroOrMore{Value: op.Or{Letter, Digit, "_"}}}
	S          = op.ZeroOrMore{Value: op.Or{' ', '\n', '\t', '\r', '\f', '\b'}}
	Terminal   = op.Or{
		op.And{'\'', op.Not{Value: '\''}, Character, op.ZeroOrMore{Value: op.And{op.Not{Value: '\''}, Character}}, '\''},
		op.And{'"', op.Not{Value: '"'}, Character, op.ZeroOrMore{Value: op.And{op.Not{Value: '"'}, Character}}, '"'},
	}
	Terminator = op.Or{';', '.'}
	Term       = op.Or{
		op.And{'(', S, op.Reference{Name: "Alternation"}, S, ')'},
		op.And{'[', S, op.Reference{Name: "Alternation"}, S, ']'},
		op.And{'{', S, op.Reference{Name: "Alternation"}, S, '}'},
		Terminal,
		Identifier,
	}
	Factor = op.Or{
		op.And{Term, S, "?"},
		op.And{Term, S, "*"},
		op.And{Term, S, "+"},
		op.And{Term, S, "-", S, Term},
		op.And{Term, S},
	}
	Concatenation = op.And{S, Factor, S, op.ZeroOrMore{Value: op.And{',', S, Factor, S}}}
	Alternation   = op.And{S, Concatenation, S, op.ZeroOrMore{Value: op.And{'|', S, Concatenation, S}}}
	Rule          = op.And{Identifier, S, '=', S, Alternation, S, Terminator}
	Grammar       = op.ZeroOrMore{Value: op.And{S, Rule, S}}
)
