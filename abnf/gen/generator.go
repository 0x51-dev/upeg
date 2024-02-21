package gen

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/0x51-dev/upeg/abnf"
	"github.com/0x51-dev/upeg/abnf/ir"
	"github.com/0x51-dev/upeg/parser/op"
	"io"
	"io/fs"
	"log"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

const (
	templatesDir = "templates"
)

var (
	//go:embed templates/*
	files     embed.FS
	templates map[string]*template.Template
)

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		panic(err)
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(files, templatesDir+"/"+tmpl.Name())
		if err != nil {
			panic(err)
		}

		templates[strings.TrimSuffix(tmpl.Name(), ".gotmpl")] = pt
	}
}

type DependencyTree = map[string]map[string]struct{}

type ExternalDependency struct {
	Name string
	Path string
}

type Generator struct {
	PackageName          string
	IgnoreAll            bool
	Ignore               map[string]struct{}
	ImportCore           bool
	CustomOperators      map[string]struct{}
	ExternalDependencies map[string]ExternalDependency

	usedExternalDependencies map[string]struct{}
	references               []string
}

func (g *Generator) GenerateOperators(input []rune) (string, error) {
	g.usedExternalDependencies = make(map[string]struct{})
	g.references = nil // Reset.

	p, err := abnf.NewParser(input)
	if err != nil {
		return "", err
	}
	n, err := p.Parse(op.And{abnf.Rulelist, op.EOF{}})
	if err != nil {
		return "", err
	}
	list, err := ir.ParseRulelist(n)
	if err != nil {
		return "", err
	}
	return g.generateOperators(list)
}

func (g *Generator) cyclic(dependencies map[string]map[string]struct{}, key string) map[string]struct{} {
	flat := make(map[string]struct{})
	for k := range dependencies[key] {
		flat[k] = struct{}{}
	}
	cyclic := make(map[string]struct{})
	var l = 0
	for l != len(flat) {
		l = len(flat)
		for dep := range flat {
			for d := range dependencies[dep] {
				if d == key {
					cyclic[dep] = struct{}{}
				}
				flat[d] = struct{}{}
			}
		}
	}
	return cyclic
}

func (g *Generator) dependencies(v any) map[string]struct{} {
	switch v := v.(type) {
	case *ir.Alternation:
		deps := make(map[string]struct{})
		for _, c := range v.Concatenations {
			for dep := range g.dependencies(c) {
				deps[dep] = struct{}{}
			}
		}
		return deps
	case *ir.Concatenation:
		deps := make(map[string]struct{})
		for _, c := range v.Repetitions {
			for dep := range g.dependencies(c) {
				deps[dep] = struct{}{}
			}
		}
		return deps
	case *ir.Repetition:
		return g.dependencies(v.Value)
	case *ir.Option:
		return g.dependencies(v.Alternation)
	case *ir.Rulename:
		return map[string]struct{}{g.ruleName(string(*v)): {}}
	case *ir.ProseVal:
		name := string(*v)
		name = name[1 : len(name)-1] // Remove "<>"
		return map[string]struct{}{g.ruleName(name): {}}
	case *ir.NumVal, *ir.CharVal: // Ignore
		return nil
	default:
		panic(fmt.Errorf("unknown type %T", v))
	}
}

func (g *Generator) dependencyTree(list *ir.Rulelist) DependencyTree {
	deps := make(map[string]map[string]struct{})
	for _, rule := range list.Rules {
		deps[g.ruleName(rule.Rulename)] = g.dependencies(rule.Alternation)
	}
	return deps
}

func (g *Generator) generateOperators(list *ir.Rulelist) (string, error) {
	t, ok := templates["abnf"]
	if !ok {
		return "", fmt.Errorf("template not found")
	}

	rules := g.rulelistToGo(list)
	var dependencies []string
	for k := range g.usedExternalDependencies {
		dependencies = append(dependencies, k)
	}
	sort.Slice(dependencies, func(i, j int) bool {
		return dependencies[i] < dependencies[j]
	})

	var tmpl bytes.Buffer
	if err := t.Execute(&tmpl, &abnfTemplateData{
		PackageName:  g.PackageName,
		ImportCore:   g.ImportCore,
		References:   g.references,
		Dependencies: dependencies,
		Rules:        rules,
	}); err != nil {
		return "", err
	}
	raw, err := io.ReadAll(&tmpl)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func (g *Generator) ruleName(s string) string {
	ss := strings.Split(s, "-")
	for i, s := range ss {
		r := []rune(s)
		ss[i] = string(unicode.ToUpper(r[0])) + string(r[1:])
	}
	return strings.Join(ss, "")
}

func (g *Generator) rulelistToGo(list *ir.Rulelist) []abnfRule {
	reference := make(map[string][]string)
	deps := g.dependencyTree(list)
	for custom := range g.CustomOperators {
		// Remove any custom operators form the (cyclic) dependency list.
		delete(deps, custom)
	}

	var names []string
	for _, rule := range list.Rules {
		name := g.ruleName(rule.Rulename)
		c := g.cyclic(deps, name)
		if len(c) != 0 {
			g.references = append(g.references, name)
			log.Printf("[CYCLIC REF] %q.", name)
			for dep := range c {
				delete(deps[dep], name)
				reference[dep] = append(reference[dep], name)
			}
		}
		names = append(names, name)
	}

	var rules []abnfRule
	for i, rule := range list.Rules {
		operator := g.toGo(rule.Alternation, reference[names[i]])
		if _, ok := g.Ignore[rule.Rulename]; !g.IgnoreAll && !ok {
			operator = fmt.Sprintf("op.Capture{Name: %q, Value: %s}", names[i], operator)
		}
		if _, ok := g.CustomOperators[rule.Rulename]; ok {
			log.Printf("[CUSTOM] %q.", rule.Rulename)
			operator = fmt.Sprintf("%sOperator{} // %s", g.ruleName(rule.Rulename), operator)
		}
		rules = append(rules, abnfRule{
			Name:     names[i],
			Operator: operator,
		})

	}
	return rules
}

func (g *Generator) toGo(v any, references []string) string {
	switch v := v.(type) {
	case *ir.Alternation:
		switch len(v.Concatenations) {
		case 0:
			return ""
		case 1:
			return g.toGo(v.Concatenations[0], references)
		default:
			var s []string
			for _, c := range v.Concatenations {
				s = append(s, g.toGo(c, references))
			}
			return fmt.Sprintf("op.Or{%s}", strings.Join(s, ", "))
		}
	case *ir.Concatenation:
		switch len(v.Repetitions) {
		case 0:
			return ""
		case 1:
			return g.toGo(v.Repetitions[0], references)
		default:
			var s []string
			for _, c := range v.Repetitions {
				s = append(s, g.toGo(c, references))
			}
			return fmt.Sprintf("op.And{%s}", strings.Join(s, ", "))
		}
	case *ir.Repetition:
		if v.Repeat == nil {
			return g.toGo(v.Value, references)
		}
		if v.Repeat.Min == nil && v.Repeat.Max == nil {
			return fmt.Sprintf("op.ZeroOrMore{Value: %s}", g.toGo(v.Value, references))
		}
		if v.Repeat.Min != nil && v.Repeat.Max == nil {
			if *v.Repeat.Min == "1" {
				return fmt.Sprintf("op.OneOrMore{Value: %s}", g.toGo(v.Value, references))
			}
			return fmt.Sprintf("op.Repeat{Min: %s, Value: %s}", *v.Repeat.Min, g.toGo(v.Value, references))
		}
		if v.Repeat.Min == nil && v.Repeat.Max != nil {
			return fmt.Sprintf("op.Repeat{Max: %s, Value: %s}", *v.Repeat.Max, g.toGo(v.Value, references))
		}
		return fmt.Sprintf("op.Repeat{Min: %s, Max: %s, Value: %s}", *v.Repeat.Min, *v.Repeat.Max, g.toGo(v.Value, references))
	case *ir.NumVal:
		switch []rune(*v)[0] {
		case 'x':
			x := strings.TrimPrefix(string(*v), "x")
			if strings.Contains(x, "-") {
				s := strings.Split(x, "-")
				switch len(s) {
				case 2:
					return fmt.Sprintf("op.RuneRange{Min: 0x%s, Max: 0x%s}", s[0], s[1])
				default:
					panic(fmt.Errorf("invalid range %s", *v))
				}
			}
			if strings.Contains(x, ".") {
				var xs []string
				for _, x := range strings.Split(x, ".") {
					xs = append(xs, fmt.Sprintf("rune(0x%s)", x))
				}
				return fmt.Sprintf("op.And{%s}", strings.Join(xs, ", "))
			}
			return fmt.Sprintf("rune(0x%s)", x)
		}
		panic(fmt.Errorf("invalid range %s", *v))
	case *ir.CharVal:
		s := strings.TrimPrefix(strings.TrimSuffix(string(*v), "\""), "\"")
		switch len(s) {
		case 0:
			return ""
		case 1:
			if s == "\\" {
				return "'\\\\'"
			}
			if s == "'" {
				return "'\\''"
			}
			return fmt.Sprintf("'%s'", s)
		default:
			return string(*v)
		}
	case *ir.Option:
		return fmt.Sprintf("op.Optional{Value: %s}", g.toGo(v.Alternation, references))
	case *ir.Rulename:
		name := g.ruleName(string(*v))
		var ref bool
		for _, r := range references {
			if r == name {
				ref = true
				break
			}
		}
		if ref {
			return fmt.Sprintf("op.Reference{Name: %q}", name)
		}
		if _, ok := g.CustomOperators[string(*v)]; ok {
			log.Printf("[CUSTOM] %q.", string(*v))
			return fmt.Sprintf("%sOperator{}", name)
		}
		if e, ok := g.ExternalDependencies[string(*v)]; ok {
			g.usedExternalDependencies[e.Path] = struct{}{}
			log.Printf("[DEPENDENCY] %q", name)
			return fmt.Sprintf("%s.%s", e.Name, name)
		}
		return name
	case *ir.ProseVal:
		name := string(*v)
		log.Printf("[PROSE VAL] %q.", g.ruleName(name))
		rname := ir.Rulename(name[1 : len(name)-1])
		return g.toGo(&rname, references)
	default:
		panic(fmt.Errorf("unsupported type %T", v))
	}
}

type abnfRule struct {
	Name     string
	Operator string
}

type abnfTemplateData struct {
	PackageName  string
	ImportCore   bool
	Dependencies []string
	References   []string
	Rules        []abnfRule
}
