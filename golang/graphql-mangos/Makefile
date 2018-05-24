all: server.exe client.exe
dep:
	go get -u ./...
server.exe: server/server.go
	go build -o server.exe server/server.go
client.exe: client/client.go
	go build -o client.exe client/client.go
clean:
	rm -f server.exe
	rm -f client.exe
