package abnf_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/0x51-dev/upeg/abnf"
	"github.com/0x51-dev/upeg/ir"
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
)

var (
	//go:embed definition.abnf
	definitionSpec string

	//go:embed core/core.abnf
	coreSpec string
)

/*
 *	IPv6address   =                            6( h16 ":" ) ls32
 *	              /                       "::" 5( h16 ":" ) ls32
 *	              / [               h16 ] "::" 4( h16 ":" ) ls32
 *	              / [ *1( h16 ":" !":" ) h16 ] "::" 3( h16 ":" ) ls32
 *	              / [ *2( h16 ":" !":" ) h16 ] "::" 2( h16 ":" ) ls32
 *	              / [ *3( h16 ":" !":" ) h16 ] "::"    h16 ":"   ls32
 *	              / [ *4( h16 ":" !":" ) h16 ] "::"              ls32
 *	              / [ *5( h16 ":" !":" ) h16 ] "::"              h16
 *	              / [ *6( h16 ":" !":" ) h16 ] "::"
 */
var ipv6addressOperator = op.Capture{
	Name: "IPv6address",
	Value: op.Or{
		op.And{op.Repeat{Min: 6, Max: 6, Value: op.And{H16, ':'}}, Ls32},
		op.And{"::", op.Repeat{Min: 5, Max: 5, Value: op.And{H16, ':'}}, Ls32},
		op.And{op.Optional{Value: H16}, "::", op.Repeat{Min: 4, Max: 4, Value: op.And{H16, ':'}}, Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 1, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", op.Repeat{Min: 3, Max: 3, Value: op.And{H16, ':'}}, Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 2, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", op.Repeat{Min: 2, Max: 2, Value: op.And{H16, ':'}}, Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 3, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", H16, ':', Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 4, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", Ls32},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 5, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::", H16},
		op.And{op.Optional{Value: op.And{op.Repeat{Max: 6, Value: op.And{H16, ':', op.Not{Value: ':'}}}, H16}}, "::"}},
}

func TestIPv6address(t *testing.T) {
	for _, address := range []string{
		"1:2:3:4:5:6:7:8", //                            6( h16 ":" ) ls32
		"::2:3:4:5:6:7:8", //                       "::" 5( h16 ":" ) ls32
		"::3:4:5:6:7:8",   // [               h16 ] "::" 4( h16 ":" ) ls32
		"1::3:4:5:6:7:8",  // [               h16 ] "::" 4( h16 ":" ) ls32
		"::4:5:6:7:8",     // [ *1( h16 ":" ) h16 ] "::" 3( h16 ":" ) ls32
		"1:2::4:5:6:7:8",  // [ *1( h16 ":" ) h16 ] "::" 3( h16 ":" ) ls32
		"::8",             // [ *5( h16 ":" ) h16 ] "::"              h16
		"1:2:3:4:5:6::8",  // [ *5( h16 ":" ) h16 ] "::"              h16
		"1:2:3:4:5:6:7::", // [ *6( h16 ":" ) h16 ] "::"
	} {
		p, err := abnf.NewParser([]rune(address))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.ParseEOF(IPv6address); err != nil {
			t.Error(err)
		}
	}
}

func TestSpecifications(t *testing.T) {
	for _, spec := range []string{coreSpec, definitionSpec} {
		p, err := abnf.NewParser([]rune(spec))
		if err != nil {
			t.Fatal(err)
		}
		n, err := p.ParseEOF(abnf.Rulelist)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := ir.ParseRulelist(n); err != nil {
			t.Fatal(err)
		}
	}
}

type IPv6addressOperator struct{}

func (I IPv6addressOperator) Match(_ parser.Cursor, p *parser.Parser) (end parser.Cursor, err error) {
	return p.Match(ipv6addressOperator)
}

func (I IPv6addressOperator) Parse(p *parser.Parser) (end *parser.Node, err error) {
	return p.Parse(ipv6addressOperator)
}

func (I IPv6addressOperator) String() string {
	return fmt.Sprintf("{%s}", op.StringAny(ipv6addressOperator))
}
