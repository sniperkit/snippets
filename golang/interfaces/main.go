package main

import (
	"github.com/xor-gate/go-by-example/interfaces/rpc"
	"github.com/xor-gate/go-by-example/interfaces/rpczmq"
)

func MyMsgHandler(w rpc.Writer, m *rpc.Msg) {
	w.Write(*m)
}

func main() {
	sock := rpczmq.New()

	mux := rpc.NewServeMux(sock)
	mux.HandleFunc("001fc23154ff65065182565553470187", "device:ping", 1, MyMsgHandler)

	msg := rpc.Msg{}
	_ = sock.Write(msg)
	sock.Close()
}
