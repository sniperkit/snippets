package main

import (
	"os"
	"flag"
	"io/ioutil"

	"gopkg.in/gomail.v2"
)

var from *string
var to *string
var subject *string
var server *string
var port *int
var username *string
var password *string
var attachment *string

func init() {
	from = flag.String("from", "", "From email address")
	to = flag.String("to", "", "To email address")
	subject = flag.String("subject", "", "To email subject")
	server = flag.String("server", "smtp.ziggozakelijk.nl", "The SMTP server host")
	port = flag.Int("port", 25, "The SMTP server port")
	username = flag.String("username", "", "The SMTP username")
	password = flag.String("password", "", "The SMTP password")
	attachment = flag.String("attachment", "", "File attachment")

	flag.Parse()
}

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", *from)
	m.SetHeader("To", *to)
	m.SetHeader("Subject", *subject)

	msg, _ := ioutil.ReadAll(os.Stdin)
	m.SetBody("text/plain", string(msg))

	if *attachment != "" {
		m.Attach(*attachment)
	}

	d := gomail.NewPlainDialer(*server, *port, *username, *password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
