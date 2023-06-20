package empty

import "reflect"

func Empty[T any]() T {
	var v T

	rt := reflect.TypeOf(v)
	rv := reflect.Zero(rt)

	return rv.Interface().(T)
}
