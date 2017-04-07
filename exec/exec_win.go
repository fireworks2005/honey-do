//  +build windows

//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/exec/exec.go
//
package exec

import (
	"os"
	"os/exec"
)

var (
	shell  string
	exists bool
)

func init() {
	var err error
	shell, err = exec.LookPath("sh")
	exists = err == nil
}

func Command(cmd string) *Cmd {
	c := &Cmd{exec.Command(shell, "-c", cmd)}
	c.Env = os.Environ()
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if !exists {
		c.Cmd = exec.Command("cmd", "/C", cmd)
	}

	return c
}
