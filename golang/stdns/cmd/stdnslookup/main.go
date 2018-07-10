package main

import (
	"fmt"
	"context"
	"github.com/xor-gate/snippets/golang/stdns"
)

func main() {
	r := stdns.NewStdResolver()
	entries, err := r.Lookup(context.Background(), "syncthing://syncthing@dorpstraat.xor-gate.org")
	fmt.Println(err)
	for _, entry := range entries {
		fmt.Println("Username", entry.Username())
		fmt.Println("Hostname", entry.Hostname())
		fmt.Println("DeviceID", entry.DeviceID)
	}
}
