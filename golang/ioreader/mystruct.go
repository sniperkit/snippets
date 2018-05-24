package main

import (
	"io"
	"fmt"
)

type MyStruct struct {
	//
}

var MyStructVar = []byte{'a', 'b', 'c', 'd', '\n'}

func (s *MyStruct) Read(p []byte) (n int, err error) {
	fmt.Println("MyStruct Read:")
	copy(p, MyStructVar)
	return len(MyStructVar), io.EOF
}

func (s *MyStruct) Write(p []byte) (n int, err error) {
	fmt.Println("MyStruct Write:", p)
	return len(p), nil
}
