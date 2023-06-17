package parser

import (
	"fmt"
)

// Capture is the interface that wraps the Parse method.
// Parse parses the input and returns a node.
type Capture interface {
	Parse(p *Parser) (end *Node, err error)

	Operator
	fmt.Stringer
}

// Operator is the interface that wraps the Match method.
// March check whether the interface matches the input.
type Operator interface {
	// Match the given value. Returns the end cursor if the match was successful.
	// Returns the start if the match failed.
	Match(start Cursor, p *Parser) (end Cursor, err error)

	fmt.Stringer
}
