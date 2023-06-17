package parser_test

import (
	"fmt"
	"github.com/0x51-dev/upeg/parser"
)

func ExampleInvalidTypeError() {
	fmt.Println(parser.NewInvalidTypeError('0'))
	// Output:
	// invalid type: int32
}

func ExampleNoMatchError() {
	p, _ := parser.New([]rune("test"))
	_, err := p.Match("testify")
	fmt.Println(err)
	// Output:
	// error stack:
	// 2) [1:4] "�" | no match: "testify"
	// test
	// ----^
	// 1) [1:4] "�" | no match: 'i'
	// test
	// ----^
}
