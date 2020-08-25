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

	func hoge() {}

	func Fuga() {}

	func unExported() {}

	func Exported() {}

	func main() {
		return
	}
	`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "sample.go", src, 0)
	if err != nil {
		log.Fatalf("parse file failed: %v", err)
	}
	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.FuncDecl:
			if d.Recv != nil {
				continue
			}

			if !d.Name.IsExported() {
				fmt.Println(d.Name)
			}
		}
	}
}
