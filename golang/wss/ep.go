package main

import (
	"time"
	"net/http"
)

func epHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	time.Sleep(time.Second)
}
