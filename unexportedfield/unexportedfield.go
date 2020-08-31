package unexportedfield

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "unexportedfield is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "unexportedfield",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.StructType:
			for _, f := range n.Fields.List {
				if f.Tag == nil {
					continue
				}
				for _, name := range f.Names {
					if !name.IsExported() {
						pass.Reportf(f.Pos(), "%s field is unexported though it has a json tag", name.String())
					}
				}
			}
		}
	})

	return nil, nil
}
