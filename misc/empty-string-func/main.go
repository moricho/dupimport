package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"
)

func main() {
	src := `
	package main

	func hoge() string {
		return ""
	}

	func fuga() {}

	func main() {
		res := hoge()
		fuga()
		fmt.Println(res)
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
			for _, stmt := range d.Body.List {
				switch stmt := stmt.(type) {
				case *ast.ReturnStmt:
					for _, v := range stmt.Results {
						switch v := v.(type) {
						case *ast.BasicLit:
							res, err := strconv.Unquote(v.Value)
							if err != nil {
								log.Fatalf("unquote failed: %v", err)
							}
							if res == "" {
								fmt.Println("function which returns empty string exists")
							}
						}
					}
				}
			}
		}
	}
}
