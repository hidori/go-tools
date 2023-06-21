package empty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyOfInt(t *testing.T) {
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
			got := Of[int]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyOfIntPtr(t *testing.T) {
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
			got := Of[*int]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyOfString(t *testing.T) {
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
			got := Of[string]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyOfStringPtr(t *testing.T) {
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
			got := Of[*string]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyOfStruct(t *testing.T) {
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
			got := Of[TestStruct]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}

func TestEmptyOfStructPtr(t *testing.T) {
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
			got := Of[*TestStruct]()
			if !assert.Equal(t, tt.want, got) {
				return
			}
		})
	}
}
