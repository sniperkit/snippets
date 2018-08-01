#!/bin/sh
go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u -v github.com/golang/protobuf/protoc-gen-go

protoc -I/usr/local/include -I. \
  -I`go env GOPATH`/src \
  -I`go env GOPATH`/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  petstore.proto

protoc -I/usr/local/include -I. \
  -I`go env GOPATH`/src \
  -I`go env GOPATH`/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  petstore.proto

protoc -I/usr/local/include -I. \
  -I`go env GOPATH`/src \
  -I`go env GOPATH`/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
  petstore.proto
