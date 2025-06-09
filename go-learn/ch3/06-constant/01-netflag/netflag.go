package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // 00001 (1)  // is up
	FlagBroadcast                      // 00010 (2)  // supports broadcast access capability
	FlagLoopback                       // 00100 (4)  // is a loopback interface
	FlagPointToPoint                   // 01000 (8)  // belongs to a point-to-point link
	FlagMulticast                      // 10000 (16) // supports multicast access capability
)

// 检查标志位
func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

// v & FlagUp：保留 FlagUp 对应的位（最低位），其他位清零
// 若结果等于 FlagUp (1)，说明该标志已设置

// 清除标志位
func TurnDown(v *Flags) { *v &^= FlagUp }

// &^ 运算符：按位清除（AND NOT）
//   *v &^= FlagUp 等价于 *v = *v & ^FlagUp
//   将 FlagUp 对应的位强制清零，其他位保持不变

// 设置标志位
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }

// |= 运算符：按位或赋值
// 	 将 FlagBroadcast 对应的位设为 1，不影响其他位

// 检查多个标志
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }

// 组合检测：先组合掩码 FlagBroadcast|FlagMulticast（二进制 10010）
// 通过 v & mask 检查是否至少有一个标志被设置
// 结果非 0 表示至少有一个标志存在

func main() {
	var v Flags = FlagMulticast | FlagUp // 10000 | 00001 = 10001 (17)
	fmt.Printf("%b %t\n", v, IsUp(v))    // 输出: 10001 true

	TurnDown(&v)                      // 清除 FlagUp: 10001 → 10000
	fmt.Printf("%b %t\n", v, IsUp(v)) // 输出: 10000 false

	SetBroadcast(&v)                    // 设置 FlagBroadcast: 10000 | 00010 = 10010
	fmt.Printf("%b %t\n", v, IsUp(v))   // 输出: 10010 false
	fmt.Printf("%b %t\n", v, IsCast(v)) // 输出: 10010 true
}
