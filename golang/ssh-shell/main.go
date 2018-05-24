package main

import (
	"fmt"
	"log"
	"time"
	"io/ioutil"
	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	key []byte
	c *ssh.Client
	host string
	cfg *ssh.ClientConfig
}

func NewSSHClient() *SSHClient {
	return &SSHClient{
		cfg: &ssh.ClientConfig {
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	}
}

func (c *SSHClient) SetPrivateKeyFromFile(filename string) error {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	pkey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}
	c.cfg.Auth = []ssh.AuthMethod {
		ssh.PublicKeys(pkey),
	}
	return nil
}

func (c *SSHClient) connect() error {
	sc, err := ssh.Dial("tcp", c.host, c.cfg)
	if err != nil {
		return err
	}
	c.c = sc
	return nil
}

func (c *SSHClient) SetUser(username string) {
	c.cfg.User = username
}

func (c *SSHClient) SetHost(hostname string) {
	c.host = hostname
}

func (c *SSHClient) Run(cmd string) error {
	if c.c == nil {
		if err := c.connect(); err != nil {
			return err
		}
	}

	s, err := c.c.NewSession()
	if err != nil {
		return err
	}
	defer s.Close()

	res, err := s.CombinedOutput(cmd)
	fmt.Printf("%s",string(res))
	return err
}

func (c *SSHClient) RunAsync(cmd string) func () {
	if c.c == nil {
		if err := c.connect(); err != nil {
			return func(){}
		}
	}

	s, err := c.c.NewSession()
	if err != nil {
		return func(){}
	}

modes := ssh.TerminalModes{
    ssh.ECHO:          0,     // disable echoing
    ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
    ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
}
// Request pseudo terminal
if err := s.RequestPty("xterm", 40, 80, modes); err != nil {
    log.Fatal("request for pseudo terminal failed: ", err)
	return func(){}
}

	go func() {
		fmt.Println("... running async")
		res, err := s.CombinedOutput(cmd)
		fmt.Printf("... done async (%v): %s",err,string(res))
	}()
	return func() {
		fmt.Println("killing the async cmd...")
		//s.Signal(ssh.SIGKILL)
		s.Close()
		fmt.Println("closed...")
	}
}

func main() {
	s := NewSSHClient()
	s.SetPrivateKeyFromFile("/Users/jerry/.ssh/id_rsa_muci")
	s.SetHost("192.168.1.201:22")
	s.SetUser("muci")
	s.Run(`echo '-- environment';env; echo '---'`)
	s.Run(`echo "hello world from go ssh client!"`)
	s.Run(`echo -- working for 300 seconds...`)
	killer := s.RunAsync(`echo "doing some stuff....";sleep 300`)
	time.Sleep(3*time.Second)
	killer()
}
