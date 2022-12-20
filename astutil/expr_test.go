package astutil

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStarExpr(t *testing.T) {
	type args struct {
		x ast.Expr
	}
	tests := []struct {
		name string
		args args
		want *ast.StarExpr
	}{
		{
			name: "success: returns *ast.StarExpr",
			args: args{
				x: NewIdent("t"),
			},
			want: &ast.StarExpr{
				X: NewIdent("t"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStarExpr(tt.args.x)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestNewSelectorExpr(t *testing.T) {
	type args struct {
		x   ast.Expr
		sel *ast.Ident
	}
	tests := []struct {
		name string
		args args
		want *ast.SelectorExpr
	}{
		{
			name: "success: returns *ast.SelectorExpr",
			args: args{
				x:   NewIdent("x"),
				sel: NewIdent("f"),
			},
			want: &ast.SelectorExpr{
				X:   NewIdent("x"),
				Sel: NewIdent("f"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSelectorExpr(tt.args.x, tt.args.sel); !assert.Equal(t, tt.want, got) {
				t.Errorf("NewSelectorExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
