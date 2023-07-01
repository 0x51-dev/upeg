package upeg

//go:generate go run github.com/0x51-dev/upeg/cmd/abnf --in=abnf/core.abnf --out=abnf/core.go --ignoreAll
//go:generate go run github.com/0x51-dev/upeg/cmd/abnf --in=abnf/definition.abnf --out=abnf/definition.go --ignore=defined-as,elements,c-wsp,c-nl,element,group
