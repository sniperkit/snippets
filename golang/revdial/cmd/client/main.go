package main

import (
	"io"
	"os"
	"log"
	"net"
	"fmt"
	"bufio"
	"github.com/golang/build/revdial"
)

func main() {
	c, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	rcl := revdial.NewListener(bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c)))
	defer rcl.Close()

	for {
		rc, err := rcl.Accept()
		if err != nil {
			log.Println("rcl.Accept err:", err)
			return
		}

		fmt.Println("new reverse connection:", rc.RemoteAddr(), rc.LocalAddr())

		go func() {
			io.Copy(os.Stdout, rc)
		}()
	}
}
