package main

import (
	"fmt"
	"log"
	"net"
	"flag"
	pb "github.com/xor-gate/snippets/golang/grpc-pubsub"
	"google.golang.org/grpc"
)

var (
	index = flag.Int("index", 0, "RPC port is 36061+index; debug port is 36661+index")
)

type server struct {}

func (s *server) Publish(ps pb.PublishService_PublishServer) error {
	for {
		msg, err := ps.Recv()
		if err != nil {
			return err
		}
		fmt.Println(string(msg.Data))
	}
	return nil
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
	pb.RegisterPublishServiceServer(g, s)
	fmt.Println("consumer listening on", addr)
	g.Serve(lis)
}
