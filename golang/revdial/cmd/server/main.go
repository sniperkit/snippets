package main

import (
//	"io"
//	"os"
	"bufio"
	"log"
	"time"
	"net"
	"github.com/golang/build/revdial"
)

const listenAddr = ":1234"

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()

	log.Printf("listen on tcp://%s\n", listenAddr)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}

		log.Println("new connection:", c.RemoteAddr())

		go func() {
			rd := revdial.NewDialer(bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c)), c)

			rc, err := rd.Dial()
			if err != nil {
				log.Println()
				return
			}
			defer rc.Close()

			for {
				_, err = rc.Write([]byte("hello reverse\n"))
				if err != nil {
					log.Println("rc.Write failed:", err)
					return
				}
				time.Sleep(time.Second)
			}
		}()
	}
}
