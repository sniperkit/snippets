# gRPC Petstore

This is hand-translate Swagger Petstore example in gRPC with gRPC-Gateway (REST)
The web gui is build using Vue.js

## Usage

Download, compile and install

```
go get github.com/xor-gate/snippets/golang/grpc-petstore
go install github.com/xor-gate/snippets/golang/grpc-petstore/cmd/grpc-petstore-server
go install github.com/xor-gate/snippets/golang/grpc-petstore/cmd/grpc-petstore-gateway
```

Run server and REST gateway

```
grpc-petstore-gateway &
grpc-petstore-server
```

From another terminal just create a Pet

```
curl --request POST http://127.0.0.1:8080/v1/pet \
     --header "Content-Type: application/json" \
     --data '{"name":"sheep"}'
```
