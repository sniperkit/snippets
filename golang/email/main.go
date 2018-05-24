package main

import (
	"io/ioutil"

	"github.com/flosch/pongo2"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Email *EmailConfig
	SMTP  *SMTPConfig
}

type EmailConfig struct {
	From     string
	ReplyTo  string
	Subject  string
	BCC      []string
	Template string
}

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

func main() {
	var config Config

	source, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}

	tpl, err := pongo2.FromFile(config.Email.Template)
	if err != nil {
		panic(err)
	}

	body, err := tpl.Execute(nil)
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.Email.From)
	m.SetHeader("Reply-To", config.Email.ReplyTo)
	m.SetHeader("To", "media@vtvblixembosch.nl")
	if len(config.Email.BCC) > 0 {
		m.SetHeader("Bcc", config.Email.BCC...)
	}
	m.SetHeader("Subject", config.Email.Subject)

	m.SetBody("text/html", body)
	//m.Attach(config.Email.Template)

	d := gomail.NewDialer(config.SMTP.Host, config.SMTP.Port, config.SMTP.User, config.SMTP.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
