package astutil

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIdent(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *ast.Ident
	}{
		{
			name: "success: returns *ast.Ident",
			args: args{
				name: "name",
			},
			want: ast.NewIdent("name"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIdent(tt.args.name)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
