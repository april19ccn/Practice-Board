// 等角投影
// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height float64 = 600, 320                     // canvas size in pixels
	xyscale       float64 = float64(width / 2 / xyrange) // pixels per x or y unit
	zscale        float64 = float64(height) * 0.4        // pixels per z unit
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// 练习3.1
func isInvalid(num float64) bool {
	return math.IsNaN(num) || math.IsInf(num, 0)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")

		queryWidth, err := strconv.ParseFloat(r.URL.Query().Get("width"), 64)
		if err == nil {
			width = queryWidth
		}
		queryHeight, err := strconv.ParseFloat(r.URL.Query().Get("height"), 64)
		if err == nil {
			height = queryHeight
		}
		xyscale = float64(width / 2 / xyrange) // pixels per x or y unit
		zscale = float64(height) * 0.4         // pixels per z unit

		getSvg(w)
	})
	fmt.Println("Server Start: localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getSvg(out io.Writer) {
	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%f' height='%f'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, isPeek1 := corner(i+1, j)
			bx, by, isPeek2 := corner(i, j)
			cx, cy, isPeek3 := corner(i, j+1)
			dx, dy, isPeek4 := corner(i+1, j+1)

			// 练习3.1
			if isInvalid(ax) || isInvalid(ay) {
				continue
			}

			if isInvalid(bx) || isInvalid(by) {
				continue
			}

			if isInvalid(cx) || isInvalid(cy) {
				continue
			}

			if isInvalid(dx) || isInvalid(dy) {
				continue
			}

			var color = "red"
			if isPeek1 || isPeek2 || isPeek3 || isPeek4 {
				color = "red"
			} else {
				color = "blue"
			}

			svg += fmt.Sprintf("<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svg += fmt.Sprintf("</svg>")

	out.Write([]byte(svg))
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z > 0
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
