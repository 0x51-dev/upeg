package parser

import "testing"

func TestParser_Reset(t *testing.T) {
	p, err := New([]rune("abc"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := p.Match('a'); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Reset().Match("ab"); err != nil {
		t.Fatal(err)
	}
}
