// +build ignore

package main

import (
	"log"

	"github.com/xor-gate/snippets/golang/sysdetect"
)

func main() {
	sd := sysdetect.NewLocal()
	result := sd.Detect()
	for key, value := range result {
		log.Printf("%s=%s", key, value)
	}
}
