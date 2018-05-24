package main

type noCopy struct{}
func (*noCopy) Lock() {}

type MyStruct struct {
	noCopy noCopy
}

func main() {
	var orig MyStruct
	cc := orig
	_ = cc
}
