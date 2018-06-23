package main

import (
	"github.com/miekg/dns"
    "net"
    "fmt"
    "log"
	"bytes"
	"golang.org/x/crypto/ssh"
	"net/url"
	"crypto/sha256"
	"encoding/hex"
)

type SSHFPResolver struct {
}

func (r *SSHFPResolver) HostKeyCallback(hostname string, remote net.Addr, key ssh.PublicKey) error {
	fmt.Println("SSH check:", hostname, remote, key)
	l, err := r.Lookup(hostname)
	if err != nil {
		return err
	}

	keyFpSHA256 := sha256.Sum256(key.Marshal())
	fmt.Println("SSH pubkey SHA256:", keyFpSHA256)

	fmt.Println("DNS:")
	for _, sshfp := range l {
            fmt.Printf("%v\n", sshfp)
		raw, _ := hex.DecodeString(sshfp.FingerPrint)
		fmt.Println("raw", raw)

		// Check if there is a match
		if bytes.Equal(keyFpSHA256[:], raw) {
			fmt.Println("sshfp: good to go!")
			return nil
		}
	}

	fmt.Println("sshfp: poor setup, no SSHFP...")
	return fmt.Errorf("sshfp: no host key found")
}

func (r *SSHFPResolver) Lookup(host string) ([]*dns.SSHFP, error) {
	c := new(dns.Client)

	m := new(dns.Msg)

	// TODO: to ugly hack to be able to parse "shulgin.xor-gate.org:6222" ...
	hostURL, err := url.Parse("tcp://"+host)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	hostname := hostURL.Hostname()

	fmt.Println("lookup", hostname, hostURL)

	m.SetQuestion(dns.Fqdn(hostURL.Hostname()), dns.TypeSSHFP)
	m.RecursionDesired = true
	resp, _, err := c.Exchange(m, net.JoinHostPort("ns1.transip.nl", "53"))

	if err != nil {
		log.Print(err)
		return nil, err
	}

    if resp.Rcode != dns.RcodeSuccess {
            log.Fatalf(" *** invalid answer name after SSHFP query\n")
    }

	var l []*dns.SSHFP

    // Stuff must be in the answer section
    for _, a := range resp.Answer {
		sshfp, ok := a.(*dns.SSHFP)
		if !ok {
			continue
		}
		l = append(l, sshfp)
    }

	return l, nil
}
