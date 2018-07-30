package main

import (
	"fmt"
	"log"
	"flag"
	"time"
	"context"
	pb "github.com/xor-gate/snippets/golang/grpc-pubsub"
	"google.golang.org/grpc"
)

var (
	server = flag.String("server", "localhost:36060", "consumer server address")
)

func main() {
	flag.Parse()

	// Connect to the server.
	conn, err := grpc.Dial(*server, grpc.WithInsecure()) // HL
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPublishServiceClient(conn) // HL

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cnt := 1

	publisher, err := client.Publish(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for {
		err := publisher.Send(&pb.Message{Data: []byte(fmt.Sprintf(`%d`, cnt))})
		if err != nil {
			log.Fatal(err)
		}
		cnt++
		time.Sleep(time.Second)
	}
}
