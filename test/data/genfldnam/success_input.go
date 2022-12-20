package data

import "go/ast"

type SuccessStruct struct {
	ignored1 int
	ignored2 int `genfldnam:""`
	ignored3 int `genfldnam:"-"`

	Int1 int `genfldnam:"+"`
	Int2 int `genfldnam:"+"`

	OtherSuccessStruct1 OtherSuccessStruct  `genfldnam:"+"`
	OtherSuccessStruct2 *OtherSuccessStruct `genfldnam:"+"`

	AstFile1 ast.File  `genfldnam:"+"`
	AstFile2 *ast.File `genfldnam:"+"`
}

type OtherSuccessStruct struct {
	OtherInt1 int `genfldnam:"+"`
	OtherInt2 int `genfldnam:"+"`
}

type EmptyStruct struct{}

type String string
