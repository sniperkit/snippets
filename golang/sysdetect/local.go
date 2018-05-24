package sysdetect

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// LocalSysDetector for local system information detection
type LocalSysDetector struct{}

// NewLocal creates a SysDetector for the local system
func NewLocal() *LocalSysDetector {
	return &LocalSysDetector{}
}

// Detect local system information
func (sd *LocalSysDetector) Detect() map[string]string {
	return Detect(sd)
}

// ReadFile from filename
func (sd *LocalSysDetector) ReadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// RunCommand executes a command with name and optional arg
func (sd *LocalSysDetector) RunCommand(name string, arg ...string) (string, error) {
	// #nosec
	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// LookupEnv retrieves the value of the environment variable named by the key
func (sd *LocalSysDetector) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

var _ SysDetector = &LocalSysDetector{}
