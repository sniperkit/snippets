package main

import (
	"errors"
	"crypto"
	"crypto/x509"
	"github.com/syncthing/syncthing/lib/protocol"
)

type CertificateStore interface {
	Set(deviceID protocol.DeviceID, cert *x509.Certificate) error
	Get(deviceID protocol.DeviceID) (*x509.Certificate, error)
	GetPublicKey(deviceID protocol.DeviceID) (crypto.PublicKey, error)
}

type certStore struct {
	certs map[string]*x509.Certificate
}

func NewMemCertificateStore() CertificateStore {
	cs := &certStore{
		certs: make(map[string]*x509.Certificate),
	}
	return cs
}

func (cs *certStore) Set(deviceID protocol.DeviceID, cert *x509.Certificate) error {
	cs.certs[deviceID.String()] = cert
	return nil
}

func (cs *certStore) Get(deviceID protocol.DeviceID) (*x509.Certificate, error) {
	cert, ok := cs.certs[deviceID.String()]
	if !ok {
		return nil, errors.New("not found")
	}
	return cert, nil
}

func (cs *certStore) GetPublicKey(deviceID protocol.DeviceID) (crypto.PublicKey, error) {
	cert, err := cs.Get(deviceID)
	if err != nil {
		return nil, err
	}
	return cert.PublicKey, nil
}
