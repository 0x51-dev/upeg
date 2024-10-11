package upeg

//go:generate go run github.com/0x51-dev/upeg/cmd/abnf --in=bnf/definition.abnf --out=bnf/definition.go --ignore=defined-as,elements,c-wsp,c-nl,element,group,rulename-br,literal-double,literal-single --importCore --package=bnf
