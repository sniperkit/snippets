package transip

import (
	"io"
	"net/http"
	"os"
)

func (a *APISettings) WsdlDownloadToFile(filename, serviceName string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(a.uriWsdl(serviceName))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
