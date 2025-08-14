// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	doPaint(1)
	doPaint(4)
	doPaint(8)
	doPaint(16)
	doPaint(32)
	doPaint(64)
	doPaint(128)

}

// 剔除了读写文件，大量时间被节省，只是为了验证并行计算的差异
func doPaint(routinues int) {
	work := make(chan int, cells)
	var wg sync.WaitGroup
	for r := routinues; r > 0; r-- {
		wg.Add(1)
		go func() {
			for i := range work {
				for j := 0; j < cells; j++ {
					corner(i+1, j)
					corner(i, j)
					corner(i, j+1)
					corner(i+1, j+1)
				}
			}
			wg.Done()
		}()
	}
	start := time.Now()
	for i := 0; i < cells; i++ {
		work <- i
	}
	close(work)
	wg.Wait()
	fmt.Printf("routines is %d, timecostis %s\n ", routinues, time.Since(start).String())
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
