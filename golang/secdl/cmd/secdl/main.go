// Tool to generate expire-able and hashed urls for files
package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/xor-gate/secdl"
)

var filename = flag.String("filename", "secret-page.html", "File to generate create url for (without leading slash!)")
var prefix = flag.String("prefix", "/s/", "Prefix of secure download area")
var secret = flag.String("secret", "1234", "Secret for validating the token")
var expire = flag.String("expire", "n", "Expire time: (n)ever, (10m)inutes, (h)our, (d)ay, (1w)eek, (2w)eeks, (m)onth)")
var listen = flag.String("listen", "", "Listen HTTP Server at \":<port>\"")
var root = flag.String("root", ".", "Root folder for HTTP Server")

func main() {
	flag.Parse()

	if *listen != "" {
		*root, _ = filepath.Abs(*root)
		fmt.Printf("[secdl] Serving HTTP at %s from path \"%s\" (prefix \"%s\", secret \"%s\")\n", *listen, *root, *prefix, *secret)
		http.Handle(*prefix, secdl.FileServer(*secret, *prefix, *root))
		http.ListenAndServe(*listen, nil)
	} else {
		c := secdl.New()

		c.SetSecret(*secret)
		c.SetPrefix(*prefix)
		c.SetFilename("/" + *filename)

		e, err := secdl.ParseExpire(*expire)
		if err != nil {
			flag.Usage()
			return
		}

		c.Encode(e)

		if e != secdl.ExpireNever {
			expireAt := time.Now().Add(time.Duration(e))
			fmt.Printf("url (expires at %s):\n\t%s\n", expireAt.String(), c.URL)
		} else {
			fmt.Printf("url (expires never):\n\t%s\n", c.URL)
		}
	}
}
