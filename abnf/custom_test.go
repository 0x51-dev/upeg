// DO NOT EDIT: This file is autogenerated by https://github.com/0x51-dev/upeg.
package abnf_test

import (
	. "github.com/0x51-dev/upeg/abnf/core"
	"github.com/0x51-dev/upeg/parser/op"
)

var (
	IPv6address = IPv6addressOperator{} // op.Capture{Name: "IPv6address", Value: op.Or{op.And{op.Repeat{Min: 6, Max: 6, Value: op.And{H16, ':'}}, Ls32}, op.And{"::", op.Repeat{Min: 5, Max: 5, Value: op.And{H16, ':'}}, Ls32}, op.And{op.Optional{Value: H16}, "::", op.Repeat{Min: 4, Max: 4, Value: op.And{H16, ':'}}, Ls32}, op.And{op.Optional{Value: op.And{op.Repeat{Max: 1, Value: op.And{H16, ':'}}, H16}}, "::", op.Repeat{Min: 3, Max: 3, Value: op.And{H16, ':'}}, Ls32}, op.And{op.Optional{Value: op.And{op.Repeat{Max: 2, Value: op.And{H16, ':'}}, H16}}, "::", op.Repeat{Min: 2, Max: 2, Value: op.And{H16, ':'}}, Ls32}, op.And{op.Optional{Value: op.And{op.Repeat{Max: 3, Value: op.And{H16, ':'}}, H16}}, "::", H16, ':', Ls32}, op.And{op.Optional{Value: op.And{op.Repeat{Max: 4, Value: op.And{H16, ':'}}, H16}}, "::", Ls32}, op.And{op.Optional{Value: op.And{op.Repeat{Max: 5, Value: op.And{H16, ':'}}, H16}}, "::", H16}, op.And{op.Optional{Value: op.And{op.Repeat{Max: 6, Value: op.And{H16, ':'}}, H16}}, "::"}}}
	H16         = op.Capture{Name: "H16", Value: op.Repeat{Min: 1, Max: 4, Value: HEXDIG}}
	Ls32        = op.Capture{Name: "Ls32", Value: op.Or{op.And{H16, ':', H16}, IPv4address}}
	IPv4address = op.Capture{Name: "IPv4address", Value: op.And{DecOctet, '.', DecOctet, '.', DecOctet, '.', DecOctet}}
	DecOctet    = op.Capture{Name: "DecOctet", Value: op.Or{DIGIT, op.And{op.RuneRange{Min: 0x31, Max: 0x39}, DIGIT}, op.And{'1', op.Repeat{Min: 2, Max: 2, Value: DIGIT}}, op.And{'2', op.RuneRange{Min: 0x30, Max: 0x34}, DIGIT}, op.And{"25", op.RuneRange{Min: 0x30, Max: 0x35}}}}
)
