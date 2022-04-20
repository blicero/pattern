// /home/krylon/go/src/github.com/blicero/pattern/render/render.go
// -*- mode: go; coding: utf-8; -*-
// Created on 09. 04. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2022-04-20 11:43:44 krylon>

// Package render implements rendering images.
package render

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
	"github.com/blicero/pattern/render/color"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const (
	size      = 2000
	path      = "pattern.svg"
	lineStyle = "fill:black;color:black;stroke:black"
)

// Pt represents a point in a two-dimensional coordinate system
type Pt struct {
	X int
	Y int
}

func rndPt(x, y int) Pt {
	return Pt{
		X: rand.Intn(x),
		Y: rand.Intn(y),
	}
}

// RenderRays renders a pattern of rays.
func RenderRays() error {
	const (
		stepCnt  = 200
		stepSize = size / stepCnt
		pCnt     = 2
	)

	var (
		err    error
		fh     *os.File
		canvas *svg.SVG
		points = make([]Pt, pCnt)
	)

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
		// points[i].X = rand.Intn(size)
		// points[i].Y = rand.Intn(size)
		points[i] = rndPt(size, size)
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
			canvas.Line(v.X, v.Y, p.X, p.Y, lineStyle)
		}
	}

	return nil
} // func RenderRays() error

func RenderCircles() error {
	const (
		circleCnt = 25
		maxRadius = size / 8
	)

	var (
		err    error
		fh     *os.File
		canvas *svg.SVG
	)

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

	for i := 0; i < circleCnt; i++ {
		var (
			radius = rand.Intn(maxRadius)
			ct     = rndPt(size, size)
			col    = color.Colors[rand.Intn(len(color.Colors))]
			style  = fmt.Sprintf("fill:%s;stroke:%s",
				col,
				col)
		)

		fmt.Printf("Circle #%d: %s\n",
			i+1,
			col)

		canvas.Circle(ct.X, ct.Y, radius, style)
	}

	return nil
} // func RenderCircles() error
