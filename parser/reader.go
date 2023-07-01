package parser

import (
	"fmt"
	"unicode/utf8"
)

// Reader is the input reader.
type Reader struct {
	// input is the input.
	input []rune
	// cursor is the current position in the input.
	cursor Cursor
}

// NewReader creates a new reader.
func NewReader(input []rune) (*Reader, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("input is empty")
	}
	return &Reader{
		input: input,
		cursor: Cursor{
			character: input[0],
		},
	}, nil
}

// Cursor the current cursor.
func (r *Reader) Cursor() Cursor {
	return r.cursor
}

// Done returns true if the reader is done.
func (r *Reader) Done() bool {
	return r.cursor.position >= uint(len(r.input))
}

// GetInputRange returns the input range from start to end.
func (r *Reader) GetInputRange(start Cursor, end Cursor) []rune {
	return r.input[start.position:end.position]
}

func (r *Reader) GetLine(end Cursor) []rune {
	position := end.position
	if position == uint(len(r.input)) {
		// Return the last line.
		return r.input[end.lastNewline:]
	}

	character := r.input[position]
	for position < uint(len(r.input)) && character != '\n' && character != '\r' {
		// Reached the end of the input.
		if position++; position == uint(len(r.input)) {
			return r.input[end.lastNewline:]
		}
		character = r.input[position]
	}
	return r.input[end.lastNewline:position]
}

// Jump to the given cursor. Ignore if cursor is nil or the reader is done.
func (r *Reader) Jump(marker Cursor) {
	r.cursor = marker
}

// Next reads the next rune.
func (r *Reader) Next() *Reader {
	r.cursor.position++
	if r.Done() {
		r.cursor.character = utf8.RuneError
		return r
	}

	// Check whether the previous character was a newline.
	r.cursor.column++
	if r.cursor.character == '\n' {
		// Covers both \n and \r\n.
		r.cursor.line++
		r.cursor.column = 0
		r.cursor.lastNewline = r.cursor.position
	}

	next := r.input[r.cursor.position]
	if next != '\n' && r.Cursor().character == '\r' {
		// Covers \r.
		r.cursor.line++
		r.cursor.column = 0
		r.cursor.lastNewline = r.cursor.position
	}

	r.cursor.character = next
	return r
}

// Rune returns the current rune.
func (r *Reader) Rune() rune {
	if r.Done() {
		return utf8.RuneError
	}
	return r.cursor.character
}
