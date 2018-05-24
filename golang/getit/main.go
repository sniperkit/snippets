package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func downloadSFTP(dir, prefix string, u *url.URL) (string, error) {
	config := &ssh.ClientConfig{
		User:            u.User.Username(),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if password, ok := u.User.Password(); ok {
		config.Auth = append(config.Auth, ssh.Password(password))
	}

	port := u.Port()
	if port == "" {
		port = "22"
	}

	sshc, err := ssh.Dial("tcp", net.JoinHostPort(u.Hostname(), port), config)
	if err != nil {
		return "", err
	}

	sftpc, err := sftp.NewClient(sshc)
	if err != nil {
		sshc.Close()
		return "", err
	}
	defer sftpc.Close()

	wd, err := sftpc.Getwd()
	if err != nil {
		return "", err
	}

	fmt.Println("path", u.Path)

	rf, err := sftpc.Open(sftp.Join(wd, u.Path))
	if err != nil {
		return "", err
	}
	defer rf.Close()

	out, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, rf)
	if err != nil {
		return "", err
	}

	return out.Name(), nil
}

func downloadHTTP(dir, prefix, url string) (string, error) {
	out, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return out.Name(), nil
}

func Download(dir, prefix, URI string) (string, error) {
	u, err := url.Parse(URI)
	if err != nil {
		return "", err
	}

	switch u.Scheme {
	case "http", "https":
		return downloadHTTP(dir, prefix, URI)
	case "sftp":
		return downloadSFTP(dir, prefix, u)
	default:
		return "", fmt.Errorf("unsupported protocol scheme %s", u.Scheme)
	}
}

func main() {
/*
	_, err := Download(".", "boem", "https://github.com/golang/dep/archive/v0.4.1.zip")
	if err != nil {
		log.Fatal(err)
	}
*/
	_, err := Download(".", "bam", "sftp://build:build@192.168.1.201/myfile")
	if err != nil {
		log.Fatal(err)
	}
}
