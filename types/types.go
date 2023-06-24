package types

func Empty[T any]() T {
	var v T

	return v
}

func AsOrEmpty[T any](v any) T {
	t, ok := v.(T)
	if !ok {
		return Empty[T]()
	}

	return t
}
