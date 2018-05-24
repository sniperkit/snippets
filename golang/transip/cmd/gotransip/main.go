package main

import (
	"flag"
	"fmt"

	"github.com/xor-gate/go-transip"
)

var login string
var privateKey string
var domain string

func init() {
	_login := flag.String("login", "foo", "login name of the api")
	_privateKey := flag.String("private-key", "key.pem", "private key for login (pem file)")
	_domain := flag.String("domain-info", "foo.com", "domain to get info for")

	flag.Parse()

	login = *_login
	privateKey = *_privateKey
	domain = *_domain
}

func main() {
	api := transip.APISettingsDefaults()
	api.SetLogin(login)
	api.SetPrivateKey(privateKey)
	fmt.Println(api)
	if domain != "" {
		fmt.Println(api.DomainServiceGetInfo(domain))
	}
}
