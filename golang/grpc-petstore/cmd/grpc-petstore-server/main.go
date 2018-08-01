package main

import (
	"fmt"
	"log"
	"net"
	"context"
	"flag"
	pb "github.com/xor-gate/snippets/golang/grpc-petstore"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9090, "Petstore gRPC port")
)

type server struct {}

func (s *server) CreatePet(ctx context.Context, pet *pb.Pet) (*pb.Empty, error) {
	fmt.Printf("CreatePet:\n%+v\n", pet)
	return &pb.Empty{}, nil
}

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	s := &server{}
	g := grpc.NewServer()
	pb.RegisterPetstoreServiceServer(g, s)
	fmt.Println("gRPC Petstore listening on", addr)
	g.Serve(lis)
}
