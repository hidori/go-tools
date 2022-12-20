package astutil

import "go/ast"

func NewValueSpec(names []*ast.Ident, values []ast.Expr) *ast.ValueSpec {
	return &ast.ValueSpec{
		Names:  names,
		Values: values,
	}
}
