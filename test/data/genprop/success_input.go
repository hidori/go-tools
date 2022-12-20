package data

import (
	"fmt"
	"go/ast"
	"regexp"
)

type SuccessStruct struct {
	ignored1 int
	ignored2 int `genprop:""`
	ignored3 int `genprop:"-"`

	int1 int  `genprop:"get,set"`
	int2 *int `genprop:"get,set"`

	string1 string  `genprop:"get,set"`
	string2 *string `genprop:"get,set"`

	interface1 interface{}  `genprop:"get,set"`
	interface2 *interface{} `genprop:"get,set"`

	otherSuccessStruct1 OtherSuccessStruct  `genprop:"get,set"`
	otherSuccessStruct2 *OtherSuccessStruct `genprop:"get,set"`

	astFile1 ast.File  `genprop:"get,set"`
	astFile2 *ast.File `genprop:"get,set"`

	api         string `genprop:"get"`
	apiEndpoint string `genprop:"get"`
}

type OtherSuccessStruct struct {
	otherInt1 int `genprop:"get"`
	otherInt2 int `genprop:"get"`
}

type EmptyStruct struct{}

type String string

func Func() {
	type IgnoredInnerStruct struct {
		Int1 int `genprop:"get"`
	}

	fmt.Println("Hello, world")
}

const IgnoredConst = 1

var IgnoredVar = regexp.MustCompile(`^$`)

var IgnoredAnonymousStruct = struct {
	Int1 int `genprop:"get"`
}{
	Int1: 1,
}
