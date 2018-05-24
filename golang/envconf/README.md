# envconf

[![License][License-Image]][License-Url]
[![Godoc][Godoc-Image]][Godoc-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]

The envconf package allows you to load [environment variables](https://en.wikipedia.org/wiki/Environment_variable) into Go structs.

## Installation and usage

To install, run:

```
$ go get github.com/xor-gate/envconf
$ go install github.com/xor-gate/envconf/cmd/jsonenv
$ $GOPATH/bin/jsonenv
{
	"user": "jerry",
	"shell": "/bin/bash",
	"path": "/Users/jerry/go/bin:/usr/local/bin:/bin:/sbin:/usr/sbin:/usr/local/sbin:/usr/bin:/usr/local/bin:/usr/local/texlive/2017/bin/x86_64-darwin/",
	"pwd": "/Users/jerry/go/src/github.com/xor-gate/envconf"
}
```

And import using:

```go
import "github.com/xor-gate/envconf"
```

Usage is very similar to the `encoding/json` package:

```
package main

import (
	"os"
	"fmt"
	"github.com/xor-gate/envconf"
)

func main() {
	type Env struct {
		User string
		Shell string
		Path string
		Pwd string
	}

	var env Env

	osEnviron := os.Environ()
	// Environ takes a []string argument formatted key=value
	//  there is also envconf.Marshal which is compatible with json.Marshal
	envconf.Environ(osEnviron, &env)

	fmt.Println(env)
}
```

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[Godoc-Url]: https://godoc.org/github.com/xor-gate/envconf
[Godoc-Image]: https://godoc.org/github.com/xor-gate/envconf?status.svg
[ReportCard-Url]: http://goreportcard.com/report/xor-gate/envconf
[ReportCard-Image]: https://goreportcard.com/badge/github.com/xor-gate/envconf
