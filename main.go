// /home/krylon/go/src/github.com/blicero/pattern/main.go
// -*- mode: go; coding: utf-8; -*-
// Created on 09. 04. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2022-04-09 09:28:33 krylon>

package main

import (
	"fmt"
	"os"

	"github.com/blicero/pattern/render"
)

func main() {
	var err error

	if err = render.Render(); err != nil {
		fmt.Fprintf(os.Stdout,
			"Cannot render pattern: %s\n",
			err.Error())
		os.Exit(1)
	}
} // func main()
