//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/doer.go
//
package honey

import (
	"time"

	"github.com/elliottpolk/honey-do/exec"
)

type Doer struct {
	Cmd    *exec.Cmd
	Raw    string
	start  time.Time
	finish time.Time
}

func NewDoer(raw string) *Doer {
	return &Doer{
		Cmd: exec.Command(raw),
		Raw: raw,
	}
}
