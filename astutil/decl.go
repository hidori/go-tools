package astutil

import (
	"go/ast"
	"go/token"
)

func NewGenDecl(tok token.Token, specs []ast.Spec) *ast.GenDecl {
	return &ast.GenDecl{
		Tok:   tok,
		Specs: specs,
	}
}

func NewFuncDecl(recv *ast.FieldList, name *ast.Ident, funcType *ast.FuncType, body *ast.BlockStmt) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: recv,
		Name: name,
		Type: funcType,
		Body: body,
	}
}
