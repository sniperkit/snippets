package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"time"
	"net"
	pb "github.com/xor-gate/go-by-example/revdial-grpc"
	"github.com/golang/build/revdial"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

const listenAddr = ":1234"

func Command(gc pb.GreeterClient, name string, arg ...string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ls, err := gc.Exec(ctx, &pb.ExecCommand{Name: name, Arg: arg})
	if err != nil {
		return
	}

	for {
		ll, err := ls.Recv()
		if err != nil {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", ll.GetStdout())
		fmt.Fprintf(os.Stderr, "%s", ll.GetStderr())
	}
}

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()

	log.Printf("listen on tcp://%s\n", listenAddr)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}

		log.Println("new connection:", c.RemoteAddr())

		go func() {
			rd := revdial.NewDialer(bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c)), c)
			defer rd.Close()

			dialer := func (network string, timeout time.Duration) (net.Conn, error) {
					return rd.Dial()
			}

			conn, err := grpc.Dial("doesntmatter", grpc.WithInsecure(), grpc.WithDialer(dialer))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}

			rpcc := pb.NewGreeterClient(conn)
			greet, err := rpcc.Greet(context.Background(), &pb.GreetMsg{Name: "i'm your commander"})
			if err != nil {
				log.Println("rpcc.Greet err:", err)
				return
			}
			fmt.Println("greet", greet)

			Command(rpcc, "apt-get", "update")
			Command(rpcc, "apt-get", "install", "-y", "build-essential", "git-core", "cmake")
			Command(rpcc, "git", "clone", "https://github.com/xor-gate/eresp.git")
			Command(rpcc, "cd", "eresp")
			Command(rpcc, "make", "-C", "eresp")
		}()
	}
}
