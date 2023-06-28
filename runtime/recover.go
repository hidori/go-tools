package runtime

import (
	"fmt"
)

func handle(err *error) {
	e := recover()
	if e != nil {
		*err = fmt.Errorf("runtime.Recover(): %v", e)
	}
}

func Recover(fc func() error) error {
	var err error

	func() {
		defer handle(&err)
		err = fc()
	}()

	return err
}

func Recover1[T1 any](fc func() (T1, error)) (T1, error) {
	var (
		v1  T1
		err error
	)

	func() {
		defer handle(&err)
		v1, err = fc()
	}()

	return v1, err
}

func Recover2[T1 any, T2 any](fc func() (T1, T2, error)) (T1, T2, error) {
	var (
		v1  T1
		v2  T2
		err error
	)

	func() {
		defer handle(&err)
		v1, v2, err = fc()
	}()

	return v1, v2, err
}

func Recover3[T1 any, T2 any, T3 any](fc func() (T1, T2, T3, error)) (T1, T2, T3, error) {
	var (
		v1  T1
		v2  T2
		v3  T3
		err error
	)

	func() {
		defer handle(&err)
		v1, v2, v3, err = fc()
	}()

	return v1, v2, v3, err
}
