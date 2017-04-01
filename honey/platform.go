//  Created by Elliott Polk on 01/04/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/platform.go
//
package honey

import (
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

type Platform struct {
	os      string
	Actions []string          `json:"actions,omitempty" yaml:"actions,omitempty"`
	Vars    map[string]string `json:"vars,omitempty" yaml:"vars,omitempty"`
}

func (p *Platform) Run() error {
	if !p.runnable() {
		return nil
	}

	doers, err := p.PrepDoers()
	if err != nil {
		return errors.Wrap(err, "unable to prep actions")
	}

	for _, d := range doers {
		if err := d.Do(); err != nil {
			return errors.Wrap(err, "unable to run platform action")
		}
	}

	return nil
}

func (p *Platform) PrepDoers() ([]*Doer, error) {
	doers := make([]*Doer, 0)
	for _, a := range p.Actions {
		raw, err := enrich(a, vars(p.Vars))
		if err != nil {
			return nil, errors.Wrap(err, "unable to enrich action")
		}

		doers = append(doers, NewDoer(raw))
	}

	return doers, nil
}

func (p *Platform) runnable() bool {
	os := runtime.GOOS
	if strings.HasPrefix(p.os, "!") && strings.HasSuffix(p.os, os) {
		return false
	}

	if !strings.HasPrefix(p.os, "!") && p.os != os {
		return false
	}

	return true
}
