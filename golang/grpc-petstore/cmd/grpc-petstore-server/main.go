package main

import (
	"fmt"
	"log"
	"net"
	"time"
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

func (s *server) ReadPet(req *pb.ReadPetRequest, rps pb.PetstoreService_ReadPetServer) error {
	fmt.Printf("ReadPet:\n%+v\n", req)
	for i := int64(0); i < 11; i++ {
		err := rps.Send(&pb.Pet{Id: i})
		if err != nil {
			log.Println(err)
			return nil
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) UpdatePet(ctx context.Context, pet *pb.Pet) (*pb.Empty, error) {
	fmt.Printf("UpdatePet:\n%+v\n", pet)
	return &pb.Empty{}, nil
}

func (s *server) DeletePet(ctx context.Context, pet *pb.Pet) (*pb.Empty, error) {
	fmt.Printf("DeletePet:\n%+v\n", pet)
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
