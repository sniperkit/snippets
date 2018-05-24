# sysdetect

[![GoDoc](https://godoc.org/github.com/xor-gate/sysdetect?status.svg)](https://godoc.org/github.com/xor-gate/sysdetect)
[![Sourcegraph](https://sourcegraph.com/github.com/xor-gate/sysdetect/-/badge.svg)](https://sourcegraph.com/github.com/xor-gate/sysdetect)
[![Build Status](https://travis-ci.org/xor-gate/sysdetect.svg?branch=master)](https://travis-ci.org/xor-gate/sysdetect)
[![Go Report Card](https://goreportcard.com/badge/github.com/xor-gate/sysdetect)](https://goreportcard.com/report/github.com/xor-gate/sysdetect)
[![codecov](https://codecov.io/gh/xor-gate/sysdetect/branch/master/graph/badge.svg)](https://codecov.io/gh/xor-gate/sysdetect)

System Detection in pure Golang (loosely based on Ansible system module)

Requires at least Go 1.6.

## Usage

```
go get github.com/xor-gate/sysdetect/cmd/sysdetect
go run $(go env GOPATH)/src/github.com/xor-gate/sysdetect/examples/local.go
```

## Links

* [os-release](https://www.freedesktop.org/software/systemd/man/os-release.html)

## License

[MIT](LICENSE)
