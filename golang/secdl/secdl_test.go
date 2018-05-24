// Copyright 2015 Jerry Jacobs. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package secdl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

var expectExpire = []struct {
	v   Expire
	e   string
	abr string
}{
	{Expire10Minutes, "10 minutes", "10m"},
	{ExpireHour, "hour", "h"},
	{ExpireDay, "day", "d"},
	{Expire1Week, "1 week", "1w"},
	{Expire2Weeks, "2 weeks", "2w"},
	{ExpireMonth, "month", "m"},
	{ExpireNever, "never", "n"},
}

func TestExpireString(t *testing.T) {
	for _, tt := range expectExpire {
		r, err := Expire.String(tt.v)
		if r != tt.e || err != nil {
			t.Errorf("Expire.String(): expected \"%s\", actual \"%s\" (err: \"%s\")", tt.e, r, err)
		}
	}

	// Inject invalid Expire value
	_, err := Expire.String(Expire(1234))
	if err == nil {
		t.Errorf("Expire.String(1234) expecting error, got nil")
	}
}

func TestParseExpire(t *testing.T) {
	// Full strings
	for _, tt := range expectExpire {
		r, err := ParseExpire(tt.e)
		if r != tt.v || err != nil {
			estr, _ := Expire.String(tt.v)
			rstr, _ := Expire.String(r)
			t.Errorf("ParseExpire(\"%s\"): expected \"%s\", actual \"%s\", (err: \"%s\")", tt.e, estr, rstr, err)
		}
	}

	// Abbreviated string
	for _, tt := range expectExpire {
		r, err := ParseExpire(tt.abr)
		if r != tt.v || err != nil {
			estr, _ := Expire.String(tt.v)
			rstr, _ := Expire.String(r)
			t.Errorf("ParseExpire(\"%s\"): expected \"%s\", actual \"%s\", (err: \"%s\")", tt.e, estr, rstr, err)
		}
	}

	// Inject invalid Expire string
	_, err := ParseExpire("aaa")
	if err == nil {
		t.Errorf("ParseExpire(\"aaa\") expecting error, got nil")
	}
}

func TestEncode(t *testing.T) {
	c := &SecDl{
		Secret:   "1234",
		Prefix:   "/dl/",
		Filename: "/dir1/dir2/dir3/secret-file.txt"}

	c.Encode(ExpireHour)
}

func TestEncodeInvalidExpire(t *testing.T) {
	c := &SecDl{
		Secret:   "1234",
		Prefix:   "/dl/",
		Filename: "/dir1/dir2/dir3/secret-file.txt"}

	_, err := c.Encode(Expire(ExpireHour + 1))
	if err == nil {
		t.Error("Invalid expire message expected")
	}
}

func TestDecodeInvalidURL(t *testing.T) {
	req := &SecDl{
		URL: "/a/dir3/secret-file.txt"}

	err := req.Decode()
	if err != fmt.Errorf("URL is malformed, unexpected field count") && req.Status != StatusError {
		t.Error("Expecting StatusError")
	}
}

func TestDecodeMonth(t *testing.T) {
	c := New()

	c.SetSecret("1234")
	c.SetPrefix("/dl/")
	c.SetFilename("/dir1/dir2/dir3/secret-file.txt")

	c.Encode(ExpireMonth)
	c.Decode()

	if c.Status != StatusValid {
		t.Error("Unable to decode")
	}
}

func TestDecodeExpired(t *testing.T) {
	c := &SecDl{
		Secret:   "1234",
		Prefix:   "/dl/",
		Filename: "/dir1/dir2/dir3/secret-file.txt"}

	c.Encode(expireSecond)
	time.Sleep(2 * time.Second)
	c.Decode()

	if c.Status != StatusExpired {
		t.Error("Token is not expired, we did wait 2 seconds...!")
	}
}

func TestDecodeTrimLongPrefix(t *testing.T) {
	c := &SecDl{
		Secret: "1234",
		Prefix: "/very/long/prefix/",
		URL:    "/very/long/prefix/0a3c6c93ee1693147335d14c4ce44867/5669eac1/secret-page.html"}

	c.Decode()

	if c.Filename != "/secret-page.html" {
		t.Error("Unexpected filename " + c.Filename)
		return
	}

	if c.Status == StatusExpired {
		t.Error("This url should never expire...")
		return
	}
}

func BenchmarkEncode(b *testing.B) {
	c := &SecDl{
		Secret:   "1234",
		Prefix:   "/dl/",
		Filename: "/dir1/dir2/dir3/secret-file.txt"}

	for i := 0; i < b.N; i++ {
		c.Encode(ExpireHour)
	}
}

func BenchmarkDecode(b *testing.B) {
	c := &SecDl{
		Secret:   "1234",
		Prefix:   "/dl/",
		Filename: "/dir1/dir2/dir3/secret-file.txt"}
	c.Encode(ExpireNever)

	for i := 0; i < b.N; i++ {
		c.Decode()
	}
}

func ExampleFileServer() {
	secret := "1234"
	prefix := "/"
	root, _ := os.Getwd()

	// Create the fileserver with prefix "/" and run in background
	http.Handle(prefix, FileServer(secret, prefix, root))
	go http.ListenAndServe(":1234", nil)

	// Wait a little so the HTTP server is listening
	time.Sleep(1 * time.Second)

	// Try to HTTP GET the LICENSE file behind a never-expirable link (secret: "1234")
	res, err := http.Get("http://127.0.0.1:1234/1ce1150e64d190a585e68ca630ddd634/568ae08b/LICENSE")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read content into variable
	readme, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("LICENSE: %d bytes\n", len(readme))

	// Output: LICENSE: 1107 bytes
}
