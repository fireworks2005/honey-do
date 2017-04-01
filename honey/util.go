//  Created by Elliott Polk on 01/04/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/util.go
//
package honey

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

func vars(local map[string]string) map[string]string {
	res := make(map[string]string)

	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		res[kv[0]] = kv[1]
	}

	for k, v := range hf.Vars {
		res[k] = v
	}

	for k, v := range local {
		res[k] = v
	}

	return res
}

func enrich(content string, vars interface{}) (string, error) {
	tpl := template.Must(template.New("cmd").Parse(content))

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		return "", errors.Wrap(err, "unable to execute text template")
	}

	return buf.String(), nil
}
