package abnf_test

import (
	_ "embed"
	"github.com/0x51-dev/upeg/abnf"
	"github.com/0x51-dev/upeg/parser/op"
	"testing"
)

var (
	//go:embed definition.abnf
	definitionSpec string

	//go:embed core.abnf
	coreSpec string
)

func TestSpecifications(t *testing.T) {
	for _, spec := range []string{coreSpec, definitionSpec} {
		p, err := abnf.NewParser([]rune(spec))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := p.Parse(op.And{abnf.Rulelist, op.EOF{}}); err != nil {
			t.Fatal(err)
		}
	}
}
