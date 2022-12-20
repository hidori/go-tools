package genfldnam

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"testing"

	"github.com/hidori/go-tools/must"
	"github.com/stretchr/testify/assert"
)

const tagName = "genfldnam"

func TestGenerator_Generate(t *testing.T) {
	type fields struct {
		config *GeneratorConfig
	}
	tests := []struct {
		name           string
		input          string
		output         string
		fields         fields
		wantErr        bool
		wantErrMessage string
	}{
		{
			name:   "success: returns []ast.Decl",
			input:  "../test/data/genfldnam/success_input.go",
			output: "../test/data/genfldnam/success_output.txt",
			fields: fields{
				config: &GeneratorConfig{
					TagName:        "genfldnam",
					Skewer:         "FieldName",
					AllNames:       true,
					AllNamesSuffix: "AllFieldNames",
				},
			},
		},
		{
			name:  "fail: returns []ast.Decl",
			input: "../test/data/genfldnam/fail_input.go",
			fields: fields{
				config: &GeneratorConfig{
					TagName:        "genfldnam",
					Skewer:         "FieldName",
					AllNames:       true,
					AllNamesSuffix: "AllFieldNames",
				},
			},
			wantErr:        true,
			wantErrMessage: "invalid tag value 'undefined'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fset := token.NewFileSet()
			f := must.Must(func() (*ast.File, error) {
				return parser.ParseFile(token.NewFileSet(), tt.input, nil, parser.AllErrors)
			})

			got, err := NewGenerator(tt.fields.config).Generate(fset, f)
			if err != nil && tt.wantErr {
				assert.Contains(t, err.Error(), tt.wantErrMessage)
				return
			}

			if !assert.NoError(t, err) {
				return
			}

			_got := bytes.NewBuffer([]byte{})
			format.Node(_got, fset, got)

			_want := bytes.NewBuffer([]byte{})
			func() {
				f := must.Must(func() (*ast.File, error) {
					return parser.ParseFile(token.NewFileSet(), tt.output, nil, parser.AllErrors)
				})
				format.Node(_want, fset, f.Decls)
			}()

			if !assert.Equal(t, _want.String(), _got.String()) {
				return
			}
		})
	}
}
