package astutil

import "go/ast"

func NewArrayType(elt ast.Expr) *ast.ArrayType {
	return &ast.ArrayType{
		Elt: elt,
	}
}

func NewFuncType(typeParams *ast.FieldList, params *ast.FieldList, results *ast.FieldList) *ast.FuncType {
	return &ast.FuncType{
		TypeParams: typeParams,
		Params:     params,
		Results:    results,
	}
}
