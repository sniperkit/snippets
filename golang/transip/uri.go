package transip

func (a *APISettings) uriWsdl(service string) string {
	return "https://" + a.Endpoint + "/wsdl/?service=" + service
}

func (a *APISettings) uriSoap(service string) string {
	return "https://" + a.Endpoint + "/soap/?service=" + service
}
