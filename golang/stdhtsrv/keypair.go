package main

import (
	"fmt"
	"math/big"
	"errors"
	"io/ioutil"
	"encoding/pem"
	"crypto"
	"crypto/rand"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/ecdsa"
	"crypto/x509"
)

// SignatureScheme identifies a supported signature algorithm based on JWT see https://tools.ietf.org/html/rfc7518
type KeyPair struct {
	Hash crypto.Hash
	PublicKey crypto.PublicKey
	PrivateKey crypto.PrivateKey
}

func NewKeyPair() *KeyPair {
	return &KeyPair{
		Hash: crypto.SHA256,
	}
}

func (kp *KeyPair) Alg() string {
	switch kp.PublicKey.(type) {
	case *rsa.PublicKey:
		return fmt.Sprintf("RS%v",kp.Hash.Size()*8)
	case *ecdsa.PublicKey:
		return fmt.Sprintf("ES%v",kp.Hash.Size()*8)
	default:
		return ""
	}
}

func (kp *KeyPair) PrivateKeyFile(filename string) error {
	pemBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var privateKey crypto.PrivateKey

	block, _ := pem.Decode(pemBytes)
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		kp.PrivateKey = privateKey
		kp.PublicKey  = privateKey.(*rsa.PrivateKey).Public()
		return nil
	}

	privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err == nil {
		kp.PrivateKey = privateKey
		kp.PublicKey  = privateKey.(*ecdsa.PrivateKey).Public()

		keyBits := kp.PrivateKey.(*ecdsa.PrivateKey).Curve.Params().BitSize
		switch keyBits {
		case 256:
			kp.Hash = crypto.SHA256
		case 384:
			kp.Hash = crypto.SHA384
		case 512:
			kp.Hash = crypto.SHA512
		default:
			return errors.New("invalid ecdsa key length")
		}
		return nil
	}

	return errors.New("unable to load private key")
}

func (kp *KeyPair) Sign(data []byte) ([]byte, error) {
	switch kp.PrivateKey.(type) {
	case *rsa.PrivateKey:
		return []byte{}, errors.New("rsa not supported")
	case *ecdsa.PrivateKey:
		return kp.ecdsaSign(data)
	}
	return nil, errors.New("private key not loaded")
}

func (kp *KeyPair) Verify(data []byte, signature []byte) bool {
	switch publicKey := kp.PublicKey.(type) {
	case *rsa.PublicKey:
		panic("rsa not supported")
		return false
	case *ecdsa.PublicKey:
		return kp.ecdsaVerify(publicKey, data, signature)
	default:
	}
	return false
}

func (kp *KeyPair) ecdsaEncodeRS(params *elliptic.CurveParams, r *big.Int, s *big.Int) []byte {
	curveBits := params.BitSize

	// TODO must be calculated beforehand
	keyBytes := curveBits / 8
	if curveBits%8 > 0 {
		keyBytes += 1
	}

	// We serialize the outpus (r and s) into big-endian byte arrays and pad
	// them with zeros on the left to make sure the sizes work out. Both arrays
	// must be keyBytes long, and the output must be 2*keyBytes long.
	rBytes := r.Bytes()
	rBytesPadded := make([]byte, keyBytes)
	copy(rBytesPadded[keyBytes-len(rBytes):], rBytes)

	sBytes := s.Bytes()
	sBytesPadded := make([]byte, keyBytes)
	copy(sBytesPadded[keyBytes-len(sBytes):], sBytes)

	return append(rBytesPadded, sBytesPadded...)
}

func (kp *KeyPair) ecdsaDecodeRS(keySizeBytes int, rsBytes []byte) (r *big.Int, s *big.Int) {
	r = big.NewInt(0).SetBytes(rsBytes[:keySizeBytes])
	s = big.NewInt(0).SetBytes(rsBytes[keySizeBytes:])
	return
}

func (kp *KeyPair) ecdsaSign(data []byte) ([]byte, error) {
	digest, _ := kp.hash(data)

	r, s, err := ecdsa.Sign(rand.Reader, kp.PrivateKey.(*ecdsa.PrivateKey), digest[:])
	if err != nil {
		return nil, err
	}

	sig := kp.ecdsaEncodeRS(kp.PrivateKey.(*ecdsa.PrivateKey).Curve.Params(), r, s)
	return sig, nil
}

func (kp *KeyPair) hash(data []byte) ([]byte, error) {
	if !kp.Hash.Available() {
		return nil, errors.New("hash unavailable")
	}
	hasher := kp.Hash.New()
	hasher.Write(data)
	return hasher.Sum(nil), nil
}

func (kp *KeyPair) ecdsaVerify(publicKey *ecdsa.PublicKey, data, signature []byte) bool {
	// TODO check if keysizes match
	// just like https://github.com/dgrijalva/jwt-go/blob/master/ecdsa.go#L120-L123
	digest, err := kp.hash(data)
	if err != nil {
		return false
	}
	// TODO keySizeBytes must be calculated beforehand, now it slows down verify
	keySizeBytes := kp.PrivateKey.(*ecdsa.PrivateKey).Curve.Params().BitSize / 8
	r, s := kp.ecdsaDecodeRS(keySizeBytes, signature)
	return ecdsa.Verify(publicKey, digest[:], r, s)
}
