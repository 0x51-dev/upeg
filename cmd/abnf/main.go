package main

import (
	"flag"
	"github.com/0x51-dev/upeg/abnf/gen"
	"os"
	"strings"
)

func main() {
	var input = flag.String("in", "definition.abnf", "input file")
	var output = flag.String("out", "definition.go", "output file")
	var ignoreAll = flag.Bool("ignoreAll", false, "do not create capture groups for all rules")
	var ignoreList = flag.String("ignore", "", "comma separated list of rules to ignore")
	var importCore = flag.Bool("importCore", false, "import core rules")
	var packageName = flag.String("package", "abnf", "package name")
	var customList = flag.String("custom", "", "comma separated list of custom operators")
	var dependencyList = flag.String("dependencies", "", "semicolon seperated list of dependencies")
	flag.Parse()

	ignore := make(map[string]struct{})
	for _, v := range strings.Split(*ignoreList, ",") {
		ignore[v] = struct{}{}
	}

	custom := make(map[string]struct{})
	for _, v := range strings.Split(*customList, ",") {
		custom[v] = struct{}{}
	}

	dependencies := make(map[string]gen.ExternalDependency)
	for _, dependency := range strings.Split(*dependencyList, ";") {
		dep := strings.Split(dependency, ",")
		if len(dep) < 2 {
			continue
		}
		path := dep[0]
		name := dep[0][strings.LastIndex(path, "/")+1:]
		for _, o := range dep[1:] {
			dependencies[o] = gen.ExternalDependency{
				Path: path,
				Name: name,
			}
		}
	}

	rawInput, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	g := gen.Generator{
		PackageName:          *packageName,
		IgnoreAll:            *ignoreAll,
		Ignore:               ignore,
		ImportCore:           *importCore,
		CustomOperators:      custom,
		ExternalDependencies: dependencies,
	}
	out, err := g.GenerateOperators([]rune(string(rawInput)))
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(*output, []byte(out), 0644); err != nil {
		panic(err)
	}
}
