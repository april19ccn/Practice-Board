package main

import "fmt"

type Point struct {
	X, Y int
	z    int // 不导出
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8, 0}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8, z:0}, Radius:5}, Spokes:20}

	w.X = 42

	w.z = 12 // 可以使用

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8, z:12}, Radius:5}, Spokes:20}

}
