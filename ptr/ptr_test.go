package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "success: returns pointer",
			args: args{
				v: 1,
			},
			want: func(v int) *int {
				return &v
			}(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pointer(tt.args.v)
			if !assert.Equal(t, *tt.want, *got) {
				return
			}
		})
	}
}

func TestValueOrDefault(t *testing.T) {
	type args struct {
		p        *int
		_default int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success: returns value",
			args: args{
				p: func(v int) *int {
					return &v
				}(1),
				_default: 10,
			},
			want: 1,
		},
		{
			name: "success: returns default",
			args: args{
				p:        nil,
				_default: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValueOrDefault(tt.args.p, tt.args._default)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
