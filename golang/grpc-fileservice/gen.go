//go:generate protoc -I . --go_out=plugins=grpc:. service.proto

package fileservice
