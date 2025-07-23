// 练习 7.6： 对tempFlag加入支持开尔文温度。
package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

/*
//!+flagvalue
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
//!-flagvalue
*/

// *celsiusFlag satisfies the flag.Value interface.
// celsiusFlag内嵌了一个Celsius类型（§2.5），因此不用实现本身就已经有String方法了。
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	// 调用fmt.Sscanf函数从输入s中解析一个浮点数（value）和一个字符串（unit）。
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	fmt.Println(value, unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	// CommandLine 是默认的命令行标志集，从 os.Args 解析。顶级函数（如 [BoolVar]、[Arg] 等）是 CommandLine 方法的包装器。
	flag.CommandLine.Var(&f, name, usage)
	// flag.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

// go run .\ex-tempconv.go
// 20°C

// go run .\ex-tempconv.go -temp -18C
// -18°C

// go run .\ex-tempconv.go -temp 212°F
// 100°C

// go run .\ex-tempconv.go -temp 273.15K
// 0°C
