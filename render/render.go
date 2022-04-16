// /home/krylon/go/src/github.com/blicero/pattern/render/render.go
// -*- mode: go; coding: utf-8; -*-
// Created on 09. 04. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2022-04-16 15:38:09 krylon>

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
		size     = 2000
		stepCnt  = 16
		stepSize = size / stepCnt
		path     = "pattern.svg"
		style    = "fill:black;color:black;stroke:black"
		pCnt     = 2
	)

	var (
		err    error
		fh     *os.File
		canvas *svg.SVG
		points = make([]Pt, pCnt)
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

	for i := 0; i < pCnt; i++ {
		points[i].X = rand.Intn(size)
		points[i].Y = rand.Intn(size)
	}

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
		for _, v := range points {
			canvas.Line(v.X, v.Y, p.X, p.Y, style)
		}
	}

	return nil
}
