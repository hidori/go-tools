package ptr

func Pointer[T any](v T) *T {
	return &v
}

func ValueOrDefault[T any](p *T, _default T) T {
	if p == nil {
		return _default
	}

	return *p
}
