package openapi

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hidori/go-tools/pointer"
	"github.com/stretchr/testify/assert"
)

func Test_validateNumber(t *testing.T) {
	type args struct {
		schema *openapi3.Schema
		name   string
		value  reflect.Value
	}
	tests := []struct {
		name string
		args args
		want []*ValidationIssue
	}{
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					Min: pointer.Of(float64(1)),
					Max: pointer.Of(float64(3)),
				},
				name:  "x",
				value: reflect.ValueOf(2),
			},
			want: nil,
		},
		{
			name: "success: returns ValidationIssueCodeMin",
			args: args{
				schema: &openapi3.Schema{
					Min: pointer.Of(float64(1)),
					Max: pointer.Of(float64(3)),
				},
				name:  "x",
				value: reflect.ValueOf(0),
			},
			want: []*ValidationIssue{
				{
					Code:     ValidationIssueCodeMin,
					Name:     "x",
					Value:    reflect.ValueOf(0),
					Expected: float64(1),
					Actual:   float64(0),
				},
			},
		},
		{
			name: "success: returns ValidationIssueCodeMax",
			args: args{
				schema: &openapi3.Schema{
					Min: pointer.Of(float64(1)),
					Max: pointer.Of(float64(3)),
				},
				name:  "x",
				value: reflect.ValueOf(4),
			},
			want: []*ValidationIssue{
				{
					Code:     ValidationIssueCodeMax,
					Name:     "x",
					Value:    reflect.ValueOf(4),
					Expected: float64(3),
					Actual:   float64(4),
				},
			},
		},
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					Enum: []interface{}{1, 2, 3},
				},
				name:  "x",
				value: reflect.ValueOf(2),
			},
			want: nil,
		},
		{
			name: "success: returns ValidationIssueCodeEnum",
			args: args{
				schema: &openapi3.Schema{
					Enum: []interface{}{1, 2, 3},
				},
				name:  "x",
				value: reflect.ValueOf(4),
			},
			want: []*ValidationIssue{
				{
					Code:     ValidationIssueCodeEnum,
					Name:     "x",
					Value:    reflect.ValueOf(4),
					Expected: []interface{}{1, 2, 3},
					Actual:   reflect.ValueOf(4),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&validator{}).validateNumber(tt.args.schema, tt.args.name, tt.args.value)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func Test_validateString(t *testing.T) {
	type args struct {
		schema *openapi3.Schema
		name   string
		value  reflect.Value
	}
	tests := []struct {
		name string
		args args
		want []*ValidationIssue
	}{
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					MinLength: 1,
					MaxLength: pointer.Of(uint64(3)),
				},
				name:  "x",
				value: reflect.ValueOf("AB"),
			},
			want: nil,
		},
		{
			name: "success: returns ValidationIssueCodeMinLength",
			args: args{
				schema: &openapi3.Schema{
					MinLength: 1,
					MaxLength: pointer.Of(uint64(3)),
				},
				name:  "x",
				value: reflect.ValueOf(""),
			},
			want: []*ValidationIssue{
				{
					Code:     ValidationIssueCodeMinLength,
					Name:     "x",
					Value:    reflect.ValueOf(""),
					Expected: uint64(1),
					Actual:   uint64(0),
				},
			},
		},
		{
			name: "success: returns ValidationIssueCodeMaxLength",
			args: args{
				schema: &openapi3.Schema{
					MinLength: 1,
					MaxLength: pointer.Of(uint64(3)),
				},
				name:  "x",
				value: reflect.ValueOf("ABCD"),
			},
			want: []*ValidationIssue{
				{
					Code:     ValidationIssueCodeMaxLength,
					Name:     "x",
					Value:    reflect.ValueOf("ABCD"),
					Expected: uint64(3),
					Actual:   uint64(4),
				},
			},
		},
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					Pattern: `ABCD`,
				},
				name:  "x",
				value: reflect.ValueOf("ABCD"),
			},
			want: nil,
		},
		{
			name: "success: returns ValidationIssueCodePattern",
			args: args{
				schema: &openapi3.Schema{
					Pattern: `abcd`,
				},
				name:  "x",
				value: reflect.ValueOf("ABCD"),
			},
			want: []*ValidationIssue{
				{
					Code:     ValidationIssueCodePattern,
					Name:     "x",
					Value:    reflect.ValueOf("ABCD"),
					Expected: `abcd`,
					Actual:   "ABCD",
				},
			},
		},
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					Enum: []interface{}{"A", "B", "C"},
				},
				name:  "x",
				value: reflect.ValueOf("B"),
			},
			want: nil,
		},
		{
			name: "success: returns ValidationIssueCodeEnum",
			args: args{
				schema: &openapi3.Schema{
					Enum: []interface{}{"A", "B", "C"},
				},
				name:  "x",
				value: reflect.ValueOf("D"),
			},
			want: []*ValidationIssue{
				{
					Code:     ValidationIssueCodeEnum,
					Name:     "x",
					Value:    reflect.ValueOf("D"),
					Expected: []interface{}{"A", "B", "C"},
					Actual:   reflect.ValueOf("D"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&validator{}).validateString(tt.args.schema, tt.args.name, tt.args.value)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func Test_validateEnum(t *testing.T) {
	type args struct {
		schema *openapi3.Schema
		name   string
		value  reflect.Value
	}
	tests := []struct {
		name           string
		args           args
		want           []*ValidationIssue
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "fail: returns error",
			args: args{
				schema: &openapi3.Schema{
					Enum: []interface{}{"123", "456", "789"},
				},
				name:  "x",
				value: reflect.ValueOf(456),
			},
			wantErr:        true,
			wantErrMessage: "value of type string cannot be converted to type int",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&validator{}).validateEnum(tt.args.schema, tt.args.name, tt.args.value)
			if err != nil && tt.wantErr {
				assert.Contains(t, err.Error(), tt.wantErrMessage)
				return
			}

			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func Test_validatePointer(t *testing.T) {
	type args struct {
		schema *openapi3.Schema
		name   string
		value  reflect.Value
	}
	tests := []struct {
		name string
		args args
		want []*ValidationIssue
	}{
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					Nullable: true,
				},
				name: "x",
				value: func() reflect.Value {
					v := "A"
					return reflect.ValueOf(&v)
				}(),
			},
			want: nil,
		},
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					Nullable: true,
				},
				name: "x",
				value: func() reflect.Value {
					var v *string
					return reflect.ValueOf(v)
				}(),
			},
			want: nil,
		},
		{
			name: "success: returns ValidationIssueCodeNullable",
			args: args{
				schema: &openapi3.Schema{
					Nullable: false,
				},
				name: "x",
				value: func() reflect.Value {
					var v *string
					return reflect.ValueOf(v)
				}(),
			},
			want: []*ValidationIssue{
				{
					Code: ValidationIssueCodeNullable,
					Name: "x",
					Value: func() reflect.Value {
						var v *string
						return reflect.ValueOf(v)
					}(),
					Expected: false,
					Actual:   true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&validator{}).validatePointer(tt.args.schema, tt.args.name, tt.args.value)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func Test_validateSlice(t *testing.T) {
	type args struct {
		schema *openapi3.Schema
		name   string
		value  reflect.Value
	}
	tests := []struct {
		name string
		args args
		want []*ValidationIssue
	}{
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					Items: &openapi3.SchemaRef{
						Value: &openapi3.Schema{},
					},
				},
				name:  "x",
				value: reflect.ValueOf([]int{1, 2, 3}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&validator{}).validateSlice(tt.args.schema, tt.args.name, tt.args.value)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func Test_validateStruct(t *testing.T) {
	type testStruct struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
	}

	type args struct {
		schema *openapi3.Schema
		name   string
		value  reflect.Value
	}
	tests := []struct {
		name string
		args args
		want []*ValidationIssue
	}{
		{
			name: "success: returns nil",
			args: args{
				schema: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"field1": &openapi3.SchemaRef{
							Value: &openapi3.Schema{},
						},
						"field2": &openapi3.SchemaRef{
							Value: &openapi3.Schema{},
						},
					},
				},
				name: "x",
				value: reflect.ValueOf(testStruct{
					Field1: 1,
					Field2: 2,
				}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&validator{}).validateStruct(tt.args.schema, tt.args.name, tt.args.value)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func Test_getStructField(t *testing.T) {
	type testStruct struct {
		IntField int `json:"int_field"`
	}

	type args struct {
		rv           reflect.Value
		propertyName string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantNil bool
	}{
		{
			name: "success: returns *reflect.StructField",
			args: args{
				rv:           reflect.ValueOf(testStruct{IntField: 123}),
				propertyName: "int_field",
			},
			want: 123,
		},
		{
			name: "success: returns nil",
			args: args{
				rv:           reflect.ValueOf(testStruct{IntField: 123}),
				propertyName: "not_there",
			},
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := (&validator{}).getStructField(tt.args.rv, tt.args.propertyName)
			if got == nil && tt.wantNil {
				return
			}

			_got := tt.args.rv.FieldByName(got.Name).Int()
			if !assert.Equal(t, tt.want, _got) {
				return
			}
		})
	}
}
