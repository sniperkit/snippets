package main

import (
	"fmt"
	"time"
)

const (
	Nanosecond   = 1
	Microsecond  = 1000 * Nanosecond
	Millisecond  = 1000 * Microsecond
)

func main() {
	t := time.Unix( 84572, 351 * Millisecond)
	fmt.Printf("%02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())
}
