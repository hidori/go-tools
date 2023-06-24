package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInt(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "success: returns 0",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[int]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyIntPtr(t *testing.T) {
	tests := []struct {
		name string
		want *int
	}{
		{
			name: "success: returns nil",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[*int]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success: returns \"\"",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[string]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyStringPtr(t *testing.T) {
	tests := []struct {
		name string
		want *string
	}{
		{
			name: "success: nil",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[*string]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyStruct(t *testing.T) {
	type TestStruct struct{}

	tests := []struct {
		name string
		want TestStruct
	}{
		{
			name: "success: returns TestStruct{}",
			want: TestStruct{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[TestStruct]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyStructPtr(t *testing.T) {
	type TestStruct struct{}

	tests := []struct {
		name string
		want *TestStruct
	}{
		{
			name: "success: nil",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[*TestStruct]()
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
			name: "success: return 1",
			args: args{
				v: interface{}(1),
			},
			want: 1,
		},
		{
			name: "success: return 0",
			args: args{
				v: interface{}("1"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AsOrEmpty[int](tt.args.v)
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
