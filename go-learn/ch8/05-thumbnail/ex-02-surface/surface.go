// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
	start := time.Now()
	OutSVG()
	// OutGoSVG()
	// OutGoSVG2()
	elapsed := time.Since(start)

	fmt.Printf("Execution time: %v\n", elapsed)
}

func OutSVG() {
	out, err := os.Create("./out.svg")
	if err != nil {
		log.Fatal(err) // 处理创建文件时的错误
	}
	defer out.Close() // 确保函数退出时文件会被关闭

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

type polygonPoints struct {
	ax, ay, bx, by, cx, cy, dx, dy float64
	i, j                           int // 用于排序
}

func OutGoSVG() {
	out, err := os.Create("./out.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	// 创建channel用于传递计算结果
	ch := make(chan polygonPoints, cells*cells)

	// 使用WaitGroup等待所有worker完成
	var wg sync.WaitGroup

	// 启动worker goroutines
	numWorkers := 6 // 可以调整这个值来找到最优的goroutine数量
	wg.Add(numWorkers)
	for w := 0; w < numWorkers; w++ {
		go worker(ch, w, numWorkers, &wg)
	}

	// 启动一个goroutine等待所有worker完成然后关闭channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 收集结果并写入文件
	pointsSlice := make([]polygonPoints, 0, cells*cells)
	for p := range ch {
		pointsSlice = append(pointsSlice, p)
	}

	// 按原始顺序排序
	sort.Slice(pointsSlice, func(i, j int) bool {
		if pointsSlice[i].i != pointsSlice[j].i {
			return pointsSlice[i].i < pointsSlice[j].i
		}
		return pointsSlice[i].j < pointsSlice[j].j
	})

	// 写入SVG
	for _, p := range pointsSlice {
		fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy)
	}

	fmt.Fprintf(out, "</svg>")
}

// worker函数处理一部分网格点的计算
func worker(ch chan<- polygonPoints, workerID, numWorkers int, wg *sync.WaitGroup) {
	defer wg.Done() // 确保worker完成时通知WaitGroup

	for i := workerID; i < cells; i += numWorkers {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			ch <- polygonPoints{ax, ay, bx, by, cx, cy, dx, dy, i, j}
		}
	}
}

type indexedPoints struct {
	workerID int
	points   []polygonPoints
}

func OutGoSVG2() {
	out, err := os.Create("./out.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	// 使用WaitGroup等待所有worker完成
	var wg sync.WaitGroup

	// 启动worker goroutines
	numWorkers := 6 // 可以调整这个值来找到最优的goroutine数量

	// 创建channel用于传递计算结果
	ch := make(chan indexedPoints, numWorkers)

	// 计算每个worker处理的行数
	rowsPerWorker := cells / numWorkers
	if cells%numWorkers != 0 {
		rowsPerWorker++ // 如果不能整除，最后一个worker会少处理一些行
	}

	// 启动worker
	for w := 0; w < numWorkers; w++ {
		startRow := w * rowsPerWorker
		if startRow >= cells { // 添加这个检查
			break
		}
		endRow := (w + 1) * rowsPerWorker
		if endRow > cells {
			endRow = cells
		}

		wg.Add(1)
		go worker2(ch, w, startRow, endRow, &wg)
	}

	// 启动一个goroutine等待所有worker完成然后关闭channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 收集结果到切片中
	results := make([][]polygonPoints, numWorkers)
	for ip := range ch {
		results[ip.workerID] = ip.points
	}

	// 按顺序处理结果
	for _, points := range results {
		for _, p := range points {
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy)
		}
	}

	fmt.Fprintf(out, "</svg>")
}

func worker2(ch chan<- indexedPoints, workerID, startRow, endRow int, wg *sync.WaitGroup) {
	defer wg.Done()

	points := make([]polygonPoints, 0, (endRow-startRow)*cells)

	for i := startRow; i < endRow; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			points = append(points, polygonPoints{ax, ay, bx, by, cx, cy, dx, dy, i, j})
		}
	}

	// 发送带索引的结果
	ch <- indexedPoints{workerID: workerID, points: points}
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
