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
	f, err := parser.ParseFile(fset, "func.go", nil, 0)
	if err != nil {
		log.Fatalf("parse file failed: %v", err)
	}

	cnt := 0
	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.FuncType:
			for _, arg := range n.Params.List {
				cnt += len(arg.Names)
			}

			if cnt >= 5 {
				fmt.Println("(^ ^)")
			}
		}
		return true
	})

}
