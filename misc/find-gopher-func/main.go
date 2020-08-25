package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "gopher.go", nil, 0)
	if err != nil {
		log.Fatalf("parse file failed: %v", err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name != "Gopher" {
				return true
			}
			if n.Obj == nil {
				return true
			}
			if n.Obj.Kind != ast.Fun {
				return true
			}

			if decl := n.Obj.Decl; decl != nil {
				switch decl := decl.(type) {
				case *ast.FuncDecl:
					if decl.Recv != nil {
						return true
					}
				}
			}

			fmt.Println("find!")
		}
		return true
	})

	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.FuncDecl:
			if d.Recv != nil {
				continue
			}

			if d.Name.String() == "Gopher" {
				fmt.Println("find!")
			}
		}
	}
}
