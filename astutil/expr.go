package astutil

import "go/ast"

func NewStarExpr(x ast.Expr) *ast.StarExpr {
	return &ast.StarExpr{
		X: x,
	}
}

func NewSelectorExpr(x ast.Expr, sel *ast.Ident) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   x,
		Sel: sel,
	}
}
