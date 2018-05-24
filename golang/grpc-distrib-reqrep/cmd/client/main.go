package main

import (
	"fmt"
	"log"
	"flag"
	"time"
	"context"
	pb "github.com/xor-gate/go-by-example/grpc-distrib-reqrep"
	"google.golang.org/grpc"
)

var (
	server = flag.String("server", "localhost:36060", "server address")
)

func main() {
	flag.Parse()

	// Connect to the server.
	conn, err := grpc.Dial(*server, grpc.WithInsecure()) // HL
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewDeviceClient(conn) // HL

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cnt := 1

	for {
		log.Println("ping...")
		_, err := client.Ping(ctx, &pb.Request{DeviceUid: "1337", Data: []byte(fmt.Sprintf(`boembatsknal %d`, cnt))})
		cnt++
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}
