package ir

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
	"strconv"
)

func ParseEmpty(n *parser.Node) error {
	return checkName(n, "Empty")
}

func checkName(n *parser.Node, name string) error {
	if n == nil {
		return NewInvalidNodeError(name, "nil")
	}
	if n.Name != name {
		return NewInvalidNodeError(name, n.Name)
	}
	return nil
}

type Alternation struct {
	Concatenations []*Concatenation
}

func ParseAlternation(n *parser.Node) (*Alternation, error) {
	var concatenations []*Concatenation
	for _, n := range n.Children() {
		switch n.Name {
		case "Concatenation":
			concatenation, err := ParseConcatenation(n)
			if err != nil {
				return nil, err
			}
			concatenations = append(concatenations, concatenation)
		default:
			return nil, NewInvalidNodeError("Concatenation", n.Name)
		}
	}
	return &Alternation{
		Concatenations: concatenations,
	}, nil
}

func (alt *Alternation) term() {}

type Concatenation struct {
	Factors []*Factor
}

func ParseConcatenation(n *parser.Node) (*Concatenation, error) {
	if err := checkName(n, "Concatenation"); err != nil {
		return nil, err
	}
	var factors []*Factor
	for _, n := range n.Children() {
		switch n.Name {
		case "Factor":
			factor, err := ParseFactor(n)
			if err != nil {
				return nil, err
			}
			factors = append(factors, factor)
		default:
			return nil, NewInvalidNodeError("Factor", n.Name)
		}
	}
	return &Concatenation{
		Factors: factors,
	}, nil
}

type Definition struct {
	Identifier  *Identifier
	Alternation *Alternation
}

func ParseDefinition(n *parser.Node) (*Definition, error) {
	if err := checkName(n, "Definition"); err != nil {
		return nil, err
	}
	identifier, err := ParseIdentifier(n.Children()[0])
	if err != nil {
		return nil, err
	}
	switch n := n.Children()[1]; n.Name {
	case "Alternation":
		alternation, err := ParseAlternation(n)
		if err != nil {
			return nil, err
		}
		return &Definition{
			Identifier:  identifier,
			Alternation: alternation,
		}, nil
	case "Empty":
		if err := ParseEmpty(n); err != nil {
			return nil, err
		}
		return &Definition{
			Identifier:  identifier,
			Alternation: nil,
		}, nil
	default:
		return nil, NewInvalidNodeError("Alternation", n.Name)
	}
}

type Factor struct {
	Amount     *int
	Term       Term
	Optional   bool
	ZeroOrMore bool
	OneOrMore  bool
	Except     Term
}

func ParseFactor(n *parser.Node) (*Factor, error) {
	if err := checkName(n, "Factor"); err != nil {
		return nil, err
	}
	var factor Factor
	cs := n.Children()
	if cs[0].Name == "Amount" {
		amount, err := strconv.Atoi(cs[0].Value())
		if err != nil {
			return nil, err
		}
		factor.Amount = &amount
		cs = cs[1:]
	}
	switch n := cs[0]; n.Name {
	case "Term":
		term, err := ParseTerm(n)
		if err != nil {
			return nil, err
		}
		factor.Term = term
	default:
		return nil, NewInvalidNodeError("Term", n.Name)
	}
	if len(cs) > 1 {
		switch n := cs[1]; n.Name {
		case "Optional":
			factor.Optional = true
		case "ZeroOrMore":
			factor.ZeroOrMore = true
		case "Except":
			if len(n.Children()) != 0 {
				term, err := ParseTerm(n.Children()[0])
				if err != nil {
					return nil, err
				}
				factor.Except = term
			} else {
				factor.OneOrMore = true
			}
		}
	}
	return &factor, nil
}

type Grammar struct {
	Definitions []*Definition
}

func ParseGrammar(n *parser.Node) (*Grammar, error) {
	if err := checkName(n, "Grammar"); err != nil {
		return nil, err
	}
	var definitions []*Definition
	for _, n := range n.Children() {
		switch n.Name {
		case "Definition":
			definition, err := ParseDefinition(n)
			if err != nil {
				return nil, err
			}
			definitions = append(definitions, definition)
		default:
			return nil, NewInvalidNodeError("Definition", n.Name)
		}
	}
	return &Grammar{
		Definitions: definitions,
	}, nil
}

type Identifier struct {
	Name string
}

func ParseIdentifier(n *parser.Node) (*Identifier, error) {
	if err := checkName(n, "Identifier"); err != nil {
		return nil, err
	}
	return &Identifier{Name: n.Value()}, nil
}

func (id *Identifier) term() {}

type InvalidNodeError struct {
	Expected string
	Actual   string
}

func NewInvalidNodeError(expected string, actual string) *InvalidNodeError {
	return &InvalidNodeError{
		Expected: expected,
		Actual:   actual,
	}
}

func (e *InvalidNodeError) Error() string {
	return fmt.Sprintf("invalid node: expected %s, got %s", e.Expected, e.Actual)
}

type Optional struct {
	Alternation *Alternation
}

func ParseOptional(n *parser.Node) (*Optional, error) {
	if err := checkName(n, "Optional"); err != nil {
		return nil, err
	}
	alternation, err := ParseAlternation(n.Children()[0])
	if err != nil {
		return nil, err
	}
	return &Optional{
		Alternation: alternation,
	}, nil
}

func (o *Optional) term() {}

type Repetition struct {
	Alternation *Alternation
}

func ParseRepetition(n *parser.Node) (*Repetition, error) {
	if err := checkName(n, "Repetition"); err != nil {
		return nil, err
	}
	alternation, err := ParseAlternation(n.Children()[0])
	if err != nil {
		return nil, err
	}
	return &Repetition{
		Alternation: alternation,
	}, nil
}

func (r *Repetition) term() {}

type SpecialSequence struct {
	Text string
}

func ParseSpecialSequence(n *parser.Node) (*SpecialSequence, error) {
	if err := checkName(n, "SpecialSequence"); err != nil {
		return nil, err
	}
	return &SpecialSequence{
		Text: n.Value(),
	}, nil
}

func (s *SpecialSequence) term() {}

type String struct {
	Value string
}

func ParseString(n *parser.Node) (*String, error) {
	if err := checkName(n, "String"); err != nil {
		return nil, err
	}
	return &String{
		Value: n.Value(),
	}, nil
}

func (s *String) term() {}

type Term interface {
	term()
}

func ParseTerm(n *parser.Node) (Term, error) {
	if err := checkName(n, "Term"); err != nil {
		return nil, err
	}
	switch n := n.Children()[0]; n.Name {
	case "Identifier":
		return ParseIdentifier(n)
	case "String":
		return ParseString(n)
	case "Alternation":
		return ParseAlternation(n)
	case "Repetition":
		return ParseRepetition(n)
	case "Optional":
		return ParseOptional(n)
	case "SpecialSequence":
		return ParseSpecialSequence(n)
	default:
		return nil, NewInvalidNodeError("Term", n.Name)
	}
}
