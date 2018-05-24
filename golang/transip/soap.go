package transip

import (
	"encoding/xml"
)

type SOAPEnvelope struct {
	XMLName       xml.Name    `xml:"SOAP-ENV:Envelope"`
	ENV           string      `xml:"xmlns:SOAP-ENV,attr,omitempty"`
	NS1           string      `xml:"xmlns:ns1,attr,omitempty"`
	XSD           string      `xml:"xmlns:xsd,attr,omitempty"`
	XSI           string      `xml:"xmlns:xsi,attr,omitempty"`
	EncodingStyle string      `xml:"SOAP-ENV:encodingStyle,attr,omitempty"`
	Enc           string      `xml:"xmlns:SOAP-ENC,attr"`
	Header        *SOAPHeader `xml:"http://schemas.xmlsoap.org/soap/envelope/ SOAP-ENV:Header",omitempty`
	Body          SOAPBody    `xml:"SOAP-ENV:Body"`
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ SOAP-ENV:Header"`
	Header  interface{}
}

type SOAPBody struct {
	XMLName xml.Name    `xml:"SOAP-ENV:Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}
