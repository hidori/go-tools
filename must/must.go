package must

import (
	"fmt"
)

func Must[T any](f func() (T, error)) T {
	r, err := f()
	if err != nil {
		panic(fmt.Sprintf("fail to Must[T any]() error = %s", err.Error()))
	}

	return r
}

func Must2[T1 any, T2 any](f func() (T1, T2, error)) (T1, T2) {
	r1, r2, err := f()
	if err != nil {
		panic(fmt.Sprintf("fail to ust2[T1 any, T2 any]() error = %s", err.Error()))
	}

	return r1, r2
}

func Must3[T1 any, T2 any, T3 any](f func() (T1, T2, T3, error)) (T1, T2, T3) {
	r1, r2, r3, err := f()
	if err != nil {
		panic(fmt.Sprintf("fail to ust2[T1 any, T2 any, T3 any]() error = %s", err.Error()))
	}

	return r1, r2, r3
}
