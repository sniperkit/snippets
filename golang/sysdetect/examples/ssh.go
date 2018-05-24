// +build ignore

package main

import (
	"io/ioutil"
	"log"

	"github.com/xor-gate/snippets/golang/sysdetect"
	"golang.org/x/crypto/ssh"
)

func main() {
	c, err := newSSHClient("cicd", "192.168.1.199:22", "/Users/jerry/.ssh/id_rsa")
	if err != nil {
		log.Fatal(err)
	}

	sd := sysdetect.NewSSH(c)

	result := sd.Detect()
	for key, value := range result {
		log.Printf("%s=%s", key, value)
	}
}

func newSSHClient(user, host, keyfile string) (*ssh.Client, error) {
	pkey, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	// privateKey could be read from a file, or retrieved from another storage
	// source, such as the Secret Service / GNOME Keyring
	key, err := ssh.ParsePrivateKey(pkey)
	if err != nil {
		return nil, err
	}

	// Authentication
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", host, config)
}
