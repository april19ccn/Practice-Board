// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Params struct {
	cycles  float64
	res     float64
	size    int
	nframes int
	delay   int
}

var palette = []color.Color{color.White, color.Black, color.RGBA{0x3f, 0xc1, 0xf1, 0xff}} // 一个slice切片

// 常量声明和变量声明一般都会出现在包级别，所以这些常量在整个包中都是可以共享的
const (
	whiteIndex = 0 // first color in palette 调色板中的第一种颜色
	blackIndex = 2 // next color in palette 调色板中的下一个颜色
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := getParams(r.URL.Query())
		lissajous(w, params)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getParams(q url.Values) Params {
	cycles, _ := strconv.ParseFloat(q.Get("cycles"), 64)
	res, _ := strconv.ParseFloat(q.Get("res"), 64)
	size, _ := strconv.Atoi(q.Get("size"))
	nframes, _ := strconv.Atoi(q.Get("nframes"))
	delay, _ := strconv.Atoi(q.Get("delay"))

	fmt.Fprintf(os.Stdout, "cycles=%v\n", cycles)
	fmt.Fprintf(os.Stdout, "res=%v\n", res)
	fmt.Fprintf(os.Stdout, "size=%v\n", size)
	fmt.Fprintf(os.Stdout, "nframes=%v\n", nframes)
	fmt.Fprintf(os.Stdout, "delay=%v\n", delay)

	return Params{
		cycles,
		res,
		size,
		nframes,
		delay,
	}
}

func lissajous(out io.Writer, params Params) {
	// 常量声明定义在函数体内部，那么这种常量就只能在函数体内用
	var (
		cycles  float64 = 5     // number of complete x oscillator revolutions 振荡器完整的 x 转数
		res     float64 = 0.001 // angular resolution 角度分辨率
		size    int     = 100   // image canvas covers [-size..+size] 图像画布封面
		nframes int     = 64    // number of animation frames 动画帧数
		delay   int     = 8     // delay between frames in 10ms units 帧与帧之间的延迟，以 10ms 为单位
	)

	if params.cycles != 0 {
		cycles = params.cycles
	}
	if params.res != 0 {
		res = params.res
	}
	if params.size != 0 {
		size = params.size
	}
	if params.nframes != 0 {
		nframes = params.nframes
	}
	if params.delay != 0 {
		delay = params.delay
	}

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // 一个struct结构体
	phase := 0.0                        // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
