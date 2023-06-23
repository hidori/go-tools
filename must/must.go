package must

import (
	"fmt"
)

func Get1[T any](f func() (T, error)) T {
	r, err := f()
	if err != nil {
		panic(fmt.Sprintf("fail to Get1[T]() with error %s", err.Error()))
	}

	return r
}

func Get2[T1 any, T2 any](f func() (T1, T2, error)) (T1, T2) {
	r1, r2, err := f()
	if err != nil {
		panic(fmt.Sprintf("fail to Get2[T1, T2]() with error: %s", err.Error()))
	}

	return r1, r2
}

func Get3[T1 any, T2 any, T3 any](f func() (T1, T2, T3, error)) (T1, T2, T3) {
	r1, r2, r3, err := f()
	if err != nil {
		panic(fmt.Sprintf("fail to Get2[T1, T2, T3]() with error: %s", err.Error()))
	}

	return r1, r2, r3
}
