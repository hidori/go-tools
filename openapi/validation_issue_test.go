package openapi

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidationIssue(t *testing.T) {
	type args struct {
		code     ValidationIssueCode
		name     string
		value    reflect.Value
		expected interface{}
		actual   interface{}
	}
	tests := []struct {
		name string
		args args
		want *ValidationIssue
	}{
		{
			name: "success: returns *ValidationIssue",
			args: args{
				code:     ValidationIssueCodeRequired,
				name:     "name",
				value:    reflect.ValueOf(int(123)),
				expected: "want",
				actual:   "got",
			},
			want: &ValidationIssue{
				Code:     ValidationIssueCodeRequired,
				Name:     "name",
				Value:    reflect.ValueOf(int(123)),
				Expected: "want",
				Actual:   "got",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewValidationIssue(tt.args.code, tt.args.name, tt.args.value, tt.args.expected, tt.args.actual)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
