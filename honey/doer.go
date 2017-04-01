//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/doer.go
//
package honey

import (
	"bytes"
	"strings"
	"time"

	"github.com/elliottpolk/honey-do/exec"
	"github.com/elliottpolk/honey-do/log"

	"github.com/pkg/errors"
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

func (d *Doer) Do() error {
	d.start = time.Now()

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := d.Cmd
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	d.finish = time.Now()

	if out := stdout.String(); len(out) > 0 {
		log.Info(out)
	}

	if err != nil {
		if err := stderr.String(); len(err) > 0 {
			log.NewError("%s", strings.TrimSpace(err))
		}

		return errors.Wrapf(err, "unable to run command %s", d.Raw)
	}

	return nil
}
