package main

import (
	"fmt"
	"log"
	"go/format"
	"bytes"
	gen "github.com/hooklift/gowsdl"
	"github.com/xor-gate/go-transip-api/transip"
)

func main() {
	c := transip.APISettingsDefaults()
	c.WsdlDownloadToFile("DomainService.wsdl", "DomainService")

	// load wsdl
	gowsdl, err := gen.NewGoWSDL("DomainService.wsdl", "domainservice", false, true)
	if err != nil {
		log.Fatalln(err)
	}

	gocode, err := gowsdl.Start()
	if err != nil {
		log.Fatalln(err)
	}

	data := new(bytes.Buffer)
	data.Write(gocode["header"])
	data.Write(gocode["types"])
	data.Write(gocode["operations"])
	data.Write(gocode["soap"])

	// go fmt the generated code
	src, err := format.Source(data.Bytes())
	if err != nil {
		log.Fatalln(err)
	}

	//src := data.Bytes()

	fmt.Println(string(src))
}
