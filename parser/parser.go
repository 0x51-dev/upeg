package parser

// Parser is the parser.
type Parser struct {
	// Reader is the input reader.
	Reader *Reader
	Rules  map[string]Operator
}

// New creates a new parser.
func New(input []rune) (*Parser, error) {
	r, err := NewReader(input)
	if err != nil {
		return nil, err
	}
	return &Parser{
		Reader: r,
		Rules:  make(map[string]Operator),
	}, nil
}

// Match the given value.
// Returns the end cursor if the match was successful.
// Returns an error if the match failed.
func (p *Parser) Match(v any) (Cursor, error) {
	start := p.Reader.Cursor()

	// Match operator.
	if v, ok := v.(Operator); ok {
		return v.Match(start, p)
	}

	return p.matchPrimitive(start, v)
}

// Parse the given value.
func (p *Parser) Parse(v any) (*Node, error) {
	if v, ok := v.(Capture); ok {
		return v.Parse(p)
	}
	_, err := p.Match(v)
	return nil, err
}

// matchPrimitive matches a primitive value. Supports runes and strings.
func (p *Parser) matchPrimitive(start Cursor, v any) (Cursor, error) {
	switch v := v.(type) {

	case rune:
		if start.character != v {
			return start, p.NewNoMatchError(v, start)
		}
		return p.Reader.Next().Cursor(), nil

	case string:
		end := start
		for _, r := range v {
			var err error
			if end, err = p.matchPrimitive(end, r); err != nil {
				p.Reader.Jump(start)
				return end, NewErrorStack(p.NewNoMatchError(v, end), err)
			}
		}
		return end, nil

	default:
		return start, NewInvalidTypeError(v)
	}
}
