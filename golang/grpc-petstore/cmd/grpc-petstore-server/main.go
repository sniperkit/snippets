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

type server struct {
	petID int64
	pets []*pb.Pet
}

func (s *server) CreatePet(ctx context.Context, pet *pb.Pet) (*pb.Empty, error) {
	fmt.Printf("CreatePet:\n%+v\n", pet)
	pet.Id = s.petID
	s.pets = append(s.pets, pet)
	s.petID++
	return &pb.Empty{}, nil
}

func (s *server) ReadPet(req *pb.ReadPetRequest, rps pb.PetstoreService_ReadPetServer) error {
	fmt.Printf("ReadPet:\n%+v\n", req)
	for _, pet := range s.pets {
		err := rps.Send(pet)
		if err != nil {
			log.Println(err)
			return nil
		}
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
	s := &server{petID: 1}
	g := grpc.NewServer()
	pb.RegisterPetstoreServiceServer(g, s)
	fmt.Println("gRPC Petstore listening on", addr)
	g.Serve(lis)
}
