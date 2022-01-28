package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var HandlerAnalyzer = &analysis.Analyzer{
	Name:     "extractorlint",
	Doc:      "Checks and validates ElementHandler defs.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.GenDecl)(nil),
	}

	insp.Preorder(nodeFilter, func(node ast.Node) {
		genDecl := node.(*ast.GenDecl)
		issues := lintHandler(genDecl)
		for _, issue := range issues {
			pass.Report(issue.Diagnose())
		}
	})

	return nil, nil //nolint:nilnil
}
