package astutil

import "go/ast"

func NewField(names []*ast.Ident, fieldType ast.Expr) *ast.Field {
	return &ast.Field{
		Names: names,
		Type:  fieldType,
	}
}

func NewFieldList(list []*ast.Field) *ast.FieldList {
	return &ast.FieldList{
		List: list,
	}
}
