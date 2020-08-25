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
	f, err := parser.ParseFile(fset, "struct.go", nil, 0)
	if err != nil {
		log.Fatalf("parse file failed: %v", err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.StructType:
			for _, field := range n.Fields.List {
				if field.Tag != nil {
					for _, f := range field.Names {
						if !f.IsExported() {
							fmt.Println("(> <)")
						}
					}
				}
			}
		}
		return true
	})

}
