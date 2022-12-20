package astutil

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockStmt(t *testing.T) {
	type args struct {
		list []ast.Stmt
	}
	tests := []struct {
		name string
		args args
		want *ast.BlockStmt
	}{
		{
			name: "success: returns *ast.BlockStmt",
			args: args{
				list: []ast.Stmt{
					NewReturnStmt([]ast.Expr{
						NewSelectorExpr(NewIdent("t"), NewIdent("n")),
					}),
				},
			},
			want: &ast.BlockStmt{
				List: []ast.Stmt{
					NewReturnStmt([]ast.Expr{
						NewSelectorExpr(NewIdent("t"), NewIdent("n")),
					}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBlockStmt(tt.args.list)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestNewReturnStmt(t *testing.T) {
	type args struct {
		results []ast.Expr
	}
	tests := []struct {
		name string
		args args
		want *ast.ReturnStmt
	}{
		{
			name: "success: returns *ast.ReturnStmt",
			args: args{
				results: []ast.Expr{
					NewSelectorExpr(NewIdent("t"), NewIdent("n")),
				},
			},
			want: &ast.ReturnStmt{
				Results: []ast.Expr{
					NewSelectorExpr(NewIdent("t"), NewIdent("n")),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewReturnStmt(tt.args.results)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestNewAssignStmt(t *testing.T) {
	type args struct {
		lhs []ast.Expr
		tok token.Token
		rhs []ast.Expr
	}
	tests := []struct {
		name string
		args args
		want *ast.AssignStmt
	}{
		{
			name: "success: returns *ast.AssignStmt",
			args: args{
				lhs: []ast.Expr{
					NewSelectorExpr(NewIdent("t"), NewIdent("n")),
				},
				tok: token.ASSIGN,
				rhs: []ast.Expr{
					NewIdent("v"),
				},
			},
			want: &ast.AssignStmt{
				Lhs: []ast.Expr{
					NewSelectorExpr(NewIdent("t"), NewIdent("n")),
				},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{
					NewIdent("v"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAssignStmt(tt.args.lhs, tt.args.tok, tt.args.rhs)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
