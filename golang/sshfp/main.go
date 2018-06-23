package main

import (
	"fmt"
)

func main() {
	res := &SSHFPResolver{}
	ssh := NewSSHClient(res.HostKeyCallback)
	err := ssh.Connect("shulgin.xor-gate.org")
	fmt.Println(err)
}
