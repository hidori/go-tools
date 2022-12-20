package astutil

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBasicLit(t *testing.T) {
	type args struct {
		kind  token.Token
		value string
	}
	tests := []struct {
		name string
		args args
		want *ast.BasicLit
	}{
		{
			name: "success: returns *ast.BasicLit",
			args: args{
				kind:  token.STRING,
				value: "value",
			},
			want: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBasicLit(tt.args.kind, tt.args.value)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestNewCompositeLit(t *testing.T) {
	type args struct {
		litType ast.Expr
		elts    []ast.Expr
	}
	tests := []struct {
		name string
		args args
		want *ast.CompositeLit
	}{
		{
			name: "success: returns *ast.CompositeLit",
			args: args{
				litType: &ast.ArrayType{
					Elt: NewIdent("string"),
				},
				elts: []ast.Expr{
					NewIdent("name"),
				},
			},
			want: &ast.CompositeLit{
				Type: &ast.ArrayType{
					Elt: NewIdent("string"),
				},
				Elts: []ast.Expr{
					NewIdent("name"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCompositeLit(tt.args.litType, tt.args.elts)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
