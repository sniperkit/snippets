package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"path/filepath"
	"context"
	"github.com/xor-gate/go-by-example/grpc-fileservice"
	"google.golang.org/grpc"
)

var (
	server = flag.String("server", "192.168.1.201:3600", "server address")
)

func upload(filename string, fc fileservice.FileClient) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	stream, err := fc.Upload(context.Background())
	if err != nil {
		return err
	}

	err = stream.Send(&fileservice.Chunk{Filename: filepath.Base(f.Name())})
	if err != nil {
		return err
	}
	fmt.Println("upload", filename)

	chunk := make([]byte, 64*1024)

	for {
		n, err := f.Read(chunk)
		//fmt.Println("read", n, err)
		chunk = chunk[:n]
		if len(chunk) == 0 {
			fmt.Println("done")
			return nil
		}

		//fmt.Println("write chunk of", len(chunk), "bytes")
		err = stream.Send(&fileservice.Chunk{Data: chunk})
		if err != nil {
			return err
		}

		chunk = chunk[:cap(chunk)]
	}

	return nil
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := fileservice.NewFileClient(conn)
	fmt.Println(upload("service.proto", client))
	fmt.Println(upload("/Users/jerry/bla.bin", client))
}
