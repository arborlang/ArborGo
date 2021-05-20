package main

import (
	"flag"
	"fmt"
	"go/types"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	typeNames    = flag.String("type", "", "comma-separated list of type names; must be set")
	templatePath = flag.String("template", "default.tmpl", "the template to generate")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of generate:\n")
	fmt.Fprintf(os.Stderr, "\tstringer [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tstringer [flags] -type T files... # Must be a single package\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

var visitorTmpl = `{{range .}}
type {{ . }}Visitor interface {
	Visit{{ . }}(n *{{ . }}) (Node, error)
}
{{ end }}

type GenericVisitor interface {
	VisitAnyNode(n Node) (Node, error)
}

type Visitor interface {
{{ range . }}	{{ . }}Visitor
{{ end }}
	GenericVisitor
}
`

func main() {
	flag.Usage = Usage
	flag.Parse()

	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo, Tests: false}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		log.Println("error reading package:", err)
		os.Exit(-1)
	}
	satisfied := []string{}
	packageName := pkgs[0].Types.Name()
	log.Println("looking for Interface", *typeNames, "in", packageName)
	for _, pkg := range pkgs {
		// Look for our interface in a package
		obj := pkg.Types.Scope().Lookup(*typeNames)
		if obj == nil {
			continue
		}
		iFace, ok := obj.Type().Underlying().(*types.Interface)
		if !ok {
			continue
		}

		for _, name := range pkg.Types.Scope().Names() {
			obj := pkg.Types.Scope().Lookup(name)

			ptr := types.NewPointer(obj.Type())
			imp := types.Implements(ptr.Underlying(), iFace)
			if imp {
				satisfied = append(satisfied, obj.Name())
			}
		}
	}
	visitorsBuilder := &strings.Builder{}
	vTmpl, err := template.New("VisitorTmpl").Parse(visitorTmpl)
	if err != nil {
		log.Fatalln(err)
	}
	visitorsBuilder.WriteString(fmt.Sprintf("package %s\n", packageName))
	visitorsBuilder.WriteString("\n//Written by the generator, do not over write\n")
	err = vTmpl.Execute(visitorsBuilder, satisfied)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s_visitors.go", packageName), []byte(visitorsBuilder.String()), 0744)
}
