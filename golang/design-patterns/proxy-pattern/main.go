package main

import (
	"fmt"
)

type ObjectOps interface {
	Ping()
}

type Object struct{}

func (o *Object) Ping() {
	fmt.Println("Object Ping()")
}

type ObjectProxy struct {
	ObjectOps
}

func (op *ObjectProxy) Ping() {
	fmt.Println("ObjectProxy before Ping()")
	op.ObjectOps.Ping()
	fmt.Println("ObjectProxy after Ping()")
}

func main() {
	op := ObjectProxy{ObjectOps: &Object{}}
	op.Ping()
}
