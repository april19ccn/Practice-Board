// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White} // 一个slice切片

func init() {
	for i := 0; i < 255; i++ { // 最大256
		palette = append(palette, color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 0xFF,
		})
	}
}

func main() {
	create, _ := os.Create("out.gif")
	lissajous(create)
}

func lissajous(out io.Writer) {
	// 常量声明定义在函数体内部，那么这种常量就只能在函数体内用
	const (
		cycles  = 5     // number of complete x oscillator revolutions 振荡器完整的 x 转数
		res     = 0.001 // angular resolution 角度分辨率
		size    = 100   // image canvas covers [-size..+size] 图像画布封面
		nframes = 64    // number of animation frames 动画帧数
		delay   = 8     // delay between frames in 10ms units 帧与帧之间的延迟，以 10ms 为单位
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette) // palette color.Color 类型的切片，并且其长度最大为 256
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
