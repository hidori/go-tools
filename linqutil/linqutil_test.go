package linqutil

import (
	"testing"

	"github.com/makiuchi-d/linq/v2"
	"github.com/stretchr/testify/assert"
)

func TestPassThrough(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success: returns 1",
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PassThrough(tt.args.v)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestAsOrEmptyInt(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success: returns 1",
			args: args{
				v: interface{}(1),
			},
			want: 1,
		},
		{
			name: "success: returns 0",
			args: args{
				v: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AsOrEmpty[int](tt.args.v)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestAppend(t *testing.T) {
	type args struct {
		source []int
		v      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success: returns [1, 2, 3]",
			args: args{
				source: []int{1, 2},
				v:      3,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_source := linq.FromSlice(tt.args.source)
			got := Append(_source, tt.args.v)

			_got, err := linq.ToSlice(got)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, _got) {
				return
			}
		})
	}
}
