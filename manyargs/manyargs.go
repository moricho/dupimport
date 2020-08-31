package manyargs

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "manyargs is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "manyargs",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			cnt := 0
			for _, arg := range n.Type.Params.List {
				cnt += len(arg.Names)
			}
			if cnt >= 5 {
				pass.Reportf(n.Pos(), "%s: too many args", n.Name.String())
			}
		}
	})

	return nil, nil
}
