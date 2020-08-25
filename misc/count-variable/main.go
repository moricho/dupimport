package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	src := `
	package main
	var v = 100
	var v = 10
	var v2 = 10
	const v3 = 100
	func main() {
		a := v + 1
		b := v + a
		a = 100
	}
	`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "sample.go", src, 0)
	if err != nil {
		log.Fatalf("parse file failed: %v", err)
	}

	// variables := map[string]bool{}
	cnt := 0
	variables := map[string]bool{}
	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.GenDecl:
			if d.Tok != token.VAR {
				continue
			}
			for _, spec := range d.Specs {
				switch spec := spec.(type) {
				case *ast.ValueSpec:
					for _, v := range spec.Names {
						if _, ok := variables[v.Name]; ok {
							continue
						} else {
							cnt++
							variables[v.Name] = true
						}
					}
				}
			}
		}
	}

	fmt.Println(cnt)
}
