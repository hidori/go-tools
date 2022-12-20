package astutil

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenDecl(t *testing.T) {
	type args struct {
		tok   token.Token
		specs []ast.Spec
	}
	tests := []struct {
		name string
		args args
		want *ast.GenDecl
	}{
		{
			name: "success: returns *ast.GenDecl",
			args: args{
				tok: token.CONST,
				specs: []ast.Spec{
					NewValueSpec(
						[]*ast.Ident{
							NewIdent("name"),
						},
						[]ast.Expr{
							NewBasicLit(token.STRING, strconv.Quote("value")),
						},
					),
				},
			},
			want: &ast.GenDecl{
				Tok: token.CONST,
				Specs: []ast.Spec{
					NewValueSpec(
						[]*ast.Ident{
							NewIdent("name"),
						},
						[]ast.Expr{
							NewBasicLit(token.STRING, strconv.Quote("value")),
						},
					),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenDecl(tt.args.tok, tt.args.specs)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestNewFuncDecl(t *testing.T) {
	type args struct {
		recv     *ast.FieldList
		name     *ast.Ident
		funcType *ast.FuncType
		body     *ast.BlockStmt
	}
	tests := []struct {
		name string
		args args
		want *ast.FuncDecl
	}{
		{
			name: "success: returns *ast.FuncDecl",
			args: args{
				recv: NewFieldList(
					[]*ast.Field{
						NewField(
							[]*ast.Ident{
								NewIdent("t"),
							},
							NewStarExpr(NewIdent("TestStruct")),
						),
					},
				),
				name: NewIdent("Set"),
				funcType: NewFuncType(
					nil,
					NewFieldList([]*ast.Field{
						NewField(
							[]*ast.Ident{
								ast.NewIdent("in"),
							},
							NewIdent("int"),
						),
					}),
					NewFieldList([]*ast.Field{
						NewField(
							[]*ast.Ident{
								NewIdent("out"),
							},
							NewIdent("int")),
					}),
				),
				body: NewBlockStmt([]ast.Stmt{
					NewAssignStmt(
						[]ast.Expr{
							NewIdent("l"),
						},
						token.ASSIGN,
						[]ast.Expr{
							NewSelectorExpr(NewIdent("t"), NewIdent("f")),
						},
					),
					NewAssignStmt(
						[]ast.Expr{
							NewSelectorExpr(NewIdent("t"), NewIdent("f")),
						},
						token.ASSIGN,
						[]ast.Expr{
							NewIdent("v"),
						},
					),
					NewReturnStmt([]ast.Expr{
						NewIdent("l"),
					}),
				}),
			},
			want: &ast.FuncDecl{
				Recv: NewFieldList(
					[]*ast.Field{
						NewField(
							[]*ast.Ident{
								NewIdent("t"),
							},
							NewStarExpr(NewIdent("TestStruct")),
						),
					},
				),
				Name: NewIdent("Set"),
				Type: NewFuncType(
					nil,
					NewFieldList([]*ast.Field{
						NewField(
							[]*ast.Ident{
								ast.NewIdent("in"),
							},
							NewIdent("int"),
						),
					}),
					NewFieldList([]*ast.Field{
						NewField(
							[]*ast.Ident{
								NewIdent("out"),
							},
							NewIdent("int")),
					}),
				),
				Body: NewBlockStmt([]ast.Stmt{
					NewAssignStmt(
						[]ast.Expr{
							NewIdent("l"),
						},
						token.ASSIGN,
						[]ast.Expr{
							NewSelectorExpr(NewIdent("t"), NewIdent("f")),
						},
					),
					NewAssignStmt(
						[]ast.Expr{
							NewSelectorExpr(NewIdent("t"), NewIdent("f")),
						},
						token.ASSIGN,
						[]ast.Expr{
							NewIdent("v"),
						},
					),
					NewReturnStmt([]ast.Expr{
						NewIdent("l"),
					}),
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFuncDecl(tt.args.recv, tt.args.name, tt.args.funcType, tt.args.body)
			if !assert.Equal(t, tt.want, got) {
				return
			}

			format.Node(os.Stdout, token.NewFileSet(), got)
		})
	}
}
