package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"compress/gzip"
	"github.com/xor-gate/go-snippets/templates/template"
)

func main() {
	templates := templates.Templates()

	var b bytes.Buffer
	b.Write(templates["index.html"])
	r, _ := gzip.NewReader(&b)
	defer r.Close()

	s, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	fmt.Println(string(s))
}
