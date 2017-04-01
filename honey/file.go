//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/file.go
//
package honey

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type File struct {
	Version string             `json:"version" yaml:"version"`
	Vars    map[string]string  `json:"vars,omitempty" yaml:"vars,omitempty"`
	Targets map[string]*Target `json:"targets" yaml:"targets"`
}

var hf *File

func LoadFile() error {
	hf = &File{}
	if content, err := ioutil.ReadFile(HoneyFile + ".yml"); err == nil {
		return yaml.Unmarshal(content, &hf)
	}

	if content, err := ioutil.ReadFile(HoneyFile + ".json"); err == nil {
		return json.Unmarshal(content, &hf)
	}

	return os.ErrNotExist
}
