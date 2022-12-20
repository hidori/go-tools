package astutil

import (
	"go/ast"
	"go/token"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValueSpec(t *testing.T) {
	type args struct {
		names  []*ast.Ident
		values []ast.Expr
	}
	tests := []struct {
		name string
		args args
		want *ast.ValueSpec
	}{
		{
			name: "success: returns *ast.ValueSpec",
			args: args{
				names: []*ast.Ident{
					NewIdent("name"),
				},
				values: []ast.Expr{
					NewBasicLit(token.STRING, strconv.Quote("value")),
				},
			},
			want: &ast.ValueSpec{
				Names: []*ast.Ident{
					NewIdent("name"),
				},
				Values: []ast.Expr{
					NewBasicLit(token.STRING, strconv.Quote("value")),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewValueSpec(tt.args.names, tt.args.values)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
