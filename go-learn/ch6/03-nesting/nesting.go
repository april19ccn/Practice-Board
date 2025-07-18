package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	// 使用匿名成员访问 X、Y
	cp := ColoredPoint{Point{1, 2}, color.RGBA{0xff, 0x00, 0x00, 0xff}}

	cp.X = 1
	fmt.Println(cp.Point.X)

	cp.Y = 2
	fmt.Println(cp.Point.Y)

	// 使用匿名成员访问 Distance、ScaleBy
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
}
