//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/target.go
//
package honey

import (
	"bytes"
	"strings"
	"text/template"
	"time"

	"github.com/elliottpolk/honey-do/log"

	"github.com/pkg/errors"
)

type Target struct {
	ran bool

	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Deps        []string          `json:"deps,omitempty" yaml:"deps,omitempty"`
	Actions     []string          `json:"actions,omitempty" yaml:"actions,omitempty"`
	Vars        map[string]string `json:"vars,omitempty" yaml:"vars,omitempty"`
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

	doers, err := t.PrepDoers()
	if err != nil {
		return errors.Wrap(err, "unable to prep actions")
	}

	for _, d := range doers {
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
	vars := make(map[string]string)
	for k, v := range hf.Vars {
		vars[k] = v
	}

	for k, v := range t.Vars {
		vars[k] = v
	}

	doers := make([]*Doer, 0)
	for _, a := range t.Actions {
		raw, err := enrich(a, vars)
		if err != nil {
			return nil, errors.Wrap(err, "unable to enrich action")
		}

		doers = append(doers, NewDoer(raw))
	}

	return doers, nil
}

func enrich(content string, vars interface{}) (string, error) {
	tpl := template.Must(template.New("cmd").Parse(content))

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		return "", errors.Wrap(err, "unable to execute text template")
	}

	return buf.String(), nil
}
