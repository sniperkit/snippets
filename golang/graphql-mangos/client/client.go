package main

import (
	"os"
	"fmt"
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/req"
	"github.com/go-mangos/mangos/transport/tcp"
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		die("usage: client <uri (e.g tcp://127.0.0.1:1337)> <query>")
	}

	var sock mangos.Socket
	var err error
	var msg []byte

	if sock, err = req.NewSocket(); err != nil {
		die("can't get new req socket: %s", err.Error())
	}
	sock.AddTransport(tcp.NewTransport())
	if err = sock.Dial(os.Args[1]); err != nil {
		die("can't dial on req socket: %s", err.Error())
	}
	if err = sock.Send([]byte(os.Args[2])); err != nil {
		die("can't send message on push socket: %s", err.Error())
	}
	if msg, err = sock.Recv(); err != nil {
		die("can't receive date: %s", err.Error())
	}
	fmt.Printf("%s\n", string(msg))
	sock.Close()
}
