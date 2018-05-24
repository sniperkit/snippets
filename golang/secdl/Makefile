export GOPATH?=$(PWD)/../../../../
export DESTDIR?=$(GOPATH)/bin
export GOBIN?=$(DESTDIR)

all: tools
ci: tools lint test bench

test:
	go test -v -race

lint:
	go tool vet .

fmt:
	gofmt -d -s .

bench:
	go test -bench=.

deb:
	go get github.com/xor-gate/debpkg
	go install github.com/xor-gate/debpkg/cmd/debpkg
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o secdl cmd/secdl/main.go; debpkg -c .debpkg.yml; rm secdl

tools: $(GOBIN)/secdl $(GOBIN)
install: tools
	mkdir -p $(PKGDIR)/usr/bin
	cp $(GOPATH)/bin/secdl $(PKGDIR)/usr/bin

$(GOBIN)/secdl: secdl.go cmd/secdl/main.go
	go install -v github.com/xor-gate/secdl/cmd/secdl

lighttpd:
	lighttpd -D -f etc/lighttpd.conf
