package upeg

//go:generate go run github.com/0x51-dev/upeg/cmd/abnf --in=abnf/core/core.abnf --out=abnf/core/core.go --ignoreAll --package=core
//go:generate go run github.com/0x51-dev/upeg/cmd/abnf --in=abnf/definition.abnf --out=abnf/definition.go --ignore=defined-as,elements,c-wsp,c-nl,element,group --importCore

//go:generate go run github.com/0x51-dev/upeg/cmd/abnf --in=abnf/testdata/ipv6.abnf --out=abnf/custom_test.go --package=abnf_test --custom=IPv6address --importCore
