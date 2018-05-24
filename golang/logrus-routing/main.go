package main

import (
	"os"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	logf, _ := os.Create("main.log")

	log := logrus.New()
	log.Out = logf
	logerrw := log.WriterLevel(logrus.ErrorLevel)

	log.Print("Hello print")
	fmt.Fprintln(logerrw, "Hello error")
}
