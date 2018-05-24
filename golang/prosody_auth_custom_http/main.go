package main

import (
	"os"
	"io"
	"fmt"
	"net/http"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "text/json")
	fmt.Fprintf(os.Stdout, "req:")
	io.Copy(os.Stdout, r.Body)
	fmt.Fprintf(w, "true")
	//fmt.Fprintf(w, "false")
	fmt.Println()
}

func main() {
	http.HandleFunc("/", authHandler)
	http.ListenAndServe(":8080", nil)
}
