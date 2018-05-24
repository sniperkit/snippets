package transip

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
)

// Input: URL encoded params string
// Output: Base64 signed string
func (a *APISettings) authSign(params string) string {
	digest := sha512.Sum512([]byte(params))
	sig, err := rsa.SignPKCS1v15(nil, a.privateKey.(*rsa.PrivateKey), crypto.SHA512, digest[:])
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(sig)
}
