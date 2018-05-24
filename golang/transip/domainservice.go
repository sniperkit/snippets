package transip

import (
	"encoding/xml"
)

const domainServiceName = "DomainService"

type domainServiceGetInfo struct {
	XMLName    xml.Name   `xml:"ns1:getInfo"`
	DomainName domainName `xml:","`
}

type domainName struct {
	XMLName xml.Name `xml:"domainName"`
	Type    string   `xml:"xsi:type,attr"`
	TLD     string   `xml:",chardata"`
}

func (a *APISettings) DomainServiceGetInfo(TLD string) error {
	bla := SOAPEnvelope{
		ENV:           "http://schemas.xmlsoap.org/soap/envelope/",
		NS1:           "http://www.transip.nl/soap",
		XSD:           "http://www.w3.org/2001/XMLSchema",
		XSI:           "http://www.w3.org/2001/XMLSchema-instance",
		Enc:           "http://schemas.xmlsoap.org/soap/encoding/",
		EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
	}
	bla.Body.Content = &domainServiceGetInfo{DomainName: domainName{Type: "xsd:string", TLD: TLD}}

	b, _ := xml.Marshal(&bla)
	s := xml.Header + string(b)

	return a.request(domainServiceName, "getInfo", s, TLD)
}
