package op_test

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

func ExampleAny() {
	p, _ := parser.New([]rune("abc"))
	_, err := p.Match(op.Repeat{Min: 4, Max: 4, Value: op.Any{}})
	fmt.Println(err)
	// Output:
	// [1:3] "�" | no match: .
	// abc
	// ---^
}

func TestAny_Match(t *testing.T) {
	input := "abc"
	p, err := parser.New([]rune(input))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range input {
		if character := p.Reader.Cursor().Character(); character != c {
			t.Fatalf("expected %c, got %c", c, character)
		}
		if _, err := p.Match(op.Any{}); err != nil {
			t.Fatal(err)
		}
	}
	// No more characters to match.
	if _, err := p.Match(op.Any{}); err == nil {
		t.Fatal("expected error")
	}
}
