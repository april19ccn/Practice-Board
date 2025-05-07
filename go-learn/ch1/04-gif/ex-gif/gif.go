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

var palette = []color.Color{color.White, color.Black, color.RGBA{0x3f, 0xc1, 0xf1, 0xff}} // 一个slice切片

// 常量声明和变量声明一般都会出现在包级别，所以这些常量在整个包中都是可以共享的
const (
	whiteIndex = 0 // first color in palette 调色板中的第一种颜色
	blackIndex = 2 // next color in palette 调色板中的下一个颜色
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.

	// rand.Seed(time.Now().UTC().UnixNano())
	// lissajous(os.Stdout)

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

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // 一个struct结构体
	phase := 0.0                        // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
