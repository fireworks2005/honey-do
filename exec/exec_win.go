//  +build windows

//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/exec/exec.go
//
package exec

import "os/exec"

var (
	shell  string
	exists bool
)

func init() {
	shell, err := exec.LookPath("sh")
	exists = err == nil
}

func Command(name string, args ...string) *Cmd {
	if !exists {
		return &Cmd{exec.Command("cmd", "/C", name, args...)}
	}

	return &Cmd{exec.Command(shell, "-c", name, args...)}
}
