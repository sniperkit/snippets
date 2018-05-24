package main

import (
	"net"
	"log"
	"io"
	"io/ioutil"
	"time"
	"bytes"
	"fmt"
	"github.com/dustin/go-humanize"
	"google.golang.org/grpc"
	"github.com/xor-gate/go-by-example/grpc-fileservice"
)

type server struct {}

func (s *server) Upload(stream fileservice.File_UploadServer) error {
	chunk, err := stream.Recv()
	if err != nil {
		return err
	}

	if chunk.Filename == "" {
		return fmt.Errorf("first chunk has no filename")
	}

	f, err := ioutil.TempFile("", chunk.Filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	fmt.Println("created", f.Name())

	tt := time.NewTicker(time.Second)
	defer tt.Stop()
	var byteCnt uint64

	for {
		chunk, err = stream.Recv()
		if err != nil {
			fmt.Println(humanize.Bytes(byteCnt), "per sec")
			fmt.Println("finished", f.Name())
			return err
		}

		select {
		case <-tt.C:
			fmt.Println(humanize.Bytes(byteCnt), "per sec")
			byteCnt = 0
		default:
			byteCnt += uint64(len(chunk.Data))
		}

		//fmt.Println("write chunk", len(chunk.Data), "bytes")
		_, err := io.Copy(f, bytes.NewReader(chunk.Data))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *server) Download(req *fileservice.DownloadRequest, stream fileservice.File_DownloadServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3600")
	if err != nil {
		log.Fatal(err)
	}
	g := grpc.NewServer()
	fileservice.RegisterFileServer(g, &server{})
	g.Serve(lis)
}
