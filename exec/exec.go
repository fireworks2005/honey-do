//  +build !windows

//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/exec/exec.go
//
package exec

import "os/exec"

var shell string

func init() {
	var err error
	if shell, err = exec.LookPath("bash"); err != nil {
		shell, _ = exec.LookPath("sh")
	}
}

func Command(cmd string) *Cmd {
	return &Cmd{exec.Command(shell, "-c", cmd)}
}