package ir

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
	"strings"
)

func ParseRulename(n *parser.Node) (string, error) {
	if err := checkName(n, "Rulename"); err != nil {
		return "", err
	}
	return n.Value(), nil
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

func checkParent(n *parser.Node, name string) error {
	if n == nil {
		return NewInvalidNodeError(name, "nil")
	}
	if n.Name != name || n.Value() != "" {
		return NewInvalidNodeError(name, n.Name)
	}
	return nil
}

type Alternation struct {
	Concatenations []*Concatenation
}

func ParseAlternation(n *parser.Node) (*Alternation, error) {
	if err := checkParent(n, "Alternation"); err != nil {
		return nil, err
	}
	var concatenations []*Concatenation
	for _, n := range n.Children() {
		switch n.Name {
		case "":
			for _, n := range n.Children() {
				concatenation, err := ParseConcatenation(n)
				if err != nil {
					return nil, err
				}
				concatenations = append(concatenations, concatenation)
			}
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

func (a Alternation) String() string {
	switch len(a.Concatenations) {
	case 0:
		return ""
	case 1:
		return a.Concatenations[0].String()
	default:
		var s []string
		for _, c := range a.Concatenations {
			s = append(s, c.String())
		}
		return fmt.Sprintf("( %s )", strings.Join(s, " / "))
	}
}

func (a Alternation) isElement() {}

type CharVal string

func (c *CharVal) String() string {
	return string(*c)
}

func (c *CharVal) isElement() {}

type Concatenation struct {
	Repetitions []*Repetition
}

func ParseConcatenation(n *parser.Node) (*Concatenation, error) {
	if err := checkParent(n, "Concatenation"); err != nil {
		return nil, err
	}
	var repetitions []*Repetition
	for _, n := range n.Children() {
		switch n.Name {
		case "":
			for _, n := range n.Children() {
				repetition, err := ParseRepetition(n)
				if err != nil {
					return nil, err
				}
				repetitions = append(repetitions, repetition)
			}
		case "Repetition":
			repetition, err := ParseRepetition(n)
			if err != nil {
				return nil, err
			}

			repetitions = append(repetitions, repetition)
		default:
			return nil, NewInvalidNodeError("Repetition", n.Name)
		}
	}
	return &Concatenation{
		Repetitions: repetitions,
	}, nil
}

func (c *Concatenation) String() string {
	switch len(c.Repetitions) {
	case 0:
		return ""
	case 1:
		return c.Repetitions[0].String()
	default:
		var s []string
		for _, r := range c.Repetitions {
			s = append(s, r.String())
		}
		return fmt.Sprintf("( %s )", strings.Join(s, " "))
	}
}

type Element interface {
	fmt.Stringer
	isElement()
}

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

type NumVal string

func (n *NumVal) String() string {
	return string(*n)
}

func (n *NumVal) isElement() {}

type Option struct {
	Alternation *Alternation
}

func ParseOption(n *parser.Node) (*Option, error) {
	if err := checkParent(n, "Option"); err != nil {
		return nil, err
	}
	if len(n.Children()) != 1 {
		return nil, NewInvalidNodeError("Option", "")
	}
	a, err := ParseAlternation(n.Children()[0])
	if err != nil {
		return nil, err
	}
	return &Option{
		Alternation: a,
	}, nil
}

func (o *Option) String() string {
	return fmt.Sprintf("[ %s ]", o.Alternation.String())
}

func (o *Option) isElement() {}

type ProseVal string

func (p *ProseVal) String() string {
	return string(*p)
}

func (p *ProseVal) isElement() {}

type Repeat struct {
	Min *string
	Max *string
}

func ParseRepeat(n *parser.Node) (*Repeat, error) {
	if err := checkName(n, "Repeat"); err != nil {
		return nil, err
	}
	v := n.Value()
	if !strings.ContainsRune(v, '*') {
		return &Repeat{
			Min: &v,
			Max: &v,
		}, nil
	}
	s := strings.Split(v, "*")
	if len(s) != 2 {
		return nil, fmt.Errorf("invalid repeat: %s", v)
	}
	if s[0] == "" && s[1] == "" {
		return &Repeat{}, nil
	}
	if s[0] == "" {
		return &Repeat{
			Max: &s[1],
		}, nil
	}
	if s[1] == "" {
		return &Repeat{
			Min: &s[0],
		}, nil
	}
	return &Repeat{
		Min: &s[0],
		Max: &s[1],
	}, nil
}

func (r *Repeat) String() string {
	if r.Min == nil && r.Max == nil {
		return "*"
	}
	if r.Min == nil {
		return fmt.Sprintf("*%s", *r.Max)
	}
	if r.Max == nil {
		return fmt.Sprintf("%s*", *r.Min)
	}
	if *r.Min == *r.Max {
		return *r.Min
	}
	return fmt.Sprintf("%s*%s", *r.Min, *r.Max)
}

type Repetition struct {
	Repeat *Repeat
	Value  Element
}

func ParseRepetition(n *parser.Node) (*Repetition, error) {
	if err := checkParent(n, "Repetition"); err != nil {
		return nil, err
	}
	var r *Repeat
	for i, n := range n.Children() {
		switch n.Name {
		case "Repeat":
			if i != 0 {
				return nil, NewInvalidNodeError("Repeat", n.Name)
			}
			repeat, err := ParseRepeat(n)
			if err != nil {
				return nil, err
			}
			r = repeat
		case "Rulename":
			v := Rulename(n.Value())
			return &Repetition{
				Repeat: r,
				Value:  &v,
			}, nil
		case "Alternation":
			a, err := ParseAlternation(n)
			if err != nil {
				return nil, err
			}
			return &Repetition{
				Repeat: r,
				Value:  a,
			}, nil
		case "Option":
			o, err := ParseOption(n)
			if err != nil {
				return nil, err
			}
			return &Repetition{
				Repeat: r,
				Value:  o,
			}, nil
		case "CharVal":
			v := CharVal(n.Value())
			return &Repetition{
				Repeat: r,
				Value:  &v,
			}, nil
		case "NumVal":
			v := NumVal(n.Children()[0].Value())
			return &Repetition{
				Repeat: r,
				Value:  &v,
			}, nil
		case "ProseVal":
			v := ProseVal(n.Value())
			return &Repetition{
				Repeat: r,
				Value:  &v,
			}, nil
		default:
			return nil, NewInvalidNodeError("Element / Repeat", n.Name)
		}
	}
	return nil, NewInvalidNodeError("Element / Repeat", "")
}

func (r *Repetition) String() string {
	if r.Repeat == nil {
		return r.Value.String()
	}
	return fmt.Sprintf("%s%s", r.Repeat, r.Value)
}

type Rule struct {
	Rulename    string
	Alternation *Alternation
}

func ParseRule(n *parser.Node) (*Rule, error) {
	if err := checkParent(n, "Rule"); err != nil {
		return nil, err
	}
	r, err := ParseRulename(n.Children()[0])
	if err != nil {
		return nil, err
	}
	a, err := ParseAlternation(n.Children()[1])
	if err != nil {
		return nil, err
	}
	return &Rule{
		Rulename:    r,
		Alternation: a,
	}, nil
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s = %s", r.Rulename, r.Alternation)
}

type Rulelist struct {
	Rules []*Rule
}

func ParseRulelist(n *parser.Node) (*Rulelist, error) {
	if err := checkParent(n, "Rulelist"); err != nil {
		return nil, err
	}
	var rules []*Rule
	for _, n := range n.Children() {
		switch n.Name {
		case "Rule":
			rule, err := ParseRule(n)
			if err != nil {
				return nil, err
			}
			rules = append(rules, rule)
		case "Comment": // Ignore these.
		default:
			return nil, NewInvalidNodeError("Rule", n.Name)
		}
	}
	return &Rulelist{
		Rules: rules,
	}, nil
}

func (r *Rulelist) String() string {
	var rules []string
	for _, rule := range r.Rules {
		rules = append(rules, rule.String())
	}
	return strings.Join(rules, "\n")
}

type Rulename string

func (r *Rulename) String() string {
	return string(*r)
}

func (r *Rulename) isElement() {}
