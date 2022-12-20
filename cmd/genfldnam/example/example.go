package example

type ExampleStruct struct {
	Value0 int // ignored
	Value1 int `genfldnam:"+"`
	Value2 int `genfldnam:"+"`
}
