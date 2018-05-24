package main

import (
	"io"
	"fmt"
	"os"
	"sync"
	"os/exec"
	"log"
	"flag"
	"net"
	"bufio"
	"github.com/golang/build/revdial"

	"github.com/xor-gate/go-by-example/revdial-grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var address string

func init() {
	flag.StringVar(&address, "host", "127.0.0.1:1234", "Host address")
	flag.Parse()
}

type server struct {}

func (s *server) Greet(ctx context.Context, gm *revdialgrpc.GreetMsg) (*revdialgrpc.GreetReply, error) {
	hostname, _ := os.Hostname()
	return &revdialgrpc.GreetReply{Hostname: hostname}, nil
}

func (s *server) Exec(cmd *revdialgrpc.ExecCommand, out revdialgrpc.Greeter_ExecServer) error {
	pr, pw := io.Pipe()

	ecmd := exec.Command(cmd.GetName(), cmd.GetArg()...)
	ecmd.Stdout = pw
	ecmd.Stderr = pw

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer pr.Close()
		r := bufio.NewReader(pr)
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				return
			}
			if err := out.Send(&revdialgrpc.ExecStream{Log: &revdialgrpc.ExecStream_Stdout{line}}); err != nil {
				return
			}
		}
		fmt.Println("done")
	}()

	ecmd.Run()
	pw.Close()
	wg.Wait()
	return nil
}

func main() {
	c, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	rcl := revdial.NewListener(bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c)))
	defer rcl.Close()

	rpcs := grpc.NewServer()
	revdialgrpc.RegisterGreeterServer(rpcs, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(rpcs)
	if err := rpcs.Serve(rcl); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
