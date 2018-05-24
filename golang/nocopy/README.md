# Marking types as non-copyable

A nice trick from https://github.com/golang/go/issues/8005#issuecomment-190753527

Which uses the `go vet` tool to detect if `Lock()` is copied which is not allowed.

```
$ go vet
main.go:12: assignment copies lock value to cc: main.MyStruct contains main.noCopy
main.go:13: assignment copies lock value to _: main.MyStruct contains main.noCopy
```
