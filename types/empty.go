package types

func EmptyOf[T any]() T {
	var v T

	return v
}
