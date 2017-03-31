//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/file_test.go
//
package honey

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var TestFile = &File{
	Version: "1",
	Vars: map[string]string{
		"FOO": "bar",
		"BIZ": "baz",
	},
	Targets: map[string]*Target{
		"test1": &Target{
			Description: "Test Description 1",
			Actions:     []string{"echo 'test 1 {{.FOO}}'"},
			Vars: map[string]string{
				"FOO": "bar",
				"BIZ": "baz",
			},
		},
		"test2": &Target{
			Description: "Test Description 2",
			Deps:        []string{"test1"},
			Actions:     []string{"echo 'test 2 {{.BIZ}}'"},
			Vars: map[string]string{
				"FOO": "bar",
				"BIZ": "baz",
			},
		},
	},
}

func writeYml() error {
	out, err := yaml.Marshal(&TestFile)
	if err != nil {
		return errors.Wrap(err, "unable to marshal yaml file")
	}

	if err := ioutil.WriteFile(HoneyFile+".yml", out, 0655); err != nil {
		return errors.Wrap(err, "unable to write yaml file")
	}

	return nil
}

func TestYamlFile(t *testing.T) {
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
		t.Error(errors.Wrap(err, "unable to load yaml test file"))
		return
	}

	if err := checkFileLoad(); err != nil {
		t.Error(err)
	}
}

func writeJson() error {
	out, err := json.Marshal(&TestFile)
	if err != nil {
		return errors.Wrap(err, "unable to marshal json file")
	}

	if err := ioutil.WriteFile(HoneyFile+".json", out, 0655); err != nil {
		return errors.Wrap(err, "unable to write json file")
	}

	return nil
}

func TestJsonFile(t *testing.T) {
	defer func() {
		if err := os.RemoveAll(HoneyFile + ".json"); err != nil && !os.IsNotExist(err) {
			t.Error(errors.Wrap(err, "unable to remove json test file"))
			t.FailNow()
		}
	}()

	if err := writeJson(); err != nil {
		t.Error(errors.Wrap(err, "unable to generate test json file"))
		return
	}

	if err := LoadFile(); err != nil {
		t.Error(errors.Wrap(err, "unable to load json test file"))
		return
	}

	if err := checkFileLoad(); err != nil {
		t.Error(err)
	}
}

func checkFileLoad() error {
	if want, got := TestFile.Version, hf.Version; want != got {
		return errors.Errorf("invalid version: wanted - %s ... got - %s", want, got)
	}

	if err := checkFileVars(TestFile.Vars, hf.Vars); err != nil {
		return err
	}

	for k, v := range TestFile.Targets {
		hft, ok := hf.Targets[k]
		if !ok {
			return errors.Errorf("invalid target: wanted - %+v ... got - %+v", v, hft)
		}

		if want, got := v.Description, hft.Description; want != got {
			return errors.Errorf("invalid description: wanted - %s ... got - %s", want, got)
		}

		if err := checkFileVars(v.Vars, hft.Vars); err != nil {
			return err
		}

		if want, got := len(v.Actions), len(hft.Actions); want != got {
			return errors.Errorf("invalid action count: wanted - %d ... got - %d", want, got)
		}

		for i, a := range v.Actions {
			if want, got := a, hft.Actions[i]; want != got {
				return errors.Errorf("invalid action: wanted - %s ... got - %s", want, got)
			}
		}
	}

	return nil
}

func checkFileVars(want, got map[string]string) error {
	for k, v := range want {
		if w, g := v, got[k]; w != g {
			return errors.Errorf("invalid var: wanted - %s ... got - %s", w, g)
		}
	}

	return nil
}
