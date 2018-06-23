package main

import (
"github.com/miekg/dns"
    "net"
    "fmt"
    "log"
)

func LookupSSHFP(host string) ([]string, error) {
	c := new(dns.Client)

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(host), dns.TypeSSHFP)
	m.RecursionDesired = true
	r, _, err := c.Exchange(m, net.JoinHostPort("ns1.transip.nl", "53"))

	if err != nil {
		log.Print(err)
		return nil, err
	}

    if r.Rcode != dns.RcodeSuccess {
            log.Fatalf(" *** invalid answer name after SSHFP query\n")
    }
    // Stuff must be in the answer section
    for _, a := range r.Answer {
            fmt.Printf("%v\n", a)
    }
	fmt.Println("done")

	return nil, nil
}
