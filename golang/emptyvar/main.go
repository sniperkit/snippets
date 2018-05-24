package main

import (
	"fmt"
)

type Boem struct {
	a int
}

func (b *Boem) IsEmpty() bool {
	return (*b == Boem{})
}

func main() {
	a := Boem{a: 1}

	if a.IsEmpty() {
		fmt.Println("empty")
	} else {
		fmt.Println("not empty")
	}
}
