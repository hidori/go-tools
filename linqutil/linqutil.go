package linqutil

import (
	"github.com/hidori/go-tools/types"
	"github.com/makiuchi-d/linq/v2"
)

func PassThrough[T any](v T) (T, error) {
	return v, nil
}

func AsOrEmpty[T any](v any) (T, error) {
	return types.AsOrEmpty[T](v), nil
}

func Append[T any](source linq.Enumerable[T], v T) linq.Enumerable[T] {
	return linq.Concat(source, linq.FromSlice([]T{v}))
}
