package dupimport

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "dupimport is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "dupimport",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		importsMap := map[string]bool{}
		for _, is := range f.Imports {
			path, err := strconv.Unquote(is.Path.Value)
			if err != nil {
				return nil, err
			}

			if _, ok := importsMap[path]; ok {
				pass.Reportf(is.Pos(), "%s package is duplicated", path)
			} else {
				importsMap[path] = true
			}
		}
	}

	return nil, nil
}
