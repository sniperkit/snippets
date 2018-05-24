package sysdetect

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

// SSHSysDetector for local system information detection
type SSHSysDetector struct {
	c *ssh.Client
}

// NewSSH creates a SysDetector on a remote system the over SSH protocol
func NewSSH(c *ssh.Client) *SSHSysDetector {
	return &SSHSysDetector{c: c}
}

// Close will close the underlaying remote SSH connection
func (sd *SSHSysDetector) Close() error {
	return sd.c.Close()
}

// ReadFile from filename
func (sd *SSHSysDetector) ReadFile(filename string) (string, error) {
	return sd.RunCommand("cat", filename)
}

// Detect system information on a SSH remote
func (sd *SSHSysDetector) Detect() map[string]string {
	return Detect(sd)
}

// RunCommand executes a command with name and optional arg
func (sd *SSHSysDetector) RunCommand(name string, arg ...string) (string, error) {
	s, err := sd.c.NewSession()
	if err != nil {
		return "", err
	}
	defer s.Close()

	var cmd string

	if len(arg) == 0 {
		cmd = name
	} else {
		cmd = fmt.Sprintf("%s %s", name, strings.Join(arg, " "))
	}

	out, err := s.Output(cmd)
	return string(out), err
}

// LookupEnv retrieves the value of the environment variable named by the key
func (sd *SSHSysDetector) LookupEnv(key string) (string, bool) {
	prefix := fmt.Sprintf("%s=", key)
	env, err := sd.RunCommand("env", "|", "grep", prefix)
	if err != nil {
		return "", false
	}
	if env == "" {
		return "", false
	}
	return strings.TrimPrefix(env, prefix), true
}

var _ SysDetector = &SSHSysDetector{}
