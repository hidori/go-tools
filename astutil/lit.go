package astutil

import (
	"go/ast"
	"go/token"
)

func NewBasicLit(kind token.Token, value string) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: value,
	}
}

func NewCompositeLit(litType ast.Expr, elts []ast.Expr) *ast.CompositeLit {
	return &ast.CompositeLit{
		Type: litType,
		Elts: elts,
	}
}
