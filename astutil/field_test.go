package astutil

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	type args struct {
		names     []*ast.Ident
		fieldType ast.Expr
	}
	tests := []struct {
		name string
		args args
		want *ast.Field
	}{
		{
			name: "success: returns *ast.Field",
			args: args{
				names: []*ast.Ident{
					NewIdent("t"),
				},
				fieldType: NewStarExpr(NewIdent("n")),
			},
			want: &ast.Field{
				Names: []*ast.Ident{
					NewIdent("t"),
				},
				Type: NewStarExpr(NewIdent("n")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewField(tt.args.names, tt.args.fieldType)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestNewFieldList(t *testing.T) {
	type args struct {
		list []*ast.Field
	}
	tests := []struct {
		name string
		args args
		want *ast.FieldList
	}{
		{
			name: "success: returns *ast.FieldList",
			args: args{
				list: []*ast.Field{
					NewField(
						[]*ast.Ident{
							NewIdent("t"),
						},
						NewStarExpr(NewIdent("n")),
					),
				},
			},
			want: &ast.FieldList{
				List: []*ast.Field{
					NewField(
						[]*ast.Ident{
							NewIdent("t"),
						},
						NewStarExpr(NewIdent("n")),
					),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFieldList(tt.args.list)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
