//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/target.go
//
package honey

import (
	"github.com/elliottpolk/honey-do/log"

	"github.com/pkg/errors"
)

type Target struct {
	ran bool

	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Deps        []string             `json:"deps,omitempty" yaml:"deps,omitempty"`
	Platforms   map[string]*Platform `json:"platforms,omitempty" yaml:"platforms,omitempty"`
	Actions     []string             `json:"actions,omitempty" yaml:"actions,omitempty"`
	Vars        map[string]string    `json:"vars,omitempty" yaml:"vars,omitempty"`
}

func (t *Target) Run() error {
	if t.ran {
		return nil
	}

	defer func() {
		t.ran = true
	}()

	if err := t.RunDeps(); err != nil {
		return errors.Wrap(err, "unable to run deps")
	}

	for n, p := range t.Platforms {
		p.os = n
		if err := p.Run(); err != nil {
			return errors.Wrap(err, "unable to run platform")
		}
	}

	doers, err := t.PrepDoers()
	if err != nil {
		return errors.Wrap(err, "unable to prep target actions")
	}

	for _, d := range doers {
		if err := d.Do(); err != nil {
			return errors.Wrap(err, "unable to run target action")
		}
	}

	return nil
}

func (t *Target) RunDeps() error {
	for _, dep := range t.Deps {
		t, ok := hf.Targets[dep]
		if !ok {
			log.NewError("invalid dep %s", dep)
			continue
		}

		if err := t.Run(); err != nil {
			return errors.Wrapf(err, "unable to run dep %s", dep)
		}

		log.Infof("dependency target %s complete", dep)
	}

	return nil
}

func (t *Target) PrepDoers() ([]*Doer, error) {
	doers := make([]*Doer, 0)
	for _, a := range t.Actions {
		raw, err := enrich(a, vars(t.Vars))
		if err != nil {
			return nil, errors.Wrap(err, "unable to enrich action")
		}

		doers = append(doers, NewDoer(raw))
	}

	return doers, nil
}
