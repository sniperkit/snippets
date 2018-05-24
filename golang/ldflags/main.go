package main

import (
	"fmt"
)

var (
	version   string
	buildDate string
)

func main() {
	fmt.Printf("version: %s (%s)\n", version, buildDate)
}
