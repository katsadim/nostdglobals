package analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	Analyzer = &analysis.Analyzer{
		Name:     "nostdglobals",
		Doc:      "Checks no global variables that are potentially dangerous are used.",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	// restrictedGlobals is a package to a list of variables map. For example the following map
	// restricts http.DefaultClient and http.DefaultTransport. The report checks are done against this map.
	restrictedGlobals = map[string][]string{
		"http": {"DefaultClient", "DefaultTransport"},
	}
)

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.SelectorExpr)(nil),
	}

	insp.Preorder(nodeFilter, func(node ast.Node) {
		selExpr, _ := node.(*ast.SelectorExpr)

		x, ok := selExpr.X.(*ast.Ident)
		if !ok {
			return
		}

		varNames, found := restrictedGlobals[x.Name]
		if !found {
			return
		}

		for _, varName := range varNames {
			if varName == selExpr.Sel.Name {
				pass.Reportf(node.Pos(), "should not make use of '%s.%s'", x.Name, varName)
			}
		}
	})

	return nil, nil
}
