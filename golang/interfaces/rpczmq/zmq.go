package rpczmq

import (
	"github.com/xor-gate/go-by-example/interfaces/rpc"
)

type ZMQ struct {
}

func New() *ZMQ {
	return &ZMQ{}
}

func (z *ZMQ) Read() (*rpc.Msg, error) {
	return nil,nil
}

func (z *ZMQ) Write(m rpc.Msg) error {
	return nil
}

func (z *ZMQ) Close() error {
	return nil
}
