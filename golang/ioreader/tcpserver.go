package main

import (
	"io"
	"net"
	"fmt"
)

type TCPServer struct {
	conn net.Listener
}

func TCPServerNew(host, port string) (srv *TCPServer, err error) {
	s := &TCPServer{}

	s.conn, err = net.Listen("tcp", host+":"+port)
	if err != nil {
		return
	}

	fmt.Println("Listening on " + host + ":" + port)

	return s, nil
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	defer conn.Close()

	msg := &MyStruct{}
	for {
		//reqLen, err := io.Copy(msg, conn)
		//if err != nil {
		//	fmt.Println("Error reading:", err.Error())
		//	break
		//}
		//fmt.Println("len:", reqLen, msg)
		_, err := io.Copy(msg, conn)
		if err != nil {
			fmt.Println(err)
			break
		}

		_, err = io.Copy(conn, msg)
		if err != nil {
			fmt.Println("err:", err)
			break
		}
	}
}

func (srv *TCPServer) Serve() {
	for {
		conn, err := srv.conn.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		fmt.Println("New connection:", conn.RemoteAddr())

		go handleRequest(conn)
	}
}

func (srv *TCPServer) Close() {
	srv.conn.Close()
}
