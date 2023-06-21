package types

import "reflect"

func EmptyOf[T any]() T {
	var v T

	rt := reflect.TypeOf(v)
	rv := reflect.Zero(rt)
	e, _ := rv.Interface().(T)

	return e
}
