//  Created by Elliott Polk on 31/03/2017
//  Copyright © 2017. All rights reserved.
//  honey-do/cmd/honey/main.go
//
package main

import (
	"github.com/elliottpolk/honey-do/honey"

	"github.com/spf13/pflag"
)

func main() {
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()

	honey.Run(pflag.Args())
}
