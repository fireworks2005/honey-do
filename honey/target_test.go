//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/target_test.go
//
package honey

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestRun(t *testing.T) {
	defer func() {
		if err := os.RemoveAll(HoneyFile + ".yml"); err != nil && !os.IsNotExist(err) {
			t.Error(errors.Wrap(err, "unable to remove yaml test file"))
			t.FailNow()
		}
	}()

	if err := writeYml(); err != nil {
		t.Error(errors.Wrap(err, "unable to generate test yaml file"))
		return
	}

	if err := LoadFile(); err != nil {
		t.Error(errors.Wrap(err, "unable to load test file"))
		return
	}

	for k, target := range hf.Targets {
		if err := target.Run(); err != nil {
			t.Error(errors.Wrapf(err, "unable to run target %s", k))
			continue
		}
	}
}

func TestRunDeps(t *testing.T) {
	defer func() {
		if err := os.RemoveAll(HoneyFile + ".yml"); err != nil && !os.IsNotExist(err) {
			t.Error(errors.Wrap(err, "unable to remove yaml test file"))
			t.FailNow()
		}
	}()

	if err := writeYml(); err != nil {
		t.Error(errors.Wrap(err, "unable to generate test yaml file"))
		return
	}

	if err := LoadFile(); err != nil {
		t.Error(errors.Wrap(err, "unable to load test file"))
		return
	}

	for k, target := range hf.Targets {
		if err := target.RunDeps(); err != nil {
			t.Error(errors.Wrapf(err, "unable to run deps for target %s", k))
			continue
		}
	}
}

func TestEnrich(t *testing.T) {
	res, err := enrich("{{.FOO}}", map[string]string{"FOO": "bar"})
	if err != nil {
		t.Error(errors.Wrap(err, "unable to enrich test content"))
		return
	}

	if want, got := "bar", res; want != got {
		t.Error("failed to enrich: wanted - %s ... got - %s", want, got)
	}
}
