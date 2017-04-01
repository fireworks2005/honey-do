//  +build !windows

//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/exec/exec.go
//
package exec

import (
	"os"
	"os/exec"
)

var shell string

func init() {
	var err error
	if shell, err = exec.LookPath("bash"); err != nil {
		shell, _ = exec.LookPath("sh")
	}
}

func Command(cmd string) *Cmd {
	c := &Cmd{exec.Command(shell, "-c", cmd)}
	c.Env = os.Environ()
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c
}
