package main

import (
	"fmt"
	"net/http"
	"context"
	"golang.org/x/net/context/ctxhttp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	client := &http.Client{}
	cancel()
	_, err := ctxhttp.Get(ctx, client, "www.google.com")
	fmt.Println("err:",err)
}
