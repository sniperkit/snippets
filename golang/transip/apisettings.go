package transip

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

type APIMode string

const (
	APIModeReadOnly  APIMode = "readonly"
	APIModeReadWrite         = "readwrite"
)

const APISettingsDefaultClientVersion = "5.4"
const APISettingsDefaultEndpoint = "api.transip.nl"
const APISettingsDefaultMode = APIModeReadOnly

type APISettings struct {
	ClientVersion string
	Mode          APIMode
	Endpoint      string
	Login         string
	PrivateKey    string
	privateKey    interface{}
}

func (a *APISettings) String() string {
	summary := `
APISettings:
	 Client version: %v
	 Endpoint: %v
	 Mode: %v
	 Login: %v
	 PrivateKey: %v`

	return fmt.Sprintf(summary, a.ClientVersion, a.Endpoint, a.Mode, a.Login, a.PrivateKey)
}

func APISettingsDefaults() *APISettings {
	return &APISettings{
		ClientVersion: APISettingsDefaultClientVersion,
		Endpoint:      APISettingsDefaultEndpoint,
		Mode:          APISettingsDefaultMode,
	}
}

func (a *APISettings) SetLogin(login string) {
	a.Login = login
}

func (a *APISettings) SetPrivateKey(filename string) error {
	if err := a.loadPrivateKey(filename); err != nil {
		return err
	}
	a.PrivateKey = filename
	return nil
}

func (a *APISettings) loadPrivateKey(filename string) error {
	c, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(c)
	if block == nil {
		return nil
	}

	a.privateKey, _ = x509.ParsePKCS8PrivateKey(block.Bytes)

	return nil
}
