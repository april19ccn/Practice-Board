// 验证 map 迭代顺序
// map 的迭代顺序并不确定，从实践来看，该顺序随机，每次运行都会变化。
// 这种设计是有意为之的，因为能防止程序依赖特定遍历顺序，而这是无法保证的。
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, 74.39967,
	}
	m["test"] = Vertex{
		12.00, 100,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}
