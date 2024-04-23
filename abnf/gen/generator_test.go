package gen_test

import (
	"fmt"
	"github.com/0x51-dev/upeg/abnf/gen"
)

func ExampleGenerator_hexConcat() {
	g := gen.Generator{
		PackageName: "hexc",
	}
	raw, _ := g.GenerateOperators([]rune("grammar = %x74.72.75.65\n"))
	fmt.Println(raw)
	// Output:
	// // Package hexc is autogenerated by https://github.com/0x51-dev/upeg. DO NOT EDIT.
	// package hexc
	//
	// import (
	//     "github.com/0x51-dev/upeg/parser/op"
	// )
	//
	// var (
	//     Grammar = op.Capture{Name: "Grammar", Value: op.And{rune(0x74), rune(0x72), rune(0x75), rune(0x65)}}
	// )
}

func Example_custom() {
	g := gen.Generator{
		PackageName: "custom",
		CustomOperators: map[string]struct{}{
			"custom":   {},
			"alpha":    {},
			"alphaNum": {},
			"HEXDIG":   {},
		},
	}
	raw, _ := g.GenerateOperators([]rune("grammar = alpha alphaNum HEXDIG\ncustom = alpha\n"))
	fmt.Println(raw)
	// Output:
	// // Package custom is autogenerated by https://github.com/0x51-dev/upeg. DO NOT EDIT.
	// package custom
	//
	// import (
	//     "github.com/0x51-dev/upeg/parser/op"
	// )
	//
	// var (
	//     Grammar = op.Capture{Name: "Grammar", Value: op.And{AlphaOperator{}, AlphaNumOperator{}, HEXDIGOperator{}}}
	//     Custom = CustomOperator{} // op.Capture{Name: "Custom", Value: AlphaOperator{}}
	// )
}

func Example_externalDependencies() {
	g := gen.Generator{
		PackageName: "external",
		ExternalDependencies: map[string]gen.ExternalDependency{
			"alpha": {
				Name: "a",
				Path: "example.com/a",
			},
			"alphaNum": {
				Name: "an",
				Path: "example.com/an",
			},
			"HEXDIG": {
				Name: "hd",
				Path: "example.com/hd",
			},
		},
	}
	raw, _ := g.GenerateOperators([]rune("grammar = alpha alphaNum HEXDIG\n"))
	fmt.Println(raw)
	// Output:
	// // Package external is autogenerated by https://github.com/0x51-dev/upeg. DO NOT EDIT.
	// package external
	//
	// import (
	//     "example.com/a"
	//     "example.com/an"
	//     "example.com/hd"
	//     "github.com/0x51-dev/upeg/parser/op"
	// )
	//
	// var (
	//     Grammar = op.Capture{Name: "Grammar", Value: op.And{a.Alpha, an.AlphaNum, hd.HEXDIG}}
	// )
}
