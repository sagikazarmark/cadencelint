package determinism

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `check Cadence workflows for non-deterministic code

The deterministic checker walks Cadence workflows
ensuring that the code is deterministic.

See https://cadenceworkflow.io/docs/06_goclient/02_create_workflows#implementation for more details.`

// nolint: gochecknoglobals,golint
var Analyzer = &analysis.Analyzer{
	Name:     "deterministic",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		function := n.(*ast.FuncDecl)

		// Filter cadence workflow functions
		if len(function.Type.Params.List) < 1 ||
			pass.TypesInfo.TypeOf(function.Type.Params.List[0].Type).String() != "go.uber.org/cadence/internal.Context" {
			return
		}

		ast.Inspect(function.Body, func(node ast.Node) bool {
			if forStmt, ok := node.(*ast.RangeStmt); ok {
				if _, ok := pass.TypesInfo.TypeOf(forStmt.X).(*types.Map); ok {
					pass.Reportf(forStmt.X.Pos(), "workflows must not range over maps")
				}
			}

			if call, ok := node.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if id, ok := sel.X.(*ast.Ident); ok && id.Name == "time" && (sel.Sel.Name == "Now" || sel.Sel.Name == "Sleep") {
						pass.Reportf(call.Fun.Pos(), "workflows must not call time.Now and time.Sleep")
					}
				}
			}

			if stmt, ok := node.(*ast.GoStmt); ok {
				pass.Reportf(stmt.Pos(), "workflows must not start goroutines")
			}

			if stmt, ok := node.(*ast.SelectStmt); ok {
				pass.Reportf(stmt.Pos(), "workflows must not select directly from channels")
			}

			if stmt, ok := node.(*ast.ChanType); ok {
				pass.Reportf(stmt.Pos(), "workflows must not create channels directly")
			}

			return true
		})
	})

	return nil, nil
}
