package main

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"strings"
	"log"
	"io/ioutil"
	"golang.org/x/crypto/scrypt"
)

func main() {
	rc, err := encrypt("boembats", strings.NewReader("boembats"))
	if err != nil {
		log.Fatal(err)
	}
	enc, err := ioutil.ReadAll(rc)
	fmt.Printf("enc %x (len %d)\n", enc, len(enc))
}

func encrypt(password string, r io.Reader) (io.ReadCloser, error) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}

	hpassword, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	fmt.Printf("hpassword %x\n", hpassword)
	fmt.Printf("salt %x\n", salt)
	fmt.Printf("iv %x\n", iv)

	block, err := aes.NewCipher(hpassword)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)

	pr, pw := io.Pipe()
	go func() {
		defer pr.Close()
		defer pw.Close()
		writer := cipher.StreamWriter{S: stream, W: pw}
		io.Copy(writer, r)
	}()

	return pr, nil
}
