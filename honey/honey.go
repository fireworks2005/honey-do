//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/honey/honey.go
//
package honey

import (
	"github.com/elliottpolk/honey-do/log"
)

const HoneyFile string = "Honeyfile"

func Run(targets []string) {
	if err := LoadFile(); err != nil {
		log.Errorf(err, "unable to load file Honeyfile(.yml|.json)")
		return
	}

	//  if no target specified, run all
	if len(targets) == 0 || (len(targets) == 1 && targets[0] == "all") {
		for n, t := range hf.Targets {
			if err := t.Run(); err != nil {
				log.Errorf(err, "unable to run target %s", n)
				continue
			}

			log.Infof("target %s complete", n)
		}

		return
	}

	for _, want := range targets {
		t, ok := hf.Targets[want]
		if !ok {
			log.NewError("invalid target %s", want)
			continue
		}

		if err := t.Run(); err != nil {
			log.Errorf(err, "unable to run target %s", want)
			continue
		}

		log.Infof("target %s complete", want)
	}
}
