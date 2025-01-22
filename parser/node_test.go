package parser_test

import (
	"fmt"

	"github.com/0x51-dev/upeg/parser"
)

func ExampleNode_String() {
	a := parser.NewNode("a", "0")
	x := parser.NewParentNode("x", []*parser.Node{a})
	fmt.Println(x)
	// Output:
	// {"x": [{"a": "0"}]}
}
