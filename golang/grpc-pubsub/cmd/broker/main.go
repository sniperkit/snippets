package main

import (
	"fmt"
	"log"
	"strings"
	"context"
	"net"
	"flag"
	pb "github.com/xor-gate/snippets/golang/grpc-pubsub"
	"google.golang.org/grpc"
)

var (
	backends = flag.String("consumers", "127.0.0.1:36061,127.0.0.1:36062", "comma-separated backend server addresses")
)

type server struct {
	backends []pb.PublishService_PublishClient
}

func (s *server) Publish(ps pb.PublishService_PublishServer) error {
	for {
		msg, err := ps.Recv()
		if err != nil {
			return err
		}

		fmt.Println(string(msg.Data))

		for _, backend := range s.backends {
			go func(be pb.PublishService_PublishClient, msg *pb.Message) {
				err := be.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			}(backend, msg)
		}
	}
	return nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", ":36060") // RPC port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &server{}

	for _, addr := range strings.Split(*backends, ",") {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Printf("fail to dial: %v", err)
			continue
		}
		fmt.Println("backend", addr)
		client := pb.NewPublishServiceClient(conn)
		publisher, err := client.Publish(context.Background())
		if err != nil {
			log.Printf("fail to dial: %v", err)
			continue
		}
		s.backends = append(s.backends, publisher)
	}

	g := grpc.NewServer()
	pb.RegisterPublishServiceServer(g, s)
	fmt.Println("broker listening on :36060")
	g.Serve(lis)
}
