package main

import (
    "fmt"
    "os"
    "path"
)

func func1() {
	fmt.Printf("Hello world from func1(args: %+v)\n", os.Args)
}

func main() {
	if path.Base(os.Args[0]) == "func1" {
		func1()
	}
}
