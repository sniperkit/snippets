package main

import (
	"fmt"
	"context"
	"github.com/xor-gate/snippets/golang/stdns"
)

func main() {
	r := stdns.NewStdResolver()
	entries, _ := r.Lookup(context.Background(), "syncthing://syncthing@dorpstraat.xor-gate.org")
	for _, entry := range entries {
		fmt.Println("URL", entry.URL)
		fmt.Println("Username", entry.Username())
		fmt.Println("Hostname", entry.Hostname())
		fmt.Println("DeviceID", entry.DeviceID)
	}
}
