package pointer

import "github.com/hidori/go-tools/types"

func Of[T any](v T) *T {
	return &v
}

func ValueOrDefault[T any](p *T, _default T) T {
	if p == nil {
		return _default
	}

	return *p
}

func ValueOrEmpty[T any](p *T) T {
	return ValueOrDefault(p, types.Empty[T]())
}
