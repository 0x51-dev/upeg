package parser

import "fmt"

// Cursor represents the current position in the input.
type Cursor struct {
	// character is the current character rune.
	character rune
	// position is the absolute position in the input.
	position uint

	// lastNewline is the last newline absolute position.
	lastNewline uint

	// line is the current line.
	line uint
	// column is the current column.
	column uint
}

// Character returns the current character rune.
func (c Cursor) Character() rune {
	return c.character
}

// Position returns the absolute position in the input.
func (c Cursor) Position() uint {
	return c.position
}

// LastNewLine returns the last end of line absolute position.
func (c Cursor) LastNewLine() uint {
	return c.lastNewline
}

// Line returns the current line.
func (c Cursor) Line() (uint, uint) {
	return c.line, c.column
}

// String returns the string representation of the cursor.
func (c Cursor) String() string {
	return fmt.Sprintf("[%d:%d] %q", c.line+1, c.column+1, string(c.character))
}
