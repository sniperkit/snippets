package main

import (
	"io"
	"crypto/sha1"
	"crypto/x509"
	"fmt"
	"encoding/pem"
	"os"
	"errors"
	"log"
	"net/url"
	"crypto/tls"
	"github.com/syncthing/syncthing/lib/config"
	"github.com/syncthing/syncthing/lib/discover"
	"github.com/syncthing/syncthing/lib/protocol"
)

const DeviceID = "LBR4YOY-CUZ2YXQ-AMRQ7ZG-RJSRM7F-FXEZXXE-UFB64PV-5NFPYLZ-CGW46Q7"

// CertLoadPem
func CertLoadPem(pemBytes []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(pemBytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

// CertWritePem writes the certificate with PEM encoding
func CertWritePem(out io.Writer, cert *x509.Certificate) error {
        block := &pem.Block{"CERTIFICATE", nil, cert.Raw}
        return pem.Encode(out, block)
}

// CertDownload downloads the tls certificate downloads and verifies for deviceID on uri
func CertDownload(deviceID protocol.DeviceID, uri *url.URL) (*x509.Certificate, error) {
	// Connect to the syncthing instance
	cfg := &tls.Config{
	        InsecureSkipVerify: true,
	}

	conn, err := tls.Dial(uri.Scheme, uri.Host, cfg)
	if err != nil {
		return nil, err
	}

	// We expect exactly one certificate.
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) != 1 {
		return nil, fmt.Errorf("unexpected number of certificates (%d != 1)", len(certs))
	}

	// Validate the certificate is from the expected device ID 
	hash := sha1.New()
	hash.Write(certs[0].Raw)
	devID := protocol.NewDeviceID(certs[0].Raw)

	if !deviceID.Equals(devID) {
		return nil, errors.New("unexpected certificate")
	}

	return certs[0], nil
}

// Lookup the deviceID on the global discovery server on which uris it is reachable
func Lookup(deviceID protocol.DeviceID) ([]*url.URL, error) {
	gdisco, err := discover.NewGlobal(config.DefaultDiscoveryServers[0], tls.Certificate{}, nil)
	if err != nil {
		return nil, err
	}

	str, err := gdisco.Lookup(deviceID)
	if err != nil {
		return nil, err
	}

	var uris []*url.URL

	for _, addr := range str {
		uri, err := url.Parse(addr)
		if err != nil {
			continue
		}
		uris = append(uris, uri)
	}

	return uris, nil
}

func main() {
	certStore := NewMemCertificateStore()

	deviceID, err := protocol.DeviceIDFromString(DeviceID)
	if err != nil {
		log.Fatal(err)
	}

	uris, err := Lookup(deviceID)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range uris {
		// TODO we only support direct connection for now
		//  need a way to use relay servers
		if u.Scheme != "tcp" {
			continue
		}

		fmt.Println("Download certificate on", u)
		cert, err := CertDownload(deviceID, u)
		if err == nil {
			fmt.Println("Certificate for", deviceID)
			certStore.Set(deviceID, cert)
			CertWritePem(os.Stdout, cert)
		}
	}

	publicKey, err := certStore.GetPublicKey(deviceID)
	if err != nil {
		log.Fatal(err)
	}

	keyPair := NewKeyPair()
	keyPair.PrivateKeyFile("/Users/jerry/Library/Application Support/Syncthing/key.pem")

	fmt.Printf("cert publickey for %s %T\n", deviceID, publicKey)
	fmt.Printf("     publickey %T\n", keyPair.PublicKey)
	fmt.Printf("     privatekey %T\n", keyPair.PrivateKey)

	data := []byte("jerry.jacobs@xor-gate.org")
	fmt.Println("alg:", keyPair.Alg())
	signature, err := keyPair.Sign(data)
	fmt.Printf("signature %X %v\n", signature, len(signature))
	fmt.Println("valid:", keyPair.Verify(data, signature))
}
