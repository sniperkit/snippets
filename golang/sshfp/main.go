package main

import (
	"fmt"
)

func main() {
	res := &SSHFPResolver{}
	ssh := NewSSHClient(res.HostKeyCallback)
	ssh.SetPrivateKeyFromFile("/Users/jerry/.ssh/id_rsa")
	err := ssh.Connect("jerry", "shulgin.xor-gate.org:6222")
	fmt.Println(err)
}
