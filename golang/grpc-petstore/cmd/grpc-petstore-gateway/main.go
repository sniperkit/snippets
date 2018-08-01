package main

import (
  "flag"
  "log"
  "net/http"

  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"
  gw "github.com/xor-gate/snippets/golang/grpc-petstore"
)

var (
  petstoreEndpoint = flag.String("petstore_endpoint", "localhost:9090", "endpoint of PetstoreService")
)

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterPetstoreServiceHandlerFromEndpoint(ctx, mux, *petstoreEndpoint, opts)
  if err != nil {
    return err
  }

  log.Println("Listening grpc-petstore-gateway on :8080 proxy to", *petstoreEndpoint)
  return http.ListenAndServe(":8080", mux)
}

func main() {
  flag.Parse()

  if err := run(); err != nil {
    log.Fatal(err)
  }
}
