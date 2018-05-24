package main

import (
	"fmt"
	"os"
	"time"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"encoding/base64"
	"gopkg.in/macaron.v1"
)

func decrypt(key []byte, ciphertext []byte) []byte {
	n := len(ciphertext)
	d := make([]byte, n)

	for i := range ciphertext {
		d[i] = ciphertext[i] ^ key[i%len(key)]
	}

	return d
}

func encrypt(key []byte, plaintext []byte) []byte {
	n := len(plaintext)
	e := make([]byte, n)

	for i := range plaintext {
		e[i] = plaintext[i] ^ key[i%len(key)]
	}

	return e
}

func main() {
	b, err := ioutil.ReadFile("main.key")
	if err != nil {
		fmt.Print(err)
	}

	str := "main.go"

	encStr := encrypt(b, []byte(str))
	enc := base64.StdEncoding.EncodeToString(encStr)

	fmt.Printf("%s\n", enc)

	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/v1/:id", func(ctx *macaron.Context) {
		sDec, _ := base64.StdEncoding.DecodeString(ctx.Params(":id"))
		sDec = decrypt(b, sDec)

		if _, err := os.Stat(string(sDec)); os.IsNotExist(err) {
			ctx.WriteHeader(400)
			return
		}

		f, err := os.Open(string(sDec))
		if err != nil {
			ctx.WriteHeader(404)
			return
		}

		// Only the first 512 bytes are used to sniff the content type.
		buffer := make([]byte, 512)
		_, err = f.Read(buffer)
		if err != nil {
			return
		}

		// Reset the read pointer if necessary.
		f.Seek(0, 0)

		// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
		contentType := http.DetectContentType(buffer)

		filename := filepath.Base(string(sDec))
		ctx.Header().Set("Content-Type", fmt.Sprintf(`%s; filename="%s"`, contentType, filename))
		ctx.Header().Set("Content-Security-Policy", "unsafe-inline")
		ctx.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filename))

		http.ServeContent(ctx.Resp, ctx.Req.Request, filename, time.Now(), f)
	})

	m.Run()
}
