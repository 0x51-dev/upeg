package parser

import (
	"fmt"
	"strings"
)

// ErrorStack is a stack of errors.
type ErrorStack struct {
	errors []error
}

// NewErrorStack returns a new ErrorStack.
func NewErrorStack(context, err error) *ErrorStack {
	if e, ok := err.(*ErrorStack); ok {
		return e.AddError(context)
	}
	return &ErrorStack{
		errors: []error{err, context},
	}
}

// AddError adds an error to the stack.
func (e *ErrorStack) AddError(err error) *ErrorStack {
	e.errors = append(e.errors, err)
	return e
}

// Error returns the error message.
func (e *ErrorStack) Error() string {
	var errors []string
	for i := len(e.errors) - 1; i >= 0; i-- {
		errors = append(errors, fmt.Sprintf("%d) %s", i+1, e.errors[i].Error()))
	}
	return fmt.Sprintf("error stack:\n%s", strings.Join(errors, "\n"))
}

// InvalidTypeError is returned when an invalid type is passed to a function.
type InvalidTypeError struct {
	v any
}

// NewInvalidTypeError returns a new InvalidTypeError.
func NewInvalidTypeError(v any) *InvalidTypeError {
	return &InvalidTypeError{v}
}

// Error returns the error message.
func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf("invalid type: %T", e.v)
}

// NoMatchError is returned when a rule does not match.
type NoMatchError struct {
	v   any
	end Cursor
	p   *Parser
	err error
}

// Error returns the error message.
func (e *NoMatchError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s | no match: %v: %v", e.end, e.v, e.err)
	}

	l := string(e.p.Reader.GetLine(e.end))
	p := strings.Repeat("-", len(l))
	switch e.v.(type) {
	case rune, string:
		return fmt.Sprintf("%s | no match: %q\n%s\n%s^", e.end, e.v, l, p)
	default:
		return fmt.Sprintf("%s | no match: %v\n%s\n%s^", e.end, e.v, l, p)
	}
}

// NewNoMatchError returns a new NoMatchError.
func (p *Parser) NewNoMatchError(v any, end Cursor) *NoMatchError {
	return &NoMatchError{
		v:   v,
		end: end,
		p:   p,
	}
}
