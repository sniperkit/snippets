package main

func main() {
	srv, _ := TCPServerNew("localhost", "3333")
	defer srv.Close()
	srv.Serve()
}
