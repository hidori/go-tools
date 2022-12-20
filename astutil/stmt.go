package astutil

import (
	"go/ast"
	"go/token"
)

func NewBlockStmt(list []ast.Stmt) *ast.BlockStmt {
	return &ast.BlockStmt{
		List: list,
	}
}

func NewReturnStmt(results []ast.Expr) *ast.ReturnStmt {
	return &ast.ReturnStmt{
		Results: results,
	}
}

func NewAssignStmt(lhs []ast.Expr, tok token.Token, rhs []ast.Expr) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: lhs,
		Tok: tok,
		Rhs: rhs,
	}
}
