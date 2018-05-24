#!/bin/sh
# Build a linux client and shoot it to a machine over ssh
export GOOS=linux
export GOARCH=amd64
DOCKERHOST=do
go build -ldflags="-s -w" -o client cmd/client/main.go
scp client remote.sh do: && rm client
ssh do ./remote.sh
