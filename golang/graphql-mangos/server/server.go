package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/rep"
	"github.com/go-mangos/mangos/transport/tcp"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		die("usage: server <url (e.g tcp://127.0.0.1:1337)>")
	}

	fmt.Println("server is running on", os.Args[1])

	var sock mangos.Socket
	var err error
	if sock, err = rep.NewSocket(); err != nil {
		die("can't get new rep socket: %s", err)
	}
	sock.AddTransport(tcp.NewTransport())
	if err = sock.Listen(os.Args[1]); err != nil {
		die("can't listen on rep socket: %s", err.Error())
	}
	for {
		req, err := sock.Recv()
		result := graphql.Do(graphql.Params{
			Schema:        testutil.StarWarsSchema,
			RequestString: string(req),
		})

		rep, err := json.Marshal(result)
		if err != nil {
			continue
		}

		err = sock.Send(rep)
		if err != nil {
			die("can't send reply: %s", err.Error())
		}
	}
}
