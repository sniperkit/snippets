package main

import (
  "flag"
  "log"
  "strings"
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

   s := &http.Server{
		Addr:    ":8080",
		Handler: allowCORS(mux),
  }

  log.Println("Listening grpc-petstore-gateway on :8080 proxy to", *petstoreEndpoint)
  return s.ListenAndServe()
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Printf("preflight request for %s", r.URL.Path)
	return
}

func main() {
  flag.Parse()

  if err := run(); err != nil {
    log.Fatal(err)
  }
}
