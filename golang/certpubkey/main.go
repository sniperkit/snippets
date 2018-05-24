package main

import (
	"log"
	"bytes"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"crypto/sha1"
	"encoding/pem"
	"fmt"
	"math/big"
)

func main() {
	rsaPubkeyCert()
	rsaPubkeyEncryption()
	rsaPubkeyTest()

}

// https://play.golang.org/p/pn7CKVVdQF
func rsaPubkeyCert() {
	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
/iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
-----END CERTIFICATE-----`

	block, _ := pem.Decode([]byte(rootPEM))
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
	fmt.Println(rsaPublicKey.N)
	fmt.Println(rsaPublicKey.E)
}

// https://gist.github.com/hansstimer/3517835
func rsaPubkeyEncryption() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		fmt.Printf("rsa.GenerateKey: %v\n", err)
	}

	message := "jerry.jacobs@xor-gate.org"
	messageBytes := bytes.NewBufferString(message)
	sha1 := sha1.New()

	encrypted, err := rsa.EncryptOAEP(sha1, rand.Reader, &privateKey.PublicKey, messageBytes.Bytes(), nil)
	if err != nil {
		fmt.Printf("EncryptOAEP: %s\n", err)
	}

	decrypted, err := rsa.DecryptOAEP(sha1, rand.Reader, privateKey, encrypted, nil)
	if err != nil {
		fmt.Printf("decrypt: %s\n", err)
	}

	decryptedString := bytes.NewBuffer(decrypted).String()
	fmt.Printf("message: %v\n", message)
	fmt.Printf("encrypted: %v\n", encrypted)
	fmt.Printf("decryptedString: %v\n", decryptedString)
}

// https://stackoverflow.com/questions/24311575/why-does-the-go-function-encryptoaep-in-the-crypto-rsa-library-require-a-random?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
func rsaPubkeyTest() {
    sha1 := sha1.New()
    n := new(big.Int)
    d := new(big.Int)

    rsa_modulus := "a8b3b284af8eb50b387034a860f146c4919f318763cd6c5598c8ae4811a1e0abc4c7e0b082d693a5e7fced675cf4668512772c0cbc64a742c6c630f533c8cc72f62ae833c40bf25842e984bb78bdbf97c0107d55bdb662f5c4e0fab9845cb5148ef7392dd3aaff93ae1e6b667bb3d4247616d4f5ba10d4cfd226de88d39f16fb"
    rsa_d := "53339cfdb79fc8466a655c7316aca85c55fd8f6dd898fdaf119517ef4f52e8fd8e258df93fee180fa0e4ab29693cd83b152a553d4ac4d1812b8b9fa5af0e7f55fe7304df41570926f3311f15c4d65a732c483116ee3d3d2d0af3549ad9bf7cbfb78ad884f84d5beb04724dc7369b31def37d0cf539e9cfcdd3de653729ead5d1"

    n.SetString(rsa_modulus, 16)
    d.SetString(rsa_d, 16)
    public := rsa.PublicKey{n, 65537}
    d.SetString(rsa_d, 16)
    private := new(rsa.PrivateKey)
    private.PublicKey = public
    private.D = d

    seed := []byte{0x18, 0xb7, 0x76, 0xea, 0x21, 0x06, 0x9d, 0x69,
        0x77, 0x6a, 0x33, 0xe9, 0x6b, 0xad, 0x48, 0xe1, 0xdd,
        0xa0, 0xa5, 0xef,
    }
    randomSource := bytes.NewReader(seed)

    in := []byte("Hello World")

    encrypted, err := rsa.EncryptOAEP(sha1, randomSource, &public, in, nil)
    if err != nil {
        log.Println("error: %s", err)
    }

    plain, err := rsa.DecryptOAEP(sha1, nil, private, encrypted, nil)
    if err != nil {
        log.Println("error: %s", err)
    }

    log.Println(string(plain))
}
