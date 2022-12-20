package astutil

import "go/ast"

func NewIdent(name string) *ast.Ident {
	return ast.NewIdent(name)
}
