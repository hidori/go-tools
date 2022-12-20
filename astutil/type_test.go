package astutil

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArrayType(t *testing.T) {
	type args struct {
		elt ast.Expr
	}
	tests := []struct {
		name string
		args args
		want *ast.ArrayType
	}{
		{
			name: "success: returns *ast.ArrayType",
			args: args{
				elt: NewIdent("name"),
			},
			want: &ast.ArrayType{
				Elt: NewIdent("name"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewArrayType(tt.args.elt)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestNewFuncType(t *testing.T) {
	type args struct {
		typeParams *ast.FieldList
		params     *ast.FieldList
		results    *ast.FieldList
	}
	tests := []struct {
		name string
		args args
		want *ast.FuncType
	}{
		{
			name: "success: returns *ast.FuncType",
			args: args{
				typeParams: nil,
				params: NewFieldList([]*ast.Field{
					NewField(
						[]*ast.Ident{
							ast.NewIdent("in"),
						},
						NewIdent("int"),
					),
				}),
				results: NewFieldList([]*ast.Field{
					NewField(
						[]*ast.Ident{
							NewIdent("out"),
						},
						NewIdent("int")),
				}),
			},
			want: &ast.FuncType{
				TypeParams: nil,
				Params: NewFieldList([]*ast.Field{
					NewField(
						[]*ast.Ident{
							ast.NewIdent("in"),
						},
						NewIdent("int"),
					),
				}),
				Results: NewFieldList([]*ast.Field{
					NewField(
						[]*ast.Ident{
							NewIdent("out"),
						},
						NewIdent("int")),
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFuncType(tt.args.typeParams, tt.args.params, tt.args.results)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
