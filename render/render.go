// /home/krylon/go/src/github.com/blicero/pattern/render/render.go
// -*- mode: go; coding: utf-8; -*-
// Created on 09. 04. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2022-04-09 15:54:22 krylon>

package render

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
)

type Pt struct {
	X int
	Y int
}

func Render() error {
	const (
		size     = 1000
		stepCnt  = 25
		stepSize = size / stepCnt
		path     = "pattern.svg"
	)

	var (
		err    error
		fh     *os.File
		canvas *svg.SVG
		p1, p2 Pt
	)

	rand.Seed(time.Now().Unix())

	if fh, err = os.Create(path); err != nil {
		fmt.Fprintf(os.Stderr,
			"Cannot open %s: %s\n",
			path,
			err.Error())
		os.Exit(1)
	}

	defer fh.Close()

	canvas = svg.New(fh)
	canvas.Start(size, size)
	defer canvas.End()

	p1.X = rand.Intn(size)
	p1.Y = rand.Intn(size)
	p2.X = rand.Intn(size)
	p2.Y = rand.Intn(size)

	// Gather edge points
	var (
		edgePoints = make([]Pt, 0, stepCnt*4)
		p          Pt
	)

	for i := 0; i <= stepCnt; i++ {
		p = Pt{X: i * stepSize, Y: 0}
		edgePoints = append(edgePoints, p)
	}

	for i := 0; i <= stepCnt; i++ {
		p = Pt{X: i * stepSize, Y: size}
		edgePoints = append(edgePoints, p)
	}

	for i := 0; i <= stepCnt; i++ {
		p = Pt{X: 0, Y: i * stepSize}
		edgePoints = append(edgePoints, p)
	}

	for i := 0; i <= stepCnt; i++ {
		p = Pt{X: size, Y: i * stepSize}
		edgePoints = append(edgePoints, p)
	}

	for _, p := range edgePoints {
		canvas.Line(p1.X, p1.Y, p.X, p.Y, "fill:black;color:black;stroke:black")
		canvas.Line(p2.X, p2.Y, p.X, p.Y, "fill:black;color:black;stroke:black")
	}

	return nil
}
