package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHAgent struct {
	name      string
	workspace string
	host      string
	user      string
	key       []byte
	resources *Resources
}

type SSHAgentInfo struct {
	a *SSHAgent
}

func NewSSHAgent(name string) *SSHAgent {
	return &SSHAgent{
		name:      name,
		resources: NewResources(),
	}
}

func (a *SSHAgent) Name() string {
	return a.name
}

func (a *SSHAgent) SetUser(user string) {
	a.user = user
}

func (a *SSHAgent) SetHost(host string) {
	a.host = host
}

func (a *SSHAgent) Host() string {
	return a.host
}

func (a *SSHAgent) User() string {
	return a.user
}

func (a *SSHAgent) SetWorkspace(w string) {
	a.workspace = w
}

func (a *SSHAgent) Workspace() string {
	return a.workspace
}

func (a *SSHAgent) AddResource(r *Resource) {
	a.resources.Add(r)
}

func (a *SSHAgent) NewResource(name string) *Resource {
	return a.resources.NewResource(name)
}

func (a *SSHAgent) SetPrivateKeyFromFile(filename string) error {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	a.key = key
	return nil
}

func (a *SSHAgent) RunTask(t *Task) error {
	// privateKey could be read from a file, or retrieved from another storage
	// source, such as the Secret Service / GNOME Keyring
	key, err := ssh.ParsePrivateKey(a.key)
	if err != nil {
		return err
	}
	// Authentication
	config := &ssh.ClientConfig{
		User: a.user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//alternatively, you could use a password
		/*
		   Auth: []ssh.AuthMethod{
		       ssh.Password("PASSWORD"),
		   },
		*/
	}
	// Connect
	client, err := ssh.Dial("tcp", a.host, config)
	if err != nil {
		return err
	}
	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// you can also pass what gets input to the stdin, allowing you to pipe
	// content from client to server
	//      session.Stdin = bytes.NewBufferString("My input")

	// Finally, run the command
	err = session.Run(t.CommandString())
	fmt.Println(b.String())
	return err
}

func (a *SSHAgent) PrivateKey() []byte {
	return a.key
}

func (a *SSHAgent) Stat() AgentInfo {
	return &SSHAgentInfo{a: a}
}

func (ai *SSHAgentInfo) Type() AgentType {
	return AgentTypeSSH
}

func (ai *SSHAgentInfo) OS() string {
	return AgentOSUnknown
}

func (ai *SSHAgentInfo) Freespace() uint64 {
	return 0
}

func (ai *SSHAgentInfo) LastSeen() time.Time {
	return time.Time{}
}

func (ai *SSHAgentInfo) Workspace() string {
	return ""
}

func (ai *SSHAgentInfo) Resources() []string {
	return []string{}
}

var _ Agent = &SSHAgent{}
var _ AgentInfo = &SSHAgentInfo{}
