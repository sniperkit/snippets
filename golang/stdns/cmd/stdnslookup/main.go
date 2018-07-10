package main

import (
	"context"
	"github.com/xor-gate/snippets/golang/stdns"
)

func main() {
	r := stdns.NewStdResolver()
	r.Lookup(context.Background(), "dorpstraat.xor-gate.org")
}
