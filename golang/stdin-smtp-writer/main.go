package main

import (
	"os"
	"log"
	"flag"
	"net/smtp"
	"io/ioutil"
)

var from *string
var to *string
var subject *string
var server *string

func init() {
	from = flag.String("from", "", "From email address")
	to = flag.String("to", "", "To email address")
	subject = flag.String("subject", "", "To email subject")
	server = flag.String("server", "smtp.ziggozakelijk.nl:25", "The SMTP server")
	flag.Parse()
}

func main() {
	c, err := smtp.Dial(*server)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.Mail(*from)
	c.Rcpt(*to)

	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	wc.Write([]byte("From: " + *from + "\r\n"))
	wc.Write([]byte("To: " + *to + "\r\n"))
	wc.Write([]byte("Subject: " + *subject + "\r\n\r\n"))

	msg, err := ioutil.ReadAll(os.Stdin)
	wc.Write(msg)
}
