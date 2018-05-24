package main

import (
	"fmt"
	"log"
	"time"
	"strings"
	"net"
	"flag"
	"context"
	pb "github.com/xor-gate/go-by-example/grpc-distrib-reqrep"
	"google.golang.org/grpc"
)

var (
	backends = flag.String("backends", "127.0.0.1:36061,127.0.0.1:36062", "comma-separated backend server addresses")
)

type server struct {
	backends []pb.DeviceClient
}

func (s *server) Ping(ctx context.Context, req *pb.Request) (*pb.Empty, error) {
	fmt.Println("ping", req)
	c := make(chan pingResult)
	for _, b := range s.backends {
		go func(backend pb.DeviceClient) {
			res, err := backend.Ping(ctx, req)
			if err != nil {
				return
			}
			c <- pingResult{res, err}
		}(b)
	}
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
		return &pb.Empty{},fmt.Errorf("boem!")
	case first := <-c:
		fmt.Println("rep", first.res)
		return first.res, first.err
	case <-ctx.Done():
		fmt.Println("ctx", ctx.Err())
		return &pb.Empty{},ctx.Err()
	}
}

type pingResult struct {
	res *pb.Empty
	err error
}

func main() {
	lis, err := net.Listen("tcp", ":36060") // RPC port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := new(server)
	for _, addr := range strings.Split(*backends, ",") {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Printf("fail to dial: %v", err)
			continue
		}
		fmt.Println("backend", addr)
		client := pb.NewDeviceClient(conn)
		s.backends = append(s.backends, client)
	}
	g := grpc.NewServer()
	pb.RegisterDeviceServer(g, s)
	fmt.Println("frontend listening on :36060")
	g.Serve(lis)
}
