// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 234.

// Package cake provides a simulation of
// a concurrent cake shop with numerous parameters.
//
// Use this command to run the benchmarks:
//
//	$ go test -bench=. gopl.io/ch8/cake

// 生产流程
// 蛋糕制作分为三个阶段：

// 烘焙 - 由baker函数完成
// 上糖衣 - 由icer函数完成
// 刻字 - 由inscriber函数完成
package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool          // 是否打印详细日志
	Cakes          int           // number of cakes to bake 烘烤的蛋糕数量
	BakeTime       time.Duration // time to bake one cake 烘焙蛋糕的时间
	BakeStdDev     time.Duration // standard deviation of baking time 烘焙蛋糕的标准差
	BakeBuf        int           // buffer slots between baking and icing  烘焙和上糖衣之间的缓冲区大小
	NumIcers       int           // number of cooks doing icing 负责上糖衣的厨师数量
	IceTime        time.Duration // time to ice one cake 上糖衣的时间
	IceStdDev      time.Duration // standard deviation of icing time  上糖衣的标准差
	IceBuf         int           // buffer slots between icing and inscribing 上糖衣和刻字之间的缓冲区大小
	InscribeTime   time.Duration // time to inscribe one cake 刻字的时间
	InscribeStdDev time.Duration // standard deviation of inscribing time 刻字的标准差
}

type cake int

// 烘焙
func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i) // 创建对应编号的蛋糕
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev) // 每个蛋糕经过随机时间烘焙（基于BakeTime和BakeStdDev）
		baked <- c
	}
	close(baked)
}

// 上糖衣
func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked { // 从baked中接收蛋糕
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

// 刻字
func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

// work blocks the calling goroutine for a period of time
// that is normally distributed around d
// with a standard deviation of stddev.
// 辅助函数，模拟工作所需的时间
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}

// Work runs the simulation 'runs' times.
// 运行模拟的主函数
// baker -> [baked channel] -> icer(s) -> [iced channel] -> inscriber
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)

		// 多个goroutine同时运行（一个baker，多个icer，一个inscriber）
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}
