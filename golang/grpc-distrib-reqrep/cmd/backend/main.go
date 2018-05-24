package main

import (
	"fmt"
	"log"
	"net"
	"flag"
	"context"
	pb "github.com/xor-gate/go-by-example/grpc-distrib-reqrep"
	"google.golang.org/grpc"
)

var (
	index = flag.Int("index", 0, "RPC port is 36061+index; debug port is 36661+index")
)

type server struct {}

func (s *server) Ping(ctx context.Context, req *pb.Request) (*pb.Empty, error) {
	log.Println("ping", req)
	return &pb.Empty{}, nil
}

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%d", 36061+*index)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	s := &server{}
	g := grpc.NewServer()
	pb.RegisterDeviceServer(g, s)
	fmt.Println("backend listening on", addr)
	g.Serve(lis)
}
