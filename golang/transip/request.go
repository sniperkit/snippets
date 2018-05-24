package transip

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func generateNonce() string {
	b := make([]byte, 7)
	rand.Read(b)
	return fmt.Sprintf("%x.%d", b, time.Now().Unix())
}

func (a *APISettings) generateRequestCookie(timestamp int64, nonce, signature string) string {
	return fmt.Sprintf("login=%s;mode=%s;timestamp=%d;nonce=%s;clientVersion=%s;signature=%s;",
		a.Login, a.Mode, timestamp, nonce, a.ClientVersion, url.QueryEscape(signature))
}

func (a *APISettings) request(soapService, soapMethod, body, tld string) error {
	if a.privateKey == nil {
		return ErrPrivateKeyNotLoaded
	}

	reader := strings.NewReader(body)
	req, _ := http.NewRequest("POST", a.uriSoap(soapService), reader)

	timestamp := time.Now().Unix()
	nonce := generateNonce()
	params := fmt.Sprintf("0="+tld+"&__method="+soapMethod+"&__service="+soapService+"&__hostname="+a.Endpoint+"&__timestamp=%d&__nonce=%s", timestamp, nonce)
	signature := a.authSign(params)
	cookie := a.generateRequestCookie(timestamp, nonce, signature)

	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-type", "text/xml; charset=utf-8")

	client := &http.Client{}
	resp, _ := client.Do(req)
	data, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(data))

	return nil
}
